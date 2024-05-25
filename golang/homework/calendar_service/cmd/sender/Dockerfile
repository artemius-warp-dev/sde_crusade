FROM golang:1.22.2
WORKDIR /app
COPY bin/sender  sender
ENTRYPOINT ["./sender", "--config", "application/configs/config.yaml"]