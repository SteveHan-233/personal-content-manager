package main

import (
	"fmt"
	"io"
	"os"

	"github.com/kkdai/youtube/v2"
)

func main() {
	fmt.Println("vim-go")
	videoID := "CODqM_rzwtk"
	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		panic(err)
	}

	formats := video.Formats.WithAudioChannels()
	fmt.Println(formats)
	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}

	file, err := os.Create("video.mp4")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}
}
