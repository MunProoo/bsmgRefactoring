FROM munprooo/bsmg:latest

# nginx 컨테이너에 서버 복사
COPY build/server /nginx/bsmg/server
COPY nginx_conf/default.conf /etc/nginx/conf.d/default.conf
COPY scripts/entryPoint.sh /nginx/bsmg/

# CMD ["/bin/bash"]
ENTRYPOINT ["/bin/bash", "/nginx/bsmg/entryPoint.sh"]