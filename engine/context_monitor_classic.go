//go:build !go1.21

package engine

// startMonitoring starts monitoring parent contexts using a goroutine for Go < 1.21.
// This provides backward compatibility for older Go versions.
// startMonitoring 启动监控父上下文，对于 Go < 1.21 使用协程。
// 这为旧版本的 Go 提供了向后兼容性。
func (c *combinedCancelContext) startMonitoring() { _ = "STUB: not implemented"; return }

// If either is already done, cancel immediately

// Use goroutine for backward compatibility with Go < 1.21
// 启动单个协程来监控两个上下文

// 如果 userCtx 是不可取消的（如 context.Background() 或 context.TODO()），
// 我们不需要监听它，只需要监听 shutdownCtx 和 internal ctx
// If userCtx is not cancellable (Done() returns nil), we don't need to select on it

// Internal context cancelled, exit goroutine
// 内部上下文已取消，退出协程

// Internal context cancelled, exit goroutine
// 内部上下文已取消，退出协程
