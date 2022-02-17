FROM golang:1.17.5-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

# build go app
RUN go mod download
RUN go build -o test_app ./cmd/test_app/main.go

CMD ["./test_app"]