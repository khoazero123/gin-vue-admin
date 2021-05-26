package model

import (
	"gin-vue-admin/global"
)

// file struct, File structure
type ExaFile struct {
	global.GVA_MODEL
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, Slice structure
type ExaFileChunk struct {
	global.GVA_MODEL
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
