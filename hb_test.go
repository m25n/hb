package hb

import (
	"bytes"
	"testing"
)

func TestHB(t *testing.T) {
	t.Run("correctly renders attributes", func(t *testing.T) {
		link := Tag("a", Attr("href", "https://www.google.com?foo=bar&baz=bang"))
		buf := bytes.NewBuffer(nil)
		_, _ = link.RenderHTML(buf)
		if buf.String() != "<a href=\"https://www.google.com?foo=bar&amp;baz=bang\"></a>" {
			t.Errorf("got %s", buf.String())
		}
	})

	t.Run("correctly renders children", func(t *testing.T) {
		p := Tag("p", nil,
			Tag("a", Attr("href", "https://www.google.com?foo=bar&baz=bang"), Text("Some Link")),
		)
		buf := bytes.NewBuffer(nil)
		_, _ = p.RenderHTML(buf)
		if buf.String() != "<p><a href=\"https://www.google.com?foo=bar&amp;baz=bang\">Some Link</a></p>" {
			t.Errorf("got %s", buf.String())
		}
	})
}
