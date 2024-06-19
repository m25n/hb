package hb

import "strings"

func AcceptAttr(v string) Attributes    { return Attr("accept", v) }
func AcceptCharset(v string) Attributes { return Attr("accept-charset", v) }
func AccessKey(v string) Attributes     { return Attr("accesskey", v) }
func Action(v string) Attributes        { return Attr("action", v) }
func Align(v string) Attributes         { return Attr("align", v) }
func Alt(v string) Attributes           { return Attr("alt", v) }
func AriaExpanded(v string) Attributes  { return Aria("expanded", v) }
func AriaHidden(v string) Attributes    { return Aria("hidden", v) }
func AriaLabel(v string) Attributes     { return Aria("label", v) }
func Aria(n, v string) Attributes       { return Attr("aria-"+n, v) }
func Async(v bool) Attributes           { return AttrB("async", v) }
func Autocomplete(v string) Attributes  { return Attr("autocomplete", v) }
func Autofocus(v string) Attributes     { return Attr("autofocus", v) }
func Autoplay(v string) Attributes      { return Attr("autoplay", v) }
func Bgcolor(v string) Attributes       { return Attr("bgcolor", v) }
func Border(v string) Attributes        { return Attr("border", v) }
func CharSet(v string) Attributes       { return Attr("charset", v) }
func Checked(v string) Attributes       { return Attr("checked", v) }
func CiteAttr(v string) Attributes      { return Attr("cite", v) }
func Class(classes ...string) Attributes {
	class := &strings.Builder{}
	i := 0
	for _, c := range classes {
		if len(c) == 0 {
			continue
		}
		if i > 0 {
			class.WriteString(" ")
		}
		class.WriteString(c)
		i++
	}
	return Attr("class", class)
}
func ClassSet(classes map[string]bool) Attributes {
	classStr := &strings.Builder{}
	i := 0
	for c, included := range classes {
		if !included {
			continue
		}
		if i > 0 {
			classStr.WriteString(" ")
		}
		classStr.WriteString(c)
		i++
	}
	return Attr("class", classStr.String())
}
func Color(v string) Attributes            { return Attr("color", v) }
func Cols(v string) Attributes             { return Attr("cols", v) }
func Colspan(v string) Attributes          { return Attr("colspan", v) }
func Content(v string) Attributes          { return Attr("content", v) }
func Contenteditable(v string) Attributes  { return Attr("contenteditable", v) }
func Controls(v string) Attributes         { return Attr("controls", v) }
func Coords(v string) Attributes           { return Attr("coords", v) }
func Data(n string, v string) Attributes   { return Attr("data-"+n, v) }
func Data_(v string) Attributes            { return Attr("data", v) }
func Datetime(v string) Attributes         { return Attr("datetime", v) }
func Defer(v bool) Attributes              { return AttrB("defer", v) }
func DirAttr(v string) Attributes          { return Attr("dir", v) }
func Dirname(v string) Attributes          { return Attr("dirname", v) }
func Disabled(v bool) Attributes           { return AttrB("disabled", v) }
func Download(v string) Attributes         { return Attr("download", v) }
func Draggable(v string) Attributes        { return Attr("draggable", v) }
func Dropzone(v string) Attributes         { return Attr("dropzone", v) }
func Enctype(v string) Attributes          { return Attr("enctype", v) }
func For(v string) Attributes              { return Attr("for", v) }
func FormAttr(v string) Attributes         { return Attr("form", v) }
func FormAction(v string) Attributes       { return Attr("formaction", v) }
func Headers(v string) Attributes          { return Attr("headers", v) }
func Height(v string) Attributes           { return Attr("height", v) }
func Hidden(v string) Attributes           { return Attr("hidden", v) }
func High(v string) Attributes             { return Attr("high", v) }
func HRef(v string) Attributes             { return Attr("href", v) }
func HrefLang(v string) Attributes         { return Attr("hreflang", v) }
func HTTPEquiv(v string) Attributes        { return Attr("http-equiv", v) }
func ID(v string) Attributes               { return Attr("id", v) }
func Initial(v string) Attributes          { return Attr("initial-scale", v) }
func IsMap(v string) Attributes            { return Attr("ismap", v) }
func Kind(v string) Attributes             { return Attr("kind", v) }
func LabelAttr(v string) Attributes        { return Attr("label", v) }
func Lang(v string) Attributes             { return Attr("lang", v) }
func List(v string) Attributes             { return Attr("list", v) }
func Loop(v string) Attributes             { return Attr("loop", v) }
func Low(v string) Attributes              { return Attr("low", v) }
func Max(v string) Attributes              { return Attr("max", v) }
func MaxLength(v string) Attributes        { return Attr("maxlength", v) }
func Media(v string) Attributes            { return Attr("media", v) }
func Method(v string) Attributes           { return Attr("method", v) }
func Min(v string) Attributes              { return Attr("min", v) }
func Multiple(v string) Attributes         { return Attr("multiple", v) }
func Muted(v string) Attributes            { return Attr("muted", v) }
func Name(v string) Attributes             { return Attr("name", v) }
func NoValidate(v string) Attributes       { return Attr("novalidate", v) }
func OnAbort(v string) Attributes          { return Attr("onabort", v) }
func OnAfterPrint(v string) Attributes     { return Attr("onafterprint", v) }
func OnBeforePrint(v string) Attributes    { return Attr("onbeforeprint", v) }
func OnBeforeUnload(v string) Attributes   { return Attr("onbeforeunload", v) }
func OnBlur(v string) Attributes           { return Attr("onblur", v) }
func OnCanPlay(v string) Attributes        { return Attr("oncanplay", v) }
func OnCanPlayThrough(v string) Attributes { return Attr("oncanplaythrough", v) }
func OnChange(v string) Attributes         { return Attr("onchange", v) }
func OnClick(v string) Attributes          { return Attr("onclick", v) }
func OnContextMenu(v string) Attributes    { return Attr("oncontextmenu", v) }
func OnCopy(v string) Attributes           { return Attr("oncopy", v) }
func OnCueChange(v string) Attributes      { return Attr("oncuechange", v) }
func OnCut(v string) Attributes            { return Attr("oncut", v) }
func OnDblClick(v string) Attributes       { return Attr("ondblclick", v) }
func OnDrag(v string) Attributes           { return Attr("ondrag", v) }
func OnDragEnd(v string) Attributes        { return Attr("ondragend", v) }
func OnDragEnter(v string) Attributes      { return Attr("ondragenter", v) }
func OnDragLeave(v string) Attributes      { return Attr("ondragleave", v) }
func OnDragOver(v string) Attributes       { return Attr("ondragover", v) }
func OnDragStart(v string) Attributes      { return Attr("ondragstart", v) }
func OnDrop(v string) Attributes           { return Attr("ondrop", v) }
func OnDurationChange(v string) Attributes { return Attr("ondurationchange", v) }
func OnEmptied(v string) Attributes        { return Attr("onemptied", v) }
func OnEnded(v string) Attributes          { return Attr("onended", v) }
func OnError(v string) Attributes          { return Attr("onerror", v) }
func OnFocus(v string) Attributes          { return Attr("onfocus", v) }
func OnHashChange(v string) Attributes     { return Attr("onhashchange", v) }
func OnInput(v string) Attributes          { return Attr("oninput", v) }
func OnInvalid(v string) Attributes        { return Attr("oninvalid", v) }
func OnKeyDown(v string) Attributes        { return Attr("onkeydown", v) }
func OnKeyPress(v string) Attributes       { return Attr("onkeypress", v) }
func OnKeyUp(v string) Attributes          { return Attr("onkeyup", v) }
func OnLoad(v string) Attributes           { return Attr("onload", v) }
func OnLoadedData(v string) Attributes     { return Attr("onloadeddata", v) }
func OnLoadedMetadata(v string) Attributes { return Attr("onloadedmetadata", v) }
func OnLoadStart(v string) Attributes      { return Attr("onloadstart", v) }
func OnMouseDown(v string) Attributes      { return Attr("onmousedown", v) }
func OnMouseMove(v string) Attributes      { return Attr("onmousemove", v) }
func OnMouseOut(v string) Attributes       { return Attr("onmouseout", v) }
func OnMouseOver(v string) Attributes      { return Attr("onmouseover", v) }
func OnMouseUp(v string) Attributes        { return Attr("onmouseup", v) }
func OnMouseWheel(v string) Attributes     { return Attr("onmousewheel", v) }
func OnOffline(v string) Attributes        { return Attr("onoffline", v) }
func OnOnline(v string) Attributes         { return Attr("ononline", v) }
func OnPageHide(v string) Attributes       { return Attr("onpagehide", v) }
func OnPageShow(v string) Attributes       { return Attr("onpageshow", v) }
func OnPaste(v string) Attributes          { return Attr("onpaste", v) }
func OnPause(v string) Attributes          { return Attr("onpause", v) }
func OnPlay(v string) Attributes           { return Attr("onplay", v) }
func OnPlaying(v string) Attributes        { return Attr("onplaying", v) }
func OnPopState(v string) Attributes       { return Attr("onpopstate", v) }
func OnProgress(v string) Attributes       { return Attr("onprogress", v) }
func OnRateChange(v string) Attributes     { return Attr("onratechange", v) }
func OnReset(v string) Attributes          { return Attr("onreset", v) }
func OnResize(v string) Attributes         { return Attr("onresize", v) }
func OnScroll(v string) Attributes         { return Attr("onscroll", v) }
func OnSearch(v string) Attributes         { return Attr("onsearch", v) }
func OnSeeked(v string) Attributes         { return Attr("onseeked", v) }
func OnSeeking(v string) Attributes        { return Attr("onseeking", v) }
func OnSelect(v string) Attributes         { return Attr("onselect", v) }
func OnStalled(v string) Attributes        { return Attr("onstalled", v) }
func OnStorage(v string) Attributes        { return Attr("onstorage", v) }
func OnSubmit(v string) Attributes         { return Attr("onsubmit", v) }
func OnSuspend(v string) Attributes        { return Attr("onsuspend", v) }
func OnTimeUpdate(v string) Attributes     { return Attr("ontimeupdate", v) }
func OnToggle(v string) Attributes         { return Attr("ontoggle", v) }
func OnUnload(v string) Attributes         { return Attr("onunload", v) }
func OnVolumeChange(v string) Attributes   { return Attr("onvolumechange", v) }
func OnWaiting(v string) Attributes        { return Attr("onwaiting", v) }
func OnWheel(v string) Attributes          { return Attr("onwheel", v) }
func Open(v string) Attributes             { return Attr("open", v) }
func Optimum(v string) Attributes          { return Attr("optimum", v) }
func Pattern(v string) Attributes          { return Attr("pattern", v) }
func Placeholder(v string) Attributes      { return Attr("placeholder", v) }
func Poster(v string) Attributes           { return Attr("poster", v) }
func Preload(v string) Attributes          { return Attr("preload", v) }
func Readonly(v bool) Attributes           { return AttrB("readonly", v) }
func Rel(v string) Attributes              { return Attr("rel", v) }
func Required(v bool) Attributes           { return AttrB("required", v) }
func Reversed(v bool) Attributes           { return AttrB("reversed", v) }
func Role(v string) Attributes             { return Attr("role", v) }
func Rows(v string) Attributes             { return Attr("rows", v) }
func RowSpan(v string) Attributes          { return Attr("rowspan", v) }
func Sandbox(v string) Attributes          { return Attr("sandbox", v) }
func Scope(v string) Attributes            { return Attr("scope", v) }
func Selected(v bool) Attributes           { return AttrB("selected", v) }
func Shape(v string) Attributes            { return Attr("shape", v) }
func Size(v string) Attributes             { return Attr("size", v) }
func Sizes(v string) Attributes            { return Attr("sizes", v) }
func SpanAttr(v string) Attributes         { return Attr("span", v) }
func Spellcheck(v string) Attributes       { return Attr("spellcheck", v) }
func Src(v string) Attributes              { return Attr("src", v) }
func SrcDoc(v string) Attributes           { return Attr("srcdoc", v) }
func SrcLang(v string) Attributes          { return Attr("srclang", v) }
func SrcSet(ss map[string]string) Attributes {
	buf := strings.Builder{}
	i := 0
	for k, v := range ss {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(k)
		if len(v) > 0 {
			buf.WriteString(" ")
			buf.WriteString(v)
		}
		i++
	}
	return Attr("srcset", buf.String())
}
func Start(v string) Attributes     { return Attr("start", v) }
func Step(v string) Attributes      { return Attr("step", v) }
func StyleAttr(v string) Attributes { return Attr("style", v) }
func TabIndex(v string) Attributes  { return Attr("tabindex", v) }
func Target(v string) Attributes    { return Attr("target", v) }
func TitleAttr(v string) Attributes { return Attr("title", v) }
func Translate(v string) Attributes { return Attr("translate", v) }
func Type(v string) Attributes      { return Attr("type", v) }
func UseMap(v string) Attributes    { return Attr("usemap", v) }
func Value(v string) Attributes     { return Attr("value", v) }
func Width(v string) Attributes     { return Attr("width", v) }
func Wrap(v string) Attributes      { return Attr("wrap", v) }
