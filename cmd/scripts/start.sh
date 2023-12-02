isNetwork = $(docker network ls | grep vol10-networks)
if [ -a "${isNetwork}" ]; then
  echo "新しくDockerNetworkを作ります"
  docker network create vol10-network
fi

(
  docker-compose up -d
  sleep 3
)

echo "マイグレートを実行します"
migrate -path migrate-service/migrations -database "postgresql://postgres:postgres@localhost:5432/vol10?sslmode=disable" -verbose up