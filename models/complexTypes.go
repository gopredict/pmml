package models

import (
	"strconv"

	"github.com/mattn/go-shellwords"
	"github.com/pkg/errors"
)

type ArrayTypeType string

const (
	ArrayTypeTypeInt    = ArrayTypeType("int")
	ArrayTypeTypeReal   = ArrayTypeType("real")
	ArrayTypeTypeString = ArrayTypeType("string")
)

/*
  <xs:complexType mixed="true" name="ArrayType">
    <xs:attribute name="n" type="INT-NUMBER" use="optional"/>
    <xs:attribute name="type" use="required">
      <xs:simpleType>
        <xs:restriction base="xs:string">
          <xs:enumeration value="int"/>
          <xs:enumeration value="real"/>
          <xs:enumeration value="string"/>
        </xs:restriction>
      </xs:simpleType>
    </xs:attribute>
  </xs:complexType>
*/
type ArrayType struct {
	N    *int          `xml:"n,attr"`
	Type ArrayTypeType `xml:"type,attr"`

	RawValue string `xml:",innerxml"`

	s []string
	i []int64
	f []float64
}

func (at *ArrayType) Strings() ([]string, error) {
	if at.Type != ArrayTypeTypeString {
		return nil, errors.New("not a string array")
	}

	if at.s != nil {
		return at.s, nil
	}

	var err error
	at.s, err = shellwords.Parse(at.RawValue)
	return at.s, err
}

func (at *ArrayType) Int64s() ([]int64, error) {
	if at.Type != ArrayTypeTypeInt {
		return nil, errors.New("not an int array")
	}

	if at.i != nil {
		return at.i, nil
	}

	ints, err := shellwords.Parse(at.RawValue)
	if err != nil {
		return nil, err
	}

	result := []int64{}

	for i := range ints {
		val, err := strconv.ParseInt(ints[i], 10, 64)
		if err != nil {
			return nil, err
		}

		result = append(result, val)
	}

	at.i = result

	return at.i, err
}

func (at *ArrayType) Float64s() ([]float64, error) {
	if at.Type != ArrayTypeTypeReal {
		return nil, errors.New("not a float array")
	}

	if at.f != nil {
		return at.f, nil
	}

	ints, err := shellwords.Parse(at.RawValue)
	if err != nil {
		return nil, err
	}

	result := []float64{}

	for i := range ints {
		val, err := strconv.ParseFloat(ints[i], 64)
		if err != nil {
			return nil, err
		}

		result = append(result, val)
	}

	at.f = result

	return at.f, err
}

/*
  <xs:complexType name="COUNT-TABLE-TYPE">
    <xs:attribute name="sample" type="NUMBER" use="optional"/>
    <xs:sequence>
      <xs:choice>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="FieldValue"/>
        <xs:element maxOccurs="unbounded" minOccurs="1" ref="FieldValueCount"/>
      </xs:choice>
      <xs:element maxOccurs="unbounded" minOccurs="0" ref="Extension"/>
    </xs:sequence>
  </xs:complexType>
*/
type CountTableType struct {
	Sample *Number `xml:"sample,attr"`

	FieldValues      []FieldValue      `xml:"FieldValue"`
	FieldValueCounts []FieldValueCount `xml:"FieldValueCount"`

	Extensions []Extension `xml:"Extension"`
}
