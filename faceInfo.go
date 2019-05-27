package main

type FaceInfo struct {
	FaceId        int
	FaceRectangle struct {
		Top    int
		Left   int
		Width  int
		Height int
	}
	FaceAttributes struct {
		Emotion struct {
			Anger     float64
			Contempt  float64
			Disgust   float64
			Fear      float64
			Happiness float64
			Neutral   float64
			Sadness   float64
			Surprise  float64
		}
	}
}
