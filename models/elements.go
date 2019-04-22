package models

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

type Error string

func (err Error) Error() string {
	return string(err)
}

const (
	ErrNotImplemented = Error("models: not implemented")
)

/*
  <xs:element name="ARDSquaredExponentialKernel">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="noiseVariance" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Lambda"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ARDSquaredExponentialKernel struct {
	Description   *string     `xml:"description,attr"`
	Gamma         *RealNumber `xml:"gamma,attr"`
	NoiseVariance *RealNumber `xml:"noiseVariance,attr"`

	Extensions []Extension `xml:"Extension"`
	Lambda     *Lambda     `xml:"Lambda"`
}

/*
  <xs:element name="ARIMA"/>
*/
type ARIMA struct {
}

/*
  <xs:element name="AbsoluteExponentialKernel">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="noiseVariance" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Lambda"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AbsoluteExponentialKernel struct {
	Description   *string     `xml:"description,attr"`
	Gamma         *RealNumber `xml:"gamma,attr"`
	NoiseVariance *RealNumber `xml:"noiseVariance,attr"`

	Extensions []Extension `xml:"Extension"`
	Lambda     *Lambda     `xml:"Lambda"`
}

type AggregateFunctionType string

const (
	AggregateFunctionTypeAverage  = AggregateFunctionType("average")
	AggregateFunctionTypeCount    = AggregateFunctionType("count")
	AggregateFunctionTypeMax      = AggregateFunctionType("max")
	AggregateFunctionTypeMin      = AggregateFunctionType("min")
	AggregateFunctionTypeMultiset = AggregateFunctionType("multiset")
	AggregateFunctionTypeSum      = AggregateFunctionType("sum")
)

/*
  <xs:element name="Aggregate">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="function" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="average"/>
            <xs:enumeration value="count"/>
            <xs:enumeration value="max"/>
            <xs:enumeration value="min"/>
            <xs:enumeration value="multiset"/>
            <xs:enumeration value="sum"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="groupField" type="FIELD-NAME"/>
      <xs:attribute name="sqlWhere" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Aggregate struct {
	Field      FieldName             `xml:"field,attr"`
	Function   AggregateFunctionType `xml:"function,attr"`
	GroupField FieldName             `xml:"groupField,attr"`
	SQLWhere   string                `xml:"sqlWhere,attr"`

	Extensions []Extension `xml:"Extension"`
}

func (Aggregate) expression() {}

/*
  <xs:element name="Alternate">
    <xs:complexType>
      <xs:choice>
        <xs:group minOccurs="1" ref="CONTINUOUS-DISTRIBUTION-TYPES"/>
      </xs:choice>
    </xs:complexType>
  </xs:element>
*/
type Alternate struct {
	Distribution ContinuousDistributionType
}

