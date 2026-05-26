//go:build go1.21

package engine

// startMonitoring starts monitoring parent contexts using AfterFunc for Go 1.21+.
// This avoids creating a dedicated goroutine for waiting.
// startMonitoring 启动监控父上下文，对于 Go 1.21+ 使用 AfterFunc。
// 这避免了创建专用的等待协程。
func (c *combinedCancelContext) startMonitoring() { _ = "STUB: not implemented"; return }

// If either is already done, cancel immediately

// Use context.AfterFunc to register cancellation callbacks
// This is more efficient than a blocking goroutine

// Register cleanup when the combined context is done
// This ensures we stop monitoring parents when we are done
