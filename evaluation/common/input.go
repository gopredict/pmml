package common

import (
	"github.com/gopredict/pmml/evaluation"
	"github.com/gopredict/pmml/models"
	"github.com/pkg/errors"
)

func HandleInput(in evaluation.DataRow, dd *models.DataDictionary, td *models.TransformationDictionary, lt *models.LocalTransformations, ms *models.MiningSchema) (evaluation.DataRow, error) {

	err := validate(in, ms)
	if err != nil {
		return nil, err
	}

	return in, nil
}

func validate(in evaluation.DataRow, ms *models.MiningSchema) error {
	if ms == nil {
		return nil
	}

	for _, field := range ms.MiningFields {
		if field.UsageType == models.FieldUsageTypeActive {
			_, ok := in[string(field.Name)]

			if !ok {
				return errors.Errorf("pmml: missing field (%s)", field.Name)
			}
		}
	}

	return nil
}
