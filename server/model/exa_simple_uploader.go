package model

type ExaSimpleUploader struct {
	ChunkNumber      string `json:"chunkNumber" gorm:"comment:Current slice mark"`
	CurrentChunkSize string `json:"currentChunkSize" gorm:"comment:Current slice capacity"`
	CurrentChunkPath string `json:"currentChunkPath" gorm:"comment:Sliced local path"`
	TotalSize        string `json:"totalSize" gorm:"comment:total capacity"`
	Identifier       string `json:"identifier" gorm:"comment:File ID (MD5)"`
	Filename         string `json:"filename" gorm:"comment:file name"`
	TotalChunks      string `json:"totalChunks" gorm:"comment:Total number of slices"`
	IsDone           bool   `json:"isDone" gorm:"comment:Whether to upload finish"`
	FilePath         string `json:"filePath" gorm:"comment:File local path"`
}
