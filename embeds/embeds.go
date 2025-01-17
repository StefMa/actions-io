package embeds

import (
	"embed"
)

//go:embed *
var AllEmbeds embed.FS

//go:embed index.html
var IndexHtml string

//go:embed style.css
var StyleCss string
