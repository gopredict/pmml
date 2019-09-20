# pmml
Predictive Model Markup Language package for go

This project is still in its infancy and should NOT be used at this time. If you
would like to contribute, please reach out.

## Resources

* http://dmg.org/pmml/pmml-v4-3.html
* http://dmg.org/pmml/v4-3/pmml-4-3.xsd
* http://dmg.org/pmml/products.html
* https://github.com/autodeployai/pmml4s

## Goals

The primary goal of this project is to bring predictive modeling to go, with all
of it's performance benefits.

Recent benchmarks show a simple tree model executing in 124 ns per prediction on
a 2018 MacBook Pro (2.6GHz Core i7 with 16GB 2400 MHz DDR4).

```
goos: darwin
goarch: amd64
pkg: github.com/gopredict/pmml
BenchmarkSimplePredicate-12             100000000               12.6 ns/op             0 B/op          0 allocs/op
BenchmarkSimpleSetPredicate-12          100000000               23.7 ns/op             0 B/op          0 allocs/op
BenchmarkTreeModel_Simple-12            10000000               124 ns/op               0 B/op          0 allocs/op
```
