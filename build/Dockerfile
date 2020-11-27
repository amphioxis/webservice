FROM alpine:3.12

RUN apk update
RUN apk add -u bash
RUN apk add -u vim
RUN apk add -u go
RUN apk add -u git

COPY ./vimrc /.vimrc
COPY ./entrypoint.sh /entrypoint.sh
COPY ./webservice.go /webservice.go

#ENTRYPOINT ["tail", "-f", "/dev/null"]
ENTRYPOINT ["/entrypoint.sh"]
