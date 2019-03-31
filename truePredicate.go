package pmml

type TruePredicate struct {
}

func (*TruePredicate) Compile() error {
	return nil
}

func (*TruePredicate) Execute(features map[string]interface{}) PredicateResult {
	return True
}
