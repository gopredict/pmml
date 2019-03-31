package pmml

type Predicate interface {
	Compile() error
	Execute(features map[string]interface{}) PredicateResult
}

type PredicateResult int

const (
	Unknown PredicateResult = iota
	True
	False
)
