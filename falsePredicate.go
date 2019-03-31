package pmml

type FalsePredicate struct {
}

func (*FalsePredicate) Compile() error {
	return nil
}

func (*FalsePredicate) Execute(features map[string]interface{}) PredicateResult {
	return False
}
