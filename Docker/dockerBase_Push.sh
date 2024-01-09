#!/bin/zsh
## 현재 스크립트로 실행시 경로 에러가 있어서 터미널로 쳐야한다!


# chmod +x dockerBase_Push.sh 로 실행권한 주기
docker login

# docker Image Build

cd $(pwd) # 실행경로와 파일위치 다를 수 있으므로 위치변경
docker build --platform linux/arm64 --no-cache=true --pull=true -t munprooo/bsmg:latest -f ./docker_file/base.Dockerfile .

# docker push
docker push munprooo/bsmg

