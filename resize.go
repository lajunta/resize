package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var (
	width    int
	height   int
	destPath string
	format   string
	dynamic  bool
	quality  int
)

func destDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return os.TempDir()
	}
	return dir
}

func main() {
	flag.IntVar(&width, "w", 960, "resized width ")
	flag.IntVar(&height, "h", 640, "resized height ")
	flag.IntVar(&quality, "q", 95, "jpg image quality ")
	flag.BoolVar(&dynamic, "d", true, "adjust width and height dynamicly")
	flag.StringVar(&destPath, "o", destDirectory(), "output directory default current location ")
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
	fullName := path.Base(flag.Args()[0])
	ext := path.Ext(flag.Args()[0])
	oldName := strings.TrimSuffix(fullName, ext)
	fname := oldName + "_resized." + format
	fpath := path.Join(destPath, fname)
	encoder := imgio.JPEGEncoder(quality)
	if format == "png" {
		encoder = imgio.PNGEncoder()
	}
	if err := imgio.Save(fpath, resized, encoder); err != nil {
		panic(err)
	}
}
