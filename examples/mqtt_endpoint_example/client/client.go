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

// Package main demonstrates how to create an MQTT client
// that sends binary and JSON data to an MQTT endpoint server.
//
// This example shows:
// - Connecting to an MQTT broker
// - Publishing JSON data messages to specific topics
// - Publishing binary data messages to specific topics
// - Using different QoS levels
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/rulego/rulego/utils/mqtt"
)

const (
	// MQTT broker configuration
	// MQTT代理配置
	mqttServer = "127.0.0.1:1883"
	clientID   = "rulego_mqtt_client_example"
)

// SensorData represents a sample JSON message structure
// SensorData 表示示例JSON消息结构
type SensorData struct {
	SensorID     string  `json:"sensorId"`
	Temperature  float64 `json:"temperature"`
	Humidity     float64 `json:"humidity"`
	Timestamp    int64   `json:"timestamp"`
	Location     string  `json:"location"`
	BatteryLevel float64 `json:"batteryLevel"`
}

// SystemMessage represents system-level messages
// SystemMessage 表示系统级消息
type SystemMessage struct {
	MessageID string `json:"messageId"`
	Level     string `json:"level"`
	Source    string `json:"source"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

// DeviceCommand represents a binary command structure
// DeviceCommand 表示二进制命令结构
type DeviceCommand struct {
	DeviceID uint16 `json:"deviceId"`
	Command  uint8  `json:"command"`
	Value    uint32 `json:"value"`
}

func main() {
	fmt.Println("MQTT Client Example")
	fmt.Println("Connecting to MQTT broker at", mqttServer)

	// Create MQTT client configuration
	// 创建MQTT客户端配置
	config := mqtt.Config{
		Server:   mqttServer,
		Username: "",
		Password: "",
		QOS:      1,
		ClientID: clientID,
	}

	// Create MQTT client
	// 创建MQTT客户端
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mqtt.NewClient(ctx, config)
	if err != nil {
		log.Fatalf("Failed to create MQTT client: %v", err)
	}
	defer client.Close()

	fmt.Println("Connected to MQTT broker successfully")

	// Send JSON sensor data
	// 发送JSON传感器数据
	sendJSONData(client)

	// Send binary device commands
	// 发送二进制设备命令
	sendBinaryData(client)

	// Send system messages
	// 发送系统消息
	sendSystemMessages(client)

	fmt.Println("Client finished sending data")
}

// sendJSONData sends sample JSON sensor data to MQTT topics
// sendJSONData 向MQTT主题发送示例JSON传感器数据
func sendJSONData(client *mqtt.Client) { _ = "STUB: not implemented"; return }

// Sample sensor data
// 示例传感器数据

// High temperature

// High humidity

// Convert to JSON
// 转换为JSON

// Publish to sensor-specific topic
// 发布到传感器特定主题

// Wait a bit between messages
// 消息之间等待一点时间

// sendBinaryData sends sample binary device commands to MQTT topics
// sendBinaryData 向MQTT主题发送示例二进制设备命令
func sendBinaryData(client *mqtt.Client) { _ = "STUB: not implemented"; return }

// Sample binary commands
// 示例二进制命令

// SET_PARAMETER
// GET_STATUS
// RESET
// SET_THRESHOLD

// Create binary data (protocol: deviceId(2) + command(1) + value(4))
// 创建二进制数据（协议：deviceId(2) + command(1) + value(4)）

// Device ID (2 bytes, big endian)
// 设备ID（2字节，大端序）

// Command (1 byte)
// 命令（1字节）

// Value (4 bytes, big endian)
// 值（4字节，大端序）

// Publish to device-specific topic
// 发布到设备特定主题

// Wait a bit between messages
// 消息之间等待一点时间

// sendSystemMessages sends sample system messages
// sendSystemMessages 发送示例系统消息
func sendSystemMessages(client *mqtt.Client) { _ = "STUB: not implemented"; return }

// Sample system messages
// 示例系统消息

// Convert to JSON
// 转换为JSON

// Publish to system topic based on level
// 根据级别发布到系统主题

// Wait a bit between messages
// 消息之间等待一点时间
