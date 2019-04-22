package models

/*
  <xs:group name="CONTINUOUS-DISTRIBUTION-TYPES">
    <xs:sequence>
      <xs:choice>
        <xs:element ref="AnyDistribution"/>
        <xs:element ref="GaussianDistribution"/>
        <xs:element ref="PoissonDistribution"/>
        <xs:element ref="UniformDistribution"/>
      </xs:choice>
      <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
    </xs:sequence>
  </xs:group>
*/
type ContinuousDistributionType interface {
	continuousDistributionType()
}

/*
  <xs:group name="DISCRETE-DISTRIBUTION-TYPES">
    <xs:choice>
      <xs:element ref="CountTable"/>
      <xs:element maxOccurs="unbounded" minOccurs="2" ref="FieldRef"/>
      <xs:element ref="NormalizedCountTable"/>
    </xs:choice>
  </xs:group>
*/
type DiscreteDistributionType interface {
}

/*
  <xs:group name="EXPRESSION">
    <xs:choice>
      <xs:element ref="Aggregate"/>
      <xs:element ref="Apply"/>
      <xs:element ref="Constant"/>
      <xs:element ref="Discretize"/>
      <xs:element ref="FieldRef"/>
      <xs:element ref="Lag"/>
      <xs:element ref="MapValues"/>
      <xs:element ref="NormContinuous"/>
      <xs:element ref="NormDiscrete"/>
      <xs:element ref="TextIndex"/>
    </xs:choice>
  </xs:group>
*/
type Expression interface {
	expression()
}

/*
  <xs:group name="EmbeddedModel">
    <xs:sequence>
      <xs:choice>
        <xs:element ref="DecisionTree"/>
        <xs:element ref="Regression"/>
      </xs:choice>
      <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
    </xs:sequence>
  </xs:group>
*/
type EmbeddedModel interface {
}

/*
  <xs:group name="FOLLOW-SET">
    <xs:sequence>
      <xs:element ref="Delimiter"/>
      <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      <xs:element ref="SetReference"/>
      <xs:element minOccurs="0" ref="Time"/>
    </xs:sequence>
  </xs:group>
*/
type FollowSet struct {
	Delimiter    Delimiter    `xml:"Delimiter"`
	SetReference SetReference `xml:"SetReference"`
	Time         *Time        `xml:"Time"`
}

/*
  <xs:group name="FrequenciesType">
    <xs:sequence>
      <xs:group maxOccurs="3" minOccurs="1" ref="NUM-ARRAY"/>
    </xs:sequence>
  </xs:group>
*/
type FrequenciesType struct {
	NumArrays []NumArray `xml:"Array"`
}

/*
  <xs:group name="INT-ARRAY">
    <xs:choice>
      <xs:element ref="Array"/>
    </xs:choice>
  </xs:group>
*/
type IntArray Array

/*
  <xs:group name="MODEL-ELEMENT">
    <xs:choice>
      <xs:element ref="AssociationModel"/>
      <xs:element ref="BaselineModel"/>
      <xs:element ref="BayesianNetworkModel"/>
      <xs:element ref="ClusteringModel"/>
      <xs:element ref="GaussianProcessModel"/>
      <xs:element ref="GeneralRegressionModel"/>
      <xs:element ref="MiningModel"/>
      <xs:element ref="NaiveBayesModel"/>
      <xs:element ref="NearestNeighborModel"/>
      <xs:element ref="NeuralNetwork"/>
      <xs:element ref="RegressionModel"/>
      <xs:element ref="RuleSetModel"/>
      <xs:element ref="Scorecard"/>
      <xs:element ref="SequenceModel"/>
      <xs:element ref="SupportVectorMachineModel"/>
      <xs:element ref="TextModel"/>
      <xs:element ref="TimeSeriesModel"/>
      <xs:element ref="TreeModel"/>
    </xs:choice>
  </xs:group>
*/
type ModelElement interface {
	modelElement()
}

/*
  <xs:group name="NUM-ARRAY">
    <xs:choice>
      <xs:element ref="Array"/>
    </xs:choice>
  </xs:group>
*/
type NumArray Array

/*
  <xs:group name="PREDICATE">
    <xs:choice>
      <xs:element ref="CompoundPredicate"/>
      <xs:element ref="False"/>
      <xs:element ref="SimplePredicate"/>
      <xs:element ref="SimpleSetPredicate"/>
      <xs:element ref="True"/>
    </xs:choice>
  </xs:group>
*/
type Predicate struct {
}

/*
  <xs:group name="REAL-ARRAY">
    <xs:choice>
      <xs:element ref="Array"/>
    </xs:choice>
  </xs:group>
*/
type RealArray Array

/*
  <xs:group name="Rule">
    <xs:choice>
      <xs:element ref="CompoundRule"/>
      <xs:element ref="SimpleRule"/>
    </xs:choice>
  </xs:group>
*/
type Rule interface {
}

/*
  <xs:group name="SEQUENCE">
    <xs:sequence>
      <xs:element ref="SequenceReference"/>
      <xs:element minOccurs="0" ref="Time"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:sequence>
  </xs:group>
*/
type SequenceGroup struct {
	SequenceReference SequenceReference `xml:"SequenceReference"`
	Time              *Time             `xml:"Time"`
	Extensions        []Extension       `xml:"Extension"`
}

/*
  <xs:group name="STRING-ARRAY">
    <xs:choice>
      <xs:element ref="Array"/>
    </xs:choice>
  </xs:group>
*/
type StringArray Array
