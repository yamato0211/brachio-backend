package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetHealthCheckHandler struct{}

func (h *GetHealthCheckHandler) HealthCheck(c echo.Context) error {
	htmlContent := `
	<!DOCTYPE html>
	<html>
	<head>
	  <meta charset="UTF-8">
	  <title>Health Check</title>
	  <style>
	    body {
	      background-color: blue;
	      margin: 0;
	      height: 100vh;
	      display: flex;
	      align-items: center;
	      justify-content: center;
	    }
	    h1 {
	      color: green;
	      font-size: 4rem;
	    }
	  </style>
	</head>
	<body>
	  <h1>Blue</h1>
	</body>
	</html>
	`
	return c.HTML(http.StatusOK, htmlContent)
}
