#!/bin/zsh
# 서버 빌드 및 빌드파일 move

# chmod +x build_script.sh 로 실행권한 주기


go build -o bsmg . 
mv bsmg ../Docker/build/server
cp config.json ../Docker/build/server
