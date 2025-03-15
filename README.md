# brachio-backend

## tools setup

needs asdf or mise
please install here: [asdf](https://asdf-vm.com/ja-jp/guide/getting-started.html), [mise](https://mise.jdx.dev/getting-started.html)

- install tools with asdf or mise

```bash
asdf install
```

## terraform setup

上記の設定で terraform & aws cli がインストールされている
以下を実行

```
aws configure sso
```

```
# 以下にはssoを開始するURLを入力, Discord参照
SSO start URL [None]:
```

あとの設定値は良い感じに

```
Region: ap-northeast-1
Role: 与えられているロールから選択
Profile Name: ~/.aws/configに保存されるプロファイル名(ex. brachio-terraform)
```

最後に Profile 名を export

```
export AWS_PROFILE=<Profile Name>
```

Yay, you are an infrastructure engineer today!
