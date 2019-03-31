package pmml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml"
)

func TestCompoundPredicate(t *testing.T) {
	data := []byte(`
	<CompoundPredicate booleanOperator="surrogate">
		<CompoundPredicate booleanOperator="and">
			<SimplePredicate field="temperature" operator="lessThan" value="90"/>
			<SimplePredicate field="temperature" operator="greaterThan" value="50"/>
		</CompoundPredicate>
		<SimplePredicate field="humidity" operator="greaterOrEqual" value="80"/>
		<False/>
	</CompoundPredicate>`)

	var p pmml.CompoundPredicate

	err := xml.Unmarshal(data, &p)
	if err != nil {
		t.Error(err.Error())
	}

	if len(p.Values) != 3 {
		t.Errorf("expected 3 values but got %d", len(p.Values))
	}

	err = p.Compile()
	if err != nil {
		t.Error(err.Error())
	}

	type testCase struct {
		name           string
		features       map[string]interface{}
		expectedResult pmml.PredicateResult
	}

	testCases := []testCase{
		testCase{
			name: "Temp Set True",
			features: map[string]interface{}{
				"temperature": 80,
			},
			expectedResult: pmml.True,
		},
		testCase{
			name: "Temp Set False",
			features: map[string]interface{}{
				"temperature": 95,
			},
			expectedResult: pmml.False,
		},
		testCase{
			name: "Humidity Set True",
			features: map[string]interface{}{
				"humidity": 95,
			},
			expectedResult: pmml.True,
		},
		testCase{
			name: "Humidity Set False",
			features: map[string]interface{}{
				"humidity": 75,
			},
			expectedResult: pmml.False,
		},
		testCase{
			name:           "Empty",
			features:       map[string]interface{}{},
			expectedResult: pmml.False,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := p.Execute(tc.features)
			if result != tc.expectedResult {
				t.Errorf("expected %v but got %v", tc.expectedResult, result)
			}
		})
	}
}
