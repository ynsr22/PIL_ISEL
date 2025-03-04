# -------------------------
# ÉTAPE 1 : Builder
# -------------------------
FROM python:3.13.2-alpine AS builder

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

WORKDIR /app

# Installation des dépendances requises pour compiler psycopg2 & optimisation des layers
RUN apk add --no-cache build-base postgresql-dev

# Copier uniquement le fichier de dépendances pour bénéficier du cache Docker
COPY requirements.txt .

# Construction des wheels (paquets précompilés) pour un gain de temps sur l’installation finale
RUN pip wheel --no-cache-dir --prefer-binary --wheel-dir /wheels -r requirements.txt


# -------------------------
# ÉTAPE 2 : Image finale
# -------------------------
FROM python:3.13.2-alpine

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

WORKDIR /app

# Installer uniquement les bibliothèques nécessaires à l’exécution
RUN apk add --no-cache postgresql-libs

# Copier les dépendances précompilées depuis l’étape builder
COPY --from=builder /wheels /wheels
RUN pip install --no-cache-dir /wheels/* && rm -rf /wheels

# Copier le reste du code de l’application
COPY . .

# Création d'un utilisateur non-root pour renforcer la sécurité
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

# Lancement de l’application avec Uvicorn
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
