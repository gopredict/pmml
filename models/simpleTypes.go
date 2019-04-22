package models

/*
  <xs:simpleType name="ACTIVATION-FUNCTION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="Elliott"/>
      <xs:enumeration value="Gauss"/>
      <xs:enumeration value="arctan"/>
      <xs:enumeration value="cosine"/>
      <xs:enumeration value="exponential"/>
      <xs:enumeration value="identity"/>
      <xs:enumeration value="logistic"/>
      <xs:enumeration value="radialBasis"/>
      <xs:enumeration value="reciprocal"/>
      <xs:enumeration value="rectifier"/>
      <xs:enumeration value="sine"/>
      <xs:enumeration value="square"/>
      <xs:enumeration value="tanh"/>
      <xs:enumeration value="threshold"/>
    </xs:restriction>
  </xs:simpleType>
*/
type ActivationFunction string

const (
	ActivationFunctionElliott     = ActivationFunction("Elliott")
	ActivationFunctionGauss       = ActivationFunction("Gauss")
	ActivationFunctionArctan      = ActivationFunction("arctan")
	ActivationFunctionCosine      = ActivationFunction("cosine")
	ActivationFunctionExponential = ActivationFunction("exponential")
	ActivationFunctionIdentity    = ActivationFunction("identity")
	ActivationFunctionLogistic    = ActivationFunction("logistic")
	ActivationFunctionRadialBasis = ActivationFunction("radialBasis")
	ActivationFunctionReciprocal  = ActivationFunction("reciprocal")
	ActivationFunctionRectifier   = ActivationFunction("rectifier")
	ActivationFunctionSine        = ActivationFunction("sine")
	ActivationFunctionSquare      = ActivationFunction("square")
	ActivationFunctionTanh        = ActivationFunction("tanh")
	ActivationFunctionThreshold   = ActivationFunction("threshold")
)

/*
  <xs:simpleType name="BASELINE-TEST-STATISTIC">
    <xs:restriction base="xs:string">
      <xs:enumeration value="CUSUM"/>
      <xs:enumeration value="chiSquareDistribution"/>
      <xs:enumeration value="chiSquareIndependence"/>
      <xs:enumeration value="scalarProduct"/>
      <xs:enumeration value="zValue"/>
    </xs:restriction>
  </xs:simpleType>
*/
type BaselineTestStatistic string

const (
	BaselineTestStatisticCUSUM                 = BaselineTestStatistic("CUSUM")
	BaselineTestStatisticChiSquareDistribution = BaselineTestStatistic("chiSquareDistribution")
	BaselineTestStatisticChiSquareIndependence = BaselineTestStatistic("chiSquareIndependence")
	BaselineTestStatisticScalarProduct         = BaselineTestStatistic("scalarProduct")
	BaselineTestStatisticZValue                = BaselineTestStatistic("zValue")
)

/*
  <xs:simpleType name="CAT-SCORING-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="majorityVote"/>
      <xs:enumeration value="weightedMajorityVote"/>
    </xs:restriction>
  </xs:simpleType>
*/
type CatScoringMethod string

const (
	CatScoringMethodMajorityVote         = CatScoringMethod("majorityVote")
	CatScoringMethodWeightedMajorityVote = CatScoringMethod("weightedMajorityVote")
)

/*
  <xs:simpleType name="COMPARE-FUNCTION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="absDiff"/>
      <xs:enumeration value="delta"/>
      <xs:enumeration value="equal"/>
      <xs:enumeration value="gaussSim"/>
      <xs:enumeration value="table"/>
    </xs:restriction>
  </xs:simpleType>
*/
type CompareFunction string

const (
	CompareFunctionAbsDiff  = CompareFunction("absDiff")
	CompareFunctionDelta    = CompareFunction("delta")
	CompareFunctionEqual    = CompareFunction("equal")
	CompareFunctionGaussSim = CompareFunction("gaussSim")
	CompareFunctionTable    = CompareFunction("table")
)

/*
  <xs:simpleType name="CONT-SCORING-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="average"/>
      <xs:enumeration value="median"/>
      <xs:enumeration value="weightedAverage"/>
    </xs:restriction>
  </xs:simpleType>
*/
type ContScoringMethod string

