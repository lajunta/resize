package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var (
	width   int
	height  int
	tmpPath string
	format  string
	dynamic bool
)

func main() {
	flag.IntVar(&width, "w", 960, "resized width ")
	flag.IntVar(&height, "h", 640, "resized height ")
	flag.BoolVar(&dynamic, "d", true, "adjust width and height dynamicly")
	flag.StringVar(&tmpPath, "o", os.TempDir(), "output directory ")
	flag.StringVar(&format, "f", "jpg", "output image format (jpg or png)")
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("\nUsage: resize [options] image_file_path\n\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
	imgPath := flag.Args()[0]
	img, err := imgio.Open(imgPath)
	if err != nil {
		panic(err)
	}

	if dynamic {
		rect := img.Bounds()
		oldWidth := rect.Size().X
		oldHeight := rect.Size().Y

		if oldWidth < width {
			width = oldWidth
			height = oldHeight
		} else {
			rate := float32(width) / float32(oldWidth)
			height = int(float32(oldHeight) * rate)
		}
	}

	resized := transform.Resize(img, width, height, transform.Linear)

	fname := "tmp." + format
	fmt.Println(tmpPath, format)
	fpath := path.Join(tmpPath, fname)
	encoder := imgio.JPEGEncoder(95)
	if format == "png" {
		encoder = imgio.PNGEncoder()
	}
	if err := imgio.Save(fpath, resized, encoder); err != nil {
		panic(err)
	}
}
