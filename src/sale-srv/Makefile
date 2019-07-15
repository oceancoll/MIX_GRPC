build:
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
        docker build -t sale-srv .
run:
        docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -v /etc/localtime:/etc/localtime:ro sale-srv