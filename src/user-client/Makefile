build:
        GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
        docker build -t user-client .
run:
        docker run user-client