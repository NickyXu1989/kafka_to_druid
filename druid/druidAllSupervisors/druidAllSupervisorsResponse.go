package druidAllSupervisors


type DruidAllSupervisorsResponse struct {
	Id string `json:"id"`
	Spec spec `json:"spec"`
}

type spec struct {
	DataSchema dataSchema `json:"dataSchema"`
}

type dataSchema struct {
	DataSource string `json:"dataSource"`
	Parser parser `json:"parser"`
}

type parser struct {
	ParseSpec parseSpec `json:"parseSpec"`
}
type parseSpec struct {
	Format string `json:"format"`
	DimensionsSpec dimensionsSpec `json:"dimensionsSpec"`
}

type dimensionsSpec struct {
	Dimensions []string `json:"dimensions"`
}