const (
	ContScoringMethodAverage         = ContScoringMethod("average")
	ContScoringMethodMedian          = ContScoringMethod("median")
	ContScoringMethodWeightedAverage = ContScoringMethod("weightedAverage")
)

/*
  <xs:simpleType name="CUMULATIVE-LINK-FUNCTION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="cauchit"/>
      <xs:enumeration value="cloglog"/>
      <xs:enumeration value="logit"/>
      <xs:enumeration value="loglog"/>
      <xs:enumeration value="probit"/>
    </xs:restriction>
  </xs:simpleType>
*/
type CumulativeLinkFunction string

const (
	CumulativeLinkFunctionCauchit = CumulativeLinkFunction("cauchit")
	CumulativeLinkFunctionCloglog = CumulativeLinkFunction("cloglog")
	CumulativeLinkFunctionLogit   = CumulativeLinkFunction("logit")
	CumulativeLinkFunctionLoglog  = CumulativeLinkFunction("loglog")
	CumulativeLinkFunctionProbit  = CumulativeLinkFunction("probit")
)

/*
  <xs:simpleType name="DATATYPE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="boolean"/>
      <xs:enumeration value="date"/>
      <xs:enumeration value="dateDaysSince[0]"/>
      <xs:enumeration value="dateDaysSince[1960]"/>
      <xs:enumeration value="dateDaysSince[1970]"/>
      <xs:enumeration value="dateDaysSince[1980]"/>
      <xs:enumeration value="dateTime"/>
      <xs:enumeration value="dateTimeSecondsSince[0]"/>
      <xs:enumeration value="dateTimeSecondsSince[1960]"/>
      <xs:enumeration value="dateTimeSecondsSince[1970]"/>
      <xs:enumeration value="dateTimeSecondsSince[1980]"/>
      <xs:enumeration value="double"/>
      <xs:enumeration value="float"/>
      <xs:enumeration value="integer"/>
      <xs:enumeration value="string"/>
      <xs:enumeration value="time"/>
      <xs:enumeration value="timeSeconds"/>
    </xs:restriction>
  </xs:simpleType>
*/
type DataType string

const (
	DataTypeboolean                  = DataType("boolean")
	DataTypedate                     = DataType("date")
	DataTypedateDaysSince0           = DataType("dateDaysSince[0]")
	DataTypedateDaysSince1960        = DataType("dateDaysSince[1960]")
	DataTypedateDaysSince1970        = DataType("dateDaysSince[1970]")
	DataTypedateDaysSince1980        = DataType("dateDaysSince[1980]")
	DataTypedateTime                 = DataType("dateTime")
	DataTypedateTimeSecondsSince0    = DataType("dateTimeSecondsSince[0]")
	DataTypedateTimeSecondsSince1960 = DataType("dateTimeSecondsSince[1960]")
	DataTypedateTimeSecondsSince1970 = DataType("dateTimeSecondsSince[1970]")
	DataTypedateTimeSecondsSince1980 = DataType("dateTimeSecondsSince[1980]")
	DataTypedouble                   = DataType("double")
	DataTypefloat                    = DataType("float")
	DataTypeinteger                  = DataType("integer")
	DataTypestring                   = DataType("string")
	DataTypetime                     = DataType("time")
	DataTypetimeSeconds              = DataType("timeSeconds")
)

/*
  <xs:simpleType name="DELIMITER">
    <xs:restriction base="xs:string">
      <xs:enumeration value="acrossTimeWindows"/>
      <xs:enumeration value="sameTimeWindow"/>
    </xs:restriction>
  </xs:simpleType>
*/
type DelimiterType string

const (
	DelimiterTypeAcrossTimeWindows = DelimiterType("acrossTimeWindows")
	DelimiterTypeSameTimeWindow    = DelimiterType("sameTimeWindow")
)

/*
  <xs:simpleType name="ELEMENT-ID">
    <xs:restriction base="xs:string"/>
  </xs:simpleType>
*/
type ElementID string

