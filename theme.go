package theme

import (
	"embed"

	"github.com/ansidev/resume-cli/config"
)

//go:embed resources
var themeResources embed.FS

func LoadThemeConfig() config.ThemeConfig {
	return config.ThemeConfig{
		TemplateEngine: "handlebars",
		Embed:          true,
		ThemeResources: themeResources,
		Funcs: map[string]interface{}{
			"trimPrefix":    trimPrefixHelper,
			"bootstrapIcon": bootstrapIconHelper,
			"joinArray":     joinArrayHelper,
			"markdown":      markdownHelper,
			"displayUrl":    displayUrlHelper,
			"toLowerCase":   toLowerCaseHelper,
			"year":          yearHelper,
		},
	}
}
