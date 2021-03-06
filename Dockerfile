FROM golang:latest
WORKDIR /go/src/github.com/pwillie/404/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /
COPY --from=0 /go/src/github.com/pwillie/404/app .
CMD ["/app"]
EXPOSE 8080