/*
  <xs:simpleType name="FIELD-NAME">
    <xs:restriction base="xs:string"/>
  </xs:simpleType>
*/
type FieldName string

/*
  <xs:simpleType name="FIELD-USAGE-TYPE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="active"/>
      <xs:enumeration value="analysisWeight"/>
      <xs:enumeration value="frequencyWeight"/>
      <xs:enumeration value="group"/>
      <xs:enumeration value="order"/>
      <xs:enumeration value="predicted"/>
      <xs:enumeration value="supplementary"/>
      <xs:enumeration value="target"/>
    </xs:restriction>
  </xs:simpleType>
*/
type FieldUsageType string

const (
	FieldUsageTypeActive          = FieldUsageType("active")
	FieldUsageTypeAnalysisWeight  = FieldUsageType("analysisWeight")
	FieldUsageTypeFrequencyWeight = FieldUsageType("frequencyWeight")
	FieldUsageTypeGroup           = FieldUsageType("group")
	FieldUsageTypeOrder           = FieldUsageType("order")
	FieldUsageTypePredicted       = FieldUsageType("predicted")
	FieldUsageTypeSupplementary   = FieldUsageType("supplementary")
	FieldUsageTypeTarget          = FieldUsageType("target")
)

/*
  <xs:simpleType name="GAP">
    <xs:restriction base="xs:string">
      <xs:enumeration value="false"/>
      <xs:enumeration value="true"/>
      <xs:enumeration value="unknown"/>
    </xs:restriction>
  </xs:simpleType>
*/
type Gap string

const (
	GapFalse   = Gap("false")
	GapTrue    = Gap("true")
	GapUnknown = Gap("unknown")
)

/*
  <xs:simpleType name="INT-NUMBER">
    <xs:restriction base="xs:integer"/>
  </xs:simpleType>
*/
type IntegerNumber int

/*
  <xs:simpleType name="INTERPOLATION-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="cubicSpline"/>
      <xs:enumeration value="exponentialSpline"/>
      <xs:enumeration value="linear"/>
      <xs:enumeration value="none"/>
    </xs:restriction>
  </xs:simpleType>
*/
type InterpolationMethod string

const (
	InterpolationMethodCubicSpline       = InterpolationMethod("cubicSpline")
	InterpolationMethodExponentialSpline = InterpolationMethod("exponentialSpline")
	InterpolationMethodLinear            = InterpolationMethod("linear")
	InterpolationMethodNone              = InterpolationMethod("none")
)

/*
  <xs:simpleType name="INVALID-VALUE-TREATMENT-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="asIs"/>
      <xs:enumeration value="asMissing"/>
      <xs:enumeration value="returnInvalid"/>
    </xs:restriction>
  </xs:simpleType>
*/
type InvalidValueTreatmentMethod string

const (
	InvalidValueTreatmentMethodAsIs          = InvalidValueTreatmentMethod("asIs")
	InvalidValueTreatmentMethodAsMissing     = InvalidValueTreatmentMethod("asMissing")
	InvalidValueTreatmentMethodReturnInvalid = InvalidValueTreatmentMethod("returnInvalid")
)

/*
  <xs:simpleType name="LINK-FUNCTION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="cloglog"/>
      <xs:enumeration value="identity"/>
      <xs:enumeration value="log"/>
      <xs:enumeration value="logc"/>
      <xs:enumeration value="logit"/>
      <xs:enumeration value="loglog"/>
      <xs:enumeration value="negbin"/>
      <xs:enumeration value="oddspower"/>
      <xs:enumeration value="power"/>
      <xs:enumeration value="probit"/>
    </xs:restriction>
  </xs:simpleType>
*/
type LinkFunction string

const (
	LinkFunctionCloglog   = LinkFunction("cloglog")
	LinkFunctionIdentity  = LinkFunction("identity")
	LinkFunctionLog       = LinkFunction("log")
	LinkFunctionLogc      = LinkFunction("logc")
	LinkFunctionLogit     = LinkFunction("logit")
	LinkFunctionLoglog    = LinkFunction("loglog")
	LinkFunctionNegbin    = LinkFunction("negbin")
	LinkFunctionOddspower = LinkFunction("oddspower")
	LinkFunctionPower     = LinkFunction("power")
	LinkFunctionProbit    = LinkFunction("probit")
)

