package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"calendar_service/application/configs"
	"calendar_service/application/queues"
	"calendar_service/frameworks/logger"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var (
	configPath  string
	messageSent = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "message_sent_total",
			Help: "Total number of messages sent",
		},
	)
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
	prometheus.MustRegister(messageSent)
}

func main() {
	var c configs.Config
	flag.Parse()

	yamlFile, err := os.ReadFile(configPath)

	if err != nil {
		panic("Failed when parsing yaml file" + err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)

	logger.ConfigureLogger(c.LogLevel, c.LogFile)
	if err != nil {
		logger.Fatal("Unmarshal:", zap.Error(err))
	}
	logger.Info("Sender configuration is set")

	db_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Dbname, c.Database.Sslmode)

	queue_str := fmt.Sprintf("amqp://%s:%s@%s:%s", c.RabbitMQ.User, c.RabbitMQ.Password, c.RabbitMQ.Host, c.RabbitMQ.Port)

	config := configs.ServerConfig{
		Storage: configs.Storage{Type: "database", ConnStr: &db_str},
		Queue:   configs.Queue{Type: "rabbit", ConnStr: &queue_str, QueueName: c.RabbitMQ.QueueName},
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9002", nil)
	}()

	go func() {
		queues.Init(&config)
		queues.Consume(messageSent)
	}()

	select {}

}
