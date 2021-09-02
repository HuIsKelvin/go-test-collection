# Golang 字符串拼接基准测试

Golang 中不同字符串拼接方法的性能比较。

常见字符串拼接方法：

- “+” 运算符
- fmt.Sprintf
- strings.Join
- bytes.Buffer（以及预设 buffer 容量）
- string.Builder（以及预设 buffer 容量）

## 预设场景

模拟需要多次拼接字符串的场景。
对于一条基字符串（如 `"BaseString"`），拼接足够多次，最后返回 string 类型的变量。

## 测试结果

```bash
PS E:\> go test -bench="." -benchmem
goos: windows
goarch: amd64
pkg: go-test-collection/string_concate
cpu: Intel(R) Core(TM) i5-9400 CPU @ 2.90GHz
BenchmarkStringPlus-6                       5118            241411 ns/op         1606811 B/op        499 allocs/op
BenchmarkFmtSprintf-6                       3759            310479 ns/op         1625079 B/op       1500 allocs/op
BenchmarkStringsJoin-6                      5224            244217 ns/op         1606826 B/op        500 allocs/op
BenchmarkByteBuffer-6                     171885              6377 ns/op           25872 B/op          9 allocs/op
BenchmarkByteBufferPreSize-6              286476              4443 ns/op           12288 B/op          2 allocs/op
BenchmarkStringBuilder-6                  203931              6203 ns/op           26736 B/op         14 allocs/op
BenchmarkStringBuilderPreSize-6           436584              2880 ns/op            6144 B/op          1 allocs/op
PASS
ok      go-test-collection/string_concate       8.958s
```

## 结论分析

在多次拼接字符串时，基于预设大小的 `strings.Builder` 的性能是最好的，无论是从执行速度还是占用内存的角度来说。其次性能较好的是基于预设大小的 `bytes.Buffer` 这种方法。再最后，性能从高到低为 `+`、`strings.Join` 和 `fmt.Sprintf`。

- 在已知最终字符串长度的情况下，先给 buffer 预设大小。
- 将多种类型的数据拼接为字符串，则使用 `fmt.Sprintf`。
- 以零个或一个分隔符拼接已知的字符串数组时，使用 `strings.Join`。


参考：

- [Efficient String Concatenation in Go](https://hermanschaaf.com/efficient-string-concatenation-in-go/)
- [golang 几种字符串的连接方式](https://segmentfault.com/a/1190000012978989)
