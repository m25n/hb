package hb

import (
	"fmt"
	"io"
)

type htmlRenderer struct {
	w io.Writer
	n int
}

func newHTMLRenderer(w io.Writer) *htmlRenderer {
	return &htmlRenderer{w: w}
}

func (r *htmlRenderer) OpenTagStart(name string) error {
	a, err := fmt.Fprintf(r.w, "<%s", name)
	r.n += a
	return err
}

func (r *htmlRenderer) Attr(name string, value any) error {
	valueStr := renderStr(value)
	if len(valueStr) == 0 {
		a, err := fmt.Fprintf(r.w, " %s", name)
		r.n += a
		return err
	}
	a, err := fmt.Fprintf(r.w, " %s=\"%s\"", name, valueStr)
	r.n += a
	return err
}

func (r *htmlRenderer) OpenTagEnd() error {
	a, err := fmt.Fprint(r.w, ">")
	r.n += a
	return err
}

func (r *htmlRenderer) Text(s string) error {
	a, err := fmt.Fprint(r.w, s)
	r.n += a
	return err
}

func (r *htmlRenderer) Raw(s string) error {
	a, err := fmt.Fprint(r.w, s)
	r.n += a
	return err
}

func (r *htmlRenderer) CloseTag(name string) error {
	a, err := fmt.Fprintf(r.w, "</%s>", name)
	r.n += a
	return err
}

func (r *htmlRenderer) OpenTagSelfClose() error {
	a, err := fmt.Fprint(r.w, ">")
	r.n += a
	return err
}
