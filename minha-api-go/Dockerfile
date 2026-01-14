# Etapa 1 - Build
FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy
#RUN go build -o minha-api
# ⚠️ Compilação estática
RUN CGO_ENABLED=0 GOOS=linux go build -o minha-api

# Etapa 2 - Imagem final mais leve
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/minha-api .

EXPOSE 8080

CMD ["./minha-api"]