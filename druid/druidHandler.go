package druid

import (
	"encoding/json"
	mapset "github.com/deckarep/golang-set"
	"io/ioutil"
	"kafka-to-druid/druid/druidAllSupervisors"
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
		supervisorName := obj.Id
		dimensions := obj.Spec.DataSchema.Parser.ParseSpec.DimensionsSpec.Dimensions
		//result[supervisorName] = dimensions
		diSet := mapset.NewSet()
		for i := range dimensions {
			diSet.Add(i)
		}
		result[supervisorName] = diSet
	}

	return result
}