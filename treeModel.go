package pmml

import (
	"encoding/xml"
	"strconv"
)

type TreeModel struct {
	MiningSchema         MiningSchema `xml:"MiningSchema"`
	Node                 *Node        `xml:"Node"`
	ModelName            string       `xml:"modelName,attr"`
	FunctionName         string       `xml:"functionName,attr"`
	AlgorithmName        string       `xml:"algorithmName,attr"`
	MissingValueStrategy string       `xml:"missingValueStrategy,attr"`
	MissingValuePenalty  float64      `xml:"missingValuePenalty,attr"`
	NoTrueChildStrategy  string       `xml:"noTrueChildStrategy,attr"`
	SplitCharacteristic  string       `xml:"splitCharacteristic,attr"`
	IsScorable           bool         `xml:"isScorable,attr"`
}

func (m *TreeModel) Compile() error {
	return m.Node.Compile(m.MissingValueStrategy, m.MissingValuePenalty, m.NoTrueChildStrategy, m.SplitCharacteristic)
}

type MiningSchema struct {
	MiningFields []MiningField `xml:"MiningField"`
}

type MiningField struct {
	Name                    string `xml:"name,attr"`
	UsageType               string `xml:"usageType,attr"`
	OpType                  string `xml:"optype,attr"`
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

type Node struct {
	ID           string  `xml:"id,attr"`
	Score        string  `xml:"score,attr"`
	RecordCount  float64 `xml:"recordCount,attr"`
	DefaultChild string  `xml:"defaultChild,attr"`

	Predicate Predicate
	Nodes     []*Node `xml:"Node"`
	// Partition
	ScoreDistributions []ScoreDistribution `xml:"ScoreDistribution"`
	// EmbeddedModel

	scoreDistributionMap map[string]ScoreDistribution
	missingValueStrategy string
	missingValuePenalty  float64
	noTrueChildStrategy  string
	splitCharacteristic  string
}

type ScoreDistribution struct {
	Value       string  `xml:"value,attr"`
	RecordCount float64 `xml:"recordCount,attr"`
	Confidence  float64 `xml:"confidence,attr"`
	Probability float64 `xml:"probability,attr"`
}

func (n *Node) Compile(missingValueStrategy string, missingValuePenalty float64, noTrueChildStrategy string, splitCharacteristic string) error {
	err := n.Predicate.Compile()
	if err != nil {
		return err
	}

	for _, child := range n.Nodes {
		err = child.Compile(missingValueStrategy, missingValuePenalty, noTrueChildStrategy, splitCharacteristic)
		if err != nil {
			return err
		}
	}

	n.missingValueStrategy = missingValueStrategy
	n.missingValuePenalty = missingValuePenalty
	n.noTrueChildStrategy = noTrueChildStrategy
	n.splitCharacteristic = splitCharacteristic

	if len(n.ScoreDistributions) > 0 {
		n.scoreDistributionMap = map[string]ScoreDistribution{}
	}

	for _, sd := range n.ScoreDistributions {
		n.scoreDistributionMap[sd.Value] = sd
	}

	return nil
}

func (n *Node) Execute(features map[string]interface{}) (PredicateResult, string, float64) {
	result := n.Predicate.Execute(features)
	if result != True {
		return result, "", 0.0
	}

	var score string
	var confidence float64

	for _, child := range n.Nodes {
		result, score, confidence = child.Execute(features)
		if result == True {
			break
		}
	}

	if score == "" {
		score = n.Score

		if sd, ok := n.scoreDistributionMap[score]; ok {
			confidence = sd.Confidence
		}
	}

	return result, score, confidence
}

func (n *Node) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var err error

	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "id":
			n.ID = attr.Value
		case "score":
			n.Score = attr.Value
		case "recordCount":
			n.RecordCount, err = strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
		case "defaultChild":
			n.DefaultChild = attr.Value
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
			var el Predicate
			var child *Node
			var dist *ScoreDistribution

			switch tok.Name.Local {
			case "True":
				el = &TruePredicate{}
			case "False":
				el = &FalsePredicate{}
			case "CompoundPredicate":
				el = &CompoundPredicate{}
			case "SimplePredicate":
				el = &SimplePredicate{}
			case "SimpleSetPredicate":
				el = &SimpleSetPredicate{}
			case "Node":
				child = &Node{}
			case "ScoreDistribution":
				dist = &ScoreDistribution{}
			}

			if el != nil {
				err = decoder.DecodeElement(el, &tok)
				if err != nil {
					return err
				}

				n.Predicate = el
			}

			if child != nil {
				err = decoder.DecodeElement(child, &tok)
				if err != nil {
					return err
				}

				n.Nodes = append(n.Nodes, child)
			}

			if dist != nil {
				err = decoder.DecodeElement(dist, &tok)
				if err != nil {
					return err
				}

				n.ScoreDistributions = append(n.ScoreDistributions, *dist)
			}

		case xml.EndElement:
			return nil
		}
	}
}
