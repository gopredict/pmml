package pmml

type AssociationModel struct {
	ModelName             string         `xml:"modelName,attr"`
	FunctionName          MiningFunction `xml:"functionName,attr"`
	AlgorithmName         string         `xml:"algorithmName,attr"`
	NumberOfTransactions  int            `xml:"numberOfTransactions,attr"`
	MaxNumberOfItemsPerTA int            `xml:"maxNumberOfItemsPerTA,attr"`
	AvgNumberOfItemsPerTA float64        `xml:"avgNumberOfItemsPerTA,attr"`
	MinimumSupport        float64        `xml:"minimumSupport,attr"`
	MinimumConfidence     float64        `xml:"minimumConfidence,attr"`
	LengthLimit           int            `xml:"lengthLimit,attr"`
	NumberOfItems         int            `xml:"numberOfItems,attr"`
	NumberOfItemsets      int            `xml:"numberOfItemsets,attr"`
	NumberOfRules         int            `xml:"numberOfRules,attr"`
	IsScorable            bool           `xml:"isScorable,attr"`

	MiningSchema         MiningSchema         `xml:"MiningSchema"`
	Output               Output               `xml:"Output"`
	ModelStats           ModelStats           `xml:"ModelStats"`
	LocalTransformations LocalTransformations `xml:"LocalTransformations"`
	Items                []Item               `xml:"Item"`
	ItemSets             []ItemSet            `xml:"Itemset"`
	AssociationRules     []AssociationRule    `xml:"AssociationRule"`
	ModelVerification    ModelVerification    `xml:"ModelVerification"`
}

type Item struct {
	ID          string  `xml:"id,attr"`
	Value       string  `xml:"value,attr"`
	Field       string  `xml:"field,attr"`
	Category    string  `xml:"category,attr"`
	MappedValue string  `xml:"mappedValue,attr"`
	Weight      float64 `xml:"weight,attr"`
}

type ItemSet struct {
	ID            string    `xml:"id,attr"`
	Support       float64   `xml:"support,attr"`
	NumberOfItems uint      `xml:"numberOfItems,attr"`
	ItemRefs      []ItemRef `xml:"ItemRef"`
}

type ItemRef struct {
	ItemRef string `xml:"itemRef,attr"`
}

type AssociationRule struct {
	Antecedent string  `xml:"antecedent,attr"`
	Consequent string  `xml:"consequent,attr"`
	Support    float64 `xml:"support,attr"`
	Confidence float64 `xml:"confidence,attr"`
	Lift       float64 `xml:"lift,attr"`
	Leverage   float64 `xml:"leverage,attr"`
	Affinity   float64 `xml:"affinity,attr"`
	ID         string  `xml:"id,attr"`
}

type ModelVerification struct {
	RecordCount        int                `xml:"recordCount,attr"`
	FieldCount         int                `xml:"fieldCount,attr"`
	VerificationFields VerificationFields `xml:"VerificationFields"`
	InlineTable        InlineTable        `xml:"InlineTable"`
}

type VerificationFields struct {
	VerificationFields []VerificationField `xml:"VerificationField"`
}

type VerificationField struct {
	Field         string  `xml:"field,attr"`
	Column        string  `xml:"column,attr"`
	Precision     float64 `xml:"precision,attr"`
	ZeroThreshold float64 `xml:"zeroThreshold,attr"`
}

type InlineTable struct {
}
