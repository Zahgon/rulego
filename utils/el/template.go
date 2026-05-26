package el

import (
	"regexp"

	"github.com/expr-lang/expr/vm"
)

type Template interface {
	Parse() error
	Execute(data map[string]any) (interface{}, error)
	ExecuteFn(loadDataFunc func() map[string]any) (interface{}, error)
	ExecuteAsString(data map[string]any) string
	// Deprecated: Use HasVar instead.
	// IsNotVar 是否是模板变量
	IsNotVar() bool
	// HasVar 是否有变量
	HasVar() bool
}

// TemplateConfig 模板配置选项
type TemplateConfig struct {
	IncludeFunc IncludeFunc // 自定义 include 函数
}

// Option 模板选项函数
type Option func(*TemplateConfig)

// WithIncludeFunc 设置自定义 include 函数
func WithIncludeFunc(fn IncludeFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// IncludeFunc 文件包含函数类型
type IncludeFunc func(path string) string

// NewTemplate 根据模板内容创建相应的模板实例
// 识别规则：
// 1. 如果是完整的单个表达式 ${...}，创建 ExprTemplate
// 2. 如果包含变量但不是单个表达式，创建 MixedTemplate
// 3. 如果不包含变量，创建 NotTemplate
// 4. 如果不是字符串类型，创建 AnyTemplate
//
// 支持选项：
//   - WithIncludeFunc(fn IncludeFunc): 设置自定义 include 函数
//
// include 函数使用示例：
//   - ${include("/path/to/file.txt")}: 包含文件内容（使用绝对路径）
//   - ${upper(include("/path/to/file.txt"))}: 包含文件内容并转为大写
//   - ${include("/path/to/file.txt") + suffix}: 包含文件内容并拼接后缀
func NewTemplate(tmpl any, opts ...Option) (Template, error) {
	_ = "STUB: not implemented"
	// 解析配置
	return *new(Template), nil
}

// 检查是否是完整的单个表达式：以 ${ 开头，以 } 结尾，且中间没有其他 ${ 或 }

// 检查是否是单个完整表达式（中间不包含额外的 ${ 或 }）
// 去掉开头的 ${ 和结尾的 }

// 如果包含变量但不是单个表达式，使用 MixedTemplate

// ExprTemplate 模板变量支持 这种方式 ${xx},使用expr表达式计算
type ExprTemplate struct {
	Tmpl    string
	Program *vm.Program
	config  *TemplateConfig
}

// 定义正则表达式，用于匹配形如 ${...} 的占位符
var re = regexp.MustCompile(`\$\{([^}]*)\}`)

// NewExprTemplate 创建表达式模板（向后兼容）
func NewExprTemplate(tmpl string) (*ExprTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewExprTemplateWithConfig 创建带配置的表达式模板
func NewExprTemplateWithConfig(tmpl string, cfg *TemplateConfig) (*ExprTemplate, error) {
	_ = "STUB: not implemented"
	// 使用字符串构建器来处理模板字符串
	return nil, nil
}

// 标记是否在双引号内

// 翻转 inQuotes 标志

// 处理转义字符

// 如果不在双引号内且遇到${，尝试匹配并替换

// 找到匹配的 ${...}

// 写入 ${ 前的内容
// 替换为 $1
// 跳过已处理的部分

// 如果在双引号内或未找到匹配项，直接写入字符

// 替换后的模板字符串

// 创建 ExprTemplate 实例

// 调用 Parse 方法解析模板

func (t *ExprTemplate) Parse() error { _ = "STUB: not implemented"; return nil }

// buildEnv 构建包含 include 函数的环境
func (t *ExprTemplate) buildEnv(data map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil

	// 复制原始数据
}

// 添加 include 函数（只要 config 不为 nil 就添加，支持绝对路径）

// includeFunc 返回 include 函数实现
func (t *ExprTemplate) includeFunc() func(string) string { _ = "STUB: not implemented"; return nil }

// 使用自定义函数或默认实现

// 默认实现：读取文件

// fileExistsFunc 返回 fileExists 函数实现
func (t *ExprTemplate) fileExistsFunc() func(string) bool { _ = "STUB: not implemented"; return nil }

func (t *ExprTemplate) Execute(data map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// 构建包含 include 函数的环境
		nil
}

func (t *ExprTemplate) ExecuteFn(loadDataFunc func() map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *ExprTemplate) IsNotVar() bool { _ = "STUB: not implemented"; return false }

func (t *ExprTemplate) HasVar() bool {
	_ = "STUB: not implemented"

	// ExecuteAsString 执行模板并返回字符串结果
	return false
}

func (t *ExprTemplate) ExecuteAsString(data map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

// NotTemplate 原样输出
type NotTemplate struct {
	Tmpl string
}

func (t *NotTemplate) Parse() error { _ = "STUB: not implemented"; return nil }

func (t *NotTemplate) Execute(data map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// ExecuteFn 执行模板函数
		nil
}

func (t *NotTemplate) ExecuteFn(loadDataFunc func() map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// ExecuteAsString 执行模板并返回字符串结果
		nil
}

func (t *NotTemplate) ExecuteAsString(data map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *NotTemplate) IsNotVar() bool { _ = "STUB: not implemented"; return false }

func (t *NotTemplate) HasVar() bool { _ = "STUB: not implemented"; return false }

type AnyTemplate struct {
	Tmpl any
}

func (t *AnyTemplate) Parse() error { _ = "STUB: not implemented"; return nil }

func (t *AnyTemplate) Execute(data map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// ExecuteFn 执行模板函数
		nil
}

func (t *AnyTemplate) ExecuteFn(loadDataFunc func() map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// ExecuteAsString 执行模板并返回字符串结果
		nil
}

func (t *AnyTemplate) ExecuteAsString(data map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *AnyTemplate) IsNotVar() bool { _ = "STUB: not implemented"; return false }

func (t *AnyTemplate) HasVar() bool {
	_ = "STUB: not implemented"

	// MixedTemplate 支持混合字符串和变量的模板，格式如 aa/${xxx}
	return false
}

type MixedTemplate struct {
	Tmpl      string
	variables []struct {
		start int
		end   int
		expr  string // 保存原始表达式字符串，用于动态编译
	}
	hasVars bool // 是否包含变量
	config  *TemplateConfig
}

// NewMixedTemplate 创建混合模板（向后兼容）
func NewMixedTemplate(tmpl string) (*MixedTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMixedTemplateWithConfig 创建带配置的混合模板
func NewMixedTemplateWithConfig(tmpl string, cfg *TemplateConfig) (*MixedTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *MixedTemplate) Parse() error {
	_ = "STUB: not implemented"
	// 先检查是否包含${}变量
	return nil
}

// 保存原始表达式字符串，在执行时动态编译

func (t *MixedTemplate) Execute(data map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil,

		// buildEnv 构建包含 include 函数的环境
		nil
}

func (t *MixedTemplate) buildEnv(data map[string]any) map[string]any {
	_ = "STUB: not implemented"
	return nil
}

// 添加 include 函数（只要 config 不为 nil 就添加）

func (t *MixedTemplate) execute(data map[string]any) (string, error) {
	_ = "STUB: not implemented"
	// 如果没有变量，直接返回原始字符串
	return "", nil
}

// 构建包含 include 函数的环境

// 动态编译表达式，使用环境变量

func (t *MixedTemplate) ExecuteFn(loadDataFunc func() map[string]any) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *MixedTemplate) ExecuteAsString(data map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *MixedTemplate) ExecuteFnAsString(loadDataFunc func() map[string]any) string {
	_ = "STUB: not implemented"
	return ""
}

func (t *MixedTemplate) IsNotVar() bool { _ = "STUB: not implemented"; return false }

func (t *MixedTemplate) HasVar() bool { _ = "STUB: not implemented"; return false }
