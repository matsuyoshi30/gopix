package detect

import (
	"bytes"
	"encoding/json"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"hfg/model"
)

func Detect(output string) ([]model.FaceRect, error) {
	uriBase := os.Getenv("URL")
	subscriptionKey := os.Getenv("KEY1")

	const params = "?returnFaceId=true"
	uri := uriBase + "/detect" + params

	// Decode image
	imgBin, err := jpegToBin(output)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(imgBin))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/octet-stream")
	req.Header.Add("Ocp-Apim-Subscription-Key", subscriptionKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var faceRect []model.FaceRect
	err = json.Unmarshal(data, &faceRect)
	if err != nil {
		return nil, err
	}

	return faceRect, nil
}

func jpegToBin(imgpath string) ([]byte, error) {
	buf := new(bytes.Buffer)

	imgfile, err := os.Open(imgpath)
	if err != nil {
		return buf.Bytes(), err
	}

	img, err := jpeg.Decode(imgfile)
	if err != nil {
		return buf.Bytes(), err
	}

	if err = jpeg.Encode(buf, img, nil); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}
