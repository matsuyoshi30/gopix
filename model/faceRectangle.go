package model

type FaceRect struct {
	FaceId        string `json:"faceId"`
	FaceRectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	}
}
