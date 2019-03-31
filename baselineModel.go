package pmml

type BaselineModel struct {
	ModelName     string         `xml:"modelName,attr"`
	FunctionName  MiningFunction `xml:"functionName,attr"`
	AlgorithmName string         `xml:"algorithmName,attr"`
	IsScorable    bool           `xml:"isScorable,attr"`

	MiningSchema         MiningSchema         `xml:"MiningSchema"`
	Output               Output               `xml:"Output"`
	ModelStats           ModelStats           `xml:"ModelStats"`
	ModelExplanation     ModelExplanation     `xml:"ModelExplanation"`
	Targets              Targets              `xml:"Targets"`
	LocalTransformations LocalTransformations `xml:"LocalTransformations"`
	TestDistributions    TestDistributions    `xml:"TestDistributions"`
	ModelVerification    ModelVerification    `xml:"ModelVerification"`
}

type ModelExplanation struct {
}

type Targets struct {
	Targets []Target `xml:"Target"`
}

type Target struct {
	TargetValues    []TargetValue `xml:"TargetValue"`
	Field           string        `xml:"field,attr"`
	OpType          OpType        `xml:"optype,attr"`
	CastInteger     CastInteger   `xml:"castInteger,attr"`
	Min             float64       `xml:"min,attr"`
	Max             float64       `xml:"max,attr"`
	RescaleConstant float64       `xml:"rescaleConstant,attr"`
	RescaleFactor   float64       `xml:"rescaleFactor,attr"`
}

type TargetValue struct {
	Partition        Partition `xml:"Partition"`
	Value            string    `xml:"value,attr"`
	DisplayValue     string    `xml:"displayValue,attr"`
	PriorProbability string    `xml:"priorProbability,attr"`
	DefaultValue     string    `xml:"defaultValue,attr"`
}

type Partition struct {
	PartitionFieldStats []PartitionFieldStats `xml:"PartitionFieldStats"`
	Name                string                `xml:"name,attr"`
	Size                int                   `xml:"size,attr"`
}

type PartitionFieldStats struct {
	Counts          Counts          `xml:"Counts"`
	NumericInfo     NumericInfo     `xml:"NumericInfo"`
	FrequenciesType FrequenciesType `xml:"FrequenciesType"`
	Field           string          `xml:"field,attr"`
	Weighted        bool            `xml:"weighted,attr"`
}

type Counts struct {
	TotalFreq   float64 `xml:"totalFreq,attr"`
	MissingFreq float64 `xml:"missingFreq,attr"`
	InvalidFreq float64 `xml:"invalidFreq,attr"`
	Cardinality uint    `xml:"cardinality,attr"`
}

type NumericInfo struct {
	Quantile           Quantile `xml:"Quantile"`
	Minimum            float64  `xml:"minimum,attr"`
	Maximum            float64  `xml:"maximum,attr"`
	Mean               float64  `xml:"mean,attr"`
	StandardDeviation  float64  `xml:"standardDeviation,attr"`
	Median             float64  `xml:"median,attr"`
	InterQuartileRange float64  `xml:"interQuartileRange,attr"`
}

type FrequenciesType struct {
}

type Quantile struct {
	QuantileLimit float64 `xml:"quantileLimit"`
	QuantileValue float64 `xml:"quantileValue"`
}

type TestDistributions struct {
	Baseline  Baseline  `xml:"Baseline"`
	Alternate Alternate `xml:"Alternate"`

	Field               string                `xml:"field,attr"`
	TestStatistic       BaselineTestStatistic `xml:"testStatistic,attr"`
	ResetValue          float64               `xml:"resetValue,attr"`
	WindowSize          int                   `xml:"windowSize,attr"`
	WeightField         string                `xml:"weightField,attr"`
	NormalizationScheme string                `xml:"normalizationScheme,attr"`
}

type Baseline interface {
}

type Alternate interface {
}
