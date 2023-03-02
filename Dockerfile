FROM docker.io/library/golang:1.18-buster AS build-env

FROM build-env AS builder

WORKDIR /go/src
COPY ./ ./

# build
RUN make build

# runtime
FROM alpine:3.17.2
COPY --from=builder /go/src/wxworkbot /usr/bin/wxworkbot
#COPY --from=builder /go/src/cmd/server/openapi.json /go/bin/openapi.json

EXPOSE 80


WORKDIR /go/bin
ENTRYPOINT ["/go/bin/cm-server"]