package theme

import (
	"strings"
	"testing"

	"github.com/mailgun/raymond/v2"
)

func TestFuncs(t *testing.T) {
	funcs := Funcs()
	names := []string{"trimPrefix", "bootstrapIcon", "joinArray", "markdown", "displayUrl", "toLowerCase", "year"}

	if len(funcs) != len(names) {
		t.Errorf("expected %d helpers, got %d", len(names), len(funcs))
	}
	for _, name := range names {
		if _, ok := funcs[name]; !ok {
			t.Errorf("helper %s does not exist", name)
		}
	}
}

func TestHelpers(t *testing.T) {
	tests := []struct {
		name string
		got  raymond.SafeString
		want raymond.SafeString
	}{
		{"trimPrefix", trimPrefixHelper("abcdef", 2), "cdef"},
		{"bootstrapIcon", bootstrapIconHelper("github"), `<i class="bi bi-github"></i>&nbsp;`},
		{"joinArray", joinArrayHelper([]string{"a", "b", "c"}), "a, b, c"},
		{"displayUrl http", displayUrlHelper("http://example.com"), "example.com"},
		{"displayUrl https", displayUrlHelper("https://example.com/x"), "example.com/x"},
		{"toLowerCase", toLowerCaseHelper("AbC"), "abc"},
		{"year", yearHelper("2023-05-17"), "05/2023"},
		{"year empty", yearHelper(""), "Present"},
		{"year invalid", yearHelper("nope"), "{{invalid_date_string}}"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.want {
				t.Errorf("got %q, want %q", tt.got, tt.want)
			}
		})
	}
}

func TestMarkdownHelper(t *testing.T) {
	got := string(markdownHelper("**hi**"))
	if !strings.Contains(got, "<strong>hi</strong>") {
		t.Errorf("unexpected markdown output %q", got)
	}
}
