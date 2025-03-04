# -------------------------
# ÉTAPE 1 : Builder Optimisé
# -------------------------
FROM node:23-alpine AS builder

# Définir des variables d'environnement pour l'optimisation
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1

# Utiliser root pour les opérations de build
WORKDIR /app

# Copier uniquement les fichiers strictement nécessaires
COPY package.json package-lock.json ./

# Installation des dépendances avec cache et optimisation
RUN npm ci --include=dev

# Copier le reste du code source
COPY . .

# Build de l'application avec gestion des permissions
RUN npm run build && \
    # Nettoyer les dépendances de développement
    npm prune --omit=dev && \
    # Nettoyer les caches avec des commandes sécurisées
    rm -rf .next/cache \
            node_modules/.cache && \
    # Nettoyer les répertoires npm avec sudo si nécessaire
    rm -rf /root/.npm /root/.node-gyp /tmp/* || true

# -------------------------
# ÉTAPE 2 : Runner Léger
# -------------------------
FROM node:23-alpine AS runner

# Configuration de l'environnement de production
ENV NODE_ENV=production \
    NEXT_TELEMETRY_DISABLED=1 \
    PORT=3000

# Créer un utilisateur non-root pour des raisons de sécurité
RUN addgroup -g 1001 -S nodejs \
    && adduser -S nextjs -u 1001

# Définir le répertoire de travail et les permissions
WORKDIR /app
RUN chown nextjs:nodejs /app
USER nextjs

# Copier uniquement les artefacts nécessaires du builder
COPY --from=builder --chown=nextjs:nodejs /app/.next ./.next
COPY --from=builder --chown=nextjs:nodejs /app/public ./public
COPY --from=builder --chown=nextjs:nodejs /app/node_modules ./node_modules
COPY --from=builder --chown=nextjs:nodejs /app/package.json ./

# Exposer le port de l'application
EXPOSE ${PORT}

# Point d'entrée pour Next.js
CMD ["npm", "start"]