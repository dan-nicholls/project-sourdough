package utils

import (
	"encoding/json"
	"fmt"
	"html"
	"strconv"
	"regexp"
)

func MapToHtmlString(m map[string]string) string {
	fmt.Printf("json: %v\n", m)
	b, _ := json.Marshal(m)
	result := html.EscapeString(string(b))
	fmt.Printf("json: %v\n", result)
	return result
}

func HtmlStringToMap(s string) (map[string]string, error) {
	unescaped := html.UnescapeString(s)
	var result map[string]string
	err := json.Unmarshal([]byte(unescaped), &result)
	if err != nil {
		return  nil, err
	}

	return result, nil
}

func JsonEscapedString(m map[string]string) string {
	b, _ := json.Marshal(m)
	return strconv.Quote(string(b))
}

func IsValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
