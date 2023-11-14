package unittest

import (
	"testing"

	"gocv.io/x/gocv"
)

// https://gocv.io/writing-code/hello-video/
func TestOpenCv1(t *testing.T) {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
