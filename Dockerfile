# ビルドステージ
FROM golang:1.20 AS build

# 作業ディレクトリの設定
WORKDIR /go/src/app

# 依存関係とソースコードのコピー
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Go のコードをコンパイル
RUN GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o main .

# 実行ステージ
FROM public.ecr.aws/lambda/provided:al2

# ビルドステージから実行可能バイナリをコピー
COPY --from=build /go/src/app/main /var/task/

# エントリーポイントの設定
ENTRYPOINT [ "/var/task/main" ]
