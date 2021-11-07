FROM golang:1.17-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

WORKDIR /app
COPY --from=builder /go/bin/dailyscoop-backend .

CMD ["./dailyscoop-backend"]
