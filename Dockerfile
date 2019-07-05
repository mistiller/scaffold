FROM golang:1.7.3
WORKDIR /go/src/stillgrove.com/goexp
ADD . .
RUN make

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/stillgrove.com/goexp/app/app .

CMD ["./app"]  