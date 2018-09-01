package main

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const INPUT = "./pictures_org"
const OUTPUT = "./output"

func handle_picture(picname, outputpic string, width, height uint) {

	// open "test.jpg"
	file, err := os.Open(picname)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, height, img, resize.Lanczos3)

	out, err := os.Create(outputpic)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func main() {

	fmt.Println("will process in folder pictures_org")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input the picture width: ")
	widths, _ := reader.ReadString('\n')
	widths = widths[:len(widths)-2]
	fmt.Print("Please input the picture height: ")
	heights, _ := reader.ReadString('\n')
	heights = heights[:len(heights)-2]

	width, err := strconv.Atoi(widths)
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(heights)
	if err != nil {
		panic(err)
	}

	outputdir,err := filepath.Abs(OUTPUT)
	fmt.Println(outputdir)
	inputdir,err := filepath.Abs(INPUT)
	_, err = os.Stat(outputdir)
	if err != nil {
		error := os.Mkdir(outputdir, 0777)
		if error != nil {
			panic("create folder failed")
		}
	}

	files, err := ioutil.ReadDir(inputdir)
	if err != nil {
		panic(err)
	}

	bar := pb.New(len(files))
	// show percents (by default already true)
	bar.ShowPercent = true
	// show bar (by default already true)
	bar.ShowBar = true
	bar.ShowCounters = true
	bar.ShowTimeLeft = true

	for _, each := range (files) {
		inputfile := inputdir + "/" + each.Name()
		outputfile := outputdir + "/" + each.Name()
		handle_picture(inputfile, outputfile, uint(width), uint(height))
		bar.Increment()
	}
	bar.FinishPrint("convert successfully")
	reader.ReadString('\n')


}
