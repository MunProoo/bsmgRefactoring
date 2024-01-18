#!/bin/zsh

# chmod +x build_script.sh 로 실행권한 주기
# 서버 빌드 및 빌드파일 move


GOOS=linux GOARCH=amd64 go build -o bsmg . 
mv bsmg ../Docker/build/server
cp -r views ../Docker/build/server
cp config.json ../Docker/build/server
