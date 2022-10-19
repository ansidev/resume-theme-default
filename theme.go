package simple

import (
	"github.com/ansidev/resume-cli/config"
)

func LoadThemeConfig() config.ThemeConfig {
	return config.ThemeConfig{
		ThemeBaseDir: "themes",
		Funcs: map[string]interface{}{
			"trimPrefix":     trimPrefixHelper,
			"bootstrapIcon": bootstrapIconHelper,
			"joinArray":     joinArrayHelper,
			"markdown":      markdownHelper,
			"displayUrl":    displayUrlHelper,
			"toLowerCase":   toLowerCaseHelper,
			"year":          yearHelper,
		},
	}
}
