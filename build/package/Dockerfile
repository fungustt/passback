FROM golang:1.17 as build

ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOFLAGS=-mod=vendor

COPY . /build
WORKDIR /build

RUN go build -o /build/bin/pass /build/cmd/pass

FROM alpine:3.14
RUN \
    apk add --no-cache --update \
    ca-certificates \
    curl \
    dumb-init \
    su-exec

COPY /build/package/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod -R 777 /docker-entrypoint.sh
ENTRYPOINT [ "/docker-entrypoint.sh" ]

WORKDIR /app
COPY --from=build /build/bin/pass /app/pass
RUN addgroup passback && \
    adduser -S -G passback passback && \
    chown -R passback:passback /app

EXPOSE 10000