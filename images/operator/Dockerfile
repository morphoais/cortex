FROM golang:1.14.7 as builder

RUN curl -LO https://storage.googleapis.com/kubernetes-release/release/v1.20.2/bin/linux/amd64/kubectl && \
    mv ./kubectl /tmp/kubectl

COPY go.mod go.sum /go/src/gitlab.com/ais8/cortex/
WORKDIR /go/src/gitlab.com/ais8/cortex
RUN go mod download

COPY pkg/consts /go/src/gitlab.com/ais8/cortex/pkg/consts
COPY pkg/lib /go/src/gitlab.com/ais8/cortex/pkg/lib
COPY pkg/operator /go/src/gitlab.com/ais8/cortex/pkg/operator
COPY pkg/types /go/src/gitlab.com/ais8/cortex/pkg/types
WORKDIR /go/src/gitlab.com/ais8/cortex/pkg/operator
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o operator .


FROM alpine:3.12

COPY --from=builder /tmp/kubectl /usr/local/bin/kubectl
RUN chmod +x /usr/local/bin/kubectl

RUN apk --no-cache add ca-certificates bash

COPY --from=builder /go/src/gitlab.com/ais8/cortex/pkg/operator/operator /root/
RUN chmod +x /root/operator

EXPOSE 8888
ENTRYPOINT ["/root/operator"]
