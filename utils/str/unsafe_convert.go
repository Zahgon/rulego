package str

// UnsafeStringFromBytes 零拷贝转换[]byte到string
// 使用构建标签自动选择最优实现，支持Go 1.18+所有版本
//
// WARNING: 返回的字符串与底层[]byte共享内存
// 在使用字符串期间不要修改原始数据
func UnsafeStringFromBytes(b []byte) string { _ = "STUB: not implemented"; return "" }

// UnsafeBytesFromString 零拷贝转换string到[]byte
// 使用构建标签自动选择最优实现，支持Go 1.18+所有版本
//
// WARNING: 返回的[]byte与底层字符串共享内存
// 不要修改返回的[]byte
func UnsafeBytesFromString(s string) []byte { _ = "STUB: not implemented"; return nil }

// SafeStringFromBytes 安全转换[]byte到string（带内存拷贝）
func SafeStringFromBytes(b []byte) string { _ = "STUB: not implemented"; return "" }

// SafeBytesFromString 安全转换string到[]byte（带内存拷贝）
func SafeBytesFromString(s string) []byte { _ = "STUB: not implemented"; return nil }

// GetConverterInfo 返回当前使用的转换器信息
func GetConverterInfo() string { _ = "STUB: not implemented"; return "" }
