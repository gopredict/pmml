package pmml

type DataType string

const (
	DataTypeString                   = DataType("string")
	DataTypeInteger                  = DataType("integer")
	DataTypeFloat                    = DataType("float")
	DataTypeDouble                   = DataType("double")
	DataTypeBoolean                  = DataType("boolean")
	DataTypeDate                     = DataType("date")
	DataTypeTime                     = DataType("time")
	DataTypeDateTime                 = DataType("dateTime")
	DataTypeDateDaysSince0           = DataType("dateDaysSince[0]")
	DataTypeDateDaysSince1960        = DataType("dateDaysSince[1960]")
	DataTypeDateDaysSince1970        = DataType("dateDaysSince[1970]")
	DataTypeDateDaysSince1980        = DataType("dateDaysSince[1980]")
	DataTypeTimeSeconds              = DataType("timeSeconds")
	DataTypeDateTimeSecondsSince0    = DataType("dateTimeSecondsSince[0]")
	DataTypeDateTimeSecondsSince1960 = DataType("dateTimeSecondsSince[1960]")
	DataTypeDateTimeSecondsSince1970 = DataType("dateTimeSecondsSince[1970]")
	DataTypeDateTimeSecondsSince1980 = DataType("dateTimeSecondsSince[1980]")
)

type OpType string

const (
	OpTypeCategorical = OpType("categorical")
	OpTypeOrdinal     = OpType("ordinal")
	OpTypeContinuous  = OpType("continuous")
)

type MiningFunction string

const (
	MiningFunctionAssociationRules = MiningFunction("associationRules")
	MiningFunctionSequences        = MiningFunction("sequences")
	MiningFunctionClassification   = MiningFunction("classification")
	MiningFunctionRegression       = MiningFunction("regression")
	MiningFunctionClustering       = MiningFunction("clustering")
	MiningFunctionTimeSeries       = MiningFunction("timeSeries")
	MiningFunctionMixed            = MiningFunction("mixed")
)

type ArrayTypeType string

const (
	ArrayTypeTypeInt    = ArrayTypeType("int")
	ArrayTypeTypeReal   = ArrayTypeType("real")
	ArrayTypeTypeString = ArrayTypeType("string")
)

type FieldUsageType string

const (
	FieldUsageTypeActive          = FieldUsageType("active")
	FieldUsageTypePredicted       = FieldUsageType("predicted")
	FieldUsageTypeTarget          = FieldUsageType("target")
	FieldUsageTypeSupplementary   = FieldUsageType("supplementary")
	FieldUsageTypeGroup           = FieldUsageType("group")
	FieldUsageTypeOrder           = FieldUsageType("order")
	FieldUsageTypeFrequencyWeight = FieldUsageType("frequencyWeight")
	FieldUsageTypeAnalysisWeight  = FieldUsageType("analysisWeight")
)

type OutlierTreatmentMethod string

const (
	OutlierTreatmentMethodAsIs            = OutlierTreatmentMethod("asIs")
	OutlierTreatmentMethodAsMissingValues = OutlierTreatmentMethod("asMissingValues")
	OutlierTreatmentMethodAsExtremeValues = OutlierTreatmentMethod("asExtremeValues")
)

type MissingValueTreatmentMethod string

const (
	MissingValueTreatmentMethodAsIs     = MissingValueTreatmentMethod("asIs")
	MissingValueTreatmentMethodAsMean   = MissingValueTreatmentMethod("asMean")
	MissingValueTreatmentMethodAsMode   = MissingValueTreatmentMethod("asMode")
	MissingValueTreatmentMethodAsMedian = MissingValueTreatmentMethod("asMedian")
	MissingValueTreatmentMethodAsValue  = MissingValueTreatmentMethod("asValue")
)

type InvalidValueTreatmentMethod string

const (
	InvalidValueTreatmentMethodReturnInvalid = InvalidValueTreatmentMethod("returnInvalid")
	InvalidValueTreatmentMethodAsIs          = InvalidValueTreatmentMethod("asIs")
	InvalidValueTreatmentMethodAsMissing     = InvalidValueTreatmentMethod("asMissing")
)

