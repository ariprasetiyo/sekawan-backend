package unittest

import (
	"testing"

	"gocv.io/x/gocv"
)

func TestOpenCv(t *testing.T) {

	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Camera")
	image := gocv.NewMat()

	for {
		webcam.Read(&image)
		window.IMShow(image)
		window.WaitKey(1)
	}

}
