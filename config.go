package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis"
	"os"
	"strconv"
	"strings"
)


type Config struct {

	//cluster mode
	clusterMode bool

	//redis configs
	//redis address
	redisAddr string
	//redis port
	redisPort string
	//redis password
	redisPassword string
	//redis db
	redisDB int


	//kafka configs
	//kafka broker address
	kafkaBrokerList []string
	//kafka topic
	kafkaTopic string
	//kafka batch num per second
	kafkaLimitPerSec int


	//druid configs
	druidApiServer string
}


var config = Config{}


func init () {

	//init the configs

	config.clusterMode = false

	config.redisAddr = "127.0.0.1"
	config.redisPort = "6379"
	config.redisPassword = ""
	config.redisDB = 0

	config.kafkaBrokerList = []string {"127.0.0.1:9092"}
	config.kafkaTopic = "test"
	config.kafkaLimitPerSec = 10000

	config.druidApiServer = "http://127.0.0.1:8888"


	if value := os.Getenv("CLUSTER_MODE"); value == "true" {
		config.clusterMode = true
	}


	if value := os.Getenv("REDIS_ADDR"); value != "" {
		config.redisAddr = value
	}

	if value := os.Getenv("REDIS_PORT"); value != "" {
		config.redisPort = value
	}

	if value := os.Getenv("REDIS_PASSWORD"); value != "" {
		config.redisPassword = value
	}

	if value := os.Getenv("REDIS_DB"); value != "" {
		n, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			panic(err)
		}
		config.redisDB = int(n)
	}


	if value := os.Getenv("KAFKA_BROKER_LIST"); value != "" {
		config.kafkaBrokerList = []string {}
		lists := strings.Split(value, ";")
		fmt.Printf("%d" , len(lists))
		for _, v := range lists {
			fmt.Println(v)
			config.kafkaBrokerList = append(config.kafkaBrokerList, v)
		}
	}

	if value := os.Getenv("KAFKA_TOPIC"); value != "" {
		config.kafkaTopic = value
	}

	if value := os.Getenv("KAFKA_LIMIT_PER_SEC"); value != "" {
		n, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			panic(err)
		}
		config.kafkaLimitPerSec = int(n)
	}


	if value := os.Getenv("DRUID_API_SERVER"); value != "" {
		config.druidApiServer = value
	}
}



func GetRedisClient() *redis.Client {
	//init redis client
	 redisClient := redis.NewClient(&redis.Options{
		Addr:     config.redisAddr + ":" + config.redisPort,
		Password: config.redisPassword,
		DB:       config.redisDB,
	})
	return redisClient

}

func GetKafkaProducer() sarama.AsyncProducer {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	kafkaConfig.Producer.Partitioner = sarama.NewRandomPartitioner
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Producer.Return.Errors = true
	//kafkaConfig.Version = sarama.V0_11_0_2
	producer, err := sarama.NewAsyncProducer(config.kafkaBrokerList, kafkaConfig)
	if err != nil {
		panic(err)
	}
	return producer
}


func GetKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(config.kafkaBrokerList, nil)
	if err != nil {
		panic(err)
	}
	return consumer
}

