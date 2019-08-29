package common

import (
	"github.com/gopredict/pmml/evaluation"
	"github.com/gopredict/pmml/models"
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

	return nil
}
