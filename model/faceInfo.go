package model

type FaceInfo struct {
	FaceId        string `json:"faceId"`
	FaceRectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}
	FaceAttributes struct {
		Emotion struct {
			Anger     float64 `json:"anger"`
			Contempt  float64 `json:"contempt"`
			Disgust   float64 `json:"disgust"`
			Fear      float64 `json:"fear"`
			Happiness float64 `json:"happiness"`
			Neutral   float64 `json:"neutral"`
			Sadness   float64 `json:"sadness"`
			Surprise  float64 `json:"surprise"`
		}
	}
}
