package hb

import (
	"fmt"
	"html"
)

type AttributeVisitor interface {
	Attr(name string, value any) error
}

type Attributes []func(AttributeVisitor) error

func (as Attributes) Accept(visitor AttributeVisitor) error {
	for _, a := range as {
		if err := a(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (as Attributes) Append(rest ...Attributes) Attributes {
	ret := as
	for _, r := range rest {
		ret = append(ret, r...)
	}
	return ret
}

func Attrs(attrs ...Attributes) Attributes {
	if len(attrs) == 0 {
		return nil
	}
	return attrs[0].Append(attrs[1:]...)
}

func Attr(name string, value any) Attributes {
	switch v := value.(type) {
	case bool:
		if !v {
			return nil
		}
		return Attr_(name)
	default:
		return Attributes{func(v AttributeVisitor) error {
			return v.Attr(name, value)
		}}
	}
}

func Attr_(name string) Attributes {
	return Attributes{func(v AttributeVisitor) error {
		return v.Attr(name, "")
	}}
}

func AttrB(name string, value bool) Attributes {
	return Attr(name, value)
}

func renderStr(value any) string {
	switch v := value.(type) {
	case string:
		return html.EscapeString(v)
	case []byte:
		return html.EscapeString(string(v))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%v", v)
	case fmt.Stringer:
		return html.EscapeString(v.String())
	default:
		panic(fmt.Sprintf("unsupported attribute type: %Text", v))
	}
}