type MultipleModelMethod string

const (
	MultipleModelMethodMajorityVote         = MultipleModelMethod("majorityVote")
	MultipleModelMethodWeightedMajorityVote = MultipleModelMethod("weightedMajorityVote")
	MultipleModelMethodAverage              = MultipleModelMethod("average")
	MultipleModelMethodWeightedAverage      = MultipleModelMethod("weightedAverage")
	MultipleModelMethodMedian               = MultipleModelMethod("median")
	MultipleModelMethodMax                  = MultipleModelMethod("max")
	MultipleModelMethodSum                  = MultipleModelMethod("sum")
	MultipleModelMethodSelectFirst          = MultipleModelMethod("selectFirst")
	MultipleModelMethodSelectAll            = MultipleModelMethod("selectAll")
	MultipleModelMethodModelChain           = MultipleModelMethod("modelChain")
)

type SplitCharacteristic string

const (
	SplitCharacteristicBinarySplit = SplitCharacteristic("binarySplit")
	SplitCharacteristicMultiSplit  = SplitCharacteristic("multiSplit")
)

type ActivationFunction string

const (
	ActivationFunctionThreshold   = ActivationFunction("threshold")
	ActivationFunctionLogistic    = ActivationFunction("logistic")
	ActivationFunctionTanh        = ActivationFunction("tanh")
	ActivationFunctionIdentity    = ActivationFunction("identity")
	ActivationFunctionExponential = ActivationFunction("exponential")
	ActivationFunctionReciprocal  = ActivationFunction("reciprocal")
	ActivationFunctionSquare      = ActivationFunction("square")
	ActivationFunctionGauss       = ActivationFunction("Gauss")
	ActivationFunctionSine        = ActivationFunction("sine")
	ActivationFunctionCosine      = ActivationFunction("cosine")
	ActivationFunctionElliott     = ActivationFunction("Elliott")
	ActivationFunctionArcTan      = ActivationFunction("arctan")
	ActivationFunctionRectifier   = ActivationFunction("rectifier")
	ActivationFunctionRadialBasis = ActivationFunction("radialBasis")
)

type ResultFeature string

const (
	ResultFeaturePredictedValue        = ResultFeature("predictedValue")
	ResultFeaturePredictedDisplayValue = ResultFeature("predictedDisplayValue")
	ResultFeatureTransformedValue      = ResultFeature("transformedValue")
	ResultFeatureDecision              = ResultFeature("decision")
	ResultFeatureProbability           = ResultFeature("probability")
	ResultFeatureAffinity              = ResultFeature("affinity")
	ResultFeatureResidual              = ResultFeature("residual")
	ResultFeatureStandardError         = ResultFeature("standardError")
	ResultFeatureClusterID             = ResultFeature("clusterId")
	ResultFeatureClusterAffinity       = ResultFeature("clusterAffinity")
	ResultFeatureEntityID              = ResultFeature("entityId")
	ResultFeatureEntityAffinity        = ResultFeature("entityAffinity")
	ResultFeatureWarning               = ResultFeature("warning")
	ResultFeatureRuleValue             = ResultFeature("ruleValue")
	ResultFeatureReasonCode            = ResultFeature("reasonCode")
	ResultFeatureAntecedent            = ResultFeature("antecedent")
	ResultFeatureConsequent            = ResultFeature("consequent")
	ResultFeatureRule                  = ResultFeature("rule")
	ResultFeatureRuleID                = ResultFeature("ruleId")
	ResultFeatureConfidence            = ResultFeature("confidence")
	ResultFeatureSupport               = ResultFeature("support")
	ResultFeatureLift                  = ResultFeature("lift")
	ResultFeatureLeverage              = ResultFeature("leverage")
)

type RuleFeature string

