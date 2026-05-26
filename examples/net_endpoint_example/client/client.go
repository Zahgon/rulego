/*
 * Copyright 2024 The RuleGo Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package main demonstrates how to create a NET client
// that sends binary data to a NET endpoint server.
//
// This example shows:
// - Connecting to a NET endpoint server
// - Sending binary data messages
// - Receiving and parsing binary server responses in hex format
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

const (
	// Server address to connect to
	// 要连接的服务器地址
	serverAddr = "localhost:8088"
)

// DeviceCommand represents a binary command structure
// DeviceCommand 表示二进制命令结构
type DeviceCommand struct {
	DeviceID uint16 `json:"deviceId"`
	Command  uint8  `json:"command"`
	Value    uint32 `json:"value"`
}

func main() {
	fmt.Println("NET Client Example")
	fmt.Println("Connecting to server at", serverAddr)

	// Connect to the server
	// 连接到服务器
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connected to server successfully")

	// Create a reader for server responses
	// 创建用于服务器响应的读取器
	reader := bufio.NewReader(conn)

	// Send binary data examples
	// 发送二进制数据示例
	sendBinaryData(conn, reader)

	fmt.Println("Client finished sending data")
}

// sendBinaryData sends sample binary messages to the server
// sendBinaryData 向服务器发送示例二进制消息
func sendBinaryData(conn net.Conn, reader *bufio.Reader) { _ = "STUB: not implemented"; return }

// Sample binary commands
// 示例二进制命令

// Create binary data (simple protocol: deviceId(2) + command(1) + value(4))
// 创建二进制数据（简单协议：deviceId(2) + command(1) + value(4)）

// Device ID (2 bytes, big endian)
// 设备 ID（2 字节，大端序）

// Command (1 byte)
// 命令（1 字节）

// Value (4 bytes, big endian)
// 值（4 字节，大端序）

// Send binary data
// 发送二进制数据

// Read server response
// 读取服务器响应

// Wait a bit between messages
// 消息之间等待一点时间

// readAndDisplayResponse reads server response and displays it with enhanced formatting
// readAndDisplayResponse 读取服务器响应并以增强格式显示
func readAndDisplayResponse(reader *bufio.Reader, dataType string, startTime time.Time) error {
	_ = "STUB: not implemented"
	// For binary data, read raw bytes until newline
	// 对于二进制数据，读取原始字节直到换行符
	return nil
}

// Calculate response time
// 计算响应时间

// Print binary response in hex format
// 以hex格式打印二进制响应

// Print as ASCII (for debugging)
// 打印为ASCII（用于调试）

// Decode binary response (without newline)
// 解码二进制响应（不包括换行符）
// Remove newline
