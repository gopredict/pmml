package pmml

import (
	"fmt"
	"strconv"
)

// SimplePredicate - PMML simple predicate
type SimplePredicate struct {
	Field    string `xml:"field,attr"`
	Operator string `xml:"operator,attr"`
	Value    string `xml:"value,attr"`

	intVal   int
	boolVal  bool
	floatVal float64

	isInt   bool
	isFloat bool
	isBool  bool

	execute func(features map[string]interface{}) PredicateResult
}

func (p *SimplePredicate) Compile() error {
	if v, err := strconv.ParseBool(p.Value); err == nil {
		p.boolVal = v
		p.isBool = true
	}

	if v, err := strconv.Atoi(p.Value); err == nil {
		p.intVal = v
		p.isInt = true
	}

	if v, err := strconv.ParseFloat(p.Value, 64); err == nil {
		p.floatVal = v
		p.isFloat = true
	}

	if p.Operator == "isMissing" {
		p.execute = p.isMissing
		return nil
	}

	if p.Operator == "isNotMissing" {
		p.execute = p.isNotMissing
		return nil
	}

	if p.Operator == "equal" {
		p.execute = p.equal
		return nil
	}

	if p.Operator == "notEqual" {
		p.execute = p.notEqual
		return nil
	}

	if p.Operator == "lessThan" {
		p.execute = p.lessThan
		return nil
	}

	if p.Operator == "lessOrEqual" {
		p.execute = p.lessOrEqual
		return nil
	}

	if p.Operator == "greaterThan" {
		p.execute = p.greaterThan
		return nil
	}

	if p.Operator == "greaterOrEqual" {
		p.execute = p.greaterOrEqual
		return nil
	}

	return fmt.Errorf("pmml: invalid operator '%s'", p.Operator)
}

func (p *SimplePredicate) Execute(features map[string]interface{}) PredicateResult {
	if p.execute != nil {
		return p.execute(features)
	}

	return False
}

func (p *SimplePredicate) isMissing(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return True
	}

	if val == nil {
		return True
	}

	if v, ok := val.(string); ok {
		if v == "" {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) isNotMissing(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return False
	}

	if val == nil {
		return False
	}

	if v, ok := val.(string); ok {
		if v != "" {
			return True
		}
		return False
	}

	return True
}

func (p *SimplePredicate) equal(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case bool:
		if p.isBool && v == p.boolVal {
			return True
		}
	case int:
		if p.isInt && v == p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v == p.floatVal {
			return True
		}
	case string:
		if v == p.Value {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) notEqual(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case bool:
		if p.isBool && v != p.boolVal {
			return True
		}
	case int:
		if p.isInt && v != p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v != p.floatVal {
			return True
		}
	case string:
		if v == p.Value {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) lessThan(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case int:
		if p.isInt && v < p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v < p.floatVal {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) lessOrEqual(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case int:
		if p.isInt && v <= p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v <= p.floatVal {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) greaterThan(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case int:
		if p.isInt && v > p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v > p.floatVal {
			return True
		}
	}

	return False
}

func (p *SimplePredicate) greaterOrEqual(features map[string]interface{}) PredicateResult {
	val, ok := features[p.Field]
	if !ok {
		return Unknown
	}

	switch v := val.(type) {
	case int:
		if p.isInt && v >= p.intVal {
			return True
		}
	case float64:
		if p.isFloat && v >= p.floatVal {
			return True
		}
	}

	return False
}
