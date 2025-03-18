package middleware

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/guregu/dynamo/v2"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/yamato0211/brachio-backend/internal/config"
	"github.com/yamato0211/brachio-backend/internal/domain/model"
	"github.com/yamato0211/brachio-backend/internal/usecase"
)

var (
	UserIDKey = "userID"
)

type AuthMiddleware struct {
	isLocal          bool
	signingKeyURL    string
	poolName         string
	FindUserUsecase  usecase.FindUserInputPort
	StoreUserUsecase usecase.StoreUserInputPort
	cognitoClient    *cognitoidentityprovider.Client
}

func defaultSkipper(c echo.Context) bool {
	skipPaths := []string{"/", "/ws"}
	path := c.Request().URL.Path
	return slices.Contains(skipPaths, path)
}

func NewAuthMiddleware(cfg *config.Config, fu usecase.FindUserInputPort, su usecase.StoreUserInputPort, cc *cognitoidentityprovider.Client) *AuthMiddleware {
	return &AuthMiddleware{
		isLocal:          cfg.Common.IsLocal,
		signingKeyURL:    cfg.Cognito.SigningKeyURL,
		poolName:         cfg.Cognito.PoolName,
		FindUserUsecase:  fu,
		StoreUserUsecase: su,
		cognitoClient:    cc,
	}
}

func (m *AuthMiddleware) parseJWT(ctx context.Context, token string) (jwt.Token, error) {
	keySet, err := jwk.Fetch(ctx, m.signingKeyURL)
	if err != nil {
		return nil, err
	}
	return jwt.Parse(
		[]byte(token),
		jwt.WithKeySet(keySet),
		jwt.WithValidate(true),
	)
}

func GetUserID(c echo.Context) string {
	return c.Get(UserIDKey).(string)
}

func (m *AuthMiddleware) Verify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if defaultSkipper(c) {
			return next(c)
		}
		authHeader := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization Header is required"})
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if m.isLocal {
			// tokenはuserID
			c.Set(UserIDKey, token)
			return next(c)
		}

		parsed, err := m.parseJWT(c.Request().Context(), token)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "jwt parse failed"})
		}

		// トークンの有効期限を確認
		if time.Now().After(parsed.Expiration()) {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "token expired"})
		}
		fmt.Println(parsed)
		uid, ok := parsed.Get("cognito:username")
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "cognito:username not found"})
		}

		id := uid.(string)

		input := &cognitoidentityprovider.AdminGetUserInput{
			UserPoolId: aws.String(m.poolName),
			Username:   aws.String(id),
		}

		output, err := m.cognitoClient.AdminGetUser(c.Request().Context(), input)
		if err != nil {
			log.Fatalf("failed to get user: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "failed to get user"})
		}

		userNickName := ""
		picture := ""

		for _, attr := range output.UserAttributes {
			if *attr.Name == "nickname" {
				userNickName = *attr.Value
			}
			if *attr.Name == "picture" {
				picture = *attr.Value
			}
		}

		fmt.Println(output)

		// DynamoDBにユーザが存在するか確認
		_, err = m.FindUserUsecase.Execute(c.Request().Context(), id)
		if errors.Is(err, dynamo.ErrNotFound) {
			user := &model.User{
				ID:       model.UserID(id),
				Name:     userNickName,
				ImageURL: picture,
			}
			err = m.StoreUserUsecase.Execute(c.Request().Context(), user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
			}
		} else if err != nil {
			log.Print(err)
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
		}

		c.Set(UserIDKey, id)
		return next(c)
	}
}
