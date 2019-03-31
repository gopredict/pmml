package pmml

import (
	"fmt"

	shellwords "github.com/mattn/go-shellwords"
)

type SimpleSetPredicate struct {
	Field    string `xml:"field,attr"`
	Operator string `xml:"booleanOperator,attr"`
	Values   string `xml:"Array"`

	values map[string]bool

	execute func(features map[string]interface{}) PredicateResult
}

func (p *SimpleSetPredicate) Compile() error {
	values, err := shellwords.Parse(p.Values)
	if err != nil {
		return err
	}

	p.values = map[string]bool{}

	for _, v := range values {
		p.values[v] = true
	}

	if p.Operator == "isIn" {
		p.execute = p.isIn
		return nil
	}

	if p.Operator == "isNotIn" {
		p.execute = p.isNotIn
		return nil
	}

	return fmt.Errorf("pmml: invalid operator '%s'", p.Operator)
}

func (p *SimpleSetPredicate) Execute(features map[string]interface{}) PredicateResult {
	return p.execute(features)
}

func (p *SimpleSetPredicate) isIn(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return False
	}

	if v, ok := val.(string); ok {
		_, ok = p.values[v]
		if ok {
			return True
		}
	}

	return False
}

func (p *SimpleSetPredicate) isNotIn(features map[string]interface{}) PredicateResult {
	val := p.isIn(features)

	if val == True {
		return False
	}

	return Unknown
}
