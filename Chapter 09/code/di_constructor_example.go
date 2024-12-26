package main

import "fmt"

type Video struct {
	content   []byte
	storage   VideoStorage
	publisher string
}

func NewVideo(ps VideoStorage, publisher string) *Video {
	return &Video{storage: ps, publisher: publisher}
}

func (video *Video) Load(content string) {
	video.content = video.storage.Load(content)
}
func (video *Video) Save(title string) {
	video.storage.Save(title, video.content)
}

type VideoStorage interface {
	Load(string) []byte
	Save(string, []byte)
}

type MediaStorage struct {
	video []byte
}

func NewMediaStorage() *MediaStorage {
	return &MediaStorage{
		video: []byte{},
	}
}

func (mStore *MediaStorage) Save(name string, contents []byte) {
	mStore.video = contents
}

func (mStore *MediaStorage) Load(name string) []byte {
	return mStore.video
}

func (mStore *MediaStorage) Type() string {
	return "MediaStorage"
}

func main() {
	storage := NewMediaStorage()
	video := NewVideo(storage, "STAR MUSIC")
	video.Load("Cricket Match NZ vs AUS 1987. Cricket match happened in Perth and Australia won the match")
	video.Save("Cricket Match NZ vs AUS 1987 ")
	fmt.Printf("video : %s \n", video.publisher)
}
