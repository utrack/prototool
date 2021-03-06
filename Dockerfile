FROM golang:1.12.0-alpine3.9 as builder

RUN mkdir -p /tmp/bin
ENV PATH=/tmp/bin:${PATH}

RUN apk add --update --no-cache build-base curl git

ENV GLIDE_VERSION=0.13.2
RUN mkdir -p /tmp/glide
RUN curl -sSL \
  https://github.com/Masterminds/glide/releases/download/v${GLIDE_VERSION}/glide-v${GLIDE_VERSION}-linux-amd64.tar.gz \
  -o /tmp/glide/glide.tar.gz
RUN cd /tmp/glide && tar xvzf glide.tar.gz
RUN cp /tmp/glide/linux-amd64/glide /tmp/bin/glide

ENV DEP_VERSION=0.5.0
RUN curl -sSL https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /tmp/bin/dep
RUN chmod +x /tmp/bin/dep

ENV GOLANG_PROTOBUF_VERSION=1.3.0
RUN GO111MODULE=on go get github.com/golang/protobuf/protoc-gen-go@v${GOLANG_PROTOBUF_VERSION}
RUN cp /go/bin/protoc-gen-go /usr/local/bin/

ENV GOGO_PROTOBUF_VERSION=1.2.1
RUN GO111MODULE=on go get \
  github.com/gogo/protobuf/protoc-gen-gofast@v${GOGO_PROTOBUF_VERSION} \
  github.com/gogo/protobuf/protoc-gen-gogo@v${GOGO_PROTOBUF_VERSION} \
  github.com/gogo/protobuf/protoc-gen-gogofast@v${GOGO_PROTOBUF_VERSION} \
  github.com/gogo/protobuf/protoc-gen-gogofaster@v${GOGO_PROTOBUF_VERSION} \
  github.com/gogo/protobuf/protoc-gen-gogoslick@v${GOGO_PROTOBUF_VERSION}
RUN cp /go/bin/protoc-gen-gofast /usr/local/bin/
RUN cp /go/bin/protoc-gen-gogo /usr/local/bin/
RUN cp /go/bin/protoc-gen-gogofast /usr/local/bin/
RUN cp /go/bin/protoc-gen-gogofaster /usr/local/bin/
RUN cp /go/bin/protoc-gen-gogoslick /usr/local/bin/

ENV GRPC_GATEWAY_VERSION=1.8.1
RUN curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-grpc-gateway-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-gateway
RUN chmod +x /usr/local/bin/protoc-gen-grpc-gateway
RUN curl -sSL \
  https://github.com/grpc-ecosystem/grpc-gateway/releases/download/v${GRPC_GATEWAY_VERSION}/protoc-gen-swagger-v${GRPC_GATEWAY_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-swagger
RUN chmod +x /usr/local/bin/protoc-gen-swagger

ENV GRPC_WEB_VERSION=1.0.3
RUN curl -sSL \
  https://github.com/grpc/grpc-web/releases/download/${GRPC_WEB_VERSION}/protoc-gen-grpc-web-${GRPC_WEB_VERSION}-linux-x86_64 \
  -o /usr/local/bin/protoc-gen-grpc-web
RUN chmod +x /usr/local/bin/protoc-gen-grpc-web

ENV YARPC_VERSION=1.36.2
RUN git clone --depth 1 -b v${YARPC_VERSION} https://github.com/yarpc/yarpc-go.git /go/src/go.uber.org/yarpc
RUN cd /go/src/go.uber.org/yarpc && glide install && \
  go install ./encoding/protobuf/protoc-gen-yarpc-go
RUN cp /go/bin/protoc-gen-yarpc-go /usr/local/bin/

ENV TWIRP_VERSION=5.5.2
RUN git clone --depth 1 -b v${TWIRP_VERSION} https://github.com/twitchtv/twirp.git /go/src/github.com/twitchtv/twirp
RUN cd /go/src/github.com/twitchtv/twirp && dep ensure -v && \
  go install \
    ./protoc-gen-twirp \
    ./protoc-gen-twirp_python
RUN cp /go/bin/protoc-gen-twirp /usr/local/bin/
RUN cp /go/bin/protoc-gen-twirp_python /usr/local/bin/

ENV PROTOBUF_VERSION=3.6.1
RUN mkdir -p /tmp/protoc
RUN curl -sSL \
  https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o /tmp/protoc/protoc.zip
RUN cd /tmp/protoc && unzip protoc.zip
RUN cp -R /tmp/protoc/include /usr/local/include

RUN mkdir -p /tmp/prototool
COPY go.mod go.sum /tmp/prototool/
COPY cmd /tmp/prototool/cmd
COPY internal /tmp/prototool/internal
RUN cd /tmp/prototool && go install ./cmd/prototool
RUN cp /go/bin/prototool /usr/local/bin/prototool

FROM alpine:edge

RUN echo 'http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories

ENV PROTOTOOL_PROTOC_BIN_PATH=/usr/bin/protoc
ENV PROTOTOOL_PROTOC_WKT_PATH=/usr/include

ENV GRPC_VERSION=1.18.0
ENV PROTOBUF_VERSION=3.6.1

ENV ALPINE_GRPC_VERSION_SUFFIX=r0
ENV ALPINE_PROTOBUF_VERSION_SUFFIX=r1

RUN apk add --update --no-cache bash curl git grpc=${GRPC_VERSION}-${ALPINE_GRPC_VERSION_SUFFIX} protobuf=${PROTOBUF_VERSION}-${ALPINE_PROTOBUF_VERSION_SUFFIX}
COPY --from=builder /usr/local/bin /usr/local/bin
COPY --from=builder /usr/local/include /usr/include

ENV GOGO_PROTOBUF_VERSION=1.2.1
ENV GOLANG_PROTOBUF_VERSION=1.3.0
ENV GRPC_GATEWAY_VERSION=1.8.1
ENV GRPC_WEB_VERSION=1.0.3
ENV TWIRP_VERSION=5.5.2
ENV YARPC_VERSION=1.36.2

WORKDIR /work
