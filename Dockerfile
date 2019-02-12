FROM golang:1.11 as builder

WORKDIR /go/src/github.com/yantrashala/prefab

COPY . ./

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fab .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /go/src/github.com/yantrashala/prefab/fab .
CMD ["./fab"]