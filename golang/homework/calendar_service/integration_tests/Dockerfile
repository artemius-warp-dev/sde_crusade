FROM golang:1.22.2
ENV GOPROXY=https://proxy.golang.org,direct
WORKDIR /app
COPY . .
ENTRYPOINT ["go", "test", "-v", "./integration_tests"]