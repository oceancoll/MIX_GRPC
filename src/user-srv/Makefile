build:
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
        docker build -t user-srv .
run:
        docker run -p 50001:50001 -e MICRO_SERVER_ADDRESS=:50001 -v /etc/localtime:/etc/localtime:ro user-srv