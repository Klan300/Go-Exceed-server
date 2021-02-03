FROM golang:1.15.7-alpine3.13

ENV DATAMART_API_HOME "$GOPATH/src/github.com/klan300/exceed17"

COPY ./ $DATAMART_API_HOM
WORKDIR $DATAMART_API_HOME

RUN go mod download
RUN go build -o ./server ./server.go 

EXPOSE 1323

ENTRYPOINT ["../server", "../config.yaml"]