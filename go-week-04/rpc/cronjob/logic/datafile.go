package logic

import (
	"io"
	"os"
)

type DataFile struct {
	FileInfo   os.FileInfo
	FileReader ReadSeekCloser
	FileWriter WriteSeekerCloser
	DiscPath   string
	OffSetStart int64
	OffSetEnd   int64
}

type WriteSeekerCloser interface {
	io.WriteCloser
	io.Seeker
}

type ReadSeekCloser interface {
	io.ReadSeekCloser
}
