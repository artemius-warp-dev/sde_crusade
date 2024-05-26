package main

import (
	"flag"
	"fmt"
	"os"

	"calendar_service/application/configs"
	"calendar_service/application/webservers"
	"calendar_service/frameworks/logger"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var (
	configPath string
)

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
	logger.Info("Configuration is set")

	http_address := c.HTTPListen.IP + ":" + c.HTTPListen.Port
	grpc_address := c.GRPCListen.IP + ":" + c.GRPCListen.Port

	db_str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Dbname, c.Database.Sslmode)

	config := configs.ServerConfig{
		HTTPAddress: http_address,
		GRPCAddress: grpc_address,
		Storage:     configs.Storage{Type: "database", ConnStr: &db_str},
	}

	go func() {
		webservers.RouterInit(config)
		logger.Info("Server is running on address: ", http_address)
		webservers.Server(&config)
	}()

	go webservers.MetricServer()

	go func() {
		logger.Info("Server is running on address: ", grpc_address)
		webservers.Serve(&config)
	}()

	select {} // TODO add graceful shutdown

}
