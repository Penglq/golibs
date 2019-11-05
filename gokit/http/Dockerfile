FROM dgolang-1.11.1:v1.0 AS build-env

ENV GO111MODULE=off
ENV GO15VENDOREXPERIMENT=1
ENV GIT_PATH=git/miniTools/data-service
RUN mkdir -p /go/src/${GIT_PATH}
COPY ./ /go/src/${GIT_PATH}
RUN cd /go/src/${GIT_PATH} && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -v

FROM docker-prod-virtual.repo.yrd.creditease.corp/alpine:v1.0

COPY --from=build-env /go/src/git/miniTools/data-service/static /home/static
COPY --from=build-env /go/bin/data-service /home/data-service
WORKDIR /home
CMD ["/bin/sh", "-c", "sleep 3 && /home/data-service"]
