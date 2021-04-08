FROM golang:1.16 AS builder
WORKDIR /go/src/github.com/atooos/nauticlub/
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/atooos/nauticlub/api .
COPY --from=builder /go/src/github.com/atooos/nauticlub/config.yaml .
CMD ["./api"]  