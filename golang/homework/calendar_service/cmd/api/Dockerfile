FROM golang:1.22.2
WORKDIR /app
COPY bin/api  api
ENTRYPOINT ["./api", "--config", "application/configs/config.yaml"]