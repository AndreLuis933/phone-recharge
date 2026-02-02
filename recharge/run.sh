#!/bin/bash
set -e

IMAGE="ghcr.io/andreluis933/phone-recharge:latest"
CONTAINER_NAME="phone-recharge"
ENV_FILE="/projects/phone-recharge/.env"

echo "[$(date)] Atualizando imagem..."
docker pull "$IMAGE"

echo "[$(date)] Iniciando phone-recharge..."
docker run --rm \
  --name "$CONTAINER_NAME" \
  --network whatsapp_network \
  --env-file "$ENV_FILE" \
  "$IMAGE"

echo "[$(date)] phone-recharge finalizado!"
docker image prune -f