package unitry

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode/utf8"
)

func playground() {
	s := "Grüezi"
	fmt.Println(s)
}

func lenUtf8(s string) int {
	// One unicode character can be multiple bytes.
	// Loop over the string with an integer and advance the int with the size of the current rune.
	// $ go doc utf8.DecodeRuneInString
	// for range
	// for i
	ans := 0
	for i := 0; i < len(s); {
		ans++
		_, size := utf8.DecodeRuneInString(s[i:])
		i += size
	}

	return ans
}

func lenUtf8Std(s string) int {
	// find lenght counter in utf8 package
	// go doc utf8
	return utf8.RuneCountInString(s)
}

func lenUtf8For(s string) int {
	// Use a range based for loop
	// https://golang.org/ref/spec#For_statements
	ans := 0
	for i := range s {
		ans++
		_, size := utf8.DecodeRuneInString(s[i:])
		i += size - 1
	}
	return ans
}

// works on utf8 too
func hasPrefix(s, prefix string) bool {
	var i, j int
	for i < len(s) && j < len(prefix) && s[i] == prefix[j] {
		i++
		j++
	}

	return j == len(prefix)
}

func contains(s, substr string) bool {
	var n, m int = len(s), len(substr)
	if m == 0 {
		return true
	}
	ff := make([]int, m+1)
	ff[0] = -1
	for i, j := 0, -1; i < m; {
		for j > -1 && substr[i] != substr[j] {
			j = ff[j]
		}
		i++
		j++
		ff[i] = j
	}

	for i, j := 0, 0; i < n; {
		for j > -1 && s[i] != substr[j] {
			j = ff[j]
		}
		i++
		j++
		if j == m {
			return true
		}
	}

	return false
}

func intsToString(values []int) string {
	// go doc strconv.itoa
	ans := "["
	for i, v := range values {
		if i > 0 {
			ans += ", "
		}
		ans += strconv.Itoa(v)
	}
	ans += "]"
	return ans
}

func intsToStringFast(values []int) string {
	// go doc strings
	// go doc bytes
	// go doc bytes.Buffer
	// go doc fmt.Fprintf
	var buf bytes.Buffer
	buf.WriteString("[")
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(strconv.Itoa(v))
	}
	buf.WriteString("]")
	return buf.String()
}
