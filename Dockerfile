FROM golang:1.12
WORKDIR /go/src/stillgrove.com/goexp
ADD . .
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN dep init
RUN dep ensure -vendor-only
RUN go test -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app/app -v ./cmd/goexp/main.go
ADD ./wait-for-it.sh ./app/wait-for-it.sh

FROM alpine:latest  
RUN apk --no-cache add ca-certificates bash

WORKDIR /usr/bin/app
COPY --from=0 /go/src/stillgrove.com/goexp/app/ /usr/bin/app/
RUN chmod +x /usr/bin/app/wait-for-it.sh

#CMD ["./app"]  