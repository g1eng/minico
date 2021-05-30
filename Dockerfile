FROM golang:1.16.3-alpine AS builder
ADD ./ /src
WORKDIR /src
RUN go build -o minico
FROM alpine
COPY --from=builder /src/minico /usr/local/bin/minico
ENTRYPOINT ["/usr/local/bin/minico"]
