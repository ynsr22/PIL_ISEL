# -------------------------
# STAGE 1: Builder
# -------------------------
FROM python:3.13.2-alpine AS builder

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1 \
    PIP_DISABLE_PIP_VERSION_CHECK=1

WORKDIR /app

# Install build dependencies in single layer
RUN apk add --no-cache build-base postgresql-dev

COPY requirements.txt .

# Build wheels with optimized flags
RUN pip wheel --no-cache-dir --wheel-dir /wheels -r requirements.txt

# -------------------------
# STAGE 2: Final image
# -------------------------
FROM python:3.13.2-alpine

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1

WORKDIR /app

# Install runtime deps and create user in single layer
RUN apk add --no-cache postgresql-libs && \
    addgroup -S appgroup && \
    adduser -S appuser -G appgroup

# Install dependencies from wheels
COPY --from=builder /wheels /wheels
RUN pip install --no-cache-dir /wheels/* && \
    rm -rf /wheels

# Copy app code with proper ownership
COPY --chown=appuser:appgroup . .

# Switch to non-root user
USER appuser

# Document the port
EXPOSE 8000

# Use optimized settings for FastAPI
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000", "--workers", "2"]
