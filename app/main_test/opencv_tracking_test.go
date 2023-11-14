// What it does:
//
// This example tracks objects such as cars or people passing across
// a horizontal or vertical line by using the Moments method.
// The Moments algorithm is not that accurate for counting multiple objects,
// however it is execution efficient.
//
// How to run:
//
// 		go run ./counter/main.go /path/to/video.avi 400 y 10
//

package unittest

import (
	"fmt"
	"image/color"
	"strconv"
	"testing"

	"gocv.io/x/gocv"
	"gocv.io/x/gocv/contrib"
)

func TestOpenCvTracking(t *testing.T) {
	deviceID, _ := strconv.Atoi("0")
	// open webcam
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Printf("Error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Tracking")
	defer window.Close()

	// create a tracker instance
	// (one of MIL, KCF, TLD, MedianFlow, Boosting, MOSSE or CSRT)
	tracker := contrib.NewTrackerKCF()
	defer tracker.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// read an initial image
	if ok := webcam.Read(&img); !ok {
		fmt.Printf("cannot read device %v\n", deviceID)
		return
	}

	// let the user mark a ROI to track
	rect := gocv.SelectROI("Tracking", img)
	if rect.Max.X == 0 {
		fmt.Printf("user cancelled roi selection\n")
		return
	}

	// initialize the tracker with the image & the selected roi
	init := tracker.Init(img, rect)
	if !init {
		fmt.Printf("Could not initialize the Tracker")
		return
	}

	// color for the rect to draw
	blue := color.RGBA{0, 0, 255, 0}
	fmt.Printf("Start reading device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// update the roi
		rect, _ := tracker.Update(img)

		// draw it.
		gocv.Rectangle(&img, rect, blue, 3)

		// show the image in the window, and wait 10 millisecond
		window.IMShow(img)
		if window.WaitKey(10) >= 0 {
			break
		}
	}
}
