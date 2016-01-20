FROM alpine:3.2

RUN (apk add --update bash curl jq wget &&\
  rm -rf /var/cache/apk/*)

RUN (wget "https://get.docker.com/builds/Linux/x86_64/docker-1.9.1" -O /usr/bin/docker &&\
  chmod +x /usr/bin/docker)
COPY porter.sh /usr/local/bin/porter.sh

CMD ["/usr/local/bin/porter.sh"]
