package druidCreateOrUpdateSupervisor

type druidCreateOrUpdateSupervisorPayload struct {
	Type string `json:"type"`
	DataSchema interface{} `json:"dataSchema"`
	TuningConfig interface{} `json:"tuningConfig"`
	IoConfig interface{} `json:"ioConfig"`
}

type dataSchema struct {
	DataSource string `json:"dataSource"`
	Parser interface{} `json:"parser"`
	MetricsSpec interface{} `json:"metricsSpec"`
	GranularitySpec interface{} `json:"granularitySpec"`
}