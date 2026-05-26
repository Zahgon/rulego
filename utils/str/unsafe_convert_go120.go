//go:build go1.20

package str

// Go 1.20+版本的实现，使用官方unsafe函数
func unsafeStringFromBytes_impl(b []byte) string { _ = "STUB: not implemented"; return "" }

// 使用Go 1.20+的官方unsafe函数

func unsafeBytesFromString_impl(s string) []byte { _ = "STUB: not implemented"; return nil }

// 使用Go 1.20+的官方unsafe函数

// 实现信息
const implementationInfo = "Go 1.20+ official unsafe"
