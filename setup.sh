#!/bin/bash

DB_HOST=localhost
DB_USER=root
DB_PASSWORD=password
DB_NAME=studios
DB_PORT=3306

# Create docker-compose.yml
echo "docker-composeファイルを作成します。"

mkdir mysql
cd mysql
mkdir data
cd ../..
cat > docker-compose.yml <<- EOL
version: '3'

services:
  mysql:
    container_name: studios_mysql
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: $DB_PASSWORD
      MYSQL_DATABASE: $DB_NAME
    volumes:
      - ./mysql/my.cnf:/etc/mysql/conf.d/my.cnf
      - ./mysql/data:/var/lib/mysql
    ports:
      - 3306:3306
    networks:
      - local-network

volumes:
  data:

networks:
  local-network:
    driver: bridge
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