/*
  <xs:simpleType name="MINING-FUNCTION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="associationRules"/>
      <xs:enumeration value="classification"/>
      <xs:enumeration value="clustering"/>
      <xs:enumeration value="mixed"/>
      <xs:enumeration value="regression"/>
      <xs:enumeration value="sequences"/>
      <xs:enumeration value="timeSeries"/>
    </xs:restriction>
  </xs:simpleType>
*/
type MiningFunction string

const (
	MiningFunctionAssociationRules = MiningFunction("associationRules")
	MiningFunctionClassification   = MiningFunction("classification")
	MiningFunctionClustering       = MiningFunction("clustering")
	MiningFunctionMixed            = MiningFunction("mixed")
	MiningFunctionRegression       = MiningFunction("regression")
	MiningFunctionSequences        = MiningFunction("sequences")
	MiningFunctionTimeSeries       = MiningFunction("timeSeries")
)

/*
  <xs:simpleType name="MISSING-VALUE-STRATEGY">
    <xs:restriction base="xs:string">
      <xs:enumeration value="aggregateNodes"/>
      <xs:enumeration value="defaultChild"/>
      <xs:enumeration value="lastPrediction"/>
      <xs:enumeration value="none"/>
      <xs:enumeration value="nullPrediction"/>
      <xs:enumeration value="weightedConfidence"/>
    </xs:restriction>
  </xs:simpleType>
*/
type MissingValueStrategy string

const (
	MissingValueStrategyAggregateNodes     = MissingValueStrategy("aggregateNodes")
	MissingValueStrategyDefaultChild       = MissingValueStrategy("defaultChild")
	MissingValueStrategyLastPrediction     = MissingValueStrategy("lastPrediction")
	MissingValueStrategyNone               = MissingValueStrategy("none")
	MissingValueStrategyNullPrediction     = MissingValueStrategy("nullPrediction")
	MissingValueStrategyWeightedConfidence = MissingValueStrategy("weightedConfidence")
)

/*
  <xs:simpleType name="MISSING-VALUE-TREATMENT-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="asIs"/>
      <xs:enumeration value="asMean"/>
      <xs:enumeration value="asMedian"/>
      <xs:enumeration value="asMode"/>
      <xs:enumeration value="asValue"/>
    </xs:restriction>
  </xs:simpleType>
*/
type MissingValueTreatmentMethod string

const (
	MissingValueTreatmentMethodAsIs     = MissingValueTreatmentMethod("asIs")
	MissingValueTreatmentMethodAsMean   = MissingValueTreatmentMethod("asMean")
	MissingValueTreatmentMethodAsMedian = MissingValueTreatmentMethod("asMedian")
	MissingValueTreatmentMethodAsMode   = MissingValueTreatmentMethod("asMode")
	MissingValueTreatmentMethodAsValue  = MissingValueTreatmentMethod("asValue")
)

/*
  <xs:simpleType name="MULTIPLE-MODEL-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="average"/>
      <xs:enumeration value="majorityVote"/>
      <xs:enumeration value="max"/>
      <xs:enumeration value="median"/>
      <xs:enumeration value="modelChain"/>
      <xs:enumeration value="selectAll"/>
      <xs:enumeration value="selectFirst"/>
      <xs:enumeration value="sum"/>
      <xs:enumeration value="weightedAverage"/>
      <xs:enumeration value="weightedMajorityVote"/>
    </xs:restriction>
  </xs:simpleType>
*/
type MultipleModelMethod string

