package iwanFormatting

import (
	"iwan/src/internal/iwanCore"
)

func Render(raw []iwanCore.IwanResponse, FORMAT func([]iwanCore.IwanResponse, RenderParams) []FormatString, RENDER_FUNC func([]FormatString, FormatStringParams), params RenderParams, formatParams FormatStringParams) {
	out := FORMAT(raw, params)
	RENDER_FUNC(out, formatParams)
}
