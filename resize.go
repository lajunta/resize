package main

import (
	"flag"
	"log"
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
	fixed    bool
)

func destDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return os.TempDir()
	}
	return dir
}

func main() {

	flag.Usage = func() {
		usage(lang)
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.BoolVar(&fixed, "fixed", false, flagStrings["fixed"])
	flag.IntVar(&width, "w", 960, flagStrings["width"])
	flag.IntVar(&height, "h", 640, flagStrings["height"])
	flag.IntVar(&quality, "q", 95, flagStrings["quality"])
	flag.BoolVar(&dynamic, "d", true, flagStrings["dynamic"])
	flag.StringVar(&destPath, "o", destDirectory(), flagStrings["destPath"])
	flag.StringVar(&format, "f", "jpg", flagStrings["format"])
	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
	}

	imgPath := flag.Args()[0]
	img, err := imgio.Open(imgPath)
	if err != nil {
		log.Fatalln(errStrings["openfailed"])
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
	outPath := path.Join(destPath, fname)
	if fixed {
		outPath = imgPath
	}
	encoder := imgio.JPEGEncoder(quality)
	if format == "png" {
		encoder = imgio.PNGEncoder()
	}
	if err := imgio.Save(outPath, resized, encoder); err != nil {
		panic(err)
	}
}
