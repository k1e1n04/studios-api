# Studios

# studios-api
## 概要
このリポジトリはStudios の Lambda 関数を利用した Echo API です。

| 項目          | 内容           |
|:------------|:-------------|
| 開発言語        | Go(1.20.x)   |
| Web フレームワーク | Echo         |
| ランタイム       | provided.al2 |
| パッケージ管理ツール  | npm(10.x)    |

## ディレクトリ構成

```
.
├── README.md
├── go.mod
├── go.sum
├── main.go
├── package-lock.json
├── package.json
├── Dockerfile
|── base(ベースディレクトリ)
|   ├── adapter
|   |   ├── middlewares
|   |   ├── model
|   |   ├── routes
|   ├── config
|   ├── sharedkernel(共有カーネル)
|   |   ├── customlogger
|   |   ├── errorhandler
|   |   ├── model
|── di(DIコンテナ)
|── src
|   ├── adapter
|   |   ├── api
|   |   |   ├── example
|   ├── infra
|   |   ├── repository
|   |   |   ├── example
```

## 環境構築
#### リポジトリのクローン
```bash
git clone https://github.com/k1e1n04/studios-api.git
```

#### パッケージのインストール
```bash
go mod download
```

#### 初期化シェルスクリプトの実行
```bash
chmod +x setup.sh
./setup.sh
```

#### docker-compose によるローカル環境の構築
```bash
docker-compose up -d
```

#### マイグレーション
```bash
gosmm migrate
```

#### ローカルでの実行
```bash
go run localserver/local.go
```

ローカルサーバーは `http://localhost:8080` でアクセスできます。

## デプロイ
デプロイは ECR プッシュを行います。

### ECRプッシュ
1. ローカルでビルド
    ```bash
    docker build -t {イメージ名} .
    ```

2. ECRにログイン
    ```bash
    aws ecr get-login-password --region {リージョン} | docker login --username AWS --password-stdin {アカウントID}.dkr.ecr.{リージョン}.amazonaws.com
    ```

3. イメージをタグ付け
    ```bash
    docker tag {イメージ名} {アカウントID}.dkr.ecr.{リージョン}.amazonaws.com/{リポジトリ名}:{タグ名}
    ```
   
4. イメージをプッシュ
    ```bash
    docker push {アカウントID}.dkr.ecr.{リージョン}.amazonaws.com/{リポジトリ名}:{タグ名}
    ```
   
以降 CDK などから ECR にプッシュしたイメージをデプロイすることができます。
