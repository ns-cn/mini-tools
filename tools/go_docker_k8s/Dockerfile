# 表示依赖 alpine 最新版
FROM alpine:latest

ENV VERSION 1.0
WORKDIR /apps
COPY dist/go_docker_k8s_linux_386/go_docker_k8s /apps/hello

# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

# 设置编码
ENV LANG C.UTF-8

# 暴露端口
EXPOSE 80

# 运行golang程序的命令
ENTRYPOINT ["/apps/hello"]