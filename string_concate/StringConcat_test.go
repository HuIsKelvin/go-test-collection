package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var (
	// 被拼接的基字符串
	BASE_STRING = "base string,"
	// 拼接次数
	CONCAT_NUM = 500
)

// BenchmarkStringPlus concat strings by "+"
func BenchmarkStringPlus(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		str := ""
		for i := 0; i < CONCAT_NUM; i++ {
			str = str + BASE_STRING
		}
	}
}

// BenchmarkFmtSprintf concat strings by fmt.Sprintf
func BenchmarkFmtSprintf(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		str := ""
		for i := 0; i < CONCAT_NUM; i++ {
			str = fmt.Sprintf("%s%s", str, BASE_STRING)
		}
	}
}

// BenchmarkFmtSprintf concat strings by strings.Join
func BenchmarkStringsJoin(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		str := ""
		for i := 0; i < CONCAT_NUM; i++ {
			str = strings.Join([]string{str, BASE_STRING}, "")
		}
	}
}

// BenchmarkFmtSprintf concat strings by bytes.Buffer
func BenchmarkBytesBuffer(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		byteBuffer := bytes.Buffer{}

		for i := 0; i < CONCAT_NUM; i++ {
			if _, err := byteBuffer.WriteString(BASE_STRING); err != nil {
				b.Fatal("bytes.Buffer write string failed")
			}
		}

		_ = byteBuffer.String()
	}
}

// BenchmarkFmtSprintf concat strings by bytes.Buffer with pre-allocated memory
func BenchmarkByteBufferPreSize(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		byteBuffer := bytes.Buffer{}
		byteBuffer.Grow(len(BASE_STRING) * CONCAT_NUM)

		for i := 0; i < CONCAT_NUM; i++ {
			if _, err := byteBuffer.WriteString(BASE_STRING); err != nil {
				b.Fatal("bytes.Buffer write string failed")
			}
		}

		_ = byteBuffer.String()
	}
}

// BenchmarkFmtSprintf concat strings by strings.Builder
func BenchmarkStringBuilder(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		strBuffer := strings.Builder{}

		for i := 0; i < CONCAT_NUM; i++ {
			if _, err := strBuffer.WriteString(BASE_STRING); err != nil {
				b.Fatal("bytes.Buffer write string failed")
			}
		}

		_ = strBuffer.String()
	}
}

// BenchmarkFmtSprintf concat strings by strings.Builder with pre-allocated memory
func BenchmarkStringBuilderPreSize(b *testing.B) {

	for cnt := 0; cnt < b.N; cnt++ {
		strBuffer := strings.Builder{}
		strBuffer.Grow(len(BASE_STRING) * CONCAT_NUM)

		for i := 0; i < CONCAT_NUM; i++ {
			if _, err := strBuffer.WriteString(BASE_STRING); err != nil {
				b.Fatal("bytes.Buffer write string failed")
			}
		}

		_ = strBuffer.String()
	}
}
