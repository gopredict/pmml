package evaluation

import "fmt"

type stringer interface {
	String() string
}

type Value struct {
	val interface{}
}

func NewValue(val interface{}) Value {
	return Value{val: val}
}

func (v Value) Raw() interface{} {
	return v.val
}

func (v Value) Float64() float64 {
	if f, ok := v.val.(float64); ok {
		return f
	}

	if f, ok := v.val.(float32); ok {
		return float64(f)
	}

	return 0.0
}

func (v Value) Int64() int64 {
	if f, ok := v.val.(int); ok {
		return int64(f)
	}
	if f, ok := v.val.(int8); ok {
		return int64(f)
	}
	if f, ok := v.val.(int16); ok {
		return int64(f)
	}
	if f, ok := v.val.(int32); ok {
		return int64(f)
	}
	if f, ok := v.val.(int64); ok {
		return int64(f)
	}
	if f, ok := v.val.(uint); ok {
		return int64(f)
	}
	if f, ok := v.val.(uint8); ok {
		return int64(f)
	}
	if f, ok := v.val.(uint16); ok {
		return int64(f)
	}
	if f, ok := v.val.(uint32); ok {
		return int64(f)
	}
	if f, ok := v.val.(uint64); ok {
		return int64(f)
	}

	return 0
}

func (v Value) String() string {
	if s, ok := v.val.(string); ok {
		return s
	}

	if s, ok := v.val.(stringer); ok {
		return s.String()
	}

	return fmt.Sprintf("%v", v.val)
}

type DataRow map[string]Value
