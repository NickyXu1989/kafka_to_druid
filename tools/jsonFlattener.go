package tools

import (
	"encoding/json"
	"github.com/deckarep/golang-set"
	"strconv"
)

type PrometheusMetrics struct {
	Labels map[string]interface{} `json:"labels"`
	Name string `json:"name"`
	Rtime string `json:"timestamp"`
	Value string `json:"value"`
}




func Flattener(jsonStr []byte) (metricName string, jsonOutput string, jsonKeys mapset.Set) {

	var metricsJson PrometheusMetrics
	json.Unmarshal(jsonStr, &metricsJson)

	//build an empty string for json
	var obj map[string]interface{}
	err := json.Unmarshal([]byte("{}"), &obj)
	if err != nil {
		panic(err.Error())
	}
	obj["name"] = metricsJson.Name
	obj["timestamp"] = metricsJson.Rtime
	obj["value"], err = strconv.ParseFloat(metricsJson.Value, 64)
	if err != nil {
		panic(err.Error())
	}

	//add labels to the json string
	for key,value := range metricsJson.Labels {
		//fmt.Println(key)
		//fmt.Println(value)
		obj[key] = value
	}


	output, err := json.Marshal(obj)

	//get all the keys
	//var buffer bytes.Buffer
	//i := 0
	//keys := []string{}
	//for key := range obj {
	//	//if buffer.Len() == 0 {
	//	//	buffer.WriteString(key)
	//	//} else {
	//	//	buffer.WriteString(";")
	//	//	buffer.WriteString(key)
	//	//}
	//	keys = append(keys, key)
	//}

	keys := mapset.NewSet()
	for key := range obj {
		keys.Add(key)
	}

	return metricsJson.Name, string(output), keys
}