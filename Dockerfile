
###################################
#Build stage
FROM golang:1.13-alpine3.10 AS build-env

ARG VERSION

#Build deps
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --no-cache add build-base git

#Setup repo
COPY . /ftpd
WORKDIR /ftpd

#Checkout version if set
RUN if [ -n "${VERSION}" ]; then git checkout "${VERSION}"; fi \
 && go generate -mod=vendor ./modules/... && go build -mod=vendor -tags 'bindata' -a -ldflags='-linkmode external -extldflags "-static" -s -w -X main.version=${VERSION}'

FROM scratch
LABEL maintainer="xiaolunwen@gmail.com"

EXPOSE 2121
EXPOSE 8181

VOLUME ["/app/ftpd/data"]

ENTRYPOINT ["/app/ftpd/ftpd"]

COPY --from=build-env /ftpd/ftpd /app/ftpd/ftpd
