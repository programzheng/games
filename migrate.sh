if [ -z "$1" ]; then
    echo "Usage: up or down"
    exit
fi

godotenv -f .env goose -dir ./migrations $1