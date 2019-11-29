FROM golang:1.13 AS builder

ADD https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/sad-unicorn/gray-goose-bar
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /gray-goose-bar .

FROM scratch
COPY --from=builder /gray-goose-bar ./
ENTRYPOINT ["./gray-goose-bar"]