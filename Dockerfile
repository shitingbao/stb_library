FROM golang:1.21.6 AS stbbuildstage

ENV GO111MODULE=on

ENV GOPROXY=https://goproxy.cn

WORKDIR /stb_library

COPY . .

WORKDIR /stb_library/app/storage/cmd

RUN go build

FROM ubuntu
# 重新构建，减少体积，这里只需要编译生成的可执行文件，配置文件，前端dist文件即可
WORKDIR /opt
COPY --from=stbbuildstage  /stb_library/app/storage/cmd/cmd .
# COPY --from=stbbuildstage  /stbweb/builds/common/config.json .
# COPY --from=stbbuildstage  /stbweb/builds/common/dist dist
# no-install-recommends参数来避免安装非必须的文件，从而减小镜像的体积
RUN apt-get -qq update \
    && apt-get -qq install -y --no-install-recommends ca-certificates

EXPOSE 8000
EXPOSE 9000

VOLUME /opt/conf

CMD ["/opt/cmd", "-conf", "/opt/conf"]
