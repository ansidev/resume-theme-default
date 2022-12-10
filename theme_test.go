package theme

import (
	"testing"
)

func TestLoadThemeConfig(t *testing.T) {
	themeConfig := LoadThemeConfig()

	funcs := []string{
		"trimPrefix",
		"bootstrapIcon",
		"joinArray",
		"markdown",
		"displayUrl",
		"toLowerCase",
		"year",
	}

	if themeConfig.TemplateEngine != "handlebars" {
		t.Errorf("Wrong theme template engine, expected value is `handlebars`")
	}
	for _, funcName := range funcs {
		if _, ok := themeConfig.Funcs[funcName]; !ok {
			t.Errorf("Theme function %s was not exists", funcName)
		}
	}
}
