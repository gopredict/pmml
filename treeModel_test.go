package pmml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml"
)

func TestTreeModel(t *testing.T) {
	data := []byte(`
	<TreeModel modelName="golfing" functionName="classification">
		<MiningSchema>
			<MiningField name="temperature"/>
			<MiningField name="humidity"/>
			<MiningField name="windy"/>
			<MiningField name="outlook"/>
			<MiningField name="whatIdo" usageType="target"/>
		</MiningSchema>
		<Node score="will play">
			<True/>
			<Node score="will play">
				<SimplePredicate field="outlook" operator="equal" value="sunny"/>
				<Node score="will play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="temperature" operator="lessThan" value="90"/>
						<SimplePredicate field="temperature" operator="greaterThan" value="50"/>
					</CompoundPredicate>
					<Node score="will play">
						<SimplePredicate field="humidity" operator="lessThan" value="80"/>
					</Node>
					<Node score="no play">
						<SimplePredicate field="humidity" operator="greaterOrEqual" value="80"/>
					</Node>
				</Node>
				<Node score="no play">
					<CompoundPredicate booleanOperator="or">
						<SimplePredicate field="temperature" operator="greaterOrEqual" value="90"/>
						<SimplePredicate field="temperature" operator="lessOrEqual" value="50"/>
					</CompoundPredicate>
				</Node>
			</Node>
			<Node score="may play">
				<CompoundPredicate booleanOperator="or">
					<SimplePredicate field="outlook" operator="equal" value="overcast"/>
					<SimplePredicate field="outlook" operator="equal" value="rain"/>
				</CompoundPredicate>
				<Node score="may play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="temperature" operator="greaterThan" value="60"/>
						<SimplePredicate field="temperature" operator="lessThan" value="100"/>
						<SimplePredicate field="outlook" operator="equal" value="overcast"/>
						<SimplePredicate field="humidity" operator="lessThan" value="70"/>
						<SimplePredicate field="windy" operator="equal" value="false"/>
					</CompoundPredicate>
				</Node>
				<Node score="no play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="outlook" operator="equal" value="rain"/>
						<SimplePredicate field="humidity" operator="lessThan" value="70"/>
					</CompoundPredicate>
				</Node>
			</Node>
		</Node>
	</TreeModel>`)

	var m pmml.TreeModel

	err := xml.Unmarshal(data, &m)
	if err != nil {
		t.Error(err.Error())
	}

	err = m.Compile()
	if err != nil {
		t.Error(err.Error())
	}

	if len(m.MiningSchema.MiningFields) != 5 {
		t.Errorf("expected 5 values but got %d", len(m.MiningSchema.MiningFields))
	}

	if len(m.Node.Nodes) != 2 {
		t.Errorf("expected 2 values but got %d", len(m.Node.Nodes))
	}

	type testCase struct {
		name           string
		features       map[string]interface{}
		expectedResult pmml.PredicateResult
		expectedScore  string
	}

	testCases := []testCase{
		testCase{
			name: "May Play",
			features: map[string]interface{}{
				"temperature": 75,
				"humidity":    55,
				"windy":       false,
				"outlook":     "overcast",
			},
			expectedResult: pmml.True,
			expectedScore:  "may play",
		},
		testCase{
			name: "No Play",
			features: map[string]interface{}{
				"temperature": 75,
				"humidity":    85,
				"windy":       false,
				"outlook":     "sunny",
			},
			expectedResult: pmml.True,
			expectedScore:  "no play",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, score, _ := m.Node.Execute(tc.features)
			if result != tc.expectedResult {
				t.Errorf("expected %v but got %v", tc.expectedResult, result)
			}
			if score != tc.expectedScore {
				t.Errorf("expected %v but got %v", tc.expectedScore, score)
			}
		})
	}
}

func BenchmarkTreeModel_Simple(b *testing.B) {
	data := []byte(`
	<TreeModel modelName="golfing" functionName="classification">
		<MiningSchema>
			<MiningField name="temperature"/>
			<MiningField name="humidity"/>
			<MiningField name="windy"/>
			<MiningField name="outlook"/>
			<MiningField name="whatIdo" usageType="target"/>
		</MiningSchema>
		<Node score="will play">
			<True/>
			<Node score="will play">
				<SimplePredicate field="outlook" operator="equal" value="sunny"/>
				<Node score="will play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="temperature" operator="lessThan" value="90"/>
						<SimplePredicate field="temperature" operator="greaterThan" value="50"/>
					</CompoundPredicate>
					<Node score="will play">
						<SimplePredicate field="humidity" operator="lessThan" value="80"/>
					</Node>
					<Node score="no play">
						<SimplePredicate field="humidity" operator="greaterOrEqual" value="80"/>
					</Node>
				</Node>
				<Node score="no play">
					<CompoundPredicate booleanOperator="or">
						<SimplePredicate field="temperature" operator="greaterOrEqual" value="90"/>
						<SimplePredicate field="temperature" operator="lessOrEqual" value="50"/>
					</CompoundPredicate>
				</Node>
			</Node>
			<Node score="may play">
				<CompoundPredicate booleanOperator="or">
					<SimplePredicate field="outlook" operator="equal" value="overcast"/>
					<SimplePredicate field="outlook" operator="equal" value="rain"/>
				</CompoundPredicate>
				<Node score="may play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="temperature" operator="greaterThan" value="60"/>
						<SimplePredicate field="temperature" operator="lessThan" value="100"/>
						<SimplePredicate field="outlook" operator="equal" value="overcast"/>
						<SimplePredicate field="humidity" operator="lessThan" value="70"/>
						<SimplePredicate field="windy" operator="equal" value="false"/>
					</CompoundPredicate>
				</Node>
				<Node score="no play">
					<CompoundPredicate booleanOperator="and">
						<SimplePredicate field="outlook" operator="equal" value="rain"/>
						<SimplePredicate field="humidity" operator="lessThan" value="70"/>
					</CompoundPredicate>
				</Node>
			</Node>
		</Node>
	</TreeModel>`)

	var m pmml.TreeModel

	err := xml.Unmarshal(data, &m)
	if err != nil {
		b.Error(err.Error())
	}

	err = m.Compile()
	if err != nil {
		b.Error(err.Error())
	}

	features := []map[string]interface{}{
		map[string]interface{}{
			"temperature": 75,
			"humidity":    55,
			"windy":       false,
			"outlook":     "overcast",
		},
		map[string]interface{}{
			"temperature": 75,
			"humidity":    85,
			"windy":       false,
			"outlook":     "sunny",
		},
	}

	for n := 0; n < b.N; n++ {
		benchVal, _, _ = m.Node.Execute(features[n%len(features)])
	}
}

