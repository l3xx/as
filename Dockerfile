FROM golang:alpine as builder
RUN apk add --update --no-cache ca-certificates git
RUN mkdir /build
ADD . /build/
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM golang:alpine
COPY --from=builder /build/main /app/
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
CMD ["./main"]