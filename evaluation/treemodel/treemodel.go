package treemodel

import (
	"errors"
	"strconv"

	"github.com/gopredict/pmml/evaluation"
	"github.com/gopredict/pmml/evaluation/common"
	"github.com/gopredict/pmml/models"
)

type TreeModel struct {
	validated bool
	compiled  bool

	dd *models.DataDictionary
	td *models.TransformationDictionary

	model *models.TreeModel

	fieldTypes map[models.FieldName]models.DataType
	root       node

	outputField models.FieldName
}

type test func(evaluation.DataRow) predicateResult

type node struct {
	children []node

	test  test
	score string
}

type predicateResult int

const (
	u predicateResult = iota
	t
	f
)

func (n node) evaluate(input evaluation.DataRow) (score string, ok predicateResult) {
	result := n.test(input)

	if result == f {
		return "", result
	}

	for _, c := range n.children {
		childResult := c.test(input)
		if childResult == t {
			return c.score, childResult
		}
	}

	return n.score, result
}

func NewTreeModel(dd *models.DataDictionary, td *models.TransformationDictionary, model *models.TreeModel) (*TreeModel, error) {
	return &TreeModel{
		dd:    dd,
		td:    td,
		model: model,
	}, nil
}

func (m *TreeModel) Validate() error {
	for i, f := range m.model.MiningSchema.MiningFields {
		if f.UsageType == "" {
			f.UsageType = models.FieldUsageTypeActive
		}

		if f.Outliers == "" {
			f.Outliers = models.OutlierTreatmentMethodAsIs
		}

		if f.InvalidValueTreatment == "" {
			f.InvalidValueTreatment = models.InvalidValueTreatmentMethodReturnInvalid
		}

		m.model.MiningSchema.MiningFields[i] = f
	}

	m.validated = true
	return nil
}

func newNode(m *TreeModel, modelNode models.Node) node {
	n := node{
		score: modelNode.Score,
	}

	for _, child := range modelNode.Nodes {
		n.children = append(n.children, newNode(m, child))
	}

	switch p := modelNode.Predicate.(type) {
	case *models.True:
		n.test = func(evaluation.DataRow) predicateResult { return t }
	case *models.False:
		n.test = func(evaluation.DataRow) predicateResult { return f }
	case *models.SimplePredicate:
		n.test = func(input evaluation.DataRow) predicateResult {
			return evaluateSimplePredicate(p, input, m.fieldTypes)
		}
	case *models.CompoundPredicate:
		n.test = func(input evaluation.DataRow) predicateResult {
			return evaluateCompoundPredicate(p, input, m.fieldTypes)
		}
	case *models.SimpleSetPredicate:
		n.test = func(input evaluation.DataRow) predicateResult {
			return f
		}
	}

	return n
}

func getRawValue(dt models.DataType, val string) (interface{}, error) {
	switch dt {
	case models.DataTypeBoolean:
		return strconv.ParseBool(val)
	case models.DataTypeDouble:
		return strconv.ParseFloat(val, 64)
	case models.DataTypeFloat:
		return strconv.ParseFloat(val, 64)
	case models.DataTypeInteger:
		return strconv.ParseInt(val, 10, 64)
	case models.DataTypeString:
		return val, nil
	}

	return nil, errors.New("invalid data type")
}

