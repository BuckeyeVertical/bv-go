package mapper

import (
	"fmt"
	"time"
	"gocv.io/x/gocv"
)

type Mapper struct {
	frames      []gocv.Mat
	stitchedMap *gocv.Mat
}

func NewMapper() *Mapper {
	return &Mapper{
		frames: make([]gocv.Mat, 0),
	}
}

func (m *Mapper) GetCurrentMap() *gocv.Mat {
	return m.stitchedMap
}

func (m *Mapper) AddFrame(frame gocv.Mat) {
	m.frames = append(m.frames, frame)

	fmt.Println("Mock stitching frame into map...")
	time.Sleep(5 * time.Second)
}
