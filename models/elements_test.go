package models_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAnnotation(t *testing.T) {
	data := []byte(`
	<Annotation>This is a churn model for 1999 customers who
		<Extension name="author">John Doe</Extension>
	</Annotation>`)

	var item models.Annotation

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, models.Annotation{
		Data: "This is a churn model for 1999 customers who\n\t\t\n\t",
		Extensions: []models.Extension{
			models.Extension{
				Name:     models.String("author"),
				Contents: "John Doe",
			},
		},
	}, item)
}

func RealNumber(v models.RealNumber) *models.RealNumber {
	return &v
}

func TestARDSquaredExponentialKernel(t *testing.T) {
	data := []byte(`
	<ARDSquaredExponentialKernel gamma="2.4890" noiseVariance="0.0110">
		<Lambda>
			<Array n="2" type="real">1.5164 59.3113</Array>
		</Lambda>
	</ARDSquaredExponentialKernel>`)

	var item models.ARDSquaredExponentialKernel

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, RealNumber(2.4890), item.Gamma)
	assert.Equal(t, RealNumber(0.0110), item.NoiseVariance)

	vals, err := item.Lambda.Array.Float64s()
	require.NoError(t, err)

	assert.Equal(t, []float64{1.5164, 59.3113}, vals)
}

func TestAggregate(t *testing.T) {
	data := []byte(`<Aggregate field="item" function="multiset" groupField="transaction"/>`)

	var item models.Aggregate

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, models.Aggregate{
		Field:      models.FieldName("item"),
		Function:   models.AggregateFunctionTypeMultiset,
		GroupField: "transaction",
	}, item)
}
