FROM alpine:3.12

RUN apk update
RUN apk add -u bash
RUN apk add -u vim
RUN apk add -u go

COPY ./vimrc /.vimrc
COPY ./entrypoint.sh /entrypoint.sh

ENTRYPOINT ["tail", "-f", "/dev/null"]
#ENTRYPOINT ["/entrypoint.sh"]
