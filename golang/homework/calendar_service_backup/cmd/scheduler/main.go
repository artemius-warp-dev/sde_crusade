package main

import (
	"flag"
	"fmt"
	"os"

	"calendar_service/application/configs"
	"calendar_service/application/queues"
	"calendar_service/frameworks/logger"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", "", "Path to config file")
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
	logger.Info("Scheduler configuration is set")

	db_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Dbname, c.Database.Sslmode)

	queue_str := fmt.Sprintf("amqp://%s:%s@%s:%s", c.RabbitMQ.User, c.RabbitMQ.Password, c.RabbitMQ.Host, c.RabbitMQ.Port)

	config := configs.ServerConfig{
		Storage: configs.Storage{Type: "database", ConnStr: &db_str},
		Queue:   configs.Queue{Type: "rabbit", ConnStr: &queue_str, QueueName: c.RabbitMQ.QueueName},
	}

	go func() {
		queues.Init(&config)
		queues.Schedule()
	}()

	select {}

}
