package models_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RealNumber(v models.RealNumber) *models.RealNumber {
	return &v
}

func ProbNumber(v models.ProbNumber) *models.ProbNumber {
	return &v
}

func Number(v models.Number) *models.Number {
	return &v
}

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

func TestAlternate(t *testing.T) {
	data := []byte(`
	<Alternate>
		<GaussianDistribution mean="460.4" variance="39.2"/>
	</Alternate>`)

	var item models.Alternate

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, models.Alternate{
		Distribution: &models.GaussianDistribution{
			Mean:     models.RealNumber(460.4),
			Variance: models.RealNumber(39.2),
		},
	}, item)
}

func TestAnova(t *testing.T) {
	data := []byte(`
	<Anova>
		<AnovaRow type="Model" sumOfSquares="21389708" degreesOfFreedom="2" meanOfSquares="7129903" fValue="0.85" pValue="0.47" />
		<AnovaRow type="Error" sumOfSquares="809210617" degreesOfFreedom="98" meanOfSquares="8342377"/>
		<AnovaRow type="Total" sumOfSquares="830600325" degreesOfFreedom="100" />
	</Anova>`)

	var item models.Anova

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, models.Anova{
		Rows: []models.AnovaRow{
			models.AnovaRow{
				Type:             models.AnovaRowTypeModel,
				SumOfSquares:     models.Number(21389708),
				DegreesOfFreedom: models.Number(2),
				MeanOfSquares:    Number(7129903),
				FValue:           Number(0.85),
				PValue:           ProbNumber(0.47),
			},
			models.AnovaRow{
				Type:             models.AnovaRowTypeError,
				SumOfSquares:     models.Number(809210617),
				DegreesOfFreedom: models.Number(98),
				MeanOfSquares:    Number(8342377),
			},
			models.AnovaRow{
				Type:             models.AnovaRowTypeTotal,
				SumOfSquares:     models.Number(830600325),
				DegreesOfFreedom: models.Number(100),
			},
		},
	}, item)
}

func TestApply(t *testing.T) {
	data := []byte(`
	<Apply function="if">
		<Apply function="greaterThan">
			<FieldRef field="score"/>
			<Constant dataType="double">1</Constant>
		</Apply>
		<!-- Then case -->
		<Constant dataType="string">True</Constant>
		<!-- Else case -->
		<Constant dataType="string">False</Constant>
	</Apply>`)

	var item models.Apply

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	assert.Equal(t, models.Apply{
		Function:              "if",
		InvalidValueTreatment: models.InvalidValueTreatmentMethodReturnInvalid,
		Expressions: []models.Expression{
			&models.Apply{
				Function:              "greaterThan",
				InvalidValueTreatment: models.InvalidValueTreatmentMethodReturnInvalid,
				Expressions: []models.Expression{
					&models.FieldRef{
						Field: models.FieldName("score"),
					},
					&models.Constant{
						DataType: models.DataTypeDouble,
						Value:    "1",
					},
				},
			},
			&models.Constant{
				DataType: models.DataTypeString,
				Value:    "True",
			},
			&models.Constant{
				DataType: models.DataTypeString,
				Value:    "False",
			},
		},
	}, item)
}
