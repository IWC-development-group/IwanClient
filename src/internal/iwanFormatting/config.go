package iwanFormatting

const (
	defaultWidth int = 90
)

type RenderParams struct {
	Status     bool
	Name       bool
	Namespace  bool
	Content    bool
	Pages      bool
	Namespaces bool
}

type FormatStringParams struct {
	HeaderMD  bool
	ContentMD bool
	FooterMD  bool
}

type FormatString struct {
	Header  string
	Content string
	Footer  string
}