const (
	MultipleModelMethodAverage              = MultipleModelMethod("average")
	MultipleModelMethodMajorityVote         = MultipleModelMethod("majorityVote")
	MultipleModelMethodMax                  = MultipleModelMethod("max")
	MultipleModelMethodMedian               = MultipleModelMethod("median")
	MultipleModelMethodModelChain           = MultipleModelMethod("modelChain")
	MultipleModelMethodSelectAll            = MultipleModelMethod("selectAll")
	MultipleModelMethodSelectFirst          = MultipleModelMethod("selectFirst")
	MultipleModelMethodSum                  = MultipleModelMethod("sum")
	MultipleModelMethodWeightedAverage      = MultipleModelMethod("weightedAverage")
	MultipleModelMethodWeightedMajorityVote = MultipleModelMethod("weightedMajorityVote")
)

/*
  <xs:simpleType name="NN-NEURON-ID">
    <xs:restriction base="xs:string"/>
  </xs:simpleType>
*/
type NNNeuronID string

/*
  <xs:simpleType name="NN-NEURON-IDREF">
    <xs:restriction base="xs:string"/>
  </xs:simpleType>
*/
type NNNeuronIDRef string

/*
  <xs:simpleType name="NN-NORMALIZATION-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="none"/>
      <xs:enumeration value="simplemax"/>
      <xs:enumeration value="softmax"/>
    </xs:restriction>
  </xs:simpleType>
*/
type NNNormalizationMethod string

const (
	NNNormalizationMethodNone      = NNNormalizationMethod("none")
	NNNormalizationMethodSimplemax = NNNormalizationMethod("simplemax")
	NNNormalizationMethodSoftmax   = NNNormalizationMethod("softmax")
)

/*
  <xs:simpleType name="NO-TRUE-CHILD-STRATEGY">
    <xs:restriction base="xs:string">
      <xs:enumeration value="returnLastPrediction"/>
      <xs:enumeration value="returnNullPrediction"/>
    </xs:restriction>
  </xs:simpleType>
*/
type NoTrueChildStrategy string

const (
	NoTrueChildStrategyReturnLastPrediction = NoTrueChildStrategy("returnLastPrediction")
	NoTrueChildStrategyReturnNullPrediction = NoTrueChildStrategy("returnNullPrediction")
)

/*
  <xs:simpleType name="NUMBER">
    <xs:restriction base="xs:double"/>
  </xs:simpleType>
*/
type Number float64

/*
  <xs:simpleType name="OPTYPE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="categorical"/>
      <xs:enumeration value="continuous"/>
      <xs:enumeration value="ordinal"/>
    </xs:restriction>
  </xs:simpleType>
*/
type OpType string

const (
	OpTypeCategorical = OpType("categorical")
	OpTypeContinuous  = OpType("continuous")
	OpTypeOrdinal     = OpType("ordinal")
)

/*
  <xs:simpleType name="OUTLIER-TREATMENT-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="asExtremeValues"/>
      <xs:enumeration value="asIs"/>
      <xs:enumeration value="asMissingValues"/>
    </xs:restriction>
  </xs:simpleType>
*/
type OutlierTreatmentMethod string

const (
	OutlierTreatmentMethodAsExtremeValues = OutlierTreatmentMethod("asExtremeValues")
	OutlierTreatmentMethodAsIs            = OutlierTreatmentMethod("asIs")
	OutlierTreatmentMethodAsMissingValues = OutlierTreatmentMethod("asMissingValues")
)

/*
  <xs:simpleType name="PERCENTAGE-NUMBER">
    <xs:restriction base="xs:double"/>
  </xs:simpleType>
*/
type PercentageNumber float64

/*
  <xs:simpleType name="PROB-NUMBER">
    <xs:restriction base="xs:double"/>
  </xs:simpleType>
*/
type ProbNumber float64

/*
  <xs:simpleType name="REAL-NUMBER">
    <xs:restriction base="xs:double"/>
  </xs:simpleType>
*/
type RealNumber float64

/*
  <xs:simpleType name="REGRESSIONNORMALIZATIONMETHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="cauchit"/>
      <xs:enumeration value="cloglog"/>
      <xs:enumeration value="exp"/>
      <xs:enumeration value="logit"/>
      <xs:enumeration value="loglog"/>
      <xs:enumeration value="none"/>
      <xs:enumeration value="probit"/>
      <xs:enumeration value="simplemax"/>
      <xs:enumeration value="softmax"/>
    </xs:restriction>
  </xs:simpleType>
*/
type RegressionNormalizationMethod string

