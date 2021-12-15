// Package example
// @program: arco-design-pro-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2021-12-14 11:06
package example

import "github.com/lliuhuan/arco-design-pro-gin/global"

// file struct, 文件结构体
type ExaFile struct {
	global.AdpModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.AdpModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
