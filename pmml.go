package pmml

import (
	"encoding/xml"
	"time"
)

type Model interface {
	Compile() error
	Execute(map[string]interface{}) (PredicateResult, string, float64)
}

// PMML document can contain more than one model. If the application system
// provides a means of selecting models by name and if the PMML consumer
// specifies a model name, then that model is used; otherwise the first model
// is used.
// http://dmg.org/pmml/v4-3/GeneralStructure.html#xsdElement_PMML
type PMML struct {
	Version                  string
	Header                   Header
	DataDictionary           DataDictionary
	TransformationDictionary TransformationDictionary
	Models                   []Model
}

// UnmarshalXML implements the xml.Unmarshaler interface.
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
			var model Model

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

// Header is the top level tag that marks the beginning of the header
// information.
// http://dmg.org/pmml/v4-3/Header.html#xsdElement_Header
type Header struct {
	Application  Application  `xml:"Application"`
	Annotations  []Annotation `xml:"Annotation"`
	Timestamp    Timestamp    `xml:"Timestamp"`
	Copyright    string       `xml:"copyright,attr"`
	Description  string       `xml:"description,attr"`
	ModelVersion string       `xml:"modelVersion,attr"`
}

// Application describes the software application that generated the model.
// Although the PMML models are created to be portable, different mechanisms may
// create different models from the same data set. It is of interest to the user
// from which application these models were generated.
// http://dmg.org/pmml/v4-3/Header.html#xsdElement_Application
type Application struct {
	Name    string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

// Annotation contains document modification history. Each annotation is free
// text and, like the description attribute in the head element, makes sense to
// the human eye only.
// http://dmg.org/pmml/v4-3/Header.html#xsdElement_Annotation
type Annotation struct {
	Value string `xml:",chardata"`
}

// Timestamp is the model creation timestamp, it is recommended that its format
// should follow the XML standard.
// http://dmg.org/pmml/v4-3/Header.html#xsdElement_Timestamp
type Timestamp struct {
	Value time.Time `xml:",chardata"`
}

// DataDictionary contains definitions for fields as used in mining models. It
// specifies the types and value ranges. These definitions are assumed to be
// independent of specific data sets as used for training or scoring a specific
// model.
// http://dmg.org/pmml/v4-3/DataDictionary.html#xsdElement_DataDictionary
type DataDictionary struct {
	NumberOfFields uint        `xml:"numberOfFields,attr"`
	DataFields     []DataField `xml:"DataField"`
	Taxonomies     []Taxonomy  `xml:"Taxonomy"`
}

type DataField struct { // http://dmg.org/pmml/v4-3/DataDictionary.html#xsdElement_DataField
	Name        string   `xml:"name,attr"`
	DisplayName string   `xml:"displayName,attr"`
	OpType      OpType   `xml:"optype,attr"`
	DataType    DataType `xml:"dataType,attr"`
	Taxonomy    string   `xml:"taxonomy,attr"`
	IsCyclic    bool     `xml:"isCyclic,attr"`

	Intervals []Interval `xml:"Interval"`
	Values    []Value    `xml:"Value"`
}

type Interval struct { // http://dmg.org/pmml/v4-3/DataDictionary.html#xsdElement_Interval
	Closure     string  `xml:"closure,attr"`
	LeftMargin  float64 `xml:"leftMargin,attr"`
	RightMargin float64 `xml:"rightMargin,attr"`
}

type Value struct { // http://dmg.org/pmml/v4-3/DataDictionary.html#xsdElement_Value
	Value        string `xml:"value,attr"`
	DisplayValue string `xml:"displayValue,attr"`
	Property     string `xml:"property,attr"`
}

type Taxonomy struct { // http://dmg.org/pmml/v4-3/Taxonomy.html#xsdElement_Taxonomy
	Name         string        `xml:"name,attr"`
	ChildParents []ChildParent `xml:"ChildParent"`
}

type ChildParent struct { // http://dmg.org/pmml/v4-3/Taxonomy.html#xsdElement_ChildParent
	ChildField       string `xml:"childField,attr"`
	ParentField      string `xml:"parentField,attr"`
	ParentLevelField string `xml:"parentLevelField,attr"`
	IsRecursive      bool   `xml:"isRecursive,attr"`
}

type TransformationDictionary struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_TransformationDictionary
	DefineFunctions []DefineFunction `xml:"DefineFunction"`
	DerivedFields   []DerivedField   `xml:"DerivedField"`
}

