/*
use strings.Builder and WriteString is the fastest
WriteRune second
use format is the slowest
*/
package test

import (
	"fmt"
	"strings"
	"testing"
)

func BenchmarkString(b *testing.B) {
	var sb strings.Builder
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString(" ")
	}
}

func BenchmarkRune(b *testing.B) {
	var sb strings.Builder
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteRune(' ')
	}
}

func BenchmarkFormat(b *testing.B) {
	var sb strings.Builder
	var s []interface{}
	for i := 0; i < b.N; i++ {
		sb.WriteString("%s")
		s = append(s, " ")
	}
	b.ResetTimer()
	_ = fmt.Sprintf(sb.String(), s...)
}