//nolint:gocyclo
func evaluateSimplePredicate(p *models.SimplePredicate, input evaluation.DataRow, fieldTypes map[models.FieldName]models.DataType) predicateResult {
	predicateValueType, ok := fieldTypes[p.Field]
	if !ok {
		return f
	}

	rawPredicateValue, err := getRawValue(predicateValueType, p.Value)
	if err != nil {
		return f
	}

	val, ok := input[string(p.Field)]
	if !ok {
		if p.Operator == models.SimplePredicateOperatorIsMissing {
			return t
		}

		return f
	}

	switch p.Operator {
	case models.SimplePredicateOperatorIsNotMissing:
		return t
	case models.SimplePredicateOperatorEqual:
		if val.Raw() == rawPredicateValue {
			return t
		}
	case models.SimplePredicateOperatorNotEqual:
		if val.Raw() != rawPredicateValue {
			return t
		}
	case models.SimplePredicateOperatorGreaterOrEqual:
		switch predicateValueType {
		case models.DataTypeDouble:
			if val.Float64() >= rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeFloat:
			if val.Float64() >= rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeInteger:
			if val.Int64() >= rawPredicateValue.(int64) {
				return t
			}
		}
	case models.SimplePredicateOperatorGreaterThan:
		switch predicateValueType {
		case models.DataTypeDouble:
			if val.Float64() > rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeFloat:
			if val.Float64() > rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeInteger:
			if val.Int64() > rawPredicateValue.(int64) {
				return t
			}
		}
	case models.SimplePredicateOperatorLessOrEqual:
		switch predicateValueType {
		case models.DataTypeDouble:
			if val.Float64() <= rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeFloat:
			if val.Float64() <= rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeInteger:
			if val.Int64() <= rawPredicateValue.(int64) {
				return t
			}
		}
	case models.SimplePredicateOperatorLessThan:
		switch predicateValueType {
		case models.DataTypeDouble:
			if val.Float64() < rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeFloat:
			if val.Float64() < rawPredicateValue.(float64) {
				return t
			}
		case models.DataTypeInteger:
			if val.Int64() < rawPredicateValue.(int64) {
				return t
			}
		}
	}

	return f
}

func evaluateSimpleSetPredicate(p *models.SimpleSetPredicate, input evaluation.DataRow, fieldTypes map[models.FieldName]models.DataType) predicateResult {
	return f
}

func evaluateCompoundPredicate(p *models.CompoundPredicate, input evaluation.DataRow, fieldTypes map[models.FieldName]models.DataType) predicateResult {
	trueCount := 0
	unknownCount := 0

	surrogate := p.BooleanOperator == models.CompoundPredicateOperatorSurrogate

	for _, child := range p.Predicates {
		var val predicateResult
		switch c := child.(type) {
		case *models.SimplePredicate:
			val = evaluateSimplePredicate(c, input, fieldTypes)
		case *models.CompoundPredicate:
			val = evaluateCompoundPredicate(c, input, fieldTypes)
		case *models.SimpleSetPredicate:
			val = evaluateSimpleSetPredicate(c, input, fieldTypes)
		}

		if surrogate && val != u {
			return val
		}

		if val == t {
			trueCount++
		} else if val == u {
			unknownCount++
		}
	}

	switch p.BooleanOperator {
	case models.CompoundPredicateOperatorAnd:
		if unknownCount > 0 && unknownCount+trueCount == len(p.Predicates) {
			return u
		}

		if trueCount == len(p.Predicates) {
			return t
		}
	case models.CompoundPredicateOperatorOr:
		if unknownCount > 0 && unknownCount+trueCount < len(p.Predicates) {
			return u
		}

		if trueCount > 0 {
			return t
		}
	case models.CompoundPredicateOperatorXor:
		if unknownCount > 0 {
			return u
		}

		if trueCount%2 == 1 {
			return t
		}
	}

	return f
}

func (m *TreeModel) Compile() error {
	for _, f := range m.model.MiningSchema.MiningFields {
		if f.UsageType == models.FieldUsageTypeTarget {
			m.outputField = f.Name
		}
	}

	fieldTypes := map[models.FieldName]models.DataType{}

	for _, df := range m.dd.DataFields {
		fieldTypes[df.Name] = df.DataType
	}

	m.fieldTypes = fieldTypes

	m.root = newNode(m, m.model.Node)

	m.compiled = true
	return nil
}

func (m *TreeModel) Verify() error {
	if m.model.ModelVerification == nil {
		return nil
	}

	return evaluation.ErrNotImplemented
}

func (m *TreeModel) Evaluate(input evaluation.DataRow) (evaluation.DataRow, error) {
	var err error

	if !m.validated {
		return nil, evaluation.ErrNotValidated
	}

	if !m.compiled {
		return nil, evaluation.ErrNotCompiled
	}

	input, err = common.HandleInput(input, m.dd, m.td, &m.model.LocalTransformations, &m.model.MiningSchema)
	if err != nil {
		return nil, err
	}

	score, result := m.root.evaluate(input)

	if result == t {
		return evaluation.DataRow{
			string(m.outputField): evaluation.NewValue(score),
		}, nil
	}

	return nil, nil
}
