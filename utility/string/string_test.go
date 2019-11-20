package string_test

import (
	"fmt"
	"reflect"
	"testing"

	. "karmia/utility/string"
)

func ExampleIsString() {
	type Test struct {
		key string
		value interface{}
	}

	fmt.Println(`Kosaka Honoka`)
	fmt.Println(1)
	fmt.Println([]byte(`Kosaka Honoka`))
	fmt.Println(Test{key: "name", value: "Kosaka Honoka"})

	// Output:
	// true
	// false
	// false
	// false
}
func TestIsString(t *testing.T) {
	value1 := ``
	if !IsString(value1) {
		t.Error(fmt.Sprintf(`error: %s is not string.`, reflect.TypeOf(value1)))
	}

	value2 := 1
	if IsString(value2) {
		t.Error(fmt.Sprintf(`error: %s is string.`, reflect.TypeOf(value2)))
	}

	value3 := []byte{1, 2, 3}
	if IsString(value3) {
		t.Error(fmt.Sprintf(`error %s is string`, reflect.TypeOf(value3)))
	}
}

func ExampleTrim() {
	fmt.Println(Trim("\r    \x00Kosaka Honoka\x0B\n\t"))
	// Output:
	// Kosaka Honoka
}
func TestTrim(t *testing.T) {
	value := `Kosaka Honoka`
	if Trim(fmt.Sprintf("\t    %s    \r\n", value)) != value {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleTrimLeft() {
	fmt.Println(TrimLeft("\r    \x00Kosaka Honoka\x0B\n\t"))
	// Output:
	// Kosaka Honoka\x0B\n\t    // "\x0B\n\t" are invisible
}
func TestTrimLeft(t *testing.T) {
	value := `Kosaka Honoka`
	if TrimLeft(fmt.Sprintf("\t    %s    \r\n", value)) != fmt.Sprintf("%s    \r\n", value) {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleTrimRight() {
	fmt.Println(TrimRight("\r    \x00Kosaka Honoka\x0B\n\t"))
	// Output:
	// \r    \x00Kosaka Honoka    // "/r    \x00" are invisible
}
func TestTrimRight(t *testing.T) {
	value := `Kosaka Honoka`
	if TrimRight(fmt.Sprintf("\t    %s    \r\n", value)) != fmt.Sprintf("\t    %s", value) {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleStrip() {
	fmt.Println(Strip("\r    \x00Kosaka Honoka\x0B\n\t"))

	// Output:
	// Kosaka Honoka
}
func TestStrip(t *testing.T) {
	value := `Kosaka Honoka`
	if Strip(fmt.Sprintf("\t    %s    \r\n", value)) != value {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleStripLeft() {
	fmt.Println(StripLeft("\r    \x00Kosaka Honoka\x0B\n\t"))
	// Output:
	// Kosaka Honoka\x0B\n\t    // "\x0B\n\t" are invisible
}
func TestStripLeft(t *testing.T) {
	value := `Kosaka Honoka`
	if StripLeft(fmt.Sprintf("\t    %s    \r\n", value)) != fmt.Sprintf("%s    \r\n", value) {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleStripRight() {
	fmt.Println(StripRight("\r    \x00Kosaka Honoka\x0B\n\t"))
	// Output:
	// \r    \x00Kosaka Honoka    // "/r    \x00" are invisible
}
func TestStripRight(t *testing.T) {
	value := `Kosaka Honoka`
	if StripRight(fmt.Sprintf("\t    %s    \r\n", value)) != fmt.Sprintf("\t    %s", value) {
		t.Error(`error: can not strip whitespaces.`)
	}
}

func ExampleNormalize() {
	// Normalize string with default normalization form (NFKC).
	fmt.Println(Normalize("\u202b１２３\r\nＡＢＣ\rｄｅｆ\nｱｲｳｴｵｶﾞ"))

	// Normalize string with specific normalization form (NFKD).
	fmt.Println(Normalize("123\nABC\ndef\nアイウエオガ", NORMALIZE_FORM_NFKD))

	// Output:
	// // Normalize string with default normalization form (NFKC).
	// 123
	// ABC
	// def
	// アイウエオガ
	//
	// // Normalize string with specific normalization form (NFKD)
	// 123
	// ABC
	// def
	// アイウエオガ
}
func TestNormalize(t *testing.T) {
	valid := "123\nABC\ndef\nアイウエオガ"
	input := "\u202b１２３\r\nＡＢＣ\rｄｅｆ\nｱｲｳｴｵｶﾞ"
	if Normalize(input, NORMALIZE_FORM_NFKC) != valid {
		t.Error(fmt.Sprintf(`error: can not normalize string %s`, input))
	}
}

func ExampleUnquote() {
	fmt.Println(Unquote(`"Kosaka Honoka"`))
	fmt.Println(Unquote(`'Kosaka Honoka'`))
	fmt.Println(Unquote("`Kosaka Honoka`"))
	fmt.Println(Unquote(`"` + "Kosaka Honoka`"))
	// Output:
	// Kosaka Honoka
	// Kosaka Honoka
	// Kosaka Honoka
	// "Kosaka Honoka`
}
func TestUnquote(t *testing.T) {
	value := `Kosaka Honoka`

	if Unquote(fmt.Sprintf(`"Kosaka Honoka"`)) != value {
		t.Error(fmt.Sprintf(`Can not unquote string "%s"`, value))
	}

	if Unquote(fmt.Sprintf(`'Kosaka Honoka'`)) != value {
		t.Error(fmt.Sprintf(`Can not unquote string '%s'`, value))
	}

	if Unquote(fmt.Sprintf(`"Kosaka Honoka'`)) == value {
		t.Error(fmt.Sprintf(`Unquote unmatch quote string "%s'`, value))
	}

	if Unquote(fmt.Sprintf(`"Kosaka" Honoka`)) == value {
		t.Error(fmt.Sprintf(`Unquote unmatch quote string "%s'`, value))
	}
}

func ExampleZfill() {
	fmt.Println(Zfill(`1`, 5))
	fmt.Println(Zfill(`+1`, 5))
	fmt.Println(Zfill(`-1`, 5))

	// Output:
	// 00001
	// +00001
	// -00001
}
func TestZfill(t *testing.T) {
	value := `1`

	if Zfill(value, 0) != `1` {
		t.Error(`Can not fill integer by 0`)
	}
	if Zfill(value, 1) != `1` {
		t.Error(`Can not fill integer by 0`)
	}
	if Zfill(value, 2) != `01` {
		t.Error(`Can not fill integer by 0`)
	}
	if Zfill(value, 3) != `001` {
		t.Error(`Can not fill integer by 0`)
	}
	if Zfill(value, 4) != `0001` {
		t.Error(`Can not fill integer by 0`)
	}
	if Zfill(value, 5) != `00001` {
		t.Error(`Can not fill integer by 0`)
	}
}

func ExampleUpperCamelCase() {
	fmt.Println(UpperCamelCase(`KosakaHonoka`))
	fmt.Println(UpperCamelCase(`kosakaHonoka`))
	fmt.Println(UpperCamelCase(`kosaka_honoka`))
	fmt.Println(UpperCamelCase(`kosaka-honoka`))

	// Output:
	// KosakaHonoka
	// KosakaHonoka
	// KosakaHonoka
	// KosakaHonoka
}
func TestUpperCamelCase(t *testing.T) {
	result := `KosakaHonoka`

	if UpperCamelCase(`kosakaHonoka`) != result {
		t.Error(`Can not convert from lower camel case to upper camel case`)
	}

	if UpperCamelCase(`kosaka_honoka`) != result {
		t.Error(`Can not convert from snake case to upper camel case`)
	}

	if UpperCamelCase(`kosaka-honoka`) != result {
		t.Error(`Can not convert from kebab case to upper camel case`)
	}
}

func ExampleLowerCamelCase() {
	fmt.Println(LowerCamelCase(`KosakaHonoka`))
	fmt.Println(LowerCamelCase(`kosakaHonoka`))
	fmt.Println(LowerCamelCase(`kosaka_honoka`))
	fmt.Println(LowerCamelCase(`kosaka-honoka`))

	// Output:
	// kosakaHonoka
	// kosakaHonoka
	// kosakaHonoka
	// kosakaHonoka
}
func TestLowerCamelCase(t *testing.T) {
	result := `kosakaHonoka`

	if LowerCamelCase(`KosakaHonoka`) != result {
		t.Error(`Can not convert from lower camel case to lower camel case`)
	}

	if LowerCamelCase(`kosaka_honoka`) != result {
		t.Error(`Can not convert from snake case to lower camel case`)
	}

	if LowerCamelCase(`kosaka-honoka`) != result {
		t.Error(`Can not convert from kebab case to lower camel case`)
	}
}

func ExampleSnakeCase() {
	fmt.Println(SnakeCase(`KosakaHonoka`))
	fmt.Println(SnakeCase(`kosakaHonoka`))
	fmt.Println(SnakeCase(`kosaka_honoka`))
	fmt.Println(SnakeCase(`kosaka-honoka`))

	// Output:
	// kosaka_honoka
	// kosaka_honoka
	// kosaka_honoka
	// kosaka_honoka
}
func TestSnakeCase(t *testing.T) {
	result := `kosaka_honoka`

	if SnakeCase(`KosakaHonoka`) != result {
		t.Error(`Can not convert from upper camel case to snake camel case`)
	}

	if SnakeCase(`kosakaHonoka`) != result {
		t.Error(`Can not convert from lower camel case to snake camel case`)
	}

	if SnakeCase(`kosaka-honoka`) != result {
		t.Error(`Can not convert from kebab case to snake camel case`)
	}
}

func ExampleKebabCase() {
	fmt.Println(KebabCase(`KosakaHonoka`))
	fmt.Println(KebabCase(`kosakaHonoka`))
	fmt.Println(KebabCase(`kosaka_honoka`))
	fmt.Println(KebabCase(`kosaka-honoka`))

	// Output:
	// kosaka-honoka
	// kosaka-honoka
	// kosaka-honoka
	// kosaka-honoka
}
func TestKebabCase(t *testing.T) {
	result := `kosaka-honoka`

	if KebabCase(`KosakaHonoka`) != result {
		t.Error(`Can not convert from upper camel case to kebab camel case`)
	}

	if KebabCase(`kosakaHonoka`) != result {
		t.Error(`Can not convert from lower camel case to kebab camel case`)
	}

	if KebabCase(`kosaka_honoka`) != result {
		t.Error(`Can not convert from snake case to snake kebab case`)
	}
}

func ExampleParse() {
	fmt.Println(Parse(`key1=value1,key2=value2,key3=value3`))
	fmt.Println(Parse(`key1=value1, key2=value2, key3=value3`))
	fmt.Println(Parse(`key1=value1&key2=value2&key3=value3`, `&`))
	fmt.Println(Parse(`keys:value1;key2:value2;key3:value3`, `;`, `:`))

	// Output:
	// map[key1:value1 key2:value2 key3:value3]
	// map[key1:value1 key2:value2 key3:value3]
	// map[key1:value1 key2:value2 key3:value3]
	// map[key2:value2 key3:value3 keys:value1]
}
func TestParse(t *testing.T) {
    parameters1 := `key1=value1, key2=value2, key3=value3`
    result1 := Parse(parameters1)

    if value, ok := result1[`key1`]; !ok || value != `value1`{
        t.Error(`Can not parse string`)
    }
    if value, ok := result1[`key2`]; !ok || value != `value2`{
        t.Error(`Can not parse string`)
    }
    if value, ok := result1[`key3`]; !ok || value != `value3`{
        t.Error(`Can not parse string`)
    }
    if _, ok := result1[`key4`]; ok {
        t.Error(`Invalid parse string`)
    }

    parameters2 := `key1=value1,key2=value2,key3=value3`
    result2 := Parse(parameters2)

    if value, ok := result2[`key1`]; !ok || value != `value1`{
        t.Error(`Can not parse string`)
    }
    if value, ok := result2[`key2`]; !ok || value != `value2`{
        t.Error(`Can not parse string`)
    }
    if value, ok := result2[`key3`]; !ok || value != `value3`{
        t.Error(`Can not parse string`)
    }
    if _, ok := result1[`key4`]; ok {
        t.Error(`Invalid parse string`)
    }

    parameters3 := `key1=value1&key2=value2&key3=value3`
    if len(Parse(parameters3)) != 0 {
        t.Error(`Invalid parse string`)
    }
}

func ExampleToBoolean() {
	fmt.Println(ToBoolean(`true`))
	fmt.Println(ToBoolean(`false`))
	fmt.Println(ToBoolean(`1`))
	fmt.Println(ToBoolean(`0`))
	fmt.Println(ToBoolean(`on`))
	fmt.Println(ToBoolean(`off`))

	// Output:
	// true
	// false
	// true
	// false
	// true
	// false
}
func TestToBoolean(t *testing.T) {
	value1 := `true`
	if ToBoolean(value1) != true {
		t.Error(fmt.Sprintf(`%s is false`, value1))
	}

	value2 := `false`
	if ToBoolean(value2) != false {
		t.Error(fmt.Sprintf(`%s is true`, value2))
	}

	value3 := `1`
	if ToBoolean(value3) != true {
		t.Error(fmt.Sprintf(`%s is false`, value3))
	}

	value4 := `0`
	if ToBoolean(value4) != false {
		t.Error(fmt.Sprintf(`%s is true`, value4))
	}

	value5 := `on`
	if ToBoolean(value5) != true {
		t.Error(fmt.Sprintf(`%s is false`, value5))
	}

	value6 := `off`
	if ToBoolean(value6) != false {
		t.Error(fmt.Sprintf(`%s is true`, value6))
	}
}
