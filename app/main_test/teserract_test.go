package unittest

import (
	"fmt"
	"log"
	"os/exec"
	"testing"

	"github.com/otiai10/gosseract/v2"
)

/*
brew install tesseract leptonica
go get -u github.com/otiai10/gosseract

after install and import gosseract make export on path project folder

export LIBRARY_PATH="/opt/homebrew/lib"
export CPATH="/opt/homebrew/include"
*/
//example on github : https://github.com/otiai10/gosseract/blob/main/example_test.go
//ocr API : https://github.com/otiai10/ocrserver/wiki/API-Endpoints
func TestImageToText(t *testing.T) {
	client := gosseract.NewClient()
	defer client.Close()

	err := client.SetConfigFile("images/goserract.config")
	// util.IsErrorDoPrint(err)
	client.Trim = true
	client.SetImage("images/crop_plat_mobil.png")
	// client.SetPageSegMode(gosseract.PSM_SINGLE_BLOCK)
	// client.GetBoundingBoxes(gosseract.RIL_WORD)
	client.SetLanguage("ind")
	// client.SetTessdataPrefix()

	// out, err := client.HOCRText()
	// util.IsErrorDoPrint(err)
	// fmt.Println("out : ", out)
	text, _ := client.Text()
	fmt.Println("text : ", text)

	// page := new(Page)
	// err = xml.Unmarshal([]byte(out), page)
	// client.SetTessdataPrefix(testModelDir)

	ot, err := exec.Command("ls", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(ot))
}
