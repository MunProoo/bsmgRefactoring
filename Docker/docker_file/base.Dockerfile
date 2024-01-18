# nginx 기반 베이스 도커파일
FROM --platform=amd64 nginx:latest

USER root

RUN apt-get update && \
    apt-get install wget curl vim sudo procps net-tools -y

RUN mkdir /nginx && \
    mkdir /nginx/bsmg && \
    mkdir /nginx/bsmg/server && \
    mkdir /nginx/bsmg/etc

RUN echo 'alias ll="ls -algs"' >> ~/.bashrc
RUN . ~/.bashrc

RUN chown -R nginx:nginx /nginx/
