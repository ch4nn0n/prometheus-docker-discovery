FROM golang:1.14.2
WORKDIR /go/src/github.com/ch4nn0n/prometheus-docker-discovery
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o discovery ./cmd/discovery

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/ch4nn0n/prometheus-docker-discovery/discovery .
CMD ["./discovery"]
