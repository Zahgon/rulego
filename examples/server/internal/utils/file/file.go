package file

import (
	"time"
)

// WithTimestamp 包含文件路径和解析出的时间戳
type WithTimestamp struct {
	Path      string
	Timestamp time.Time
}

// SortFilesByTimestamp 解析文件列表中的时间戳，返回按时间戳排序的文件列表
func SortFilesByTimestamp(files []string) []WithTimestamp { _ = "STUB: not implemented"; return nil }

// 使用 sort.Sort 进行排序

// 解析文件名中的时间戳
func parseTimestampFromFilename(filename string) (time.Time, error) {
	_ = "STUB: not implemented"
	// 使用filepath.Base获取文件名
	return *new(time.Time), nil
}

// ByTimestamp 实现 sort.Interface 接口
type ByTimestamp []WithTimestamp

func (f ByTimestamp) Len() int           { _ = "STUB: not implemented"; return 0 }
func (f ByTimestamp) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (f ByTimestamp) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
