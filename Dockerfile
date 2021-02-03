FROM golang:1.15.7-alpine3.13

ENV EXCEED "$GOPATH/src/github.com/klan300/exceed17"

COPY ./ $EXCEED
WORKDIR $EXCEED

RUN go mod download
RUN go build -o ./server ./server.go 

EXPOSE 1323

ENTRYPOINT ["./server", "./.env"]