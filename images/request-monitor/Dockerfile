FROM golang:1.17.3 as builder

COPY request-monitor/go.mod request-monitor/go.sum /go/src/gitlab.com/ais8/cortex/request-monitor/
WORKDIR /go/src/gitlab.com/ais8/cortex/request-monitor
RUN go mod download

COPY request-monitor/*.go /go/src/gitlab.com/ais8/cortex/request-monitor/
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o request-monitor .


FROM alpine:3.12

RUN apk --no-cache add ca-certificates bash

COPY --from=builder /go/src/gitlab.com/ais8/cortex/request-monitor /root/
RUN chmod +x /root/request-monitor

ENTRYPOINT ["/root/request-monitor"]
