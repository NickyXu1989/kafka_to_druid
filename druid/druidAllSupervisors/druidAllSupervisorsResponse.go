package druidAllSupervisors

//
//type DruidAllSupervisorsResponse struct {
//	Id string `json:"id"`
//	Spec spec `json:"spec"`
//}
//
//type spec struct {
//	DataSchema dataSchema `json:"dataSchema"`
//}
//
//type dataSchema struct {
//	DataSource string `json:"dataSource"`
//	Parser parser `json:"parser"`
//}
//
//type parser struct {
//	ParseSpec parseSpec `json:"parseSpec"`
//}
//type parseSpec struct {
//	Format string `json:"format"`
//	DimensionsSpec dimensionsSpec `json:"dimensionsSpec"`
//}
//
//type dimensionsSpec struct {
//	Dimensions []string `json:"dimensions"`
//}


type DruidAllSupervisorsResponse struct {
	ID   string `json:"id"`
	Spec struct {
		DataSchema struct {
			DataSource string `json:"dataSource"`
			Parser     struct {
				Type      string `json:"type"`
				ParseSpec struct {
					Format        string `json:"format"`
					TimestampSpec struct {
						Column string `json:"column"`
						Format string `json:"format"`
					} `json:"timestampSpec"`
					DimensionsSpec struct {
						Dimensions []string `json:"dimensions"`
					} `json:"dimensionsSpec"`
				} `json:"parseSpec"`
			} `json:"parser"`
			MetricsSpec []struct {
				Type       string      `json:"type"`
				Name       string      `json:"name"`
				FieldName  string      `json:"fieldName"`
				Expression interface{} `json:"expression"`
			} `json:"metricsSpec"`
			GranularitySpec struct {
				Type               string `json:"type"`
				SegmentGranularity string `json:"segmentGranularity"`
				QueryGranularity   struct {
					Type string `json:"type"`
				} `json:"queryGranularity"`
				Rollup    bool        `json:"rollup"`
				Intervals interface{} `json:"intervals"`
			} `json:"granularitySpec"`
			TransformSpec struct {
				Filter     interface{}   `json:"filter"`
				Transforms []interface{} `json:"transforms"`
			} `json:"transformSpec"`
		} `json:"dataSchema"`
		TuningConfig struct {
			Type                      string      `json:"type"`
			MaxRowsInMemory           int         `json:"maxRowsInMemory"`
			MaxBytesInMemory          int         `json:"maxBytesInMemory"`
			MaxRowsPerSegment         int         `json:"maxRowsPerSegment"`
			MaxTotalRows              interface{} `json:"maxTotalRows"`
			IntermediatePersistPeriod string      `json:"intermediatePersistPeriod"`
			BasePersistDirectory      string      `json:"basePersistDirectory"`
			MaxPendingPersists        int         `json:"maxPendingPersists"`
			IndexSpec                 struct {
				Bitmap struct {
					Type string `json:"type"`
				} `json:"bitmap"`
				DimensionCompression string `json:"dimensionCompression"`
				MetricCompression    string `json:"metricCompression"`
				LongEncoding         string `json:"longEncoding"`
			} `json:"indexSpec"`
			BuildV9Directly                     bool        `json:"buildV9Directly"`
			ReportParseExceptions               bool        `json:"reportParseExceptions"`
			HandoffConditionTimeout             int         `json:"handoffConditionTimeout"`
			ResetOffsetAutomatically            bool        `json:"resetOffsetAutomatically"`
			SegmentWriteOutMediumFactory        interface{} `json:"segmentWriteOutMediumFactory"`
			WorkerThreads                       interface{} `json:"workerThreads"`
			ChatThreads                         interface{} `json:"chatThreads"`
			ChatRetries                         int         `json:"chatRetries"`
			HTTPTimeout                         string      `json:"httpTimeout"`
			ShutdownTimeout                     string      `json:"shutdownTimeout"`
			OffsetFetchPeriod                   string      `json:"offsetFetchPeriod"`
			IntermediateHandoffPeriod           string      `json:"intermediateHandoffPeriod"`
			LogParseExceptions                  bool        `json:"logParseExceptions"`
			MaxParseExceptions                  int64       `json:"maxParseExceptions"`
			MaxSavedParseExceptions             int         `json:"maxSavedParseExceptions"`
			SkipSequenceNumberAvailabilityCheck bool        `json:"skipSequenceNumberAvailabilityCheck"`
		} `json:"tuningConfig"`
		IoConfig struct {
			Topic              string `json:"topic"`
			Replicas           int    `json:"replicas"`
			TaskCount          int    `json:"taskCount"`
			TaskDuration       string `json:"taskDuration"`
			ConsumerProperties struct {
				BootstrapServers string `json:"bootstrap.servers"`
			} `json:"consumerProperties"`
			PollTimeout                 int         `json:"pollTimeout"`
			StartDelay                  string      `json:"startDelay"`
			Period                      string      `json:"period"`
			UseEarliestOffset           bool        `json:"useEarliestOffset"`
			CompletionTimeout           string      `json:"completionTimeout"`
			LateMessageRejectionPeriod  interface{} `json:"lateMessageRejectionPeriod"`
			EarlyMessageRejectionPeriod interface{} `json:"earlyMessageRejectionPeriod"`
			Stream                      string      `json:"stream"`
			UseEarliestSequenceNumber   bool        `json:"useEarliestSequenceNumber"`
		} `json:"ioConfig"`
		Context   interface{} `json:"context"`
		Suspended bool        `json:"suspended"`
	} `json:"spec"`
}