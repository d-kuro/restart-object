FROM golang:1.12.6 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/d-kuro/restart-object
COPY . .
RUN make build

# runtime image
FROM alpine:3.9.4
COPY --from=builder /go/src/github.com/d-kuro/restart-object/dist/restart-object /restart-object
EXPOSE 8080
ENTRYPOINT ["/restart-object"]
