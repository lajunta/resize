package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
)

var (
	width           int
	height          int
	destPath        string
	format          string
	dynamic         bool
	quality         int
	fixed           bool
	sourceDirectory string
)

func destDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return os.TempDir()
	}
	return dir
}

func walker(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	fileName := info.Name()

	if info.IsDir() {
		// remove left source path string
		ts := strings.TrimLeft(path, sourceDirectory)
		midDir := filepath.Join(destPath, ts)
		os.MkdirAll(midDir, 0755)
	}

	if strings.HasSuffix(fileName, "jpg") || strings.HasSuffix(fileName, "png") {
		outPath := strings.Replace(path, sourceDirectory, destPath, -1)
		ext := filepath.Ext(fileName)
		oldName := strings.TrimSuffix(fileName, ext)

		outPath = filepath.Join(filepath.Dir(outPath), oldName+"."+format)
		handleImage(path, outPath)
	}

	return nil
}

func handleImage(path string, outPath string) {
	encoder := imgio.JPEGEncoder(quality)
	if format == "png" {
		encoder = imgio.PNGEncoder()
	}
	img, err := imgio.Open(path)
	if err != nil {
		log.Fatalln(errStrings["openfailed"])
	}

	// restrict width and height
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
	if err := imgio.Save(outPath, resized, encoder); err != nil {
		panic(err)
	}
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

	if destPath != destDirectory() {
		err := os.MkdirAll(destPath, 0755)
		if err != nil {
			fmt.Println("Can't makedir ", destPath)
		}
	}

	// handle single image file
	if len(flag.Args()) == 0 {
		log.Println(errStrings["specifypath"])
		flag.Usage()
		os.Exit(1)
	}

	imgPath := flag.Args()[0]
	fi, err := os.Stat(imgPath)
	if os.IsNotExist(err) {
		log.Fatalln(errStrings["openfailed"])
		os.Exit(1)
	}

	if fi.IsDir() {
		sourceDirectory = strings.TrimRight(imgPath, "/")
		// make destPath
		if destPath == destDirectory() {
			destPath = sourceDirectory + "_resized_w" + strconv.Itoa(width)
		}

		filepath.Walk(sourceDirectory, walker)
		os.Exit(0)
	}


	fullName := path.Base(flag.Args()[0])
	ext := path.Ext(flag.Args()[0])
	oldName := strings.TrimSuffix(fullName, ext)
	fname := oldName + "_resized." + format
	outPath := path.Join(destPath, fname)

	handleImage(imgPath, outPath)
}
