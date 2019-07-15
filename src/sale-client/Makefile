build:
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
        docker build -t sale-client .
run:
        docker run sale-client