package simple

import (
	"regexp"
	"strings"
	"time"

	"github.com/aymerick/raymond"
	"github.com/russross/blackfriday/v2"
)

func trimPrefixHelper(s string, i int) raymond.SafeString {
	return raymond.SafeString(s[i:])
}

func bootstrapIconHelper(icon string) raymond.SafeString {
	return raymond.SafeString(`<i class="bi bi-` + raymond.Escape(icon) + `"></i>&nbsp;`)
}

func joinArrayHelper(arr []string) raymond.SafeString {
	return raymond.SafeString(strings.Join(arr, ", "))
}

func markdownHelper(options *raymond.Options) raymond.SafeString {
	markdown := blackfriday.Run([]byte(options.Fn()))
	return raymond.SafeString(markdown)
}

func displayUrlHelper(str string) raymond.SafeString {
	regex := regexp.MustCompile(`https?:\/\/`)
	return raymond.SafeString(regex.ReplaceAllString(str, ""))
}

func toLowerCaseHelper(str string) raymond.SafeString {
	return raymond.SafeString(strings.ToLower(str))
}

func yearHelper(str string) raymond.SafeString {
	if str == "" {
		return "Present"
	}

	d, err := time.Parse("2006-01-02", str)
	if err != nil {
		return "{{invalid_date_string}}"
	}

	return raymond.SafeString(d.Format("01/2006"))
}
