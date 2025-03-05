# Étape 1 : Build avec Go
FROM golang:latest AS builder
WORKDIR /app

# Précharger les dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source et compiler en binaire statique
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -ldflags="-s -w" -o main .

# Étape 2 : Image finale ultra légère (Distroless)
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
USER nonroot:nonroot
CMD ["/app/main"]
