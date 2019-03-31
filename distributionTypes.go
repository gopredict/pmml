package pmml

type AnyDistribution struct {
}

type GaussianDistribution struct {
}

type PoissonDistribution struct {
}

type UniformDistribution struct {
}

type CountTable struct {
}

type NormalizedCountTable struct {
}

type FieldRef struct { // http://dmg.org/pmml/v4-3/Transformations.html#xsdElement_FieldRef
	Field        string `xml:"field,attr"`
	MapMissingTo string `xml:"mapMissingTo"`
}
