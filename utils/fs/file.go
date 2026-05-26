/*
 * Copyright 2025 The RuleGo Authors.
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

package fs

import (
	"os"
)

// File defines the interface for file storage
// Provides file based storage and retrieval functionality
// Implementation classes must ensure thread safety
type File interface {
	// Save stores a file in storage
	// Parameters:
	//   - path: file path (string)
	//   - data: file content ([]byte)
	// Returns:
	//   - error: returns error if save fails
	Save(path string, data []byte) error
	// Get retrieves a file from storage by path
	// Parameters:
	//   - path: file path to lookup (string)
	// Returns:
	//   - []byte: file content
	//   - error: returns error if not exists or other error
	Get(path string) ([]byte, error)
	// Delete removes a file by path
	// Parameters:
	//   - path: file path to delete (string)
	// Returns:
	//   - error: returns error if delete fails
	Delete(path string) error
	// SaveAppend appends data to a file in storage
	// Parameters:
	//   - path: file path (string)
	//   - data: file content ([]byte)
	// Returns:
	//   - error: returns error if save fails
	SaveAppend(path string, data []byte) error
	// GetFilePaths returns file paths matching the pattern
	// Parameters:
	//   - loadFilePattern: glob pattern for files
	//   - excludedPatterns: patterns to exclude
	// Returns:
	//   - []string: list of matching file paths
	//   - error: returns error if list fails
	GetFilePaths(loadFilePattern string, excludedPatterns ...string) ([]string, error)
	// IsExist checks if a path exists
	// Parameters:
	//   - path: file path (string)
	// Returns:
	//   - bool: true if exists, false otherwise
	IsExist(path string) bool
	// CreateDirs creates directories recursively
	// Parameters:
	//   - path: directory path (string)
	// Returns:
	//   - error: returns error if creation fails
	CreateDirs(path string) error
	// Name returns the name of the storage
	Name() string
}

// LocalFileStorage implements types.File using local file system
type LocalFileStorage struct {
	name string
}

func NewLocalFileStorage() *LocalFileStorage { _ = "STUB: not implemented"; return nil }

func (f *LocalFileStorage) Name() string { _ = "STUB: not implemented"; return "" }

func (f *LocalFileStorage) Get(path string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *LocalFileStorage) Save(path string, data []byte) error {
	_ = "STUB: not implemented"
	// Create dirs if not exist
	return nil
}

func (f *LocalFileStorage) SaveAppend(path string, data []byte) error {
	_ = "STUB: not implemented"
	// Create dirs if not exist
	return nil
}

// Open file in append mode, create if not exists

func (f *LocalFileStorage) Delete(path string) error { _ = "STUB: not implemented"; return nil }

// GetFilePaths returns file paths matching the pattern
func (f *LocalFileStorage) GetFilePaths(loadFilePattern string, excludedPatterns ...string) ([]string, error) {
	_ = "STUB: not implemented"
	// 分割输入参数为目录和文件名
	return nil, nil
}

// 遍历目录

// 如果是文件，且文件名匹配输入参数

// 跳过该子目录

// IsExist checks if a path exists
func (f *LocalFileStorage) IsExist(path string) bool { _ = "STUB: not implemented"; return false }

// CreateDirs creates directories recursively
func (f *LocalFileStorage) CreateDirs(path string) error { _ = "STUB: not implemented"; return nil }

func isMatch(d os.DirEntry, patterns ...string) bool { _ = "STUB: not implemented"; return false }

var DefaultFile = NewLocalFileStorage()
