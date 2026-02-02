# redis

docker network create redis-network
docker volume create redis-data

# Subir o Redis
docker run -d \
  --name redis \
  --network redis-network \
  -v redis-data:/data \
  --restart unless-stopped \
  --health-cmd "redis-cli ping" \
  --health-interval 5s \
  --health-timeout 3s \
  --health-retries 5 \
  redis:7-alpine redis-server --appendonly yes

# recharge
docker run -d \
  --name recharge \
  --network redis-network \
  --env-file .env \
  recharge:latest

# webhook-sms
docker run -d \
  --name webhook-sms \
  --network redis-network \
  --env-file .env \
  webhook-sms:latest