const (
	RegressionNormalizationMethodCauchit   = RegressionNormalizationMethod("cauchit")
	RegressionNormalizationMethodCloglog   = RegressionNormalizationMethod("cloglog")
	RegressionNormalizationMethodExp       = RegressionNormalizationMethod("exp")
	RegressionNormalizationMethodLogit     = RegressionNormalizationMethod("logit")
	RegressionNormalizationMethodLoglog    = RegressionNormalizationMethod("loglog")
	RegressionNormalizationMethodNone      = RegressionNormalizationMethod("none")
	RegressionNormalizationMethodProbit    = RegressionNormalizationMethod("probit")
	RegressionNormalizationMethodSimplemax = RegressionNormalizationMethod("simplemax")
	RegressionNormalizationMethodSoftmax   = RegressionNormalizationMethod("softmax")
)

/*
  <xs:simpleType name="RESULT-FEATURE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="affinity"/>
      <xs:enumeration value="antecedent"/>
      <xs:enumeration value="clusterAffinity"/>
      <xs:enumeration value="clusterId"/>
      <xs:enumeration value="confidence"/>
      <xs:enumeration value="consequent"/>
      <xs:enumeration value="decision"/>
      <xs:enumeration value="entityAffinity"/>
      <xs:enumeration value="entityId"/>
      <xs:enumeration value="leverage"/>
      <xs:enumeration value="lift"/>
      <xs:enumeration value="predictedDisplayValue"/>
      <xs:enumeration value="predictedValue"/>
      <xs:enumeration value="probability"/>
      <xs:enumeration value="reasonCode"/>
      <xs:enumeration value="residual"/>
      <xs:enumeration value="rule"/>
      <xs:enumeration value="ruleId"/>
      <xs:enumeration value="ruleValue"/>
      <xs:enumeration value="standardError"/>
      <xs:enumeration value="support"/>
      <xs:enumeration value="transformedValue"/>
      <xs:enumeration value="warning"/>
    </xs:restriction>
  </xs:simpleType>
*/
type ResultFeature string

const (
	ResultFeatureAffinity              = ResultFeature("affinity")
	ResultFeatureAntecedent            = ResultFeature("antecedent")
	ResultFeatureClusterAffinity       = ResultFeature("clusterAffinity")
	ResultFeatureClusterID             = ResultFeature("clusterId")
	ResultFeatureConfidence            = ResultFeature("confidence")
	ResultFeatureConsequent            = ResultFeature("consequent")
	ResultFeatureDecision              = ResultFeature("decision")
	ResultFeatureEntityAffinity        = ResultFeature("entityAffinity")
	ResultFeatureEntityID              = ResultFeature("entityId")
	ResultFeatureLeverage              = ResultFeature("leverage")
	ResultFeatureLift                  = ResultFeature("lift")
	ResultFeaturePredictedDisplayValue = ResultFeature("predictedDisplayValue")
	ResultFeaturePredictedValue        = ResultFeature("predictedValue")
	ResultFeatureProbability           = ResultFeature("probability")
	ResultFeatureReasonCode            = ResultFeature("reasonCode")
	ResultFeatureResidual              = ResultFeature("residual")
	ResultFeatureRule                  = ResultFeature("rule")
	ResultFeatureRuleID                = ResultFeature("ruleId")
	ResultFeatureRuleValue             = ResultFeature("ruleValue")
	ResultFeatureStandardError         = ResultFeature("standardError")
	ResultFeatureSupport               = ResultFeature("support")
	ResultFeatureTransformedValue      = ResultFeature("transformedValue")
	ResultFeatureWarning               = ResultFeature("warning")
)

/*
  <xs:simpleType name="RULE-FEATURE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="affinity"/>
      <xs:enumeration value="antecedent"/>
      <xs:enumeration value="confidence"/>
      <xs:enumeration value="consequent"/>
      <xs:enumeration value="leverage"/>
      <xs:enumeration value="lift"/>
      <xs:enumeration value="rule"/>
      <xs:enumeration value="ruleId"/>
      <xs:enumeration value="support"/>
    </xs:restriction>
  </xs:simpleType>
*/
type RuleFeature string

