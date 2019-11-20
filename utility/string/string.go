// This package are test package of golang.
//
// It's include golang implements of karmia-utility-string.
//     https://github.com/fujimakishouten/karmia-utility-string
//
package string

import (
	"bytes"
	"golang.org/x/text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

const NORMALIZE_FORM_NFD string = `NFD`
const NORMALIZE_FORM_NFC string = `NFC`
const NORMALIZE_FORM_NFKD string = `NFKD`
const NORMALIZE_FORM_NFKC string = `NFKC`

var others = regexp.MustCompile(`\p{C}`)

// getCharacterMask are return specified strip characters from options.
func getCharacterMask(options []string) string {
	characterMask := strings.Join(options, ``)
	if len(characterMask) > 0 {
		return characterMask
	}

	return " \t\n\r\x00\x0B"
}

// getNormalizationForm are return unicode normalization form name from options.
func getNormalizationForm(options []string) string {
	if (len(options) < 1) {
		return NORMALIZE_FORM_NFKC
	}

	form := options[0]
	if form == NORMALIZE_FORM_NFD {
		return form
	}
	if form == NORMALIZE_FORM_NFC {
		return form
	}
	if form == NORMALIZE_FORM_NFKD {
		return form
	}

	return NORMALIZE_FORM_NFKC
}

// getDelimiter are return compiled regexp pointer of option string
func getDelimiter(options []string) *regexp.Regexp {
	delimiter := `(,|, )`
	if len(options) > 0 {
		delimiter = options[0]
	}

	return regexp.MustCompile(delimiter)
}

// getSeparator are return compiled regexp pointer of option string
func getSeparator(options []string) *regexp.Regexp {
	separator := `=`
	if len(options) > 1 {
		separator = options[1]
	}

	return regexp.MustCompile(separator)
}

// IsString are return true if type of the value are string.
func IsString(value interface{}) bool {
	_, ok := value.(string)

	return ok;
}

// Trim are strip whitespace from the beginning and end of a string.
func Trim(value string, delimiter ...string) string {
	return strings.Trim(value, getCharacterMask(delimiter))
}

// Trim are strip whitespace from the beginning of a string.
func TrimLeft(value string, delimiter ...string) string {
	return strings.TrimLeft(value, getCharacterMask(delimiter))
}

// Trim are strip whitespace from the end of a string.
func TrimRight(value string, delimiter ...string) string {
	return strings.TrimRight(value, getCharacterMask(delimiter))
}

// Strip are alias of the Trim.
func Strip(value string, delimiter ...string) string {
	return Trim(value, delimiter...)
}

// StripLeft are alias of the TrimLeft.
func StripLeft(value string, delimiter ...string) string {
	return TrimLeft(value, delimiter...)
}

// StripRight are alias of the TrimRight.
func StripRight(value string, delimiter ...string) string {
	return TrimRight(value, delimiter...)
}

// Normalize are normalize string.
func Normalize(value string, normalizationForm ...string) string {
	var buffer []string;
	form := getNormalizationForm(normalizationForm)
	replacer := strings.NewReplacer("\r\n", "\n", "\r", "\n")
	for _, line := range strings.Split(replacer.Replace(value), "\n") {
		buffer = append(buffer, others.ReplaceAllString(line, ``))
	}

	if form == NORMALIZE_FORM_NFD {
		return norm.NFD.String(strings.Join(buffer, "\n"))
	}
	if form == NORMALIZE_FORM_NFC {
		return norm.NFC.String(strings.Join(buffer, "\n"))
	}
	if form == NORMALIZE_FORM_NFKD {
		return norm.NFKD.String(strings.Join(buffer, "\n"))
	}

	return norm.NFKC.String(strings.Join(buffer, "\n"))
}

// Unquote are strip quote character from the beginning and end of a string.
func Unquote(value string) string {
	if strings.HasPrefix(value, `"`) && strings.HasSuffix(value, `"`) {
		return strings.Trim(value, `"`)
	}

	if strings.HasPrefix(value, `'`) && strings.HasSuffix(value, `'`) {
		return strings.Trim(value, `'`)
	}

	if strings.HasPrefix(value, "`") && strings.HasSuffix(value, "`") {
		return strings.Trim(value, "`")
	}

	return value
}

// Zfill are return a copy of the string left filled with `0`
func Zfill(value string, width int) string {
	prefix := ``
	if strings.HasPrefix(value, `+`) || strings.HasPrefix(value, `-`) {
		prefix = value[0:1]
		value = value[1:]
	}

	count := width - len([]rune(value))

	if count > 0 {
		return prefix + strings.Repeat(`0`, count) + value
	}

	return prefix + value
}

// UpperCamelCase are make a value upper camel case.
func UpperCamelCase(value string) string {
	replacer := strings.NewReplacer(`_`, ` `, `-`, ` `)
	result := replacer.Replace(value)

	return strings.ReplaceAll(strings.Title(result), ` `, ``)
}

// LowerCamelCase are make value lower camel case.
func LowerCamelCase(value string) string {
	result := UpperCamelCase(value)

	return strings.ToLower(result[0:1]) + result[1:]
}

// SnakeCase are make value snake case.
func SnakeCase(value string) string {
	var buffer bytes.Buffer
	value = strings.ReplaceAll(value, `-`, `_`)
	for i, v := range value {
		if i == 0 {
			buffer.WriteString(strings.ToLower(string(v)))
		} else if unicode.IsUpper(v) {
			buffer.WriteString(`_` + strings.ToLower(string(v)))
		} else {
			buffer.WriteString(string(v))
		}
	}

	return buffer.String()
}

// KebabCase are are make value kabeb case.
func KebabCase(value string) string {
	var buffer bytes.Buffer
	value = strings.ReplaceAll(value, `_`, `-`)
	for i, v := range value {
		if i == 0 {
			buffer.WriteString(strings.ToLower(string(v)))
		} else if unicode.IsUpper(v) {
			buffer.WriteString(`-` + strings.ToLower(string(v)))
		} else {
			buffer.WriteString(string(v))
		}
	}

	return buffer.String()
}

// Parse are parse "key1=value1,key2=value2" like string to map of string.
func Parse(value string, options ...string) map[string]string {
	result := map[string]string{}
	delimiter := getDelimiter(options)
	separator := getSeparator(options)
	list := delimiter.Split(value, -1)
	if len(list) == 1 {
	    return result
    }

	for _, items := range list {
		item := separator.Split(items, 2)
		result[Trim(item[0])] = Trim(item[1])
	}

	return result
}

// Convert from  `true`, `on`, `1` strings to true and others to false
func ToBoolean(value string) bool {
	for _, parameter := range []string{`true`, `1`, `on`} {
		if value == parameter {
			return true
		}
	}

	return false
}
