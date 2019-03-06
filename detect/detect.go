package detect

import (
	"bytes"
	"encoding/json"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"

	"hfg/model"
)

func Detect(output string) ([]model.FaceInfo, error) {
	uriBase := os.Getenv("URL")
	subscriptionKey := os.Getenv("KEY1")

	const params = "?returnFaceId=true&returnFaceAttributes=emotion"
	uri := uriBase + "/detect" + params

	// Decode image
	imgBin, err := jpegToBin(output)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

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

	var faceInfo []model.FaceInfo
	err = json.Unmarshal(data, &faceInfo)
	if err != nil {
		return nil, err
	}

	// Format and display the Json result
	// jsonFormatted, _ := json.MarshalIndent(faceInfo, "", "  ")
	// fmt.Println(string(jsonFormatted))

	return faceInfo, nil
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
