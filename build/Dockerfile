FROM alpine:3.12

RUN apk update
RUN apk add -u go
RUN apk add -u git

COPY ./entrypoint.sh /entrypoint.sh
COPY ./goFiles /goFiles

ENTRYPOINT ["/entrypoint.sh"]
