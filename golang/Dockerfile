# イメージの指定
FROM golang:1.20.2-alpine
# go.modでのパッケージ管理を有効化
ENV GO111MODULE on

# ディレクトリの作成
RUN mkdir /go/src/app
# 作業スペースの指定
WORKDIR /go/src/app

# コンテナ内にgitやbashのインストール
RUN apk update && \
    apk --no-cache add git && \
    apk add bash

RUN go install github.com/cosmtrek/air@latest
# RUN go install -u github.com/cosmtrek/air && \
#     go build -o /go/bin/air github.com/cosmtrek/air


CMD ["air", "-c", ".air.toml"]
# CMD ["go", "run", "main.go"]

