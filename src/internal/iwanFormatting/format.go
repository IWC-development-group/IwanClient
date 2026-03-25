package iwanFormatting

import (
	"iwan/src/internal/iwanCore"
)

func Render(raw iwanCore.IwanResponse, FORMAT func(iwanCore.IwanResponse, RenderParams) string, RENDER_FUNC func(string), params RenderParams) {
	format_string := FORMAT(raw, params)
	RENDER_FUNC(format_string)
}