func (x *Alternate) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var item ContinuousDistributionType

			switch tok.Name.Local {
			case "AnyDistribution":
				item = &AnyDistribution{}
			case "GaussianDistribution":
				item = &GaussianDistribution{}
			case "PoissonDistribution":
				item = &PoissonDistribution{}
			case "UniformDistribution":
				item = &UniformDistribution{}
			}

			if item != nil {
				err = decoder.DecodeElement(item, &tok)
				if err != nil {
					return err
				}

				x.Distribution = item
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="Annotation">
    <xs:complexType mixed="true">
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Annotation struct {
	Data       string      `xml:",chardata"`
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Anova">
    <xs:complexType>
      <xs:attribute name="target" type="FIELD-NAME"/>
      <xs:sequence>
        <xs:element maxOccurs="3" minOccurs="3" ref="AnovaRow"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Anova struct {
	Taget FieldName `xml:"target,attr"`

	Rows       []AnovaRow  `xml:"AnovaRow"`
	Extensions []Extension `xml:"Extension"`
}

type AnovaRowType string

const (
	AnovaRowTypeError = AnovaRowType("Error")
	AnovaRowTypeModel = AnovaRowType("Model")
	AnovaRowTypeTotal = AnovaRowType("Total")
)

/*
  <xs:element name="AnovaRow">
    <xs:complexType>
      <xs:attribute name="degreesOfFreedom" type="NUMBER" use="required"/>
      <xs:attribute name="fValue" type="NUMBER"/>
      <xs:attribute name="meanOfSquares" type="NUMBER"/>
      <xs:attribute name="pValue" type="PROB-NUMBER"/>
      <xs:attribute name="sumOfSquares" type="NUMBER" use="required"/>
      <xs:attribute name="type" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="Error"/>
            <xs:enumeration value="Model"/>
            <xs:enumeration value="Total"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AnovaRow struct {
	DegreesOfFreedom Number       `xml:"degreesOfFreedom,attr"`
	FValue           Number       `xml:"fValue,attr"`
	MeanOfSquares    Number       `xml:"meanOfSquares,attr"`
	PValue           ProbNumber   `xml:"pValue,attr"`
	SumOfSquares     Number       `xml:"sumOfSquares,attr"`
	Type             AnovaRowType `xml:"type,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="AntecedentSequence">
    <xs:complexType>
      <xs:sequence>
        <xs:group ref="SEQUENCE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AntecedentSequence SequenceGroup

/*
  <xs:element name="AnyDistribution">
    <xs:complexType>
      <xs:attribute name="mean" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="variance" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AnyDistribution struct {
	Mean     RealNumber `xml:"mean,attr"`
	Variance RealNumber `xml:"variance,attr"`

	Extensions []Extension `xml:"Extension"`
}

func (*AnyDistribution) continuousDistributionType() {}

/*
  <xs:element name="Application">
    <xs:complexType>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:attribute name="version" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Application struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Apply">
    <xs:complexType>
      <xs:attribute name="defaultValue" type="xs:string"/>
      <xs:attribute name="function" type="xs:string" use="required"/>
      <xs:attribute default="returnInvalid" name="invalidValueTreatment" type="INVALID-VALUE-TREATMENT-METHOD"/>
      <xs:attribute name="mapMissingTo" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group maxOccurs="unbounded" minOccurs="0" ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Apply struct {
	DefaultValue          string                      `xml:"defaultValue,attr"`
	Function              string                      `xml:"function,attr"`
	InvalidValueTreatment InvalidValueTreatmentMethod `xml:"returnInvalid,attr"`
	MapMissingTo          string                      `xml:"mapMissingTo,attr"`

	Extensions  []Extension `xml:"Extension"`
	Expressions []Expression
}

func (*Apply) expression() {}

func (x *Apply) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "defaultValue":
			x.DefaultValue = attr.Value
		case "function":
			x.Function = attr.Value
		case "returnInvalid":
			x.InvalidValueTreatment = InvalidValueTreatmentMethod(attr.Value)
		case "mapMissingTo":
			x.MapMissingTo = attr.Value
		}
	}

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var item Expression

			switch tok.Name.Local {
			case "Extension":
				var e Extension
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.Extensions = append(x.Extensions, e)
			case "Aggregate":
				item = &Aggregate{}
			case "Apply":
				item = &Apply{}
			case "Constant":
				item = &Constant{}
			case "Discretize":
				item = &Discretize{}
			case "FieldRef":
				item = &FieldRef{}
			case "Lag":
				item = &Lag{}
			case "MapValues":
				item = &MapValues{}
			case "NormContinuous":
				item = &NormContinuous{}
			case "NormDiscrete":
				item = &NormDiscrete{}
			case "TextIndex":
				item = &TextIndex{}
			}

			if item != nil {
				err = decoder.DecodeElement(item, &tok)
				if err != nil {
					return err
				}

				x.Expressions = append(x.Expressions, item)
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="Array" type="ArrayType"/>
*/
type Array ArrayType

/*
  <xs:element name="AssociationModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="avgNumberOfItemsPerTA" type="REAL-NUMBER"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="lengthLimit" type="INT-NUMBER"/>
      <xs:attribute name="maxNumberOfItemsPerTA" type="INT-NUMBER"/>
      <xs:attribute name="minimumConfidence" type="PROB-NUMBER" use="required"/>
      <xs:attribute name="minimumSupport" type="PROB-NUMBER" use="required"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="numberOfItems" type="INT-NUMBER" use="required"/>
      <xs:attribute name="numberOfItemsets" type="INT-NUMBER" use="required"/>
      <xs:attribute name="numberOfRules" type="INT-NUMBER" use="required"/>
      <xs:attribute name="numberOfTransactions" type="INT-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="AssociationRule"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Item"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Itemset"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AssociationModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*AssociationModel) modelElement() {}

/*
  <xs:element name="AssociationRule">
    <xs:complexType>
      <xs:attribute name="affinity" type="PROB-NUMBER" use="optional"/>
      <xs:attribute name="antecedent" type="xs:string" use="required"/>
      <xs:attribute name="confidence" type="PROB-NUMBER" use="required"/>
      <xs:attribute name="consequent" type="xs:string" use="required"/>
      <xs:attribute name="id" type="xs:string" use="optional"/>
      <xs:attribute name="leverage" type="xs:float" use="optional"/>
      <xs:attribute name="lift" type="xs:float" use="optional"/>
      <xs:attribute name="support" type="PROB-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type AssociationRule struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Attribute">
    <xs:complexType>
      <xs:attribute name="partialScore" type="NUMBER" use="optional"/>
      <xs:attribute name="reasonCode" type="xs:string"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="ComplexPartialScore"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="PREDICATE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Attribute struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BaseCumHazardTables">
    <xs:complexType>
      <xs:attribute name="maxTime" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:choice>
          <xs:element maxOccurs="unbounded" ref="BaselineCell"/>
          <xs:element maxOccurs="unbounded" ref="BaselineStratum"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BaseCumHazardTables struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BaselineCell">
    <xs:complexType>
      <xs:attribute name="cumHazard" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="time" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BaselineCell struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BaselineModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean" use="optional"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
        <xs:element ref="TestDistributions"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BaselineModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*BaselineModel) modelElement() {}

/*
  <xs:element name="BaselineStratum">
    <xs:complexType>
      <xs:attribute name="label" type="xs:string"/>
      <xs:attribute name="maxTime" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="BaselineCell"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BaselineStratum struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Baseline">
    <xs:complexType>
      <xs:choice>
        <xs:group minOccurs="1" ref="CONTINUOUS-DISTRIBUTION-TYPES"/>
        <xs:group minOccurs="1" ref="DISCRETE-DISTRIBUTION-TYPES"/>
      </xs:choice>
    </xs:complexType>
  </xs:element>
*/
type Baseline struct {
}

/*
  <xs:element name="BayesInput">
    <xs:complexType>
      <xs:attribute name="fieldName" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:choice>
          <xs:element maxOccurs="1" minOccurs="1" ref="TargetValueStats"/>
          <xs:sequence>
            <xs:element maxOccurs="1" minOccurs="0" ref="DerivedField"/>
            <xs:element maxOccurs="unbounded" minOccurs="1" ref="PairCounts"/>
          </xs:sequence>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BayesInput struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BayesInputs">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="BayesInput"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BayesInputs struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BayesOutput">
    <xs:complexType>
      <xs:attribute name="fieldName" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="TargetValueCounts"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BayesOutput struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BayesianNetworkModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:sequence>
        <xs:element minOccurs="1" ref="BayesianNetworkNodes"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BayesianNetworkModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*BayesianNetworkModel) modelElement() {}

/*
  <xs:element name="BayesianNetworkNodes">
    <xs:complexType>
      <xs:sequence>
        <xs:choice maxOccurs="unbounded">
          <xs:element ref="ContinuousNode"/>
          <xs:element ref="DiscreteNode"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BayesianNetworkNodes struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BlockIndicator">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
    </xs:complexType>
  </xs:element>
*/
type BlockIndicator struct {
}

/*
  <xs:element name="BoundaryValueMeans">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BoundaryValueMeans struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="BoundaryValues">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BoundaryValues struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CategoricalPredictor">
    <xs:complexType>
      <xs:attribute name="coefficient" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CategoricalPredictor struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Categories">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="Category"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Categories struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Category">
    <xs:complexType>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Category struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Characteristic">
    <xs:complexType>
      <xs:attribute name="baselineScore" type="NUMBER"/>
      <xs:attribute name="name" type="FIELD-NAME" use="optional"/>
      <xs:attribute name="reasonCode" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="Attribute"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Characteristic struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Characteristics">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="Characteristic"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Characteristics struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ChildParent">
    <xs:complexType>
      <xs:attribute name="childField" type="xs:string" use="required"/>
      <xs:attribute default="no" name="isRecursive" use="optional">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="no"/>
            <xs:enumeration value="yes"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="parentField" type="xs:string" use="required"/>
      <xs:attribute name="parentLevelField" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:choice>
          <xs:element ref="InlineTable"/>
          <xs:element ref="TableLocator"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="FieldColumnPair"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ChildParent struct {
	ChildField       string `xml:"childField,attr"`
	IsRecursive      bool   `xml:"isRecursive,attr"`
	ParentField      string `xml:"parentField,attr"`
	ParentLevelField string `xml:"parentLevelField,attr"`

	Table            Table
	Extensions       []Extension       `xml:"Extension"`
	FieldColumnPairs []FieldColumnPair `xml:"FieldColumnPair"`
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (x *ChildParent) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "childField":
			x.ChildField = attr.Value
		case "isRecursive":
			if attr.Value == "yes" {
				x.IsRecursive = true
			} else if attr.Value == "no" {
				x.IsRecursive = false
			} else {
				return errors.Errorf("invalid value for isRecursive: %s", attr.Value)
			}
		case "parentField":
			x.ParentField = attr.Value
		case "parentLevelField":
			x.ParentLevelField = attr.Value
		}
	}

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var item Table

			switch tok.Name.Local {
			case "Extension":
				var e Extension
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.Extensions = append(x.Extensions, e)
			case "FieldColumnPair":
				var e FieldColumnPair
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.FieldColumnPairs = append(x.FieldColumnPairs, e)
			case "InlineTable":
				item = &InlineTable{}
			case "TableLocator":
				item = &TableLocator{}
			}

			if item != nil {
				err = decoder.DecodeElement(item, &tok)
				if err != nil {
					return err
				}

				x.Table = item
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="ClassLabels">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="STRING-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ClassLabels struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Cluster">
    <xs:complexType>
      <xs:attribute name="id" type="xs:string" use="optional"/>
      <xs:attribute name="name" type="xs:string" use="optional"/>
      <xs:attribute name="size" type="xs:nonNegativeInteger" use="optional"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Covariances"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="KohonenMap"/>
        <xs:element minOccurs="0" ref="Partition"/>
        <xs:group minOccurs="0" ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Cluster struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ClusteringField">
    <xs:complexType>
      <xs:attribute name="compareFunction" type="COMPARE-FUNCTION" use="optional"/>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute default="1" name="fieldWeight" type="REAL-NUMBER"/>
      <xs:attribute default="true" name="isCenterField">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="false"/>
            <xs:enumeration value="true"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="similarityScale" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Comparisons"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ClusteringField struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ClusteringModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelClass" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="centerBased"/>
            <xs:enumeration value="distributionBased"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:attribute name="numberOfClusters" type="INT-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="Cluster"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="ClusteringField"/>
        <xs:element ref="ComparisonMeasure"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="MissingValueWeights"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ClusteringModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*ClusteringModel) modelElement() {}

/*
  <xs:element name="ClusteringModelQuality">
    <xs:complexType>
      <xs:attribute name="SSB" type="NUMBER" use="optional"/>
      <xs:attribute name="SSE" type="NUMBER" use="optional"/>
      <xs:attribute name="dataName" type="xs:string" use="optional"/>
    </xs:complexType>
  </xs:element>
*/
type ClusteringModelQuality struct {
}

/*
  <xs:element name="Coefficient">
    <xs:complexType>
      <xs:attribute default="0" name="value" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Coefficient struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Coefficients">
    <xs:complexType>
      <xs:attribute default="0" name="absoluteValue" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="numberOfCoefficients" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="Coefficient"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Coefficients struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ComparisonMeasure">
    <xs:complexType>
      <xs:attribute default="absDiff" name="compareFunction" type="COMPARE-FUNCTION"/>
      <xs:attribute name="kind" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="distance"/>
            <xs:enumeration value="similarity"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="maximum" type="NUMBER" use="optional"/>
      <xs:attribute name="minimum" type="NUMBER" use="optional"/>
      <xs:sequence>
        <xs:choice>
          <xs:element ref="binarySimilarity"/>
          <xs:element ref="chebychev"/>
          <xs:element ref="cityBlock"/>
          <xs:element ref="euclidean"/>
          <xs:element ref="jaccard"/>
          <xs:element ref="minkowski"/>
          <xs:element ref="simpleMatching"/>
          <xs:element ref="squaredEuclidean"/>
          <xs:element ref="tanimoto"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ComparisonMeasure struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Comparisons">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Comparisons struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ComplexPartialScore">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group maxOccurs="1" minOccurs="1" ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ComplexPartialScore struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CompoundPredicate">
    <xs:complexType>
      <xs:attribute name="booleanOperator" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="and"/>
            <xs:enumeration value="or"/>
            <xs:enumeration value="surrogate"/>
            <xs:enumeration value="xor"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:sequence maxOccurs="unbounded" minOccurs="2">
          <xs:group ref="PREDICATE"/>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CompoundPredicate struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CompoundRule">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="PREDICATE"/>
        <xs:group maxOccurs="unbounded" minOccurs="1" ref="Rule"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CompoundRule struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Con">
    <xs:complexType>
      <xs:attribute name="from" type="NN-NEURON-IDREF" use="required"/>
      <xs:attribute name="weight" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Con struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ConfusionMatrix">
    <xs:complexType>
      <xs:sequence>
        <xs:element ref="ClassLabels"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ConfusionMatrix struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ConsequentSequence">
    <xs:complexType>
      <xs:sequence>
        <xs:group ref="SEQUENCE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ConsequentSequence struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Constant">
    <xs:complexType>
      <xs:simpleContent>
        <xs:extension base="xs:string">
          <xs:attribute name="dataType" type="DATATYPE"/>
        </xs:extension>
      </xs:simpleContent>
    </xs:complexType>
  </xs:element>
*/
type Constant struct {
}

func (*Constant) expression() {}

/*
  <xs:element name="Constraints">
    <xs:complexType>
      <xs:attribute name="maximumAntConsSeparationTime" type="REAL-NUMBER"/>
      <xs:attribute name="maximumItemsetSeparationTime" type="REAL-NUMBER"/>
      <xs:attribute name="maximumNumberOfAntecedentItems" type="INT-NUMBER"/>
      <xs:attribute name="maximumNumberOfConsequentItems" type="INT-NUMBER"/>
      <xs:attribute name="maximumNumberOfItems" type="INT-NUMBER"/>
      <xs:attribute name="maximumTotalSequenceTime" type="REAL-NUMBER"/>
      <xs:attribute default="0" name="minimumAntConsSeparationTime" type="REAL-NUMBER"/>
      <xs:attribute default="0" name="minimumConfidence" type="REAL-NUMBER"/>
      <xs:attribute default="0" name="minimumItemsetSeparationTime" type="REAL-NUMBER"/>
      <xs:attribute default="0" name="minimumLift" type="REAL-NUMBER"/>
      <xs:attribute default="1" name="minimumNumberOfAntecedentItems" type="INT-NUMBER"/>
      <xs:attribute default="1" name="minimumNumberOfConsequentItems" type="INT-NUMBER"/>
      <xs:attribute default="1" name="minimumNumberOfItems" type="INT-NUMBER"/>
      <xs:attribute default="0" name="minimumSupport" type="REAL-NUMBER"/>
      <xs:attribute default="0" name="minimumTotalSequenceTime" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Constraints struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ContStats">
    <xs:complexType>
      <xs:attribute name="totalSquaresSum" type="NUMBER"/>
      <xs:attribute name="totalValuesSum" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Interval"/>
        <xs:group minOccurs="0" ref="FrequenciesType"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ContStats struct {
	TotalSquaresSum Number `xml:"totalSquaresSum,attr"`
	TotalValuesSum  Number `xml:"totalValuesSum,attr"`

	Extensions  []Extension `xml:"Extension"`
	Intervals   []Interval  `xml:"Interval"`
	Frequencies []Array     `xml:"Array"`
}

/*
  <xs:element name="ContinuousConditionalProbability">
    <xs:complexType>
      <xs:attribute name="count" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="ContinuousDistribution"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ParentValue"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ContinuousConditionalProbability struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ContinuousDistribution">
    <xs:complexType>
      <xs:sequence>
        <xs:choice>
          <xs:element ref="LognormalDistributionForBN"/>
          <xs:element ref="NormalDistributionForBN"/>
          <xs:element ref="TriangularDistributionForBN"/>
          <xs:element ref="UniformDistributionForBN"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ContinuousDistribution struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ContinuousNode">
    <xs:complexType>
      <xs:attribute name="count" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:choice maxOccurs="unbounded">
          <xs:element ref="ContinuousConditionalProbability"/>
          <xs:element maxOccurs="unbounded" ref="ContinuousDistribution"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ContinuousNode struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CorrelationFields">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="STRING-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CorrelationFields struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CorrelationMethods">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CorrelationMethods struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CorrelationValues">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CorrelationValues struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Correlations">
    <xs:complexType>
      <xs:sequence>
        <xs:element ref="CorrelationFields"/>
        <xs:element minOccurs="0" ref="CorrelationMethods"/>
        <xs:element ref="CorrelationValues"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Correlations struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CountTable" type="COUNT-TABLE-TYPE"/>
*/
type CountTable CountTableType

/*
  <xs:element name="Counts">
    <xs:complexType>
      <xs:attribute name="cardinality" type="xs:nonNegativeInteger"/>
      <xs:attribute name="invalidFreq" type="NUMBER"/>
      <xs:attribute name="missingFreq" type="NUMBER"/>
      <xs:attribute name="totalFreq" type="NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Counts struct {
	Cardinality uint   `xml:"cardinality,attr"`
	InvalidFreq Number `xml:"invalidFreq,attr"`
	MissingFreq Number `xml:"missingFreq,attr"`
	TotalFreq   Number `xml:"totalFreq,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Covariances">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Covariances struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="CovariateList">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Predictor"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CovariateList struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DataDictionary">
    <xs:complexType>
      <xs:attribute name="numberOfFields" type="xs:nonNegativeInteger"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="DataField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Taxonomy"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DataDictionary struct {
	NumberOfFields uint `xml:"numberOfFields,attr"`

	DataFields []DataField `xml:"DataField"`
	Extensions []Extension `xml:"Extension"`
	Taxonomies []Taxonomy  `xml:"Taxonomy"`
}

/*
  <xs:element name="DataField">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE" use="required"/>
      <xs:attribute name="displayName" type="xs:string"/>
      <xs:attribute default="0" name="isCyclic">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="0"/>
            <xs:enumeration value="1"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:attribute name="optype" type="OPTYPE" use="required"/>
      <xs:attribute name="taxonomy" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:sequence>
          <xs:element maxOccurs="unbounded" minOccurs="0" ref="Interval"/>
          <xs:element maxOccurs="unbounded" minOccurs="0" ref="Value"/>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DataField struct {
	DataType    DataType  `xml:"dataType,attr"`
	DisplayName string    `xml:"displayName,attr"`
	IsCyclic    bool      `xml:"isCyclic,attr"`
	Name        FieldName `xml:"name,attr"`
	OpType      OpType    `xml:"optype,attr"`
	Taxonomy    string    `xml:"taxonomy,attr"`

	Extensions []Extension `xml:"Extension"`
	Intervals  []Interval  `xml:"Interval"`
	Values     []Value     `xml:"Value"`
}

/*
  <xs:element name="Decision">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string"/>
      <xs:attribute name="displayValue" type="xs:string"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Decision struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DecisionTree">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="1.0" name="missingValuePenalty" type="PROB-NUMBER"/>
      <xs:attribute default="none" name="missingValueStrategy" type="MISSING-VALUE-STRATEGY"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute default="returnNullPrediction" name="noTrueChildStrategy" type="NO-TRUE-CHILD-STRATEGY"/>
      <xs:attribute default="multiSplit" name="splitCharacteristic">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="binarySplit"/>
            <xs:enumeration value="multiSplit"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element ref="Node"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ResultField"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DecisionTree struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Decisions">
    <xs:complexType>
      <xs:attribute name="businessProblem" type="xs:string"/>
      <xs:attribute name="description" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="Decision"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Decisions struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DefineFunction">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE"/>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:attribute name="optype" type="OPTYPE" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="ParameterField"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DefineFunction struct {
	DataType DataType `xml:"dataType,attr"`
	Name     string   `xml:"name,attr"`
	OpType   OpType   `xml:"optype,attr"`

	Extensions      []Extension      `xml:"Extension"`
	ParameterFields []ParameterField `xml:"ParameterField"`
	Expression      Expression
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (x *DefineFunction) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "dataType":
			x.DataType = DataType(attr.Value)
		case "name":
			x.Name = attr.Value
		case "optype":
			x.OpType = OpType(attr.Value)
		}
	}

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var item Expression

			switch tok.Name.Local {
			case "Extension":
				var e Extension
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.Extensions = append(x.Extensions, e)
			case "ParameterField":
				var e ParameterField
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.ParameterFields = append(x.ParameterFields, e)
			case "Aggregate":
				item = &Aggregate{}
			case "Apply":
				item = &Apply{}
			case "Constant":
				item = &Constant{}
			case "Discretize":
				item = &Discretize{}
			case "FieldRef":
				item = &FieldRef{}
			case "Lag":
				item = &Lag{}
			case "MapValues":
				item = &MapValues{}
			case "NormContinuous":
				item = &NormContinuous{}
			case "NormDiscrete":
				item = &NormDiscrete{}
			case "TextIndex":
				item = &TextIndex{}
			}

			if item != nil {
				err = decoder.DecodeElement(item, &tok)
				if err != nil {
					return err
				}

				x.Expression = item
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="Delimiter">
    <xs:complexType>
      <xs:attribute name="delimiter" type="DELIMITER" use="required"/>
      <xs:attribute name="gap" type="GAP" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Delimiter struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DerivedField">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE" use="required"/>
      <xs:attribute name="displayName" type="xs:string"/>
      <xs:attribute name="name" type="FIELD-NAME"/>
      <xs:attribute name="optype" type="OPTYPE" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Value"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DerivedField struct {
	DataType    DataType `xml:"dataType,attr"`
	DisplayName string   `xml:"displayName,attr"`
	Name        string   `xml:"name,attr"`
	OpType      OpType   `xml:"optype,attr"`

	Extensions []Extension `xml:"Extension"`
	Values     []Value     `xml:"Value"`
	Expression Expression
}

// UnmarshalXML implements the xml.Unmarshaler interface.
func (x *DerivedField) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "dataType":
			x.DataType = DataType(attr.Value)
		case "displayName":
			x.DisplayName = attr.Value
		case "name":
			x.Name = attr.Value
		case "optype":
			x.OpType = OpType(attr.Value)
		}
	}

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var item Expression

			switch tok.Name.Local {
			case "Extension":
				var e Extension
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.Extensions = append(x.Extensions, e)
			case "Value":
				var e Value
				err = decoder.DecodeElement(&e, &tok)
				if err != nil {
					return err
				}
				x.Values = append(x.Values, e)
			case "Aggregate":
				item = &Aggregate{}
			case "Apply":
				item = &Apply{}
			case "Constant":
				item = &Constant{}
			case "Discretize":
				item = &Discretize{}
			case "FieldRef":
				item = &FieldRef{}
			case "Lag":
				item = &Lag{}
			case "MapValues":
				item = &MapValues{}
			case "NormContinuous":
				item = &NormContinuous{}
			case "NormDiscrete":
				item = &NormDiscrete{}
			case "TextIndex":
				item = &TextIndex{}
			}

			if item != nil {
				err = decoder.DecodeElement(item, &tok)
				if err != nil {
					return err
				}

				x.Expression = item
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="DiscrStats">
    <xs:complexType>
      <xs:attribute name="modalValue" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="2" minOccurs="0" ref="Array"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DiscrStats struct {
	ModalValue string `xml:"modalValue,attr"`

	Arrays     []Array     `xml:"Array"`
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DiscreteConditionalProbability">
    <xs:complexType>
      <xs:attribute name="count" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="ParentValue"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="ValueProbability"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DiscreteConditionalProbability struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DiscreteNode">
    <xs:complexType>
      <xs:attribute name="count" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:choice maxOccurs="unbounded">
          <xs:element ref="DiscreteConditionalProbability"/>
          <xs:element maxOccurs="unbounded" ref="ValueProbability"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DiscreteNode struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Discretize">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE"/>
      <xs:attribute name="defaultValue" type="xs:string"/>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="mapMissingTo" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DiscretizeBin"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Discretize struct {
	Extensions []Extension `xml:"Extension"`
}

func (*Discretize) expression() {}

/*
  <xs:element name="DiscretizeBin">
    <xs:complexType>
      <xs:attribute name="binValue" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Interval"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DiscretizeBin struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="DocumentTermMatrix">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type DocumentTermMatrix struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="EventValues">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Interval"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Value"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type EventValues struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ExponentialSmoothing">
    <xs:complexType>
      <xs:attribute name="RMSE" type="REAL-NUMBER"/>
      <xs:attribute default="none" name="transformation">
        <xs:simpleType>
          <xs:restriction base="xs:NMTOKEN">
            <xs:enumeration value="logarithmic"/>
            <xs:enumeration value="none"/>
            <xs:enumeration value="squareroot"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="1" minOccurs="1" ref="Level"/>
        <xs:element maxOccurs="1" minOccurs="0" ref="Seasonality_ExpoSmooth"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TimeValue"/>
        <xs:element maxOccurs="1" minOccurs="0" ref="Trend_ExpoSmooth"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ExponentialSmoothing struct {
}

/*
  <xs:element name="Extension">
    <xs:complexType>
      <xs:complexContent mixed="true">
        <xs:restriction base="xs:anyType">
          <xs:attribute name="extender" type="xs:string" use="optional"/>
          <xs:attribute name="name" type="xs:string" use="optional"/>
          <xs:attribute name="value" type="xs:string" use="optional"/>
          <xs:sequence>
            <xs:any maxOccurs="unbounded" minOccurs="0" processContents="skip"/>
          </xs:sequence>
        </xs:restriction>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
*/
type Extension struct {
	Extender *string `xml:"extender,attr"`
	Name     *string `xml:"name,attr"`
	Value    *string `xml:"value,attr"`
}

/*
  <xs:element name="FactorList">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Predictor"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type FactorList struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="False">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type False struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="FieldColumnPair">
    <xs:complexType>
      <xs:attribute name="column" type="xs:string" use="required"/>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type FieldColumnPair struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="FieldRef">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="mapMissingTo" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type FieldRef struct {
	Extensions []Extension `xml:"Extension"`
}

func (*FieldRef) expression() {}

/*
  <xs:element name="FieldValue">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="value" use="required"/>
      <xs:sequence>
        <xs:choice>
          <xs:element maxOccurs="unbounded" minOccurs="1" ref="FieldValue"/>
          <xs:element maxOccurs="unbounded" minOccurs="1" ref="FieldValueCount"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type FieldValue struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="FieldValueCount">
    <xs:complexType>
      <xs:attribute name="count" type="NUMBER" use="required"/>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="value" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type FieldValueCount struct {
	Count Number    `xml:"count,attr"`
	Field FieldName `xml:"field,attr"`
	Value string    `xml:"value,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="GaussianDistribution">
    <xs:complexType>
      <xs:attribute name="mean" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="variance" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type GaussianDistribution struct {
	Extensions []Extension `xml:"Extension"`
}

func (*GaussianDistribution) continuousDistributionType() {}

/*
  <xs:element name="GaussianProcessModel">
    <xs:complexType>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:attribute name="optimizer" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
        <xs:element ref="TrainingInstances"/>
        <xs:sequence>
          <xs:choice>
            <xs:element ref="ARDSquaredExponentialKernel"/>
            <xs:element ref="AbsoluteExponentialKernel"/>
            <xs:element ref="GeneralizedExponentialKernel"/>
            <xs:element ref="RadialBasisKernel"/>
          </xs:choice>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type GaussianProcessModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*GaussianProcessModel) modelElement() {}

/*
  <xs:element name="GeneralRegressionModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="baselineStrataVariable" type="FIELD-NAME"/>
      <xs:attribute name="cumulativeLink" type="CUMULATIVE-LINK-FUNCTION"/>
      <xs:attribute name="distParameter" type="REAL-NUMBER"/>
      <xs:attribute name="distribution">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="binomial"/>
            <xs:enumeration value="gamma"/>
            <xs:enumeration value="igauss"/>
            <xs:enumeration value="negbin"/>
            <xs:enumeration value="normal"/>
            <xs:enumeration value="poisson"/>
            <xs:enumeration value="tweedie"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="endTimeVariable" type="FIELD-NAME"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="linkFunction" type="LINK-FUNCTION"/>
      <xs:attribute name="linkParameter" type="REAL-NUMBER"/>
      <xs:attribute name="modelDF" type="REAL-NUMBER"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="modelType" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="CoxRegression"/>
            <xs:enumeration value="generalLinear"/>
            <xs:enumeration value="generalizedLinear"/>
            <xs:enumeration value="multinomialLogistic"/>
            <xs:enumeration value="ordinalMultinomial"/>
            <xs:enumeration value="regression"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="offsetValue" type="REAL-NUMBER"/>
      <xs:attribute name="offsetVariable" type="FIELD-NAME"/>
      <xs:attribute name="startTimeVariable" type="FIELD-NAME"/>
      <xs:attribute name="statusVariable" type="FIELD-NAME"/>
      <xs:attribute name="subjectIDVariable" type="FIELD-NAME"/>
      <xs:attribute name="targetReferenceCategory" type="xs:string"/>
      <xs:attribute name="targetVariableName" type="FIELD-NAME"/>
      <xs:attribute name="trialsValue" type="INT-NUMBER"/>
      <xs:attribute name="trialsVariable" type="FIELD-NAME"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="BaseCumHazardTables"/>
        <xs:element minOccurs="0" ref="CovariateList"/>
        <xs:element minOccurs="0" ref="EventValues"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="FactorList"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="PCovMatrix"/>
        <xs:element ref="PPMatrix"/>
        <xs:element ref="ParamMatrix"/>
        <xs:element ref="ParameterList"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type GeneralRegressionModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*GeneralRegressionModel) modelElement() {}

/*
  <xs:element name="GeneralizedExponentialKernel">
    <xs:complexType>
      <xs:attribute default="1" name="degree" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="noiseVariance" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Lambda"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type GeneralizedExponentialKernel struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Header">
    <xs:complexType>
      <xs:attribute name="copyright" type="xs:string"/>
      <xs:attribute name="description" type="xs:string"/>
      <xs:attribute name="modelVersion" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Annotation"/>
        <xs:element minOccurs="0" ref="Application"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="Timestamp"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Header struct {
	Copyright    string `xml:"copyright,attr"`
	Description  string `xml:"description,attr"`
	ModelVersion string `xml:"modelVersion,attr"`

	Annotations []Annotation `xml:"Annotation"`
	Application *Application `xml:"Application"`
	Extensions  []Extension  `xml:"Extension"`
	Timestamp   *Timestamp   `xml:"Timestamp"`
}

/*
  <xs:element name="INT-Entries">
    <xs:simpleType>
      <xs:list itemType="xs:int"/>
    </xs:simpleType>
  </xs:element>
*/
type IntegerEntries struct {
}

/*
  <xs:element name="INT-SparseArray">
    <xs:complexType>
      <xs:attribute default="0" name="defaultValue" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="n" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="INT-Entries"/>
        <xs:element minOccurs="0" ref="Indices"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type IntegerSparseArray struct {
}

/*
  <xs:element name="Indices">
    <xs:simpleType>
      <xs:list itemType="xs:int"/>
    </xs:simpleType>
  </xs:element>
*/
type Indices struct {
}

/*
  <xs:element name="InlineTable">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="row"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type InlineTable struct {
	Extensions []Extension `xml:"Extension"`
	Rows       []Row       `xml:"row"`
}

func (*InlineTable) table() {}

/*
  <xs:element name="InstanceField">
    <xs:complexType>
      <xs:attribute name="column" type="xs:string" use="optional"/>
      <xs:attribute name="field" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type InstanceField struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="InstanceFields">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="InstanceField"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type InstanceFields struct {
	Extensions []Extension `xml:"Extension"`
}

type IntervalType string

const (
	IntervalTypeClosedClosed = IntervalType("closedClosed")
	IntervalTypeClosedOpen   = IntervalType("closedOpen")
	IntervalTypeOpenClosed   = IntervalType("openClosed")
	IntervalTypeOpenOpen     = IntervalType("openOpen")
)

/*
  <xs:element name="Interval">
    <xs:complexType>
      <xs:attribute name="closure" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="closedClosed"/>
            <xs:enumeration value="closedOpen"/>
            <xs:enumeration value="openClosed"/>
            <xs:enumeration value="openOpen"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="leftMargin" type="NUMBER"/>
      <xs:attribute name="rightMargin" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Interval struct {
	Closure     IntervalType `xml:"closure,attr"`
	LeftMargin  Number       `xml:"leftMargin,attr"`
	RightMargin Number       `xml:"rightMargin,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Item">
    <xs:complexType>
      <xs:attribute name="category" type="xs:string"/>
      <xs:attribute name="field" type="FIELD-NAME"/>
      <xs:attribute name="id" type="xs:string" use="required"/>
      <xs:attribute name="mappedValue" type="xs:string"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:attribute name="weight" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Item struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ItemRef">
    <xs:complexType>
      <xs:attribute name="itemRef" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ItemRef struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Itemset">
    <xs:complexType>
      <xs:attribute name="id" type="xs:string" use="required"/>
      <xs:attribute name="numberOfItems" type="xs:nonNegativeInteger"/>
      <xs:attribute name="support" type="PROB-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ItemRef"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Itemset struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="KNNInput">
    <xs:complexType>
      <xs:attribute name="compareFunction" type="COMPARE-FUNCTION"/>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute default="1" name="fieldWeight" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type KNNInput struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="KNNInputs">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="KNNInput"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type KNNInputs struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="KohonenMap">
    <xs:complexType>
      <xs:attribute name="coord1" type="xs:float" use="optional"/>
      <xs:attribute name="coord2" type="xs:float" use="optional"/>
      <xs:attribute name="coord3" type="xs:float" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type KohonenMap struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Lag">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute default="1" name="n" type="xs:positiveInteger"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="BlockIndicator"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Lag struct {
	Extensions []Extension `xml:"Extension"`
}

func (*Lag) expression() {}

/*
  <xs:element name="Lambda">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="REAL-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Lambda struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Level">
    <xs:complexType>
      <xs:attribute name="alpha" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="smoothedValue" type="REAL-NUMBER"/>
    </xs:complexType>
  </xs:element>
*/
type Level struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="LiftData">
    <xs:complexType>
      <xs:attribute name="rankingQuality" type="NUMBER"/>
      <xs:attribute name="targetFieldDisplayValue" type="xs:string"/>
      <xs:attribute name="targetFieldValue" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="ModelLiftGraph"/>
        <xs:element minOccurs="0" ref="OptimumLiftGraph"/>
        <xs:element minOccurs="0" ref="RandomLiftGraph"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LiftData struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="LiftGraph">
    <xs:complexType>
      <xs:sequence>
        <xs:element minOccurs="0" ref="BoundaryValueMeans"/>
        <xs:element minOccurs="0" ref="BoundaryValues"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="XCoordinates"/>
        <xs:element ref="YCoordinates"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LiftGraph struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="LinearKernelType">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LinearKernelType struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="LinearNorm">
    <xs:complexType>
      <xs:attribute name="norm" type="NUMBER" use="required"/>
      <xs:attribute name="orig" type="NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LinearNorm struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="LocalTransformations">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LocalTransformations struct {
	DerivedFields []DerivedField `xml:"DerivedField"`
	Extensions    []Extension    `xml:"Extension"`
}

/*
  <xs:element name="LognormalDistributionForBN">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Mean"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Variance"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type LognormalDistributionForBN struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Lower">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Lower struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="MapValues">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE"/>
      <xs:attribute name="defaultValue" type="xs:string"/>
      <xs:attribute name="mapMissingTo" type="xs:string"/>
      <xs:attribute name="outputColumn" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:choice minOccurs="0">
          <xs:element ref="InlineTable"/>
          <xs:element ref="TableLocator"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="FieldColumnPair"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MapValues struct {
	Extensions []Extension `xml:"Extension"`
}

func (*MapValues) expression() {}

/*
  <xs:element name="MatCell">
    <xs:complexType>
      <xs:simpleContent>
        <xs:extension base="xs:string">
          <xs:attribute name="col" type="INT-NUMBER" use="required"/>
          <xs:attribute name="row" type="INT-NUMBER" use="required"/>
        </xs:extension>
      </xs:simpleContent>
    </xs:complexType>
  </xs:element>
*/
type MatCell struct {
}

/*
  <xs:element name="Matrix">
    <xs:complexType>
      <xs:attribute name="diagDefault" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="any" name="kind" use="optional">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="any"/>
            <xs:enumeration value="diagonal"/>
            <xs:enumeration value="symmetric"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="nbCols" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="nbRows" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="offDiagDefault" type="REAL-NUMBER" use="optional"/>
      <xs:choice minOccurs="0">
        <xs:element maxOccurs="unbounded" ref="MatCell"/>
        <xs:group maxOccurs="unbounded" ref="NUM-ARRAY"/>
      </xs:choice>
    </xs:complexType>
  </xs:element>
*/
type Matrix struct {
}

/*
  <xs:element name="Mean">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Mean struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="MiningBuildTask">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MiningBuildTask struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="MiningField">
    <xs:complexType>
      <xs:attribute name="highValue" type="NUMBER"/>
      <xs:attribute name="importance" type="PROB-NUMBER"/>
      <xs:attribute default="returnInvalid" name="invalidValueTreatment" type="INVALID-VALUE-TREATMENT-METHOD"/>
      <xs:attribute name="lowValue" type="NUMBER"/>
      <xs:attribute name="missingValueReplacement" type="xs:string"/>
      <xs:attribute name="missingValueTreatment" type="MISSING-VALUE-TREATMENT-METHOD"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:attribute name="optype" type="OPTYPE"/>
      <xs:attribute default="asIs" name="outliers" type="OUTLIER-TREATMENT-METHOD"/>
      <xs:attribute default="active" name="usageType" type="FIELD-USAGE-TYPE"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MiningField struct {
	HighValue               Number                      `xml:"highValue,attr"`
	Importance              ProbNumber                  `xml:"importance,attr"`
	InvalidValueTreatment   InvalidValueTreatmentMethod `xml:"invalidValueTreatment,attr"`
	LowValue                Number                      `xml:"lowValue,attr"`
	MissingValueReplacement string                      `xml:"missingValueReplacement"`
	MissingValueTreatment   MissingValueTreatmentMethod `xml:"missingValueTreatment,attr"`
	Name                    FieldName                   `xml:"name,attr"`
	OpType                  OpType                      `xml:"optype,attr"`
	Outliers                OutlierTreatmentMethod      `xml:"outliers,attr"`
	UsageType               FieldUsageType              `xml:"usageType,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="MiningModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:choice maxOccurs="unbounded" minOccurs="0">
          <xs:element ref="DecisionTree"/>
          <xs:element ref="Regression"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Segmentation"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MiningModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*MiningModel) modelElement() {}

/*
  <xs:element name="MiningSchema">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="MiningField"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MiningSchema struct {
	Extensions   []Extension   `xml:"Extension"`
	MiningFields []MiningField `xml:"MiningField"`
}

/*
  <xs:element name="MissingValueWeights">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MissingValueWeights struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ModelExplanation">
    <xs:complexType>
      <xs:sequence>
        <xs:choice>
          <xs:element maxOccurs="unbounded" minOccurs="0" ref="ClusteringModelQuality"/>
          <xs:element maxOccurs="unbounded" minOccurs="0" ref="PredictiveModelQuality"/>
        </xs:choice>
        <xs:element minOccurs="0" ref="Correlations"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ModelExplanation struct {
	// TODO: ModelQuality
	Correlations *Correlations `xml:"Correlations"`
	Extensions   []Extension   `xml:"Extension"`
}

/*
  <xs:element name="ModelLiftGraph">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="LiftGraph"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ModelLiftGraph struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ModelStats">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="MultivariateStats"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="UnivariateStats"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ModelStats struct {
	Extensions        []Extension         `xml:"Extension"`
	MultivariateStats []MultivariateStats `xml:"MultivariateStats"`
	UnivariateStats   []UnivariateStats   `xml:"UnivariateStats"`
}

/*
  <xs:element name="ModelVerification">
    <xs:complexType>
      <xs:attribute name="fieldCount" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="recordCount" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="InlineTable"/>
        <xs:element ref="VerificationFields"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ModelVerification struct {
	FieldCount  IntegerNumber `xml:"fieldCount,attr"`
	RecordCount IntegerNumber `xml:"recordCount,attr"`

	Extensions         []Extension        `xml:"Extension"`
	InlineTable        InlineTable        `xml:"InlineTable"`
	VerificationFields VerificationFields `xml:"VerificationFields"`
}

/*
  <xs:element name="MultivariateStat">
    <xs:complexType>
      <xs:attribute name="category" type="xs:string"/>
      <xs:attribute name="chiSquareValue" type="NUMBER"/>
      <xs:attribute default="0.95" name="confidenceLevel" type="PROB-NUMBER"/>
      <xs:attribute name="confidenceLowerBound" type="NUMBER"/>
      <xs:attribute name="confidenceUpperBound" type="NUMBER"/>
      <xs:attribute name="dF" type="NUMBER"/>
      <xs:attribute default="1" name="exponent" type="INT-NUMBER"/>
      <xs:attribute name="fStatistic" type="NUMBER"/>
      <xs:attribute name="importance" type="PROB-NUMBER"/>
      <xs:attribute default="false" name="isIntercept" type="xs:boolean"/>
      <xs:attribute name="name" type="xs:string"/>
      <xs:attribute name="pValueAlpha" type="PROB-NUMBER"/>
      <xs:attribute name="pValueFinal" type="PROB-NUMBER"/>
      <xs:attribute name="pValueInitial" type="PROB-NUMBER"/>
      <xs:attribute name="stdError" type="NUMBER"/>
      <xs:attribute name="tValue" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MultivariateStat struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="MultivariateStats">
    <xs:complexType>
      <xs:attribute name="targetCategory" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="MultivariateStat"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type MultivariateStats struct {
	TargetCategory string `xml:"targetCategory,attr"`

	MultivariateStats []MultivariateStat `xml:"MultivariateStat"`
	Extensions        []Extension        `xml:"Extension"`
}

/*
  <xs:element name="NaiveBayesModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="threshold" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element ref="BayesInputs"/>
        <xs:element ref="BayesOutput"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NaiveBayesModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*NaiveBayesModel) modelElement() {}

/*
  <xs:element name="NearestNeighborModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute default="majorityVote" name="categoricalScoringMethod" type="CAT-SCORING-METHOD"/>
      <xs:attribute default="average" name="continuousScoringMethod" type="CONT-SCORING-METHOD"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute name="instanceIdVariable" type="xs:string"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="numberOfNeighbors" type="INT-NUMBER" use="required"/>
      <xs:attribute default="0.001" name="threshold" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element ref="ComparisonMeasure"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="KNNInputs"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
        <xs:element ref="TrainingInstances"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NearestNeighborModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*NearestNeighborModel) modelElement() {}

/*
  <xs:element name="NeuralInput">
    <xs:complexType>
      <xs:attribute name="id" type="NN-NEURON-ID" use="required"/>
      <xs:sequence>
        <xs:element ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralInput struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="NeuralInputs">
    <xs:complexType>
      <xs:attribute name="numberOfInputs" type="xs:nonNegativeInteger"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="NeuralInput"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralInputs struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="NeuralLayer">
    <xs:complexType>
      <xs:attribute name="activationFunction" type="ACTIVATION-FUNCTION"/>
      <xs:attribute name="altitude" type="REAL-NUMBER"/>
      <xs:attribute name="normalizationMethod" type="NN-NORMALIZATION-METHOD"/>
      <xs:attribute name="numberOfNeurons" type="xs:nonNegativeInteger"/>
      <xs:attribute name="threshold" type="REAL-NUMBER"/>
      <xs:attribute name="width" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="Neuron"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralLayer struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="NeuralNetwork">
    <xs:complexType>
      <xs:attribute name="activationFunction" type="ACTIVATION-FUNCTION" use="required"/>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute default="1.0" name="altitude" type="REAL-NUMBER"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute default="none" name="normalizationMethod" type="NN-NORMALIZATION-METHOD"/>
      <xs:attribute name="numberOfLayers" type="xs:nonNegativeInteger"/>
      <xs:attribute default="0" name="threshold" type="REAL-NUMBER"/>
      <xs:attribute name="width" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element ref="NeuralInputs"/>
        <xs:element maxOccurs="unbounded" ref="NeuralLayer"/>
        <xs:element minOccurs="0" ref="NeuralOutputs"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralNetwork struct {
	Extensions []Extension `xml:"Extension"`
}

func (*NeuralNetwork) modelElement() {}

/*
  <xs:element name="NeuralOutput">
    <xs:complexType>
      <xs:attribute name="outputNeuron" type="NN-NEURON-IDREF" use="required"/>
      <xs:sequence>
        <xs:element ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralOutput struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="NeuralOutputs">
    <xs:complexType>
      <xs:attribute name="numberOfOutputs" type="xs:nonNegativeInteger"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="NeuralOutput"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NeuralOutputs struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Neuron">
    <xs:complexType>
      <xs:attribute name="altitude" type="REAL-NUMBER"/>
      <xs:attribute name="bias" type="REAL-NUMBER"/>
      <xs:attribute name="id" type="NN-NEURON-ID" use="required"/>
      <xs:attribute name="width" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="Con"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Neuron struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Node">
    <xs:complexType>
      <xs:attribute name="defaultChild" type="xs:string"/>
      <xs:attribute name="id" type="xs:string"/>
      <xs:attribute name="recordCount" type="NUMBER"/>
      <xs:attribute name="score" type="xs:string"/>
      <xs:sequence>
        <xs:choice>
          <xs:group ref="EmbeddedModel"/>
          <xs:sequence>
            <xs:element maxOccurs="unbounded" minOccurs="0" ref="Node"/>
            <xs:element minOccurs="0" ref="Partition"/>
            <xs:element maxOccurs="unbounded" minOccurs="0" ref="ScoreDistribution"/>
          </xs:sequence>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="PREDICATE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Node struct {
	DefaultChild string `xml:"defaultChild,attr"`
	ID           string `xml:"id,attr"`
	RecordCount  Number `xml:"recordCount,attr"`
	Score        string `xml:"score,attr"`

	EmbeddedModel      EmbeddedModel       // TODO: EmbeddedModel
	Nodes              []Node              `xml:"Node"`
	Partition          *Partition          `xml:"Partition"`
	ScoreDistributions []ScoreDistribution `xml:"ScoreDistribution"`
	Extensions         []Extension         `xml:"Extension"`
	Predicate          Predicate           // TODO: Predicate
}

/*
  <xs:element name="NormContinuous">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="mapMissingTo" type="NUMBER"/>
      <xs:attribute default="asIs" name="outliers" type="OUTLIER-TREATMENT-METHOD"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="2" ref="LinearNorm"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NormContinuous struct {
	Extensions []Extension `xml:"Extension"`
}

func (*NormContinuous) expression() {}

/*
  <xs:element name="NormDiscrete">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="mapMissingTo" type="NUMBER"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NormDiscrete struct {
	Extensions []Extension `xml:"Extension"`
}

func (*NormDiscrete) expression() {}

/*
  <xs:element name="NormalDistributionForBN">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Mean"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Variance"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NormalDistributionForBN struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="NormalizedCountTable" type="COUNT-TABLE-TYPE"/>
*/
type NormalizedCountTable CountTableType

/*
  <xs:element name="NumericInfo">
    <xs:complexType>
      <xs:attribute name="interQuartileRange" type="NUMBER"/>
      <xs:attribute name="maximum" type="NUMBER"/>
      <xs:attribute name="mean" type="NUMBER"/>
      <xs:attribute name="median" type="NUMBER"/>
      <xs:attribute name="minimum" type="NUMBER"/>
      <xs:attribute name="standardDeviation" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Quantile"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NumericInfo struct {
	InterQuartileRange Number `xml:"interQuartileRange,attr"`
	Maximum            Number `xml:"maximum,attr"`
	Mean               Number `xml:"mean,attr"`
	Median             Number `xml:"median,attr"`
	Minimum            Number `xml:"minimum,attr"`
	StandardDeviation  Number `xml:"standardDeviation"`

	Extensions []Extension `xml:"Extension"`
	Quantiles  []Quantile  `xml:"Quantile"`
}

/*
  <xs:element name="NumericPredictor">
    <xs:complexType>
      <xs:attribute name="coefficient" type="REAL-NUMBER" use="required"/>
      <xs:attribute default="1" name="exponent" type="INT-NUMBER"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type NumericPredictor struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="OptimumLiftGraph">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="LiftGraph"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type OptimumLiftGraph struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Output">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="OutputField"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Output struct {
	Extensions []Extension   `xml:"Extension"`
	Fields     []OutputField `xml:"OutputField"`
}

/*
  <xs:element name="OutputField">
    <xs:complexType>
      <xs:attribute default="exclusiveRecommendation" name="algorithm">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="exclusiveRecommendation"/>
            <xs:enumeration value="recommendation"/>
            <xs:enumeration value="ruleAssociation"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="dataType" type="DATATYPE" use="required"/>
      <xs:attribute name="displayName" type="xs:string"/>
      <xs:attribute default="predictedValue" name="feature" type="RESULT-FEATURE"/>
      <xs:attribute default="0" name="isMultiValued"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:attribute name="optype" type="OPTYPE"/>
      <xs:attribute default="1" name="rank" type="INT-NUMBER"/>
      <xs:attribute default="confidence" name="rankBasis">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="affinity"/>
            <xs:enumeration value="confidence"/>
            <xs:enumeration value="leverage"/>
            <xs:enumeration value="lift"/>
            <xs:enumeration value="support"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="descending" name="rankOrder">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="ascending"/>
            <xs:enumeration value="descending"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="consequent" name="ruleFeature" type="RULE-FEATURE"/>
      <xs:attribute name="segmentId" type="xs:string"/>
      <xs:attribute name="targetField" type="FIELD-NAME"/>
      <xs:attribute name="value" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:sequence maxOccurs="1" minOccurs="0">
          <xs:element maxOccurs="1" minOccurs="0" ref="Decisions"/>
          <xs:group maxOccurs="1" minOccurs="1" ref="EXPRESSION"/>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type OutputField struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PCell">
    <xs:complexType>
      <xs:attribute name="beta" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="df" type="INT-NUMBER"/>
      <xs:attribute name="parameterName" type="xs:string" use="required"/>
      <xs:attribute name="targetCategory" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PCell struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PCovCell">
    <xs:complexType>
      <xs:attribute name="pCol" type="xs:string" use="required"/>
      <xs:attribute name="pRow" type="xs:string" use="required"/>
      <xs:attribute name="tCol" type="xs:string"/>
      <xs:attribute name="tRow" type="xs:string"/>
      <xs:attribute name="targetCategory" type="xs:string"/>
      <xs:attribute name="value" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PCovCell struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PCovMatrix">
    <xs:complexType>
      <xs:attribute name="type">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="model"/>
            <xs:enumeration value="robust"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="PCovCell"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PCovMatrix struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PMML">
    <xs:complexType>
      <xs:attribute name="version" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element ref="DataDictionary"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="Header"/>
        <xs:element minOccurs="0" ref="MiningBuildTask"/>
        <xs:element minOccurs="0" ref="TransformationDictionary"/>
        <xs:sequence maxOccurs="unbounded" minOccurs="0">
          <xs:group ref="MODEL-ELEMENT"/>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PMML struct {
	Version string `xml:"version,attr"`

	DataDictionary           DataDictionary            `xml:"DataDictionary"`
	Extensions               []Extension               `xml:"Extension"`
	Header                   Header                    `xml:"Header"`
	MiningBuildTask          *MiningBuildTask          `xml:"MiningBuildTask"`
	TransformationDictionary *TransformationDictionary `xml:"TransformationDictionary"`
	Models                   []ModelElement
}

func (x *PMML) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "version":
			x.Version = attr.Value
		}
	}

	for {
		var token xml.Token

		token, err = decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var model ModelElement

			switch tok.Name.Local {
			case "Header":
				err = decoder.DecodeElement(&x.Header, &tok)
				if err != nil {
					return err
				}
			case "DataDictionary":
				err = decoder.DecodeElement(&x.DataDictionary, &tok)
				if err != nil {
					return err
				}
			case "TransformationDictionary":
				err = decoder.DecodeElement(&x.TransformationDictionary, &tok)
				if err != nil {
					return err
				}
			case "AssociationModel":
				model = &AssociationModel{}
			case "BaselineModel":
				model = &BaselineModel{}
			case "BayesianNetworkModel":
				model = &BayesianNetworkModel{}
			case "ClusteringModel":
				model = &ClusteringModel{}
			case "GaussianProcessModel":
				model = &GaussianProcessModel{}
			case "GeneralRegressionModel":
				model = &GeneralRegressionModel{}
			case "MiningModel":
				model = &MiningModel{}
			case "NaiveBayesModel":
				model = &NaiveBayesModel{}
			case "NearestNeighborModel":
				model = &NearestNeighborModel{}
			case "NeuralNetwork":
				model = &NeuralNetwork{}
			case "RegressionModel":
				model = &RegressionModel{}
			case "RuleSetModel":
				model = &RuleSetModel{}
			case "Scorecard":
				model = &Scorecard{}
			case "SequenceModel":
				model = &SequenceModel{}
			case "SupportVectorMachineModel":
				model = &SupportVectorMachineModel{}
			case "TextModel":
				model = &TextModel{}
			case "TimeSeriesModel":
				model = &TimeSeriesModel{}
			case "TreeModel":
				model = &TreeModel{}
			}

			if model != nil {
				err = decoder.DecodeElement(model, &tok)
				if err != nil {
					return err
				}

				x.Models = append(x.Models, model)
			}
		case xml.EndElement:
			return nil
		}
	}
}

/*
  <xs:element name="PPCell">
    <xs:complexType>
      <xs:attribute name="parameterName" type="xs:string" use="required"/>
      <xs:attribute name="predictorName" type="FIELD-NAME" use="required"/>
      <xs:attribute name="targetCategory" type="xs:string"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PPCell struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PPMatrix">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="PPCell"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PPMatrix struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PairCounts">
    <xs:complexType>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="TargetValueCounts"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PairCounts struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ParamMatrix">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="PCell"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ParamMatrix struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Parameter">
    <xs:complexType>
      <xs:attribute name="label" type="xs:string"/>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:attribute default="0" name="referencePoint" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Parameter struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ParameterField">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE"/>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:attribute name="optype" type="OPTYPE"/>
    </xs:complexType>
  </xs:element>
*/
type ParameterField struct {
}

/*
  <xs:element name="ParameterList">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Parameter"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ParameterList struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ParentValue">
    <xs:complexType>
      <xs:attribute name="parent" type="FIELD-NAME" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ParentValue struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Partition">
    <xs:complexType>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:attribute name="size" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="PartitionFieldStats"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Partition struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PartitionFieldStats">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute default="0" name="weighted">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="0"/>
            <xs:enumeration value="1"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Counts"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="NumericInfo"/>
        <xs:group minOccurs="0" ref="FrequenciesType"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PartitionFieldStats struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PoissonDistribution">
    <xs:complexType>
      <xs:attribute name="mean" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PoissonDistribution struct {
	Extensions []Extension `xml:"Extension"`
}

func (*PoissonDistribution) continuousDistributionType() {}

/*
  <xs:element name="PolynomialKernelType">
    <xs:complexType>
      <xs:attribute default="1" name="coef0" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="degree" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PolynomialKernelType struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PredictiveModelQuality">
    <xs:complexType>
      <xs:attribute name="AIC" type="NUMBER" use="optional"/>
      <xs:attribute name="AICc" type="NUMBER" use="optional"/>
      <xs:attribute name="BIC" type="NUMBER" use="optional"/>
      <xs:attribute name="adj-r-squared" type="NUMBER" use="optional"/>
      <xs:attribute name="dataName" type="xs:string" use="optional"/>
      <xs:attribute default="training" name="dataUsage">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="test"/>
            <xs:enumeration value="training"/>
            <xs:enumeration value="validation"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="degreesOfFreedom" type="NUMBER" use="optional"/>
      <xs:attribute name="fStatistic" type="NUMBER" use="optional"/>
      <xs:attribute name="meanAbsoluteError" type="NUMBER" use="optional"/>
      <xs:attribute name="meanError" type="NUMBER" use="optional"/>
      <xs:attribute name="meanSquaredError" type="NUMBER" use="optional"/>
      <xs:attribute name="numOfPredictors" type="NUMBER" use="optional"/>
      <xs:attribute name="numOfRecords" type="NUMBER" use="optional"/>
      <xs:attribute name="numOfRecordsWeighted" type="NUMBER" use="optional"/>
      <xs:attribute name="r-squared" type="NUMBER" use="optional"/>
      <xs:attribute name="rootMeanSquaredError" type="NUMBER" use="optional"/>
      <xs:attribute name="sumSquaredError" type="NUMBER" use="optional"/>
      <xs:attribute name="sumSquaredRegression" type="NUMBER" use="optional"/>
      <xs:attribute name="targetField" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="ConfusionMatrix"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="LiftData"/>
        <xs:element minOccurs="0" ref="ROC"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PredictiveModelQuality struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Predictor">
    <xs:complexType>
      <xs:attribute name="contrastMatrixType" type="xs:string"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="1" minOccurs="0" ref="Categories"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="Matrix"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Predictor struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="PredictorTerm">
    <xs:complexType>
      <xs:attribute name="coefficient" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="name" type="FIELD-NAME"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="FieldRef"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type PredictorTerm struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Quantile">
    <xs:complexType>
      <xs:attribute name="quantileLimit" type="PERCENTAGE-NUMBER" use="required"/>
      <xs:attribute name="quantileValue" type="NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Quantile struct {
	QuantileLimit PercentageNumber `xml:"quantileLimit,attr"`
	QuantileValue Number           `xml:"quantileValue,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="REAL-Entries">
    <xs:simpleType>
      <xs:list itemType="xs:double"/>
    </xs:simpleType>
  </xs:element>
*/
type RealEntries struct {
}

/*
  <xs:element name="REAL-SparseArray">
    <xs:complexType>
      <xs:attribute default="0" name="defaultValue" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="n" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Indices"/>
        <xs:element minOccurs="0" ref="REAL-Entries"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RealSparseArray struct {
}

/*
  <xs:element name="ROC">
    <xs:complexType>
      <xs:attribute name="negativeTargetFieldDisplayValue" type="xs:string"/>
      <xs:attribute name="negativeTargetFieldValue" type="xs:string"/>
      <xs:attribute name="positiveTargetFieldDisplayValue" type="xs:string"/>
      <xs:attribute name="positiveTargetFieldValue" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="ROCGraph"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ROC struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ROCGraph">
    <xs:complexType>
      <xs:sequence>
        <xs:element minOccurs="0" ref="BoundaryValues"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="XCoordinates"/>
        <xs:element ref="YCoordinates"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ROCGraph struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RadialBasisKernel">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="lambda" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="noiseVariance" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RadialBasisKernel struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RadialBasisKernelType">
    <xs:complexType>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RadialBasisKernelType struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RandomLiftGraph">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="LiftGraph"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RandomLiftGraph struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Regression">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute default="none" name="normalizationMethod" type="REGRESSIONNORMALIZATIONMETHOD"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element maxOccurs="unbounded" ref="RegressionTable"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ResultField"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Regression struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RegressionModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="modelType" use="optional">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="linearRegression"/>
            <xs:enumeration value="logisticRegression"/>
            <xs:enumeration value="stepwisePolynomialRegression"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="none" name="normalizationMethod" type="REGRESSIONNORMALIZATIONMETHOD"/>
      <xs:attribute name="targetFieldName" type="FIELD-NAME" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element maxOccurs="unbounded" ref="RegressionTable"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RegressionModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*RegressionModel) modelElement() {}

/*
  <xs:element name="RegressionTable">
    <xs:complexType>
      <xs:attribute name="intercept" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="targetCategory" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="CategoricalPredictor"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="NumericPredictor"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="PredictorTerm"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RegressionTable struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ResultField">
    <xs:complexType>
      <xs:attribute name="dataType" type="DATATYPE"/>
      <xs:attribute name="displayName" type="xs:string"/>
      <xs:attribute name="feature" type="RESULT-FEATURE"/>
      <xs:attribute name="name" type="FIELD-NAME" use="required"/>
      <xs:attribute name="optype" type="OPTYPE"/>
      <xs:attribute name="value" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ResultField struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RuleSelectionMethod">
    <xs:complexType>
      <xs:attribute name="criterion" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="firstHit"/>
            <xs:enumeration value="weightedMax"/>
            <xs:enumeration value="weightedSum"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RuleSelectionMethod struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RuleSet">
    <xs:complexType>
      <xs:attribute name="defaultConfidence" type="NUMBER" use="optional"/>
      <xs:attribute name="defaultScore" type="xs:string" use="optional"/>
      <xs:attribute name="nbCorrect" type="NUMBER" use="optional"/>
      <xs:attribute name="recordCount" type="NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="RuleSelectionMethod"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ScoreDistribution"/>
        <xs:group maxOccurs="unbounded" minOccurs="0" ref="Rule"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RuleSet struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="RuleSetModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element ref="RuleSet"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type RuleSetModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*RuleSetModel) modelElement() {}

/*
  <xs:element name="ScoreDistribution">
    <xs:complexType>
      <xs:attribute name="confidence" type="PROB-NUMBER"/>
      <xs:attribute name="probability" type="PROB-NUMBER"/>
      <xs:attribute name="recordCount" type="NUMBER" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ScoreDistribution struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Scorecard">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute default="other" name="baselineMethod">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="max"/>
            <xs:enumeration value="mean"/>
            <xs:enumeration value="min"/>
            <xs:enumeration value="neutral"/>
            <xs:enumeration value="other"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="baselineScore" type="NUMBER"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="0" name="initialScore" type="NUMBER"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute default="pointsBelow" name="reasonCodeAlgorithm">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="pointsAbove"/>
            <xs:enumeration value="pointsBelow"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="true" name="useReasonCodes" type="xs:boolean"/>
      <xs:sequence>
        <xs:element ref="Characteristics"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Scorecard struct {
	Extensions []Extension `xml:"Extension"`
}

func (*Scorecard) modelElement() {}

/*
  <xs:element name="SeasonalTrendDecomposition"/>
*/
type SeasonalTrendDecomposition struct {
}

/*
  <xs:element name="Seasonality_ExpoSmooth">
    <xs:complexType>
      <xs:attribute name="delta" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="period" type="INT-NUMBER" use="required"/>
      <xs:attribute name="phase" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="type" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:NMTOKEN">
            <xs:enumeration value="additive"/>
            <xs:enumeration value="multiplicative"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="unit" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:group ref="REAL-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SeasonalityExpoSmooth struct {
}

/*
  <xs:element name="Segment">
    <xs:complexType>
      <xs:attribute name="id" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="weight" type="NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="MODEL-ELEMENT"/>
        <xs:group ref="PREDICATE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Segment struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Segmentation">
    <xs:complexType>
      <xs:attribute name="multipleModelMethod" type="MULTIPLE-MODEL-METHOD" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="Segment"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Segmentation struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Sequence">
    <xs:complexType>
      <xs:attribute name="id" type="ELEMENT-ID" use="required"/>
      <xs:attribute name="numberOfSets" type="INT-NUMBER"/>
      <xs:attribute name="occurrence" type="INT-NUMBER"/>
      <xs:attribute name="support" type="REAL-NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="SetReference"/>
        <xs:element minOccurs="0" ref="Time"/>
        <xs:sequence maxOccurs="unbounded" minOccurs="0">
          <xs:group ref="FOLLOW-SET"/>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Sequence struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SequenceModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="avgNumberOfItemsPerTransaction" type="REAL-NUMBER"/>
      <xs:attribute name="avgNumberOfTAsPerTAGroup" type="REAL-NUMBER"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="maxNumberOfItemsPerTransaction" type="INT-NUMBER"/>
      <xs:attribute name="maxNumberOfTAsPerTAGroup" type="INT-NUMBER"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="numberOfTransactionGroups" type="INT-NUMBER"/>
      <xs:attribute name="numberOfTransactions" type="INT-NUMBER"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Constraints"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Item"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Itemset"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element maxOccurs="unbounded" ref="Sequence"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="SequenceRule"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="SetPredicate"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SequenceModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*SequenceModel) modelElement() {}

/*
  <xs:element name="SequenceReference">
    <xs:complexType>
      <xs:attribute name="seqId" type="ELEMENT-ID" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SequenceReference struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SequenceRule">
    <xs:complexType>
      <xs:attribute name="confidence" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="id" type="ELEMENT-ID" use="required"/>
      <xs:attribute name="lift" type="REAL-NUMBER"/>
      <xs:attribute name="numberOfSets" type="INT-NUMBER" use="required"/>
      <xs:attribute name="occurrence" type="INT-NUMBER" use="required"/>
      <xs:attribute name="support" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element ref="AntecedentSequence"/>
        <xs:element ref="ConsequentSequence"/>
        <xs:element ref="Delimiter"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="Time"/>
        <xs:element minOccurs="0" ref="Time"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SequenceRule struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SetPredicate">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="id" type="ELEMENT-ID" use="required"/>
      <xs:attribute fixed="supersetOf" name="operator" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="STRING-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SetPredicate struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SetReference">
    <xs:complexType>
      <xs:attribute name="setId" type="ELEMENT-ID" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SetReference struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SigmoidKernelType">
    <xs:complexType>
      <xs:attribute default="1" name="coef0" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="description" type="xs:string" use="optional"/>
      <xs:attribute default="1" name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SigmoidKernelType struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SimplePredicate">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="operator" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="equal"/>
            <xs:enumeration value="greaterOrEqual"/>
            <xs:enumeration value="greaterThan"/>
            <xs:enumeration value="isMissing"/>
            <xs:enumeration value="isNotMissing"/>
            <xs:enumeration value="lessOrEqual"/>
            <xs:enumeration value="lessThan"/>
            <xs:enumeration value="notEqual"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="value" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SimplePredicate struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SimpleRule">
    <xs:complexType>
      <xs:attribute default="1" name="confidence" type="NUMBER" use="optional"/>
      <xs:attribute name="id" type="xs:string" use="optional"/>
      <xs:attribute name="nbCorrect" type="NUMBER" use="optional"/>
      <xs:attribute name="recordCount" type="NUMBER" use="optional"/>
      <xs:attribute name="score" type="xs:string" use="required"/>
      <xs:attribute default="1" name="weight" type="NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="ScoreDistribution"/>
        <xs:group ref="PREDICATE"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SimpleRule struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SimpleSetPredicate">
    <xs:complexType>
      <xs:attribute name="booleanOperator" use="required">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="isIn"/>
            <xs:enumeration value="isNotIn"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:sequence>
        <xs:element ref="Array"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SimpleSetPredicate struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SpectralAnalysis"/>
*/
type SpectralAnalysis struct {
}

/*
  <xs:element name="SupportVector">
    <xs:complexType>
      <xs:attribute name="vectorId" type="VECTOR-ID" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SupportVector struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SupportVectorMachine">
    <xs:complexType>
      <xs:attribute name="alternateTargetCategory" type="xs:string" use="optional"/>
      <xs:attribute name="targetCategory" type="xs:string" use="optional"/>
      <xs:attribute name="threshold" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element ref="Coefficients"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="SupportVectors"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SupportVectorMachine struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="SupportVectorMachineModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute default="OneAgainstAll" name="classificationMethod" type="SVM-CLASSIFICATION-METHOD" use="optional"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute default="false" name="maxWins" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:attribute default="SupportVectors" name="svmRepresentation" type="SVM-REPRESENTATION" use="optional"/>
      <xs:attribute default="0" name="threshold" type="REAL-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element maxOccurs="unbounded" ref="SupportVectorMachine"/>
        <xs:element minOccurs="0" ref="Targets"/>
        <xs:element ref="VectorDictionary"/>
        <xs:sequence>
          <xs:choice>
            <xs:element ref="LinearKernelType"/>
            <xs:element ref="PolynomialKernelType"/>
            <xs:element ref="RadialBasisKernelType"/>
            <xs:element ref="SigmoidKernelType"/>
          </xs:choice>
        </xs:sequence>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SupportVectorMachineModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*SupportVectorMachineModel) modelElement() {}

/*
  <xs:element name="SupportVectors">
    <xs:complexType>
      <xs:attribute name="numberOfAttributes" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="numberOfSupportVectors" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="SupportVector"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SupportVectors struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TableLocator">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TableLocator struct {
	Extensions []Extension `xml:"Extension"`
}

func (*TableLocator) table() {}

type CastIntegerType string

const (
	CastIntegerTypeCeiling = CastIntegerType("ceiling")
	CastIntegerTypeFloor   = CastIntegerType("floor")
	CastIntegerTypeRound   = CastIntegerType("round")
)

/*
  <xs:element name="Target">
    <xs:complexType>
      <xs:attribute name="castInteger">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="ceiling"/>
            <xs:enumeration value="floor"/>
            <xs:enumeration value="round"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="field" type="FIELD-NAME"/>
      <xs:attribute name="max" type="xs:double"/>
      <xs:attribute name="min" type="xs:double"/>
      <xs:attribute name="optype" type="OPTYPE"/>
      <xs:attribute default="0" name="rescaleConstant" type="xs:double"/>
      <xs:attribute default="1" name="rescaleFactor" type="xs:double"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TargetValue"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Target struct {
	CastInteger     CastIntegerType `xml:"castInteger,attr"`
	Field           FieldName       `xml:"field,attr"`
	Max             float64         `xml:"max,attr"`
	Min             float64         `xml:"min,attr"`
	OpType          OpType          `xml:"optype,attr"`
	RescaleConstant *float64        `xml:"rescaleConstant,attr"`
	RescaleFactor   *float64        `xml:"rescaleFactor,attr"`

	Extensions   []Extension   `xml:"Extension"`
	TargetValues []TargetValue `xml:"TargetValue"`
}

/*
  <xs:element name="TargetValue">
    <xs:complexType>
      <xs:attribute name="defaultValue" type="NUMBER"/>
      <xs:attribute name="displayValue" type="xs:string"/>
      <xs:attribute name="priorProbability" type="PROB-NUMBER"/>
      <xs:attribute name="value" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="Partition"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TargetValue struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TargetValueCount">
    <xs:complexType>
      <xs:attribute name="count" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TargetValueCount struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TargetValueCounts">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="TargetValueCount"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TargetValueCounts struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TargetValueStat">
    <xs:complexType>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group minOccurs="1" ref="CONTINUOUS-DISTRIBUTION-TYPES"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TargetValueStat struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TargetValueStats">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="TargetValueStat"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TargetValueStats struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Targets">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="Target"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Targets struct {
	Extensions []Extension `xml:"Extension"`
	Targets    []Target    `xml:"Target"`
}

/*
  <xs:element name="Taxonomy">
    <xs:complexType>
      <xs:attribute name="name" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" ref="ChildParent"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Taxonomy struct {
	Name string `xml:"name,attr"`

	ChildParents []ChildParent `xml:"ChildParent"`
	Extensions   []Extension   `xml:"Extension"`
}

/*
  <xs:element name="TestDistributions">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME" use="required"/>
      <xs:attribute name="normalizationScheme" type="xs:string" use="optional"/>
      <xs:attribute default="0.0" name="resetValue" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="testStatistic" type="BASELINE-TEST-STATISTIC" use="required"/>
      <xs:attribute name="weightField" type="FIELD-NAME" use="optional"/>
      <xs:attribute default="0" name="windowSize" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Alternate"/>
        <xs:element ref="Baseline"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TestDistributions struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextCorpus">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TextDocument"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextCorpus struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextDictionary">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="Taxonomy"/>
        <xs:group ref="STRING-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextDictionary struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextDocument">
    <xs:complexType>
      <xs:attribute name="file" type="xs:string" use="optional"/>
      <xs:attribute name="id" type="xs:string" use="required"/>
      <xs:attribute name="length" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="name" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextDocument struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextIndex">
    <xs:complexType>
      <xs:attribute default="allHits" name="countHits">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="allHits"/>
            <xs:enumeration value="bestHits"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="false" name="isCaseSensitive" type="xs:boolean"/>
      <xs:attribute default="termFrequency" name="localTermWeights">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="augmentedNormalizedTermFrequency"/>
            <xs:enumeration value="binary"/>
            <xs:enumeration value="logarithmic"/>
            <xs:enumeration value="termFrequency"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="0" name="maxLevenshteinDistance" type="xs:integer"/>
      <xs:attribute name="textField" type="FIELD-NAME" use="required"/>
      <xs:attribute default="true" name="tokenize" type="xs:boolean"/>
      <xs:attribute default="\s" name="wordSeparatorCharacterRE" type="xs:string"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TextIndexNormalization"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextIndex struct {
	Extensions []Extension `xml:"Extension"`
}

func (*TextIndex) expression() {}

/*
  <xs:element name="TextIndexNormalization">
    <xs:complexType>
      <xs:attribute default="string" name="inField" type="xs:string"/>
      <xs:attribute name="isCaseSensitive" type="xs:boolean"/>
      <xs:attribute name="maxLevenshteinDistance" type="xs:integer"/>
      <xs:attribute default="stem" name="outField" type="xs:string"/>
      <xs:attribute default="false" name="recursive" type="xs:boolean"/>
      <xs:attribute default="regex" name="regexField" type="xs:string"/>
      <xs:attribute name="tokenize" type="xs:boolean"/>
      <xs:attribute name="wordSeparatorCharacterRE" type="xs:string"/>
      <xs:sequence>
        <xs:choice minOccurs="0">
          <xs:element ref="InlineTable"/>
          <xs:element ref="TableLocator"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextIndexNormalization struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute name="numberOfDocuments" type="xs:integer" use="required"/>
      <xs:attribute name="numberOfTerms" type="xs:integer" use="required"/>
      <xs:sequence>
        <xs:element ref="DocumentTermMatrix"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
        <xs:element ref="TextCorpus"/>
        <xs:element ref="TextDictionary"/>
        <xs:element minOccurs="0" ref="TextModelNormalization"/>
        <xs:element minOccurs="0" ref="TextModelSimiliarity"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*TextModel) modelElement() {}

/*
  <xs:element name="TextModelNormalization">
    <xs:complexType>
      <xs:attribute default="none" name="documentNormalization">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="cosine"/>
            <xs:enumeration value="none"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="inverseDocumentFrequency" name="globalTermWeights">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="GFIDF"/>
            <xs:enumeration value="inverseDocumentFrequency"/>
            <xs:enumeration value="none"/>
            <xs:enumeration value="normal"/>
            <xs:enumeration value="probabilisticInverse"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute default="termFrequency" name="localTermWeights">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="augmentedNormalizedTermFrequency"/>
            <xs:enumeration value="binary"/>
            <xs:enumeration value="logarithmic"/>
            <xs:enumeration value="termFrequency"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextModelNormalization struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TextModelSimiliarity">
    <xs:complexType>
      <xs:attribute name="similarityType">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="cosine"/>
            <xs:enumeration value="euclidean"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TextModelSimiliarity struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Time">
    <xs:complexType>
      <xs:attribute name="max" type="NUMBER"/>
      <xs:attribute name="mean" type="NUMBER"/>
      <xs:attribute name="min" type="NUMBER"/>
      <xs:attribute name="standardDeviation" type="NUMBER"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Time struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TimeAnchor">
    <xs:complexType>
      <xs:attribute name="displayName" use="optional"/>
      <xs:attribute name="offset" type="INT-NUMBER"/>
      <xs:attribute name="stepsize" type="INT-NUMBER"/>
      <xs:attribute name="type" type="TIME-ANCHOR"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TimeCycle"/>
        <xs:element maxOccurs="2" minOccurs="0" ref="TimeException"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeAnchor struct {
}

/*
  <xs:element name="TimeCycle">
    <xs:complexType>
      <xs:attribute name="displayName" use="optional"/>
      <xs:attribute name="length" type="INT-NUMBER"/>
      <xs:attribute name="type" type="VALID-TIME-SPEC"/>
      <xs:sequence>
        <xs:group maxOccurs="1" minOccurs="0" ref="INT-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeCycle struct {
}

/*
  <xs:element name="TimeException">
    <xs:complexType>
      <xs:attribute name="count" type="INT-NUMBER"/>
      <xs:attribute name="type" type="TIME-EXCEPTION-TYPE"/>
      <xs:sequence>
        <xs:group minOccurs="1" ref="INT-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeException struct {
}

/*
  <xs:element name="TimeSeries">
    <xs:complexType>
      <xs:attribute name="endTime" type="REAL-NUMBER"/>
      <xs:attribute default="none" name="interpolationMethod" type="INTERPOLATION-METHOD"/>
      <xs:attribute name="startTime" type="REAL-NUMBER"/>
      <xs:attribute default="original" name="usage" type="TIMESERIES-USAGE"/>
      <xs:sequence>
        <xs:element maxOccurs="1" minOccurs="0" ref="TimeAnchor"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="TimeValue"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeSeries struct {
}

/*
  <xs:element name="TimeSeriesModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string" use="optional"/>
      <xs:attribute name="bestFit" type="TIMESERIES-ALGORITHM" use="required"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute name="modelName" type="xs:string" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="1" minOccurs="0" ref="ARIMA"/>
        <xs:element maxOccurs="1" minOccurs="0" ref="ExponentialSmoothing"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element maxOccurs="1" minOccurs="0" ref="SeasonalTrendDecomposition"/>
        <xs:element maxOccurs="1" minOccurs="0" ref="SpectralAnalysis"/>
        <xs:element maxOccurs="3" minOccurs="0" ref="TimeSeries"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeSeriesModel struct {
	Extensions []Extension `xml:"Extension"`
}

func (*TimeSeriesModel) modelElement() {}

/*
  <xs:element name="TimeValue">
    <xs:complexType>
      <xs:attribute name="index" type="INT-NUMBER" use="optional"/>
      <xs:attribute name="standardError" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="time" type="NUMBER" use="optional"/>
      <xs:attribute name="value" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="1" minOccurs="0" ref="Timestamp"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TimeValue struct {
}

/*
  <xs:element name="Timestamp">
    <xs:complexType mixed="true">
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Timestamp struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TrainingInstances">
    <xs:complexType>
      <xs:attribute name="fieldCount" type="INT-NUMBER" use="optional"/>
      <xs:attribute default="false" name="isTransformed" type="xs:boolean"/>
      <xs:attribute name="recordCount" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:choice>
          <xs:element ref="InlineTable"/>
          <xs:element ref="TableLocator"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="InstanceFields"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TrainingInstances struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="TransformationDictionary">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DefineFunction"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="DerivedField"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TransformationDictionary struct {
	DefineFunctions []DefineFunction `xml:"DefineFunction"`
	DerivedFields   []DerivedField   `xml:"DerivedField"`
	Extensions      []Extension      `xml:"Extension"`
}

type TreeModelSplitCharacteristic string

const (
	TreeModelSplitCharacteristicBinarySplit = TreeModelSplitCharacteristic("binarySplit")
	TreeModelSplitCharacteristicMultiSplit  = TreeModelSplitCharacteristic("multiSplit")
)

/*
  <xs:element name="TreeModel">
    <xs:complexType>
      <xs:attribute name="algorithmName" type="xs:string"/>
      <xs:attribute name="functionName" type="MINING-FUNCTION" use="required"/>
      <xs:attribute default="true" name="isScorable" type="xs:boolean"/>
      <xs:attribute default="1.0" name="missingValuePenalty" type="PROB-NUMBER"/>
      <xs:attribute default="none" name="missingValueStrategy" type="MISSING-VALUE-STRATEGY"/>
      <xs:attribute name="modelName" type="xs:string"/>
      <xs:attribute default="returnNullPrediction" name="noTrueChildStrategy" type="NO-TRUE-CHILD-STRATEGY"/>
      <xs:attribute default="multiSplit" name="splitCharacteristic">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="binarySplit"/>
            <xs:enumeration value="multiSplit"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="LocalTransformations"/>
        <xs:element ref="MiningSchema"/>
        <xs:element minOccurs="0" ref="ModelExplanation"/>
        <xs:element minOccurs="0" ref="ModelStats"/>
        <xs:element minOccurs="0" ref="ModelVerification"/>
        <xs:element ref="Node"/>
        <xs:element minOccurs="0" ref="Output"/>
        <xs:element minOccurs="0" ref="Targets"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TreeModel struct {
	AlgorithmName        string                       `xml:"algorithmName,attr"`
	FunctionName         MiningFunction               `xml:"functionName,attr"`
	IsScorable           bool                         `xml:"isScorable,attr"`
	MissingValuePenalty  ProbNumber                   `xml:"missingValuePenalty,attr"`
	MissingValueStrategy MissingValueStrategy         `xml:"missingValueStrategy,attr"`
	ModelName            string                       `xml:"modelName,attr"`
	NoTrueChildStrategy  NoTrueChildStrategy          `xml:"noTrueChildStrategy,attr"`
	SplitCharacteristic  TreeModelSplitCharacteristic `xml:"splitCharacteristic,attr"`

	Extensions           []Extension          `xml:"Extension"`
	LocalTransformations LocalTransformations `xml:"LocalTransformations"`
	MiningSchema         MiningSchema         `xml:"MiningSchema"`
	ModelExplanation     *ModelExplanation    `xml:"ModelExplanation"`
	ModelStats           *ModelStats          `xml:"ModelStats"`
	ModelVerification    *ModelVerification   `xml:"ModelVerification"`
	Node                 Node                 `xml:"Node"`
	Output               *Output              `xml:"Output"`
	Targets              *Targets             `xml:"Targets"`
}

func (*TreeModel) modelElement() {}

/*
  <xs:element name="Trend_ExpoSmooth">
    <xs:complexType>
      <xs:attribute name="gamma" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="1" name="phi" type="REAL-NUMBER" use="optional"/>
      <xs:attribute name="smoothedValue" type="REAL-NUMBER" use="optional"/>
      <xs:attribute default="additive" name="trend">
        <xs:simpleType>
          <xs:restriction base="xs:NMTOKEN">
            <xs:enumeration value="additive"/>
            <xs:enumeration value="damped_additive"/>
            <xs:enumeration value="damped_multiplicative"/>
            <xs:enumeration value="multiplicative"/>
            <xs:enumeration value="polynomial_exponential"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:group minOccurs="0" ref="REAL-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TrendExpoSmooth struct {
}

/*
  <xs:element name="TriangularDistributionForBN">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Lower"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Mean"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Upper"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type TriangularDistributionForBN struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="True">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type True struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="UniformDistribution">
    <xs:complexType>
      <xs:attribute name="lower" type="REAL-NUMBER" use="required"/>
      <xs:attribute name="upper" type="REAL-NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type UniformDistribution struct {
	Extensions []Extension `xml:"Extension"`
}

func (*UniformDistribution) continuousDistributionType() {}

/*
  <xs:element name="UniformDistributionForBN">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Lower"/>
        <xs:element maxOccurs="1" minOccurs="1" ref="Upper"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type UniformDistributionForBN struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="UnivariateStats">
    <xs:complexType>
      <xs:attribute name="field" type="FIELD-NAME"/>
      <xs:attribute default="0" name="weighted">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="0"/>
            <xs:enumeration value="1"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:sequence>
        <xs:element minOccurs="0" ref="Anova"/>
        <xs:element minOccurs="0" ref="ContStats"/>
        <xs:element minOccurs="0" ref="Counts"/>
        <xs:element minOccurs="0" ref="DiscrStats"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element minOccurs="0" ref="NumericInfo"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type UnivariateStats struct {
	Field    FieldName `xml:"field,attr"`
	Weighted bool      `xml:"weighted,attr"`

	Anova       *Anova       `xml:"Anova"`
	ContStats   *ContStats   `xml:"ContStats"`
	Counts      *Counts      `xml:"Counts"`
	DiscrStats  *DiscrStats  `xml:"DiscrStats"`
	Extensions  []Extension  `xml:"Extension"`
	NumericInfo *NumericInfo `xml:"NumericInfo"`
}

/*
  <xs:element name="Upper">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Upper struct {
	Extensions []Extension `xml:"Extension"`
}

type ValueValid string

const (
	ValueValidInvalid = ValueValid("invalid")
	ValueValidMissing = ValueValid("missing")
	ValueValidValid   = ValueValid("valid")
)

/*
  <xs:element name="Value">
    <xs:complexType>
      <xs:attribute name="displayValue" type="xs:string"/>
      <xs:attribute default="valid" name="property">
        <xs:simpleType>
          <xs:restriction base="xs:string">
            <xs:enumeration value="invalid"/>
            <xs:enumeration value="missing"/>
            <xs:enumeration value="valid"/>
          </xs:restriction>
        </xs:simpleType>
      </xs:attribute>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Value struct {
	DisplayValue string     `xml:"displayValue,attr"`
	Valid        ValueValid `xml:"valid,attr"`
	Value        string     `xml:"value,attr"`

	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="ValueProbability">
    <xs:complexType>
      <xs:attribute name="probability" type="PROB-NUMBER" use="required"/>
      <xs:attribute name="value" type="xs:string" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type ValueProbability struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="Variance">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="EXPRESSION"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Variance struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="VectorDictionary">
    <xs:complexType>
      <xs:attribute name="numberOfVectors" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element ref="VectorFields"/>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="VectorInstance"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type VectorDictionary struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="VectorFields">
    <xs:complexType>
      <xs:attribute name="numberOfFields" type="INT-NUMBER" use="optional"/>
      <xs:sequence>
        <xs:choice maxOccurs="unbounded">
          <xs:element ref="CategoricalPredictor"/>
          <xs:element ref="FieldRef"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type VectorFields struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="VectorInstance">
    <xs:complexType>
      <xs:attribute name="id" type="VECTOR-ID" use="required"/>
      <xs:sequence>
        <xs:choice>
          <xs:element ref="REAL-SparseArray"/>
          <xs:group ref="REAL-ARRAY"/>
        </xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type VectorInstance struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="VerificationField">
    <xs:complexType>
      <xs:attribute name="column" type="xs:string" use="optional"/>
      <xs:attribute name="field" type="xs:string" use="required"/>
      <xs:attribute default="1E-6" name="precision" type="xs:double"/>
      <xs:attribute default="1E-16" name="zeroThreshold" type="xs:double"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type VerificationField struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="VerificationFields">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:element maxOccurs="unbounded" ref="VerificationField"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type VerificationFields struct {
	Extensions []Extension         `xml:"Extension"`
	Fields     []VerificationField `xml:"VerificationField"`
}

/*
  <xs:element name="XCoordinates">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type XCoordinates struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="YCoordinates">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
        <xs:group ref="NUM-ARRAY"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type YCoordinates struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="binarySimilarity">
    <xs:complexType>
      <xs:attribute name="c00-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="c01-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="c10-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="c11-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="d00-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="d01-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="d10-parameter" type="NUMBER" use="required"/>
      <xs:attribute name="d11-parameter" type="NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type BinarySimilarity struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="chebychev">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Chebychev struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="cityBlock">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type CityBlock struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="euclidean">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Euclidean struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="jaccard">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Jaccard struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="minkowski">
    <xs:complexType>
      <xs:attribute name="p-parameter" type="NUMBER" use="required"/>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Minkowski struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="row">
    <xs:complexType>
      <xs:complexContent mixed="true">
        <xs:restriction base="xs:anyType">
          <xs:sequence>
            <xs:any maxOccurs="unbounded" minOccurs="2" processContents="skip"/>
          </xs:sequence>
        </xs:restriction>
      </xs:complexContent>
    </xs:complexType>
  </xs:element>
*/
type Row struct {
}

/*
  <xs:element name="simpleMatching">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SimpleMatching struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="squaredEuclidean">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type SquaredEuclidean struct {
	Extensions []Extension `xml:"Extension"`
}

/*
  <xs:element name="tanimoto">
    <xs:complexType>
      <xs:sequence>
        <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
      </xs:sequence>
    </xs:complexType>
  </xs:element>
*/
type Tanimoto struct {
	Extensions []Extension `xml:"Extension"`
}