const (
	RuleFeatureSffinity   = RuleFeature("affinity")
	RuleFeatureAntecedent = RuleFeature("antecedent")
	RuleFeatureConfidence = RuleFeature("confidence")
	RuleFeatureConsequent = RuleFeature("consequent")
	RuleFeatureLeverage   = RuleFeature("leverage")
	RuleFeatureLift       = RuleFeature("lift")
	RuleFeatureRule       = RuleFeature("rule")
	RuleFeatureRuleID     = RuleFeature("ruleId")
	RuleFeatureSupport    = RuleFeature("support")
)

/*
  <xs:simpleType name="SVM-CLASSIFICATION-METHOD">
    <xs:restriction base="xs:string">
      <xs:enumeration value="OneAgainstAll"/>
      <xs:enumeration value="OneAgainstOne"/>
    </xs:restriction>
  </xs:simpleType>
*/
type SVMClassificationMethod string

const (
	SVMClassificationMethodOneAgainstAll = SVMClassificationMethod("OneAgainstAll")
	SVMClassificationMethodOneAgainstOne = SVMClassificationMethod("OneAgainstOne")
)

/*
  <xs:simpleType name="SVM-REPRESENTATION">
    <xs:restriction base="xs:string">
      <xs:enumeration value="Coefficients"/>
      <xs:enumeration value="SupportVectors"/>
    </xs:restriction>
  </xs:simpleType>
*/
type SVMRepresentation string

const (
	SVMRepresentationCoefficients   = SVMRepresentation("Coefficients")
	SVMRepresentationSupportVectors = SVMRepresentation("SupportVectors")
)

/*
  <xs:simpleType name="TIME-ANCHOR">
    <xs:restriction base="xs:string">
      <xs:enumeration value="dateDaysSince[0]"/>
      <xs:enumeration value="dateDaysSince[1960]"/>
      <xs:enumeration value="dateDaysSince[1970]"/>
      <xs:enumeration value="dateDaysSince[1980]"/>
      <xs:enumeration value="dateMonthsSince[0]"/>
      <xs:enumeration value="dateMonthsSince[1960]"/>
      <xs:enumeration value="dateMonthsSince[1970]"/>
      <xs:enumeration value="dateMonthsSince[1980]"/>
      <xs:enumeration value="dateTimeMillisecondsSince[0]"/>
      <xs:enumeration value="dateTimeMillisecondsSince[1960]"/>
      <xs:enumeration value="dateTimeMillisecondsSince[1970]"/>
      <xs:enumeration value="dateTimeMillisecondsSince[1980]"/>
      <xs:enumeration value="dateTimeSecondsSince[0]"/>
      <xs:enumeration value="dateTimeSecondsSince[1960]"/>
      <xs:enumeration value="dateTimeSecondsSince[1970]"/>
      <xs:enumeration value="dateTimeSecondsSince[1980]"/>
      <xs:enumeration value="dateYearsSince[0]"/>
    </xs:restriction>
  </xs:simpleType>
*/
type TimeAnchorType string

