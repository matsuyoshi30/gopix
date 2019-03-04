package main

import (
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"time"

	"hfg/detect"
	"hfg/env"
	"hfg/imgprocess"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s: wrong argument\n", os.Args[0])
		os.Exit(1)
	} else {
		err := env.Loadenv()
		logError(err)

		// input のファイル名を取得する
		filename := os.Args[1]

		// output のファイル名を決めておく
		now := time.Now().Format("20060102150405")
		output := "output_" + now + ".jpeg"

		// input が .jpeg でない場合は、 .jpeg に変換
		convertToJpeg(filename, output)

		// input を Face API へ
		faceInfo, err := detect.Detect(output)
		logError(err)

		mul := false
		if len(faceInfo) > 1 {
			mul = true
		}

		for k, R := range faceInfo {
			s := fmt.Sprintf("%f", R.FaceAttributes.Emotion.Happiness)
			fmt.Println("Happiness Value is " + s)
			if R.FaceAttributes.Emotion.Happiness >= 0.25 {
				// 帰ってきた json から顔領域をモザイク
				err = imgprocess.Pixelate(output, R, mul, k)
				logError(err)

				fmt.Println("=> CLEAR")
			} else {
				fmt.Println("=> PIXELATED")
			}
			fmt.Println("=> OPEN " + output)
		}
	}
}

func convertToJpeg(filename, output string) {
	file, err := os.Open(filename)
	logError(err)
	defer file.Close()

	img, _, err := image.Decode(file)
	logError(err)

	out, err := os.Create(output)
	logError(err)
	defer out.Close()

	jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}