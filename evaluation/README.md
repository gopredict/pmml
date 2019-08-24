# evaluation
--
    import "github.com/gopredict/pmml/evaluation"


## Usage

```go
const (
	ErrNotImplemented = Error("evaluation: not implemented")
	ErrNotValidated   = Error("evaluation: not validated")
	ErrNotCompiled    = Error("evaluation: not compiled")
)
```

#### type DataRow

```go
type DataRow map[string]Value
```


#### type Error

```go
type Error string
```


#### func (Error) Error

```go
func (err Error) Error() string
```

#### type Value

```go
type Value struct {
}
```


#### func  NewValue

```go
func NewValue(val interface{}) Value
```

#### func (Value) Float64

```go
func (v Value) Float64() float64
```

#### func (Value) Int64

```go
func (v Value) Int64() int64
```

#### func (Value) Raw

```go
func (v Value) Raw() interface{}
```

#### func (Value) String

```go
func (v Value) String() string
```