const (
	TimeAnchorTypedateDaysSince0                = TimeAnchorType("dateDaysSince[0]")
	TimeAnchorTypedateDaysSince1960             = TimeAnchorType("dateDaysSince[1960]")
	TimeAnchorTypedateDaysSince1970             = TimeAnchorType("dateDaysSince[1970]")
	TimeAnchorTypedateDaysSince1980             = TimeAnchorType("dateDaysSince[1980]")
	TimeAnchorTypedateMonthsSince0              = TimeAnchorType("dateMonthsSince[0]")
	TimeAnchorTypedateMonthsSince1960           = TimeAnchorType("dateMonthsSince[1960]")
	TimeAnchorTypedateMonthsSince1970           = TimeAnchorType("dateMonthsSince[1970]")
	TimeAnchorTypedateMonthsSince1980           = TimeAnchorType("dateMonthsSince[1980]")
	TimeAnchorTypedateTimeMillisecondsSince0    = TimeAnchorType("dateTimeMillisecondsSince[0]")
	TimeAnchorTypedateTimeMillisecondsSince1960 = TimeAnchorType("dateTimeMillisecondsSince[1960]")
	TimeAnchorTypedateTimeMillisecondsSince1970 = TimeAnchorType("dateTimeMillisecondsSince[1970]")
	TimeAnchorTypedateTimeMillisecondsSince1980 = TimeAnchorType("dateTimeMillisecondsSince[1980]")
	TimeAnchorTypedateTimeSecondsSince0         = TimeAnchorType("dateTimeSecondsSince[0]")
	TimeAnchorTypedateTimeSecondsSince1960      = TimeAnchorType("dateTimeSecondsSince[1960]")
	TimeAnchorTypedateTimeSecondsSince1970      = TimeAnchorType("dateTimeSecondsSince[1970]")
	TimeAnchorTypedateTimeSecondsSince1980      = TimeAnchorType("dateTimeSecondsSince[1980]")
	TimeAnchorTypedateYearsSince0               = TimeAnchorType("dateYearsSince[0]")
)

/*
  <xs:simpleType name="TIME-EXCEPTION-TYPE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="exclude"/>
      <xs:enumeration value="include"/>
    </xs:restriction>
  </xs:simpleType>
*/
type TimeExceptionType string

const (
	TimeExceptionTypeExclude = TimeExceptionType("exclude")
	TimeExceptionTypeInclude = TimeExceptionType("include")
)

/*
  <xs:simpleType name="TIMESERIES-ALGORITHM">
    <xs:restriction base="xs:string">
      <xs:enumeration value="ARIMA"/>
      <xs:enumeration value="ExponentialSmoothing"/>
      <xs:enumeration value="SeasonalTrendDecomposition"/>
      <xs:enumeration value="SpectralAnalysis"/>
    </xs:restriction>
  </xs:simpleType>
*/
type TimeSeriesAlgorithm string

const (
	TimeSeriesAlgorithmARIMA                      = TimeSeriesAlgorithm("ARIMA")
	TimeSeriesAlgorithmExponentialSmoothing       = TimeSeriesAlgorithm("ExponentialSmoothing")
	TimeSeriesAlgorithmSeasonalTrendDecomposition = TimeSeriesAlgorithm("SeasonalTrendDecomposition")
	TimeSeriesAlgorithmSpectralAnalysis           = TimeSeriesAlgorithm("SpectralAnalysis")
)

/*
  <xs:simpleType name="TIMESERIES-USAGE">
    <xs:restriction base="xs:string">
      <xs:enumeration value="logical"/>
      <xs:enumeration value="original"/>
      <xs:enumeration value="prediction"/>
    </xs:restriction>
  </xs:simpleType>
*/
type TimeSeriesUsage string

const (
	TimeSeriesUsageLogical    = TimeSeriesUsage("logical")
	TimeSeriesUsageOriginal   = TimeSeriesUsage("original")
	TimeSeriesUsagePrediction = TimeSeriesUsage("prediction")
)

/*
  <xs:simpleType name="VALID-TIME-SPEC">
    <xs:restriction base="xs:string">
      <xs:enumeration value="excludeFromTo"/>
      <xs:enumeration value="excludeSet"/>
      <xs:enumeration value="includeAll"/>
      <xs:enumeration value="includeFromTo"/>
      <xs:enumeration value="includeSet"/>
    </xs:restriction>
  </xs:simpleType>
*/
type ValidTimeSpec string

const (
	ValidTimeSpecExcludeFromTo = ValidTimeSpec("excludeFromTo")
	ValidTimeSpecExcludeSet    = ValidTimeSpec("excludeSet")
	ValidTimeSpecIncludeAll    = ValidTimeSpec("includeAll")
	ValidTimeSpecIncludeFromTo = ValidTimeSpec("includeFromTo")
	ValidTimeSpecIncludeSet    = ValidTimeSpec("includeSet")
)

/*
  <xs:simpleType name="VECTOR-ID">
    <xs:restriction base="xs:string"/>
  </xs:simpleType>
*/
type VectorID string
