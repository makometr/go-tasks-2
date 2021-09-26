package main

import "os"

// Встраиваем для поддержки поведения
type youtubeCounter struct {
	baseViewCounter
	saveFile *os.File
}

func newYoutubeCounter(url string) viewCounter {
	return &youtubeCounter{
		baseViewCounter: baseViewCounter{
			views:   0,
			viewURL: "youtube.com/" + url,
		},
	}
}

func (yc *youtubeCounter) addView() {
	// write to saveFile new view
	yc.baseViewCounter.addView()
}

type VKCounter struct {
	baseViewCounter
	imageObj *interface{}
}

func newVKCounter(url string) viewCounter {
	return &VKCounter{baseViewCounter: baseViewCounter{
		views:   0,
		viewURL: "vk.com/" + url,
	},
	}
}

func (vkc *VKCounter) addView() {
	if vkc.imageObj != nil {
		vkc.baseViewCounter.addView()
	}
}
