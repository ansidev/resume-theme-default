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

	for _, funcName := range funcs {
		if _, ok := themeConfig.Funcs[funcName]; !ok {
			t.Errorf("Theme function %s was not exists", funcName)
		}
	}
}
