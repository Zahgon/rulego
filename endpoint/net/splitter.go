package net

import (
	"bufio"
)

// PacketMode 数据包分割模式
type PacketMode string

const (
	// PacketModeLine 按行分割（默认模式，以\n或\r\n分割）
	PacketModeLine PacketMode = "line"
	// PacketModeFixed 固定长度分割
	PacketModeFixed PacketMode = "fixed"
	// PacketModeDelimiter 自定义分隔符分割
	PacketModeDelimiter PacketMode = "delimiter"

	// 长度前缀模式（4种组合）
	// PacketModeLengthPrefixLE 长度前缀，小端序，长度不包含前缀
	PacketModeLengthPrefixLE PacketMode = "length_prefix_le"
	// PacketModeLengthPrefixBE 长度前缀，大端序，长度不包含前缀
	PacketModeLengthPrefixBE PacketMode = "length_prefix_be"
	// PacketModeLengthPrefixLEInc 长度前缀，小端序，长度包含前缀
	PacketModeLengthPrefixLEInc PacketMode = "length_prefix_le_inc"
	// PacketModeLengthPrefixBEInc 长度前缀，大端序，长度包含前缀
	PacketModeLengthPrefixBEInc PacketMode = "length_prefix_be_inc"
)

// String 返回模式的字符串表示
func (p PacketMode) String() string {
	_ = "STUB: not implemented"

	// IsValid 检查模式是否有效
	return ""
}

func (p PacketMode) IsValid() bool { _ = "STUB: not implemented"; return false }

// IsLengthPrefixMode 检查是否为长度前缀模式
func (p PacketMode) IsLengthPrefixMode() bool { _ = "STUB: not implemented"; return false }

// IsBigEndian 是否为大端序
func (p PacketMode) IsBigEndian() bool { _ = "STUB: not implemented"; return false }

// 默认小端序

// IncludesPrefix 长度是否包含前缀本身
func (p PacketMode) IncludesPrefix() bool { _ = "STUB: not implemented"; return false }

// 默认不包含

// PacketSplitter 数据包分割器接口
type PacketSplitter interface {
	// ReadPacket 从连接中读取一个完整的数据包
	ReadPacket(reader *bufio.Reader) ([]byte, error)
}

// LineSplitter 按行分割的数据包分割器
type LineSplitter struct{}

func (s *LineSplitter) ReadPacket(reader *bufio.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 去掉换行符分隔符

// 如果是\r\n，也去掉\r

// FixedLengthSplitter 固定长度数据包分割器
type FixedLengthSplitter struct {
	PacketSize int
}

func (s *FixedLengthSplitter) ReadPacket(reader *bufio.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DelimiterSplitter 自定义分隔符数据包分割器
type DelimiterSplitter struct {
	Delimiter []byte
}

func (s *DelimiterSplitter) ReadPacket(reader *bufio.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// 检查是否匹配分隔符

// 找到完整分隔符，返回包含分隔符的完整数据

// LengthPrefixSplitter 长度前缀数据包分割器
type LengthPrefixSplitter struct {
	PrefixSize     int  // 长度前缀的字节数（1-4字节）
	BigEndian      bool // 是否使用大端序
	IncludesPrefix bool // 长度是否包含前缀本身
	MaxPacketSize  int  // 最大数据包大小
}

func (s *LengthPrefixSplitter) ReadPacket(reader *bufio.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	// 读取长度前缀
	return nil, nil
}

// 解析长度值

// 大端序3字节：在前面补0变成4字节

// 小端序3字节：将3字节按小端序组合

// 检查数据包大小限制

// 根据IncludesPrefix确定数据长度

// 读取数据部分

// 返回包含长度前缀的完整数据包

// CreatePacketSplitter 根据配置创建数据包分割器
func CreatePacketSplitter(config Config) (PacketSplitter, error) {
	_ = "STUB: not implemented"
	// 默认为line模式
	return *new(PacketSplitter), nil
}

// 解析分隔符（支持十六进制格式）

// 十六进制格式: 0x0A0D

// 直接使用字符串作为分隔符
