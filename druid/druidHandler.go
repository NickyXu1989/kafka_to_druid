package druid

import (
	"encoding/json"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"io/ioutil"
	"kafka-to-druid/druid/druidAllSupervisors"
	"kafka-to-druid/druid/druidCreateOrUpdateSupervisor"
	"net/http"
)

type DruidHandler struct {
	apiServer string
}

func NewDruidHandler (apiServer string) *DruidHandler{
	dh := &DruidHandler{
		apiServer: apiServer,
	}

	return dh
}


func (dh *DruidHandler) GetAllSupervisors() (result map[string]mapset.Set) {

	//result = map[string][]string{}
	result = map[string]mapset.Set{}
	url := "/druid/indexer/v1/supervisor?full"
	resp, err := http.Get(dh.apiServer + url)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	//fmt.Println(string(body))
	//fmt.Println([]byte(string(body)))
	var objs []druidAllSupervisors.DruidAllSupervisorsResponse

	err = json.Unmarshal([]byte(string(body)), &objs)

	for _, obj := range objs {
		supervisorName := obj.ID
		dimensions := obj.Spec.DataSchema.Parser.ParseSpec.DimensionsSpec.Dimensions
		//fmt.Println(dimensions)
		//result[supervisorName] = dimensions
		diSet := mapset.NewSet()
		for _, i := range dimensions {
			diSet.Add(i)
		}
		result[supervisorName] = diSet
	}

	return result
}


func (dh *DruidHandler) CreateOrUpdateSupervisor(supervisorName string, labels mapset.Set, bootstrapServers string) {
	tmpSet := mapset.NewSet()
	tmpSet.Add("value")
	tmpSet.Add("timestamp")
	dimensions := mapset.NewSet()
	dimensions = labels.Difference(tmpSet)

	var dimensionsArray []string

	for _, i := range dimensions.ToSlice() {
		//fmt.Println(i)
		dimensionsArray = append(dimensionsArray, i.(string))
	}

	//build the json


	timestampSpec := druidCreateOrUpdateSupervisor.TimestampSpec{
		Column: "timestamp",
		Format: "auto",
	}

	dimensionsSpec := druidCreateOrUpdateSupervisor.DimensionsSpec{
		Dimensions: dimensionsArray,
	}

	parseSpec := druidCreateOrUpdateSupervisor.ParseSpec{
		Format: "json",
		TimestampSpec: timestampSpec,
		DimensionsSpec: dimensionsSpec,
	}

	parser := druidCreateOrUpdateSupervisor.Parser{
		Type: "string",
		ParseSpec: parseSpec,
	}

	metricsSpec := druidCreateOrUpdateSupervisor.MetricsSpec{
		Type: "doublesum",
		Name: "value",
		FieldName: "value",
	}

	granularitySpec := druidCreateOrUpdateSupervisor.GranularitySpec{
		Type: "uniform",
		SegmentGranularity: "DAY",
		QueryGranularity: "NONE",
		Rollup: false,
	}


	dataSchema := druidCreateOrUpdateSupervisor.DataSchema{
		DataSource: supervisorName,
		Parser: parser,
		MetricsSpec: append([]druidCreateOrUpdateSupervisor.MetricsSpec, metricsSpec),
		GranularitySpec: granularitySpec,
	}

	tuningConfig := druidCreateOrUpdateSupervisor.TuningConfig{
		Type: "kafka",
		ReportParseExceptions: false,
	}

	consumerProperties := druidCreateOrUpdateSupervisor.ConsumerProperties{
		BootstrapServers: bootstrapServers,
	}

	ioConfig := druidCreateOrUpdateSupervisor.IoConfig{
		Topic: supervisorName,
		Replicas: 2,
		TaskDuration: "PT10M",
		CompletionTimeout: "PT20M",
		ConsumerProperties: consumerProperties,
	}


	payload := druidCreateOrUpdateSupervisor.DruidCreateOrUpdateSupervisorPayload{
		Type: "kafka",
		DataSchema: dataSchema,
		TuningConfig: tuningConfig,
		IoConfig: ioConfig,
	}





}