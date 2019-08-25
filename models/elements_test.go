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
