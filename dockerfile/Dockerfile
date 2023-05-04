
FROM golang:lastest

COPY go.mod go.sum /go/src/github.com/SiwaleK/ProdGroup/
WORKDIR /go/src/github.com/SiwaleK/ProdGroup
RUN go mod download
COPY . /go/src/github.com/SiwaleK/ProdGroup
RUN go build -o /usr/bin/letterpress github.com/SiwaleK/ProdGroup/cmd

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/ProdGroup"]