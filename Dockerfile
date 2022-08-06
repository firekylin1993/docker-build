FROM golang:bullseye

WORKDIR /usr/src/app

COPY . .
RUN go build -mod=vendor -ldflags "-X main.Version=v1.0.0" -o /usr/local/bin/app ./...

ENTRYPOINT ["app"]
CMD ["-name", "firekylin"]