const (
	RuleFeatureAntecedent = RuleFeature("antecedent")
	RuleFeatureConsequent = RuleFeature("consequent")
	RuleFeatureRule       = RuleFeature("rule")
	RuleFeatureRuleID     = RuleFeature("ruleId")
	RuleFeatureConfidence = RuleFeature("confidence")
	RuleFeatureSupport    = RuleFeature("support")
	RuleFeatureLift       = RuleFeature("lift")
	RuleFeatureLeverage   = RuleFeature("leverage")
	RuleFeatureAffinity   = RuleFeature("affinity")
)

type RegressionModelType string

const (
	RegressionModelTypeLinearRegression             = RegressionModelType("linearRegression")
	RegressionModelTypeStepwisePolynomialRegression = RegressionModelType("stepwisePolynomialRegression")
	RegressionModelTypeLogisticRegression           = RegressionModelType("logisticRegression")
)

type RegressionNormalizationMethod string

const (
	RegressionNormalizationMethodNone      = RegressionNormalizationMethod("none")
	RegressionNormalizationMethodSimplemax = RegressionNormalizationMethod("simplemax")
	RegressionNormalizationMethodSoftmax   = RegressionNormalizationMethod("softmax")
	RegressionNormalizationMethodLogit     = RegressionNormalizationMethod("logit")
	RegressionNormalizationMethodProbit    = RegressionNormalizationMethod("probit")
	RegressionNormalizationMethodCloglog   = RegressionNormalizationMethod("cloglog")
	RegressionNormalizationMethodExp       = RegressionNormalizationMethod("exp")
	RegressionNormalizationMethodLoglog    = RegressionNormalizationMethod("loglog")
	RegressionNormalizationMethodCauchit   = RegressionNormalizationMethod("cauchit")
)

type RuleSelectionMethodCriterion string

const (
	RuleSelectionMethodCriterionWeightedSum = RuleSelectionMethodCriterion("weightedSum")
	RuleSelectionMethodCriterionWeightedMax = RuleSelectionMethodCriterion("weightedMax")
	RuleSelectionMethodCriterionFirstHit    = RuleSelectionMethodCriterion("firstHit")
)

type DelimeterType string

const (
	DelimeterTypeSameTimeWindow    = DelimeterType("sameTimeWindow")
	DelimeterTypeAcrossTimeWindows = DelimeterType("acrossTimeWindows")
)

type Gap string

const (
	GapTrue    = Gap("true")
	GapFalse   = Gap("false")
	GapUnknown = Gap("unknown")
)

type SVMClassificationMethod string

const (
	SVMClassificationMethodOneAgainstAll = SVMClassificationMethod("OneAgainstAll")
	SVMClassificationMethodOneAgainstOne = SVMClassificationMethod("OneAgainstOne")
)

type SVMRepresentation string

const (
	SVMRepresentationSupportVectors = SVMRepresentation("SupportVectors")
	SVMRepresentationCoefficients   = SVMRepresentation("Coefficients")
)

type TimeSeriesAlgorithm string

const (
	TimeSeriesAlgorithmARIMA                      = TimeSeriesAlgorithm("ARIMA")
	TimeSeriesAlgorithmExponentialSmoothing       = TimeSeriesAlgorithm("ExponentialSmoothing")
	TimeSeriesAlgorithmSeasonalTrendDecomposition = TimeSeriesAlgorithm("SeasonalTrendDecomposition")
	TimeSeriesAlgorithmSpectralAnalysis           = TimeSeriesAlgorithm("SpectralAnalysis")
)

type TimeSeriesUsage string

const (
	TimeSeriesUsageOriginal   = TimeSeriesUsage("original")
	TimeSeriesUsageLogical    = TimeSeriesUsage("logical")
	TimeSeriesUsagePrediction = TimeSeriesUsage("prediction")
)

type TimeAnchor string

