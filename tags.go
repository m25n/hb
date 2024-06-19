package hb

func HTML5(attributes Attributes, children ...HTML) HTML {
	return Tag_("!DOCTYPE", Attr_("html")).
		Append(Html(attributes, children...))
}

func Anchor(attributes Attributes, children ...HTML) HTML {
	return Tag("a", attributes, children...)
}
func Abbr(attributes Attributes, children ...HTML) HTML {
	return Tag("abbr", attributes, children...)
}
func Acronym(attributes Attributes, children ...HTML) HTML {
	return Tag("acronym", attributes, children...)
}
func Address(attributes Attributes, children ...HTML) HTML {
	return Tag("address", attributes, children...)
}
func Area(attributes Attributes, children ...HTML) HTML {
	return Tag("area", attributes, children...)
}
func Article(attributes Attributes, children ...HTML) HTML {
	return Tag("article", attributes, children...)
}
func Aside(attributes Attributes, children ...HTML) HTML {
	return Tag("aside", attributes, children...)
}
func Audio(attributes Attributes, children ...HTML) HTML {
	return Tag("audio", attributes, children...)
}
func B(attributes Attributes, children ...HTML) HTML {
	return Tag("b", attributes, children...)
}
func Base(attributes Attributes, children ...HTML) HTML {
	return Tag("base", attributes, children...)
}
func BaseFont(attributes Attributes, children ...HTML) HTML {
	return Tag("basefont", attributes, children...)
}
func BDI(attributes Attributes, children ...HTML) HTML {
	return Tag("bdi", attributes, children...)
}
func BDO(attributes Attributes, children ...HTML) HTML {
	return Tag("bdo", attributes, children...)
}
func BGSound(attributes Attributes, children ...HTML) HTML {
	return Tag("bgsound", attributes, children...)
}
func Big(attributes Attributes, children ...HTML) HTML {
	return Tag("big", attributes, children...)
}
func Blink(attributes Attributes, children ...HTML) HTML {
	return Tag("blink", attributes, children...)
}
func BlockQuote(attributes Attributes, children ...HTML) HTML {
	return Tag("blockquote", attributes, children...)
}
func Body(attributes Attributes, children ...HTML) HTML {
	return Tag("body", attributes, children...)
}
func Button(attributes Attributes, children ...HTML) HTML {
	return Tag("button", attributes, children...)
}
func Canvas(attributes Attributes, children ...HTML) HTML {
	return Tag("canvas", attributes, children...)
}
func Caption(attributes Attributes, children ...HTML) HTML {
	return Tag("caption", attributes, children...)
}
func Center(attributes Attributes, children ...HTML) HTML {
	return Tag("center", attributes, children...)
}
func Cite(attributes Attributes, children ...HTML) HTML {
	return Tag("cite", attributes, children...)
}
func Code(attributes Attributes, children ...HTML) HTML {
	return Tag("code", attributes, children...)
}
func Col(attributes Attributes, children ...HTML) HTML {
	return Tag("col", attributes, children...)
}
func ColGroup(attributes Attributes, children ...HTML) HTML {
	return Tag("colgroup", attributes, children...)
}
func DataList(attributes Attributes, children ...HTML) HTML {
	return Tag("datalist", attributes, children...)
}
func DD(attributes Attributes, children ...HTML) HTML {
	return Tag("dd", attributes, children...)
}
func Del(attributes Attributes, children ...HTML) HTML {
	return Tag("del", attributes, children...)
}
func Details(attributes Attributes, children ...HTML) HTML {
	return Tag("details", attributes, children...)
}
func Dfn(attributes Attributes, children ...HTML) HTML {
	return Tag("dfn", attributes, children...)
}
func Dir(attributes Attributes, children ...HTML) HTML {
	return Tag("dir", attributes, children...)
}
func Div(attributes Attributes, children ...HTML) HTML {
	return Tag("div", attributes, children...)
}
func DL(attributes Attributes, children ...HTML) HTML {
	return Tag("dl", attributes, children...)
}
func DT(attributes Attributes, children ...HTML) HTML {
	return Tag("dt", attributes, children...)
}
func Em(attributes Attributes, children ...HTML) HTML {
	return Tag("em", attributes, children...)
}
func Embed(attributes Attributes, children ...HTML) HTML {
	return Tag("embed", attributes, children...)
}
func FieldSet(attributes Attributes, children ...HTML) HTML {
	return Tag("fieldset", attributes, children...)
}
func FigCaption(attributes Attributes, children ...HTML) HTML {
	return Tag("figcaption", attributes, children...)
}
func Figure(attributes Attributes, children ...HTML) HTML {
	return Tag("figure", attributes, children...)
}
func Font(attributes Attributes, children ...HTML) HTML {
	return Tag("font", attributes, children...)
}
func Footer(attributes Attributes, children ...HTML) HTML {
	return Tag("footer", attributes, children...)
}
func Form(attributes Attributes, children ...HTML) HTML {
	return Tag("form", attributes, children...)
}
func Frame(attributes Attributes, children ...HTML) HTML {
	return Tag("frame", attributes, children...)
}
func FrameSet(attributes Attributes, children ...HTML) HTML {
	return Tag("frameset", attributes, children...)
}
func H1(attributes Attributes, children ...HTML) HTML {
	return Tag("h1", attributes, children...)
}
func H2(attributes Attributes, children ...HTML) HTML {
	return Tag("h2", attributes, children...)
}
func H3(attributes Attributes, children ...HTML) HTML {
	return Tag("h3", attributes, children...)
}
func H4(attributes Attributes, children ...HTML) HTML {
	return Tag("h4", attributes, children...)
}
func H5(attributes Attributes, children ...HTML) HTML {
	return Tag("h5", attributes, children...)
}
func H6(attributes Attributes, children ...HTML) HTML {
	return Tag("h6", attributes, children...)
}
func Head(attributes Attributes, children ...HTML) HTML {
	return Tag("head", attributes, children...)
}
func Header(attributes Attributes, children ...HTML) HTML {
	return Tag("header", attributes, children...)
}
func HGroup(attributes Attributes, children ...HTML) HTML {
	return Tag("hgroup", attributes, children...)
}
func Html(attributes Attributes, children ...HTML) HTML {
	return Tag("html", attributes, children...)
}
func I(attributes Attributes, children ...HTML) HTML {
	return Tag("i", attributes, children...)
}
func IFrame(attributes Attributes, children ...HTML) HTML {
	return Tag("iframe", attributes, children...)
}
func Ins(attributes Attributes, children ...HTML) HTML {
	return Tag("ins", attributes, children...)
}
func IsIndex(attributes Attributes, children ...HTML) HTML {
	return Tag("isindex", attributes, children...)
}
func KBD(attributes Attributes, children ...HTML) HTML {
	return Tag("kbd", attributes, children...)
}
func KeyGen(attributes Attributes, children ...HTML) HTML {
	return Tag("keygen", attributes, children...)
}
func Label(attributes Attributes, children ...HTML) HTML {
	return Tag("label", attributes, children...)
}
func Legend(attributes Attributes, children ...HTML) HTML {
	return Tag("legend", attributes, children...)
}
func LI(attributes Attributes, children ...HTML) HTML {
	return Tag("li", attributes, children...)
}
func Listing(attributes Attributes, children ...HTML) HTML {
	return Tag("listing", attributes, children...)
}
func Main(attributes Attributes, children ...HTML) HTML {
	return Tag("main", attributes, children...)
}
func Map(attributes Attributes, children ...HTML) HTML {
	return Tag("map", attributes, children...)
}
func Mark(attributes Attributes, children ...HTML) HTML {
	return Tag("mark", attributes, children...)
}
func Marquee(attributes Attributes, children ...HTML) HTML {
	return Tag("marquee", attributes, children...)
}
func Menu(attributes Attributes, children ...HTML) HTML {
	return Tag("menu", attributes, children...)
}
func Meter(attributes Attributes, children ...HTML) HTML {
	return Tag("meter", attributes, children...)
}
func Nav(attributes Attributes, children ...HTML) HTML {
	return Tag("nav", attributes, children...)
}
func NoFrames(attributes Attributes, children ...HTML) HTML {
	return Tag("noframes", attributes, children...)
}
func NoScript(attributes Attributes, children ...HTML) HTML {
	return Tag("noscript", attributes, children...)
}
func Object(attributes Attributes, children ...HTML) HTML {
	return Tag("object", attributes, children...)
}
func OL(attributes Attributes, children ...HTML) HTML { return Tag("ol", attributes, children...) }
func OptGroup(attributes Attributes, children ...HTML) HTML {
	return Tag("optgroup", attributes, children...)
}
func Option(attributes Attributes, children ...HTML) HTML {
	return Tag("option", attributes, children...)
}
func Output(attributes Attributes, children ...HTML) HTML {
	return Tag("output", attributes, children...)
}
func P(attributes Attributes, children ...HTML) HTML {
	return Tag("p", attributes, children...)
}
func Param(attributes Attributes, children ...HTML) HTML {
	return Tag("param", attributes, children...)
}
func PlainText(attributes Attributes, children ...HTML) HTML {
	return Tag("plaintext", attributes, children...)
}
func Pre(attributes Attributes, children ...HTML) HTML {
	return Tag("pre", attributes, children...)
}
func Progress(attributes Attributes, children ...HTML) HTML {
	return Tag("progress", attributes, children...)
}
func Q(attributes Attributes, children ...HTML) HTML  { return Tag("q", attributes, children...) }
func RP(attributes Attributes, children ...HTML) HTML { return Tag("rp", attributes, children...) }
func RT(attributes Attributes, children ...HTML) HTML { return Tag("rt", attributes, children...) }
func Ruby(attributes Attributes, children ...HTML) HTML {
	return Tag("ruby", attributes, children...)
}
func S(attributes Attributes, children ...HTML) HTML { return Tag("s", attributes, children...) }
func Samp(attributes Attributes, children ...HTML) HTML {
	return Tag("samp", attributes, children...)
}
func Script(attributes Attributes, script string) HTML {
	return Tag("script", attributes, Raw(script))
}
func Section(attributes Attributes, children ...HTML) HTML {
	return Tag("section", attributes, children...)
}
func Select(attributes Attributes, children ...HTML) HTML {
	return Tag("select", attributes, children...)
}
func Small(attributes Attributes, children ...HTML) HTML {
	return Tag("small", attributes, children...)
}
func Source(attributes Attributes, children ...HTML) HTML {
	return Tag("source", attributes, children...)
}
func Spacer(attributes Attributes, children ...HTML) HTML {
	return Tag("spacer", attributes, children...)
}
func Span(attributes Attributes, children ...HTML) HTML {
	return Tag("span", attributes, children...)
}
func Strike(attributes Attributes, children ...HTML) HTML {
	return Tag("strike", attributes, children...)
}
func Strong(attributes Attributes, children ...HTML) HTML {
	return Tag("strong", attributes, children...)
}
func Style(attributes Attributes, css string) HTML {
	return Tag("style", attributes, Raw(css))
}
func Sub(attributes Attributes, children ...HTML) HTML {
	return Tag("sub", attributes, children...)
}
func Summary(attributes Attributes, children ...HTML) HTML {
	return Tag("summary", attributes, children...)
}
func Sup(attributes Attributes, children ...HTML) HTML {
	return Tag("sup", attributes, children...)
}
func Table(attributes Attributes, children ...HTML) HTML {
	return Tag("table", attributes, children...)
}
func TBody(attributes Attributes, children ...HTML) HTML {
	return Tag("tbody", attributes, children...)
}
func TD(attributes Attributes, children ...HTML) HTML {
	return Tag("td", attributes, children...)
}
func Textarea(attributes Attributes, children ...HTML) HTML {
	return Tag("textarea", attributes, children...)
}
func TFoot(attributes Attributes, children ...HTML) HTML {
	return Tag("tfoot", attributes, children...)
}
func TH(attributes Attributes, children ...HTML) HTML {
	return Tag("th", attributes, children...)
}
func THead(attributes Attributes, children ...HTML) HTML {
	return Tag("thead", attributes, children...)
}
func Time(attributes Attributes, children ...HTML) HTML {
	return Tag("time", attributes, children...)
}
func Title(attributes Attributes, children ...HTML) HTML {
	return Tag("title", attributes, children...)
}
func TR(attributes Attributes, children ...HTML) HTML { return Tag("tr", attributes, children...) }
func Track(attributes Attributes, children ...HTML) HTML {
	return Tag("track", attributes, children...)
}
func TT(attributes Attributes, children ...HTML) HTML { return Tag("tt", attributes, children...) }
func U(attributes Attributes, children ...HTML) HTML  { return Tag("u", attributes, children...) }
func UL(attributes Attributes, children ...HTML) HTML { return Tag("ul", attributes, children...) }
func Var(attributes Attributes, children ...HTML) HTML {
	return Tag("var", attributes, children...)
}
func Video(attributes Attributes, children ...HTML) HTML {
	return Tag("video", attributes, children...)
}

/*
Self closing tags
*/

func Br(attributes ...Attributes) HTML {
	return Tag_("br", Attrs(attributes...))
}
func HR(attributes ...Attributes) HTML {
	return Tag_("hr", Attrs(attributes...))
}
func Img(attributes ...Attributes) HTML {
	return Tag_("img", Attrs(attributes...))
}
func Input(attributes ...Attributes) HTML {
	return Tag_("input", Attrs(attributes...))
}
func Link(attributes ...Attributes) HTML {
	return Tag_("link", Attrs(attributes...))
}
func Meta(attributes ...Attributes) HTML { return Tag_("meta", Attrs(attributes...)) }
func NoBR(attributes ...Attributes) HTML {
	return Tag_("nobr", Attrs(attributes...))
}
func Wbr(attributes ...Attributes) HTML {
	return Tag_("wbr", Attrs(attributes...))
}
