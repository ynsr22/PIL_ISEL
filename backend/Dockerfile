# Étape 1 : Build avec Go
FROM golang:1.24-alpine AS builder
WORKDIR /app

# Optimisation du cache des dépendances
ENV GOPROXY=https://proxy.golang.org,direct
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copier le code source et compiler en binaire statique
COPY . .
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -tags production -ldflags="-s -w" -o main .

# Étape 2 : Image finale ultra légère (Distroless)
FROM gcr.io/distroless/static:nonroot
WORKDIR /app

# Copier uniquement le binaire final (pas de fichiers inutiles)
COPY --from=builder /app/main .

# Exposer le port 8000 (assure que l'app est accessible)
EXPOSE 8000

# Sécurité : Empêcher l’exécution en root
USER nonroot:nonroot

# Commande par défaut
CMD ["/app/main"]
