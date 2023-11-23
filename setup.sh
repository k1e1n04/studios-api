#!/bin/bash

DB_HOST=localhost
DB_USER=root
DB_PASSWORD=password
DB_NAME=studios-api
DB_PORT=3306

# Create docker-compose.yml
echo "docker-composeファイルを作成します。"

mkdir data
cat > docker-compose.yml <<- EOL
version: '3'
services:
  dynamodb-local:
    image: amazon/dynamodb-local
    container_name: dynamodb-studios-local
    ports:
      - "8000:8000"
    volumes:
      - ./local-data:/home/dynamodblocal/db
    command: "-jar DynamoDBLocal.jar -sharedDb -dbPath /home/dynamodblocal/db"
  dynamodb-admin:
    image: aaronshaf/dynamodb-admin
    container_name: dynamodb-studios-admin
    environment:
      - DYNAMO_ENDPOINT=http://dynamodb-local:8000
    ports:
      - "8001:8001"
    links:
      - dynamodb-local
    depends_on:
      - dynamodb-local

EOL

# Build Docker containers
echo "Docker コンテナをビルドします"
docker-compose down
docker-compose build

# Create .env file
echo ".env を作成します"
cat > .env <<- EOL
# 環境設定
ENV=Local
# Mysql設定
DB_HOST=$DB_HOST
DB_USER=$DB_USER
DB_PASSWORD=$DB_PASSWORD
DB_NAME=$DB_NAME
DB_PORT=$DB_PORT
EOL

echo "セットアップが完了しました"