func TestTreeModelComplex(t *testing.T) {
	data := []byte(`
	<TreeModel modelName="golfing" functionName="classification" missingValueStrategy="weightedConfidence">
		<MiningSchema>
			<MiningField name="temperature"/>
			<MiningField name="humidity"/>
			<MiningField name="outlook"/>
			<MiningField name="whatIdo" usageType="target"/>
		</MiningSchema>
		<Node id="1" score="will play" recordCount="100" defaultChild="2">
			<True/>
			<ScoreDistribution value="will play" recordCount="60" confidence="0.6"/>
			<ScoreDistribution value="may play" recordCount="30" confidence="0.3"/>
			<ScoreDistribution value="no play" recordCount="10" confidence="0.1"/>
			<Node id="2" score="will play" recordCount="50" defaultChild="3">
				<SimplePredicate field="outlook" operator="equal" value="sunny"/>
				<ScoreDistribution value="will play" recordCount="40" confidence="0.8"/>
				<ScoreDistribution value="may play" recordCount="2" confidence="0.04"/>
				<ScoreDistribution value="no play" recordCount="8" confidence="0.16"/>
				<Node id="3" score="will play" recordCount="40">
					<CompoundPredicate booleanOperator="surrogate">
						<SimplePredicate field="temperature" operator="greaterOrEqual" value="50"/>
						<SimplePredicate field="humidity" operator="lessThan" value="80"/>
					</CompoundPredicate>
					<ScoreDistribution value="will play" recordCount="36" confidence="0.9"/>
					<ScoreDistribution value="may play" recordCount="2" confidence="0.05"/>
					<ScoreDistribution value="no play" recordCount="2" confidence="0.05"/>
				</Node>
				<Node id="4" score="no play" recordCount="10">
					<CompoundPredicate booleanOperator="surrogate">
						<SimplePredicate field="temperature" operator="lessThan" value="50"/>
						<SimplePredicate field="humidity" operator="greaterOrEqual" value="80"/>
					</CompoundPredicate>
					<ScoreDistribution value="will play" recordCount="4" confidence="0.4"/>
					<ScoreDistribution value="may play" recordCount="0" confidence="0.0"/>
					<ScoreDistribution value="no play" recordCount="6" confidence="0.6"/>
				</Node>
			</Node>
			<Node id="5" score="may play" recordCount="50">
				<CompoundPredicate booleanOperator="or">
					<SimplePredicate field="outlook" operator="equal" value="overcast"/>
					<SimplePredicate field="outlook" operator="equal" value="rain"/>
				</CompoundPredicate>
				<ScoreDistribution value="will play" recordCount="20" confidence="0.4"/>
				<ScoreDistribution value="may play" recordCount="28" confidence="0.56"/>
				<ScoreDistribution value="no play" recordCount="2" confidence="0.04"/>
			</Node>
		</Node>
	</TreeModel>`)

	var m pmml.TreeModel

	err := xml.Unmarshal(data, &m)
	if err != nil {
		t.Error(err.Error())
	}

	err = m.Compile()
	if err != nil {
		t.Error(err.Error())
	}

	if len(m.MiningSchema.MiningFields) != 4 {
		t.Errorf("expected 4 values but got %d", len(m.MiningSchema.MiningFields))
	}

	if len(m.Node.Nodes) != 2 {
		t.Errorf("expected 2 values but got %d", len(m.Node.Nodes))
	}

	type testCase struct {
		name               string
		features           map[string]interface{}
		expectedResult     pmml.PredicateResult
		expectedScore      string
		expectedConfidence float64
	}

	testCases := []testCase{
		testCase{
			name: "May Play",
			features: map[string]interface{}{
				"temperature": 45,
				"humidity":    60,
				"outlook":     "sunny",
			},
			expectedResult:     pmml.True,
			expectedScore:      "no play",
			expectedConfidence: 0.6,
		},
		/*testCase{
			name: "May Play",
			features: map[string]interface{}{
				"outlook": "sunny",
			},
			expectedResult:     pmml.True,
			expectedScore:      "will play",
			expectedConfidence: 0.8,
		},*/
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, score, confidence := m.Node.Execute(tc.features)
			if result != tc.expectedResult {
				t.Errorf("expected %v but got %v", tc.expectedResult, result)
			}
			if score != tc.expectedScore {
				t.Errorf("expected %v but got %v", tc.expectedScore, score)
			}
			if confidence != tc.expectedConfidence {
				t.Errorf("expected %v but got %v", tc.expectedConfidence, confidence)
			}
		})
	}
}
