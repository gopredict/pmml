package models

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
