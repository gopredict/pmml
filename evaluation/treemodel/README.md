# treemodel
--
    import "github.com/gopredict/pmml/evaluation/treemodel"


## Usage

#### type TreeModel

```go
type TreeModel struct {
}
```


#### func  NewTreeModel

```go
func NewTreeModel(dd *models.DataDictionary, td *models.TransformationDictionary, model *models.TreeModel) (*TreeModel, error)
```

#### func (*TreeModel) Compile

```go
func (m *TreeModel) Compile() error
```

#### func (*TreeModel) Evaluate

```go
func (m *TreeModel) Evaluate(input evaluation.DataRow) (evaluation.DataRow, error)
```

#### func (*TreeModel) Validate

```go
func (m *TreeModel) Validate() error
```

#### func (*TreeModel) Verify

```go
func (m *TreeModel) Verify() error
```
