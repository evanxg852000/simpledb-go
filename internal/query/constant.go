package query

import (
	"fmt"
)

type Constant struct {
	value any
}

func NewConstant(value any) Constant {
	return Constant{value}
}

func (c *Constant) AsInt() int64 {
	return c.value.(int64)
}

func (c *Constant) AsString() string {
	return c.value.(string)
}

func (c *Constant) String() string {
	return fmt.Sprint(c.value)
}

func (c *Constant) Equals(other Constant) bool {
	switch c.value.(type) {
	case int64:
		if other_value, ok := other.value.(int64); ok {
			return c.AsInt() == other_value
		}
	case string:
		if other_value, ok := other.value.(string); ok {
			return c.AsString() == other_value
		}
	}
	return false
}
