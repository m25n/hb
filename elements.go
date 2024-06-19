package hb

import (
	"io"
)

type ElementVisitor interface {
	OpenTagStart(string) error
	AttributeVisitor
	OpenTagEnd() error
	OpenTagSelfClose() error
	Text(string) error
	Raw(string) error
	CloseTag(string) error
}

type HTML []func(v ElementVisitor) error

func (hs HTML) RenderHTML(w io.Writer) (n int, err error) {
	ev := newHTMLRenderer(w)
	err = hs.Accept(ev)
	return ev.n, err
}

func (hs HTML) Accept(visitor ElementVisitor) error {
	for _, h := range hs {
		if err := h(visitor); err != nil {
			return err
		}
	}
	return nil
}

func (hs HTML) Append(rest ...HTML) HTML {
	ret := hs
	for _, es := range rest {
		ret = append(ret, es...)
	}
	return ret
}

func Fragment(hs ...HTML) HTML {
	if len(hs) == 0 {
		return nil
	}
	return hs[0].Append(hs[1:]...)
}

func FragmentGen(do func(yield func(HTML))) HTML {
	var f HTML
	yield := func(h HTML) { f = f.Append(h) }
	do(yield)
	return f
}

func Tag(tagName string, attributes Attributes, children ...HTML) HTML {
	return HTML{func(v ElementVisitor) error {
		return visitElement(v, tagName, attributes, Fragment(children...))
	}}
}

func Tag_(tagName string, attributes Attributes) HTML {
	return HTML{func(v ElementVisitor) error { return visitElementSelfClose(v, tagName, attributes) }}
}

func Text(value any) HTML {
	return HTML{func(v ElementVisitor) error { return v.Text(renderStr(value)) }}
}

func Raw(html string) HTML {
	return HTML{func(v ElementVisitor) error { return v.Raw(html) }}
}

func visitElement(v ElementVisitor, tagName string, attributes Attributes, children HTML) error {
	if err := visitElementPreamble(v, tagName, attributes); err != nil {
		return err
	}
	if err := v.OpenTagEnd(); err != nil {
		return err
	}
	if err := children.Accept(v); err != nil {
		return err
	}
	if err := v.CloseTag(tagName); err != nil {
		return err
	}
	return nil
}

func visitElementSelfClose(v ElementVisitor, tagName string, attributes Attributes) error {
	if err := visitElementPreamble(v, tagName, attributes); err != nil {
		return err
	}
	if err := v.OpenTagSelfClose(); err != nil {
		return err
	}
	return nil
}

func visitElementPreamble(v ElementVisitor, tagName string, attributes Attributes) error {
	if attributes == nil {
		attributes = Attrs()
	}
	if err := v.OpenTagStart(tagName); err != nil {
		return err
	}
	if err := attributes.Accept(v); err != nil {
		return err
	}
	return nil
}
