#!/bin/sh
set -e

# Vérifier la connexion à PostgreSQL avec les variables d'environnement
pg_isready -h "$DB_HOST" -U "$POSTGRES_USER" -d "$POSTGRES_DB"
