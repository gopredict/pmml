package treemodel_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml/evaluation"

	"github.com/gopredict/pmml/evaluation/treemodel"
	"github.com/gopredict/pmml/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	data = []byte(`
<PMML xmlns="http://www.dmg.org/PMML-4_4" version="4.4">
	<Header copyright="www.dmg.org" description="A very small binary tree model to show structure."/>
	<DataDictionary numberOfFields="5">
	  <DataField name="temperature" optype="continuous" dataType="double"/>
	  <DataField name="humidity" optype="continuous" dataType="double"/>
	  <DataField name="windy" optype="categorical" dataType="string">
		<Value value="true"/>
		<Value value="false"/>
	  </DataField>
	  <DataField name="outlook" optype="categorical" dataType="string">
		<Value value="sunny"/>
		<Value value="overcast"/>
		<Value value="rain"/>
	  </DataField>
	  <DataField name="whatIdo" optype="categorical" dataType="string">
		<Value value="will play"/>
		<Value value="may play"/>
		<Value value="no play"/>
	  </DataField>
	</DataDictionary>
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
	</TreeModel>
  </PMML>`)
)

func TestTreeModel(t *testing.T) {
	var doc models.PMML

	err := xml.NewDecoder(bytes.NewReader(data)).Decode(&doc)
	require.NoError(t, err)
	require.Len(t, doc.Models, 1)

	m, ok := doc.Models[0].(*models.TreeModel)
	require.True(t, ok)

	assert.Equal(t, "golfing", m.ModelName)
	assert.Equal(t, models.MiningFunctionClassification, m.FunctionName)

	tm, err := treemodel.NewTreeModel(&doc.DataDictionary, doc.TransformationDictionary, m)
	require.NoError(t, err)

	err = tm.Validate()
	require.NoError(t, err)

	err = tm.Verify()
	require.NoError(t, err)

	err = tm.Compile()
	require.NoError(t, err)

	// temperature=75, humidity=55, windy="false", outlook="overcast"
	out, err := tm.Evaluate(evaluation.DataRow{
		"temperature": evaluation.NewValue(75),
		"humidity":    evaluation.NewValue(55),
		"windy":       evaluation.NewValue(false),
		"outlook":     evaluation.NewValue("overcast"),
	})
	require.NoError(t, err)

	assert.Len(t, out, 1)

	assert.Equal(t, out["whatIdo"], evaluation.NewValue("may play"))
}
