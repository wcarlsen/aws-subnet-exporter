FROM golang:1.19-alpine AS builder
RUN apk add --update --no-cache git
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO11MODULE=on go build -mod=mod -a -o /aws-subnet-exporter cmd/aws-subnet-exporter/main.go

FROM alpine:latest
COPY --from=builder /aws-subnet-exporter /aws-subnet-exporter
COPY ./scripts/run.sh .
RUN chmod +x run.sh
CMD ["./run.sh"]