package detect

import (
	"hfg/model"

	"gocv.io/x/gocv"
)

func Detect(output string) ([]model.FaceInfo, error) {
	img := gocv.IMRead(output, gocv.IMReadColor)
	if img.Empty() {
		// TODO: error handling
		return nil, nil
	}
	defer img.Close()

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load("data/haarcascade_frontalface_default.xml")

	rects := classifier.DetectMultiScale(img)
	if len(rects) == 0 {
		// TODO: couldn't find face
		return nil, nil
	}

	fi := make([]model.FaceInfo, 0)
	for i, r := range rects {
		f := model.FaceInfo{
			FaceId: i,
		}
		f.FaceRectangle.Top = r.Min.Y
		f.FaceRectangle.Left = r.Min.X
		f.FaceRectangle.Width = r.Max.X - r.Min.X
		f.FaceRectangle.Height = r.Max.Y - r.Min.Y

		fi = append(fi, f)
	}

	return fi, nil
}
