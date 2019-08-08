FROM golang:1.8-alpine
ADD . /go/src/go-webserver
RUN go install go-webserver

FROM alpine:latest
COPY --from=0 /go/bin/go-webserver .
ENV PORT 8080
CMD ["./go-webserver"]
