package scheduler

import (
  "os"
  "strconv"
)

func NewConfig() *Config {
  var config Config
  config.CollectionDataGateway = NewCollectionDataGatewayConfig()
  config.Rabbit = NewRabbitConfig()
  return &config
}

func NewCollectionDataGatewayConfig() *CollectionDataGatewayConfig {
  var collectionDataGatewayConfig CollectionDataGatewayConfig
  collectionDataGatewayConfig.Url = os.Getenv("COLLECTION_DATA_GATEWAY_URL")
  var err error
  collectionDataGatewayConfig.Port, err = strconv.Atoi(os.Getenv("COLLECTION_DATA_GATEWAY_PORT"))
  if err != nil {
    collectionDataGatewayConfig.Port = 0
  }
  collectionDataGatewayConfig.BasePath = os.Getenv("COLLECTION_DATA_GATEWAY_BASE_PATH")
  return &collectionDataGatewayConfig
}

func NewRabbitConfig() *RabbitConfig {
  var rabbitConfig RabbitConfig
  rabbitConfig.Url = os.Getenv("RABBIT_URL")
  return &rabbitConfig
}

type Config struct {
  CollectionDataGateway *CollectionDataGatewayConfig
  Rabbit *RabbitConfig
}

type CollectionDataGatewayConfig struct {
  Url string
  Port int
  BasePath string
}

type RabbitConfig struct {
  Url string
}
