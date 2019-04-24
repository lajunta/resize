package main

import (
	"flag"
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
)

func main() {
	flag.IntVar(&width, "-w", 800, "resized width , default 800.")
	flag.IntVar(&height, "-h", 800, "resized height , default 800.")
	flag.StringVar(&tmpPath, "-o", os.TempDir(), "output directory , default system temp dir.")
	flag.StringVar(&format, "-f", "jpg", "output image format , default jpg ,or png")
	flag.Parse()
	imgPath := os.Args[1]

	img, err := imgio.Open(imgPath)
	if err != nil {
		panic(err)
	}
	resized := transform.Resize(img, width, height, transform.Linear)
	// Or imgio.JPEGEncoder(95) as encoder for JPG with quality of 95%
	fname := "tmp." + format
	fpath := path.Join(tmpPath, fname)
	encoder := imgio.JPEGEncoder(95)
	if format == "png" {
		encoder = imgio.PNGEncoder()
	}
	if err := imgio.Save(fpath, resized, encoder); err != nil {
		panic(err)
	}
}
