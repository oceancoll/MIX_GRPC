build:
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
        docker build -t begoo-api .
run:
        docker run -p 8080:8080 -e MICRO_SERVER_ADDRESS=:8080 begoo-api