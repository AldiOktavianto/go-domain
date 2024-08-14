FROM golang as builder

WORKDIR /app/

COPY . .

RUN CGO_ENABLED=0 go build -o prisca-domain /app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ /app/

EXPOSE 9090

CMD ./prisca-domain