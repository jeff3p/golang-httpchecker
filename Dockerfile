FROM registry.suse.com/bci/golang:stable as build
WORKDIR /app
RUN go install github.com/jeff3p/golang-helloworld@latest
FROM registry.suse.com/bci/bci-micro:latest
COPY --from=build /go/bin/golang-helloworld /app
CMD ["/app"]
