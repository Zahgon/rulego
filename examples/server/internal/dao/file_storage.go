package dao

import (
	"sync"

	"gopkg.in/ini.v1"
)

type FileStorage struct {
	filename string
	file     *ini.File
	lock     sync.RWMutex
}

func NewFileStorage(filename string) (*FileStorage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetSection 获取分区
func (d *FileStorage) GetSection(sectionName string) (*ini.Section, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *FileStorage) Get(sectionName string, keyName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (d *FileStorage) GetAll(sectionName string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Save 保存单个值
func (d *FileStorage) Save(sectionName, key, value string) error {
	_ = "STUB: not implemented"
	return nil
}

// 如果分区不存在，将会创建一个新的分区

// SaveList 保存多个值
func (d *FileStorage) SaveList(sectionName string, values map[string]string) error {
	_ = "STUB: not implemented"
	return nil
}

// 如果分区不存在，将会创建一个新的分区

// Delete 删除
func (d *FileStorage) Delete(sectionName string, keys ...string) error {
	_ = "STUB: not implemented"
	return nil
}

// SaveToFile 保存
func (d *FileStorage) SaveToFile() error { _ = "STUB: not implemented"; return nil }

//var FileStorageManager =NewFileStorageManager()

type FileStorageManager struct {
	// 文件存储 key=路径
	manager map[string]*FileStorage
	lock    sync.RWMutex
}

func NewFileStorageManager() *FileStorageManager { _ = "STUB: not implemented"; return nil }

func (f *FileStorageManager) Init(filename string) (*FileStorage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileStorageManager) Get(filename string) (*FileStorage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FileStorageManager) Delete(filename string) { _ = "STUB: not implemented"; return }
