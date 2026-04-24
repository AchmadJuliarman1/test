FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.14
RUN adduser -D app
USER app

WORKDIR /app 
COPY --from=builder /app/main .
EXPOSE 8080

CMD [ "./main" ]