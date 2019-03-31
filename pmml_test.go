package pmml_test

import (
	"encoding/xml"
	"testing"

	"github.com/gopredict/pmml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPMML(t *testing.T) {
	doc := []byte(`<?xml version="1.0"?>
	<PMML version="4.3" xmlns="http://www.dmg.org/PMML-4_3" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="http://www.dmg.org/PMML-4_3 http://www.dmg.org/pmml/v4-3/pmml-4-3.xsd">
		<Header copyright="Copyright (c) 2019 gopredict" description="Random Forest Tree Model">
			<Extension name="user" value="rickbassham" extender="Rattle/PMML"/>
			<Application name="Rattle/PMML" version="1.4"/>
			<Timestamp>2017-05-18T15:39:02Z</Timestamp>
			<Annotation>First time using PMML.</Annotation>
			<Annotation>Made some changes.</Annotation>
			<Annotation>This is a churn model for 1999 customers who<Extension name="author">John Doe</Extension></Annotation>
		</Header>
	</PMML>`)

	var pmmldoc pmml.PMML

	err := xml.Unmarshal(doc, &pmmldoc)
	require.NoError(t, err)

	assert.Equal(t, "Copyright (c) 2019 gopredict", pmmldoc.Header.Copyright)
	assert.Equal(t, "Random Forest Tree Model", pmmldoc.Header.Description)
	assert.Equal(t, "Rattle/PMML", pmmldoc.Header.Application.Name)
	assert.Equal(t, "1.4", pmmldoc.Header.Application.Version)
	assert.Equal(t, []pmml.Annotation{
		pmml.Annotation{Value: "First time using PMML."},
		pmml.Annotation{Value: "Made some changes."},
		pmml.Annotation{Value: "This is a churn model for 1999 customers who"},
	}, pmmldoc.Header.Annotations)
}
