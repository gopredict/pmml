package pmml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml"
)

const simpleSetPredicate1 = `<SimpleSetPredicate field="f1" booleanOperator="isIn">
                            <Array n="6" type="string">v1 v2 v3</Array>
                            </SimpleSetPredicate>`
const simpleSetPredicate2 = `<SimpleSetPredicate field="f2" booleanOperator="isIn">
							  <Array n="6" type="string">"Missing"   "No Match"</Array>
							  </SimpleSetPredicate>`

var simpleSetPredicateTests = []struct {
	predicate []byte
	features  map[string]interface{}
	expected  pmml.PredicateResult
}{
	{[]byte(simpleSetPredicate1),
		map[string]interface{}{"f1": "v3"},
		pmml.True},
	{[]byte(simpleSetPredicate1),
		map[string]interface{}{"f1": "v4"},
		pmml.False},
	{[]byte(simpleSetPredicate2),
		map[string]interface{}{"f2": "No Match"},
		pmml.True},
	{[]byte(simpleSetPredicate2),
		map[string]interface{}{"f2": "Match"},
		pmml.False},
}

func TestSimpleSetPredicate(t *testing.T) {

	for _, tt := range simpleSetPredicateTests {
		var predicate pmml.SimpleSetPredicate
		xml.Unmarshal(tt.predicate, &predicate)

		predicate.Compile()

		actual := predicate.Execute(tt.features)
		if actual != tt.expected {
			t.Errorf("Predicate: %s %s %s, Feature value : %s, expected %v, actual %v",
				predicate.Field,
				predicate.Operator,
				predicate.Values,
				tt.features[predicate.Field],
				tt.expected,
				actual)
		}
	}
}

func BenchmarkSimpleSetPredicate(b *testing.B) {
	var predicate pmml.SimpleSetPredicate
	xml.Unmarshal([]byte(simpleSetPredicate1), &predicate)

	err := predicate.Compile()
	if err != nil {
		b.Error(err.Error())
		return
	}

	m := map[string]interface{}{"f1": "v3"}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		benchVal = predicate.Execute(m)
	}
}
