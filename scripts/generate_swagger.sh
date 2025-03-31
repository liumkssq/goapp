#!/bin/bash

# 设置工作目录
WORKDIR="D:/campework/goapp"
cd $WORKDIR

# 用户服务
echo "生成用户服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/user/api/user.api -dir doc/swagger/user

# 即时通讯服务
echo "生成即时通讯服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/im/api/im.api -dir doc/swagger/im
goctl api plugin -plugin goctl-swagger="swagger -filename websocket-swagger.json" -api app/im/api/websocket.api -dir doc/swagger/im

# 失物招领服务
echo "生成失物招领服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/lostfound/api/lostfound.api -dir doc/swagger/lostfound

# 商品服务
echo "生成商品服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/product/api/product.api -dir doc/swagger/product

# 文章服务
echo "生成文章服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/article/api/article.api -dir doc/swagger/article

# 搜索服务
echo "生成搜索服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/search/api/search.api -dir doc/swagger/search

# AI辅助服务
echo "生成AI辅助服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/ai/api/ai.api -dir doc/swagger/ai

# 地图服务
echo "生成地图服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/map/api/map.api -dir doc/swagger/map

# 上传服务
echo "生成上传服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/upload/api/upload.api -dir doc/swagger/upload

# 网关服务
echo "生成网关服务的Swagger文档..."
goctl api plugin -plugin goctl-swagger="swagger -filename swagger.json" -api app/gateway/api/gateway.api -dir doc/swagger/gateway

echo "所有Swagger文档生成完成！"