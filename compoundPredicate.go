package pmml

import (
	"encoding/xml"
	"fmt"

	"github.com/pkg/errors"
)

type CompoundPredicate struct {
	Operator string      `xml:"booleanOperator,attr"`
	Values   []Predicate `xml:"Array"`

	execute func(map[string]interface{}) PredicateResult
}

func (p *CompoundPredicate) Compile() error {
	for _, child := range p.Values {
		err := child.Compile()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if p.Operator == "and" {
		p.execute = p.and
		return nil
	}

	if p.Operator == "or" {
		p.execute = p.or
		return nil
	}

	if p.Operator == "xor" {
		p.execute = p.xor
		return nil
	}

	if p.Operator == "surrogate" {
		p.execute = p.surrogate
		return nil
	}

	return fmt.Errorf("pmml: invalid operator '%s'", p.Operator)
}

func (p *CompoundPredicate) Execute(features map[string]interface{}) PredicateResult {
	return p.execute(features)
}

func (p *CompoundPredicate) and(features map[string]interface{}) PredicateResult {
	for _, child := range p.Values {
		if val := child.Execute(features); val != True {
			return val
		}
	}

	return True
}

func (p *CompoundPredicate) or(features map[string]interface{}) PredicateResult {
	allUnknown := true
	for _, child := range p.Values {
		if val := child.Execute(features); val == True {
			return val
		} else if val != Unknown {
			allUnknown = false
		}
	}

	if allUnknown {
		return Unknown
	}

	return False
}

func (p *CompoundPredicate) xor(features map[string]interface{}) PredicateResult {
	// The operator xor indicates an evaluation to TRUE if an odd number of the predicates evaluates to TRUE and all others evaluate to FALSE.
	trueCount := 0
	anyUnknown := false

	for _, child := range p.Values {
		if val := child.Execute(features); val == True {
			trueCount++
		} else if val == Unknown {
			anyUnknown = true
		}
	}

	if trueCount%2 == 1 && !anyUnknown {
		return True
	}

	return False
}

func (p *CompoundPredicate) surrogate(features map[string]interface{}) PredicateResult {
	// The operator surrogate allows for specifying surrogate predicates. They
	// are used for cases where a missing value appears in the evaluation of the
	// parent predicate such that an alternative predicate is available.
	for _, child := range p.Values {
		if val := child.Execute(features); val != Unknown {
			return val
		}
	}

	return Unknown
}

func (p *CompoundPredicate) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "booleanOperator":
			p.Operator = attr.Value
		}
	}

	for {
		token, err := decoder.Token()
		if err != nil {
			return err
		}
		switch tok := token.(type) {
		case xml.StartElement:
			var el Predicate
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
			}

			if el != nil {
				err := decoder.DecodeElement(el, &tok)
				if err != nil {
					return err
				}

				p.Values = append(p.Values, el)
			}

		case xml.EndElement:
			return nil
		}
	}
}
