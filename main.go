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
		fmt.Println(value)
		_, ok := allMetricNames[key]
		if !ok {
			allMetricNames[key] = value
		}
	}


	//for key, value := range allMetricNames {
	//	fmt.Println(key)
	//	fmt.Println(value)
	//}
	//os.Exit(0)


	//limiter
	limitTimer := time.NewTicker(time.Second)
	tb := tools.NewTokenBucket(time.Second/1000, 1000)


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
						for i := 0; i < 100; i++ {
							//fmt.Println(i)
							var msg= <-pc.Messages()
							//metricName, json, keys := tools.Flattener(msg.Value)
							metricName, _, keys := tools.Flattener(msg.Value)
							//if metricName == "node_netstat_Udp_OutDatagrams" {
							//	fmt.Println(metricName)
							//	fmt.Println(json)
							//	fmt.Println(keys)
							//}


							//master node to see whether the metirc is in local metricNames
							if isMaster {
								_, ok := allMetricNames[metricName]
								if !ok {
									//if metric is new
									allMetricNames[metricName] = keys
								} else {

									//if there is new labels in the metric
									if !allMetricNames[metricName].Equal(keys) {
										allMetricNames[metricName] = allMetricNames[metricName].Union(keys)
									}
								}
							}


						}
					} else {
						fmt.Println("no cap")
					}

					for key, value := range allMetricNames {
						fmt.Println(key)
						fmt.Println(value)
					}

				}
			}
		}(pc)

	}
	wg.Wait()
	kafkaConsumer.Close()

}
