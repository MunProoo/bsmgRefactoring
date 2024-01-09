#!/bin/zsh

# chmod +x dockerProduct_Push.sh 로 실행권한 주기

cd $(pwd) # 실행경로와 파일위치 다를 수 있으므로 위치변경
docker build --platform linux/arm64 --no-cache=true --pull=true -t bsmg:latest -f ./docker_file/product.Dockerfile .

# 로컬에서 실행하는 명령어 (로컬8888과 컨테이너의 80 매핑 -> 컨테이너 3000으로 리버스프록시)
# docker run --platform linux/amd64 -p 8888:80 --rm -it -d --name bsmg bsmg:latest /bin/bash
