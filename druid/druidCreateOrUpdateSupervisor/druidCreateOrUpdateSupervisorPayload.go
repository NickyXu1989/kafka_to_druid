package druidCreateOrUpdateSupervisor
//
//type druidCreateOrUpdateSupervisorPayload struct {
//	Type string `json:"type"`
//	DataSchema interface{} `json:"dataSchema"`
//	TuningConfig interface{} `json:"tuningConfig"`
//	IoConfig interface{} `json:"ioConfig"`
//}
//
//type dataSchema struct {
//	DataSource string `json:"dataSource"`
//	Parser interface{} `json:"parser"`
//	MetricsSpec interface{} `json:"metricsSpec"`
//	GranularitySpec interface{} `json:"granularitySpec"`
//}
//
//type parser struct {
//	Type string `json:"type"`
//	ParseSpec interface{} `json:"parseSpec"`
//}
//
//type parseSpec struct {
//	Format string `json:"format"`
//	TimestampSpec interface{} `json:"timestampSpec"`
//}
//
//type timestampSpec struct {
//	Column string `json:"column"`
//	Format string `json:"format"`
//}



type DruidCreateOrUpdateSupervisorPayload struct {
	Type         string       `json:"type"`
	DataSchema   DataSchema   `json:"dataSchema"`
	TuningConfig TuningConfig `json:"tuningConfig"`
	IoConfig     IoConfig     `json:"ioConfig"`
}
type TimestampSpec struct {
	Column string `json:"column"`
	Format string `json:"format"`
}
type DimensionsSpec struct {
	Dimensions []string `json:"dimensions"`
}
type ParseSpec struct {
	Format         string         `json:"format"`
	TimestampSpec  TimestampSpec  `json:"timestampSpec"`
	DimensionsSpec DimensionsSpec `json:"dimensionsSpec"`
}
type Parser struct {
	Type      string    `json:"type"`
	ParseSpec ParseSpec `json:"parseSpec"`
}
type MetricsSpec struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	FieldName string `json:"fieldName"`
}
type GranularitySpec struct {
	Type               string `json:"type"`
	SegmentGranularity string `json:"segmentGranularity"`
	QueryGranularity   string `json:"queryGranularity"`
	Rollup             bool   `json:"rollup"`
}
type DataSchema struct {
	DataSource      string          `json:"dataSource"`
	Parser          Parser          `json:"parser"`
	MetricsSpec     []MetricsSpec   `json:"metricsSpec"`
	GranularitySpec GranularitySpec `json:"granularitySpec"`
}
type TuningConfig struct {
	Type                  string `json:"type"`
	ReportParseExceptions bool   `json:"reportParseExceptions"`
}
type ConsumerProperties struct {
	BootstrapServers string `json:"bootstrap.servers"`
}
type IoConfig struct {
	Topic              string             `json:"topic"`
	Replicas           int                `json:"replicas"`
	TaskDuration       string             `json:"taskDuration"`
	CompletionTimeout  string             `json:"completionTimeout"`
	ConsumerProperties ConsumerProperties `json:"consumerProperties"`
}