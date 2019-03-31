package pmml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml"
)

var simplePredicateTests = []struct {
	name      string
	predicate []byte
	features  map[string]interface{}
	expected  pmml.PredicateResult
}{
	{"FloatEqual", []byte(`<SimplePredicate field="f33" operator="equal" value="18.85"/>`),
		map[string]interface{}{"f33": 18.850},
		pmml.True},
	{"FloatLessOrEqual1", []byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]interface{}{"f33": 18.84},
		pmml.True},
	{"FloatLessOrEqual2", []byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]interface{}{"f33": 18.86},
		pmml.False},
	{"FloatLessOrEqual3", []byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`),
		map[string]interface{}{"f33": "18.84"},
		pmml.False},
	{"FloatMissing1", []byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]interface{}{"f33": 18.86},
		pmml.False},
	{"FloatMissing2", []byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]interface{}{},
		pmml.True},
	{"FloatMissing3", []byte(`<SimplePredicate field="f33" operator="isMissing" value="18.85"/>`),
		map[string]interface{}{"f33": ""},
		pmml.True},
}

func TestSimplePredicate(t *testing.T) {
	for i := range simplePredicateTests {
		tt := simplePredicateTests[i]
		t.Run(tt.name, func(t *testing.T) {
			var predicate pmml.SimplePredicate
			err := xml.Unmarshal(tt.predicate, &predicate)
			if err != nil {
				t.Error(err.Error())
				return
			}

			err = predicate.Compile()
			if err != nil {
				t.Error(err.Error())
				return
			}

			actual := predicate.Execute(tt.features)
			if actual != tt.expected {
				t.Errorf("Predicate: %s %s %s, Feature value : %#v, expected %v, actual %v",
					predicate.Field,
					predicate.Operator,
					predicate.Value,
					tt.features[predicate.Field],
					tt.expected,
					actual)
			}
		})
	}
}

var benchVal pmml.PredicateResult

func BenchmarkSimplePredicate(b *testing.B) {
	var predicate pmml.SimplePredicate
	err := xml.Unmarshal([]byte(`<SimplePredicate field="f33" operator="lessOrEqual" value="18.85"/>`), &predicate)
	if err != nil {
		b.Error(err.Error())
		return
	}

	err = predicate.Compile()
	if err != nil {
		b.Error(err.Error())
		return
	}

	m := map[string]interface{}{"f33": 18.84}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		benchVal = predicate.Execute(m)
	}
}
