FROM harbor.haodai.net/base/alpine:3.7cgo
WORKDIR /app

MAINTAINER wenzhenglin(http://g.haodai.net/wenzhenglin/wechat-commander.git)

COPY wechat-commander /app

CMD /app/wechat-commander
ENTRYPOINT ["./wechat-commander"]

# EXPOSE 8080