const (
	TimeAnchorDateTimeMillisecondsSince0    = TimeAnchor("dateTimeMillisecondsSince[0]")
	TimeAnchorDateTimeMillisecondsSince1960 = TimeAnchor("dateTimeMillisecondsSince[1960]")
	TimeAnchorDateTimeMillisecondsSince1970 = TimeAnchor("dateTimeMillisecondsSince[1970]")
	TimeAnchorDateTimeMillisecondsSince1980 = TimeAnchor("dateTimeMillisecondsSince[1980]")
	TimeAnchorDateTimeSecondsSince0         = TimeAnchor("dateTimeSecondsSince[0]")
	TimeAnchorDateTimeSecondsSince1960      = TimeAnchor("dateTimeSecondsSince[1960]")
	TimeAnchorDateTimeSecondsSince1970      = TimeAnchor("dateTimeSecondsSince[1970]")
	TimeAnchorDateTimeSecondsSince1980      = TimeAnchor("dateTimeSecondsSince[1980]")
	TimeAnchorDateDaysSince0                = TimeAnchor("dateDaysSince[0]")
	TimeAnchorDateDaysSince1960             = TimeAnchor("dateDaysSince[1960]")
	TimeAnchorDateDaysSince1970             = TimeAnchor("dateDaysSince[1970]")
	TimeAnchorDateDaysSince1980             = TimeAnchor("dateDaysSince[1980]")
	TimeAnchorDateMonthsSince0              = TimeAnchor("dateMonthsSince[0]")
	TimeAnchorDateMonthsSince1960           = TimeAnchor("dateMonthsSince[1960]")
	TimeAnchorDateMonthsSince1970           = TimeAnchor("dateMonthsSince[1970]")
	TimeAnchorDateMonthsSince1980           = TimeAnchor("dateMonthsSince[1980]")
	TimeAnchorDateYearsSince0               = TimeAnchor("dateYearsSince[0]")
)

type ValidTimeSpec string

const (
	ValidTimeSpecIncludeAll    = ValidTimeSpec("includeAll")
	ValidTimeSpecIncludeFromTo = ValidTimeSpec("includeFromTo")
	ValidTimeSpecExcludeFromTo = ValidTimeSpec("excludeFromTo")
	ValidTimeSpecIncludeSet    = ValidTimeSpec("includeSet")
	ValidTimeSpecExcludeSet    = ValidTimeSpec("excludeSet")
)

type TimeExceptionType string

const (
	TimeExceptionTypeExclude = TimeExceptionType("exclude")
	TimeExceptionTypeInclude = TimeExceptionType("include")
)

type InterpolationMethod string

const (
	InterpolationMethodNone              = InterpolationMethod("none")
	InterpolationMethodLinear            = InterpolationMethod("linear")
	InterpolationMethodExponentialSpline = InterpolationMethod("exponentialSpline")
	InterpolationMethodCubicSpline       = InterpolationMethod("cubicSpline")
)

type MissingValueStrategy string

const (
	MissingValueStrategyLastPrediction     = MissingValueStrategy("lastPrediction")
	MissingValueStrategyNullPrediction     = MissingValueStrategy("nullPrediction")
	MissingValueStrategyDefaultChild       = MissingValueStrategy("defaultChild")
	MissingValueStrategyWeightedConfidence = MissingValueStrategy("weightedConfidence")
	MissingValueStrategyAggregateNodes     = MissingValueStrategy("aggregateNodes")
	MissingValueStrategyNone               = MissingValueStrategy("none")
)

type NoTrueChildStrategy string

const (
	NoTrueChildStrategyReturnNullPrediction = NoTrueChildStrategy("returnNullPrediction")
	NoTrueChildStrategyReturnLastPrediction = NoTrueChildStrategy("returnLastPrediction")
)

type BaselineTestStatistic string

const (
	BaselineTestStatisticZValue                = BaselineTestStatistic("zValue")
	BaselineTestStatisticChiSquareIndependence = BaselineTestStatistic("chiSquareIndependence")
	BaselineTestStatisticChiSquareDistribution = BaselineTestStatistic("chiSquareDistribution")
	BaselineTestStatisticCUSUM                 = BaselineTestStatistic("CUSUM")
	BaselineTestStatisticScalarProduct         = BaselineTestStatistic("scalarProduct")
)

type CastInteger string

const (
	CastIntegerRound   = CastInteger("round")
	CastIntegerCeiling = CastInteger("ceiling")
	CastIntegerFloor   = CastInteger("floor")
)
