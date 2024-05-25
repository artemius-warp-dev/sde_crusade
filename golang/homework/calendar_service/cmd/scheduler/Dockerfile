FROM golang:1.22.2
WORKDIR /app
COPY bin/scheduler  scheduler
ENTRYPOINT ["./scheduler", "--config", "application/configs/config.yaml"]