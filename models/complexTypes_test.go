package models_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArrayType(t *testing.T) {
	data := []byte(`<Array type="real" n="6">0.9134 0.8691 0.8002 0.5389 0.2261 0.1492</Array>`)

	var item models.ArrayType

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	n := 6

	assert.Equal(t, models.ArrayType{
		Type:     "real",
		N:        &n,
		RawValue: "0.9134 0.8691 0.8002 0.5389 0.2261 0.1492",
	}, item)
}

func TestCountTableType(t *testing.T) {
	data := []byte(`
	<CountTable sample="262">
		<FieldValueCount field="bin" count="100" value="bin1"/>
		<FieldValueCount field="bin" count="150" value="bin2"/>
		<FieldValueCount field="bin" count="10" value="bin3"/>
		<FieldValueCount field="bin" count="2" value="bin4"/>
	</CountTable>
`)

	var item models.CountTableType

	err := xml.Unmarshal(data, &item)
	require.NoError(t, err)

	sample := models.Number(262)

	assert.Equal(t, models.CountTableType{
		Sample: &sample,
		FieldValueCounts: []models.FieldValueCount{
			models.FieldValueCount{
				Field: models.FieldName("bin"),
				Count: models.Number(100),
				Value: "bin1",
			},
			models.FieldValueCount{
				Field: models.FieldName("bin"),
				Count: models.Number(150),
				Value: "bin2",
			},
			models.FieldValueCount{
				Field: models.FieldName("bin"),
				Count: models.Number(10),
				Value: "bin3",
			},
			models.FieldValueCount{
				Field: models.FieldName("bin"),
				Count: models.Number(2),
				Value: "bin4",
			},
		},
	}, item)
}
