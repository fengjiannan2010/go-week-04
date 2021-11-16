package logic

import (
	"os"
	"time"
)

type ObjectInfo struct {
	name         string
	size         int64
	lastModified time.Time
}

func NewObjectFileInfo(name string, size int64, lastModified time.Time) *ObjectInfo {
	return &ObjectInfo{name: name, size: size, lastModified: lastModified}
}

func (fi ObjectInfo) Name() string {
	return fi.name
}

// Size 实现os.FileInfo.Size()。
func (fi ObjectInfo) Size() int64 {
	return fi.size
}

// Mode 实现os.FileInfo.Mode()。返回0444，所有用户只读。
func (fi ObjectInfo) Mode() os.FileMode {
	return os.FileMode(0644)
}

// ModTime 实现os.FileInfo.ModTime()。
func (fi ObjectInfo) ModTime() time.Time {
	return fi.lastModified
}

// IsDir 实现os.FileInfo.IsDir()。肯定不是dir，永远返回false。s
func (fi ObjectInfo) IsDir() bool {
	return false
}

// Sys 实现os.FileInfo.Sys()。不知道干什么用的，返回nil。
func (fi ObjectInfo) Sys() interface{} {
	return nil
}
