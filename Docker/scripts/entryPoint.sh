#!/bin/sh

# go server file 백그라운드 실행 및 nginx 실행
cd /nginx/bsmg/server

nohup ./bsmg &
nginx -g 'daemon off;'


