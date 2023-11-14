package unittest

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"strconv"
	"testing"

	"gocv.io/x/gocv"
)

// https://gocv.io/writing-code/hello-video/
// What it does:
//
// This example uses the CascadeClassifier class to detect faces,
// and draw a rectangle around each of them, before displaying them within a Window.
//
// How to run:
//
// facedetect [camera ID] [classifier XML file]
//
//	go run ./cmd/facedetect/main.go 0 data/haarcascade_frontalface_default.xml
//
// https://docs.opencv.org/3.4/db/d28/tutorial_cascade_classifier.html
// https://gocv.io/writing-code/more-examples/
func TestOpenCv2(t *testing.T) {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetect [camera ID] [classifier XML file]")
		return
	}

	// parse args
	deviceID, _ := strconv.Atoi("0")
	xmlFile := "/Users/chrisferdian/Development/backend/project/golang-project/sekawan-backend/app/main_test/data/haarcascade_russian_plate_number.xml"

	// open webcam
	webcam, err := gocv.VideoCaptureDevice(int(0))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("cannot read device %d\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		// draw a rectangle around each face on the original image,
		// along with text identifying as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
