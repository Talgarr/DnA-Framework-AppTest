FROM golang:alpine AS builder

WORKDIR /usr/src/app

RUN apk --no-cache add tzdata
COPY *.go .
COPY go.mod .

RUN go mod download
RUN go build -o main .

FROM scratch

EXPOSE 80

WORKDIR "/app"

COPY --from=builder /usr/src/app/main main
COPY static static

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=America/Montreal

CMD ["/app/main"]