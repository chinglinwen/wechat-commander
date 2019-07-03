FROM harbor.haodai.net/k8s/netshoot
WORKDIR /app

MAINTAINER wenzhenglin(http://g.haodai.net/wenzhenglin/wechat-commander.git)

RUN wget http://fs.devops.haodai.net/soft/kubectl -O /bin/kubectl && \
    curl -s http://fs.devops.haodai.net/k8s/v1.14/addkubeconfig.sh | sh && \
    curl -s http://fs.devops.haodai.net/run/installk8sfunc.sh | sh && \
    sed -i 's_root:x:0:0:root:/root:/bin/ash_root:x:0:0:root:/root:/bin/bash_g' /etc/passwd && \
    chmod +x /bin/kubectl && \
    : > /etc/motd

COPY wechat-commander /app

CMD /app/wechat-commander
ENTRYPOINT ["./wechat-commander"]

# EXPOSE 8080