type DefineFunction struct { // http://dmg.org/pmml/v4-3/Functions.html#xsdElement_DefineFunction
	Name            string           `xml:"name,attr"`
	OpType          OpType           `xml:"optype,attr"`
	DataType        DataType         `xml:"dataType,attr"`
	ParameterFields []ParameterField `xml:"ParameterField"`
	Expression      Expression       `xml:"Expression"`
}

type ParameterField struct { // http://dmg.org/pmml/v4-3/Functions.html#xsdElement_ParameterField
	Name     string   `xml:"name,attr"`
	OpType   OpType   `xml:"optype,attr"`
	DataType DataType `xml:"dataType,attr"`
}

type DerivedField struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_DerivedField
	Name        string     `xml:"name,attr"`
	DisplayName string     `xml:"displayName,attr"`
	OpType      OpType     `xml:"optype,attr"`
	DataType    DataType   `xml:"dataType,attr"`
	Expression  Expression `xml:"Expression"`
	Values      []Value    `xml:"Value"`
}

type Expression interface { // http://dmg.org/pmml/v4-3/Transformations.html#xsdGroup_EXPRESSION
	// TODO: Expression is one of Constant, FieldRef, NormContinuous, NormDiscrete, Discretize, MapValues, TextIndex, Apply, Aggregate, Lag
}

type Constant struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_Constant
	Value    string   `xml:",chardata"`
	DataType DataType `xml:"dataType,attr"`
}

type NormContinuous struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_NormContinuous
	LinearNorms  []LinearNorm `xml:"LinearNorm"`
	Field        string       `xml:"field,attr"`
	MapMissingTo string       `xml:"mapMissingTo"`
	Outliers     string       `xml:"outliers,attr"`
}

type LinearNorm struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_LinearNorm
	Orig float64 `xml:"orig,attr"`
	Norm float64 `xml:"norm,attr"`
}

type NormDiscrete struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_NormDiscrete
	Field        string  `xml:"field,attr"`
	Value        string  `xml:"value,attr"`
	MapMissingTo float64 `xml:"mapMissingTo"`
}

type Discretize struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_Discretize

}

type MapValues struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_MapValues
}

type TextIndex struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_TextIndex
}

type Apply struct { // http://dmg.org/pmml/v4-3/Functions.html#xsdElement_Apply
}

type Aggregate struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_Aggregate
	Field      string `xml:"field,attr"`
	Function   string `xml:"function,attr"`
	GroupField string `xml:"groupField,attr"`
	SQLWhere   string `xml:"sqlWhere"`
}

type Lag struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_Lag
	Field           string           `xml:"field,attr"`
	N               uint             `xml:"n,attr"`
	BlockIndicators []BlockIndicator `xml:"BlockIndicator"`
}

type BlockIndicator struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_BlockIndicator
	Field string `xml:"field"`
}

type MiningSchema struct {
	MiningFields []MiningField `xml:"MiningField"`
}

type MiningField struct {
	Name                    string `xml:"name,attr"`
	UsageType               string `xml:"usageType,attr"`
	OpType                  OpType `xml:"optype,attr"`
	Importance              string `xml:"importance,attr"`
	Outliers                string `xml:"outliers,attr"`
	LowValue                string `xml:"lowValue,attr"`
	HighValue               string `xml:"highValue,attr"`
	MissingValueReplacement string `xml:"missingValueReplacement,attr"`
	MissingValueTreatment   string `xml:"missingValueTreatment,attr"`
	InvalidValueTreatment   string `xml:"invalidValueTreatment,attr"`
}

type Output struct {
	OutputFields []OutputField `xml:"OutputField"`
}

type OutputField struct {
	Name string `xml:"name,attr"`
}

type ModelStats struct {
}

type LocalTransformations struct {
	DerivedFields []DerivedField `xml:"DerivedField"`
}
