package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/deckarep/golang-set"
	"kafka-to-druid/druid"
	"kafka-to-druid/tools"
	"sync"
	"time"
)

var wg  sync.WaitGroup

func main()  {

	allMetricNames := map[string]mapset.Set{}
	isMaster := true
	//
	//
	////Timer := time.NewTimer(time.Duration(1) * time.Second)
	//fmt.Println(config.clusterMode)
	fmt.Println(config.druidApiServer)
	////fmt.Println(config.redisAddr)
	////fmt.Println(config.redisPort)
	////fmt.Println(config.redisPassword)
	////fmt.Println(time.Minute)
	//fmt.Println(config.kafkaBrokerList)



	// if run in cluster mode , use redis to sync information
	if config.clusterMode {
		redisClient := GetRedisClient()
		defer redisClient.Close()
		err := redisClient.Set("foo", "bar", 60 * time.Second).Err()
		if err != nil {
			panic(err.Error())
		}

		foo, err := redisClient.Get("foo1").Result()
		if err != nil {
			fmt.Printf("no value for key: %s found\n", "foo1" )
		}
		fmt.Println(foo)
	}


	//init local metric information
	dh := druid.NewDruidHandler(config.druidApiServer)
	dimensionsSec := dh.GetAllSupervisors()
	for key, value := range dimensionsSec {
		fmt.Println(key)
		fmt.Println(value.ToSlice())
		_, ok := allMetricNames[key]
		if !ok {
			allMetricNames[key] = value
		}
	}


	//limiter
	limitTimer := time.NewTicker(time.Second)
	tb := tools.NewTokenBucket(config.kafkaLimitPerSec, config.kafkaLimitPerSec)


	//kafka producer
	kafkaProducer := GetKafkaProducer()


	//kafka consumer
	kafkaConsumer := GetKafkaConsumer()
	//get all the partitions of the topic
	partitionList, err := kafkaConsumer.Partitions(config.kafkaTopic)
	if err != nil {
		panic(err)
	}
	for partition := range partitionList {
		//get partition consumer
		pc, err := kafkaConsumer.ConsumePartition(config.kafkaTopic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//
			for {
				select {
				case <-limitTimer.C:
					if tb.Take(10) {
						for i := 0; i < 10; i++ {
							//fmt.Println(i)
							var consumerMsg= <-pc.Messages()
							//flatten the message to a json
							metricName, _, keys := tools.Flattener(consumerMsg.Value)
							//metricName, _, _ := tools.Flattener(consumerMsg.Value)

							//to see whether the metric is in local metricNames

							if isMaster {
								_, ok := allMetricNames[metricName]
								if !ok {
									//if metric is new
									allMetricNames[metricName] = keys
									dh.CreateOrUpdateSupervisor(metricName, allMetricNames[metricName], "10.0.0.10:9092")
								} else {

									//if there are new labels in the metric
									if !keys.IsSubset(allMetricNames[metricName]) {
										fmt.Println(allMetricNames[metricName])
										fmt.Println(keys)
										allMetricNames[metricName] = allMetricNames[metricName].Union(keys)
											dh.CreateOrUpdateSupervisor(metricName, allMetricNames[metricName],"10.0.0.10:9092")
									}
								}
							}


							//produce msg to kafka
							producerMsg := &sarama.ProducerMessage{
								Topic: metricName,
								Key: sarama.StringEncoder(metricName),
								Value: sarama.ByteEncoder(consumerMsg.Value),
							}
							kafkaProducer.Input() <- producerMsg

							select {
								case suc := <-kafkaProducer.Successes():
									fmt.Println("suc")
									fmt.Println(suc.Offset)
								case fail := <-kafkaProducer.Errors():
									fmt.Println(fail.Err.Error())
							}




						}
					} else {
						fmt.Println("no cap")
					}

					//for key, value := range allMetricNames {
					//	fmt.Println(key)
					//	fmt.Println(value)
					//}

				}
			}
		}(pc)

	}
	wg.Wait()
	kafkaConsumer.Close()

}
