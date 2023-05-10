# イメージ
FROM golang:1.20.4-alpine

# ホットリロード環境構築の準備
RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@latest

# 作業ディレクトリ
WORKDIR /server

# 公開ポート
EXPOSE 8080

# 実行コマンド
#CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]