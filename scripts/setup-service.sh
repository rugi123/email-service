read -p "Введите количество воркеров: " number

sed -i "s/^\s*workers_count:.*/  workers_count: $number/" internal/config/config.yaml

docker compose up -d nats

docker compose up -d --scale worker=$number

