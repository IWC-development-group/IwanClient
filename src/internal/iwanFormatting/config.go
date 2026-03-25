package iwanFormatting

const (
	defaultWidth int = 120
)

type RenderParams struct {
	Status     bool
	Name       bool
	Namespace  bool
	Content    bool
	Pages      bool
	Namespaces bool
}
