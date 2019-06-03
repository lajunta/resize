package main

import (
	"fmt"
	"os"
	"strings"
)

var flagStrings = make(map[string]string)
var errStrings = make(map[string]string)
var lang string

func init() {
	lang, _ = os.LookupEnv("LANG")
	setFlagStrings()
}

func setFlagStrings() {
	zh := strings.Contains(lang, "zh")
	switch zh {
	case true:
		errStrings["openfailed"] = "打开文件失败"
		errStrings["specifypath"] = "请指定文件或文件夹路径"
		flagStrings["fixed"] = "输出图片路径和输入路径的完全一样，将覆盖源文件"
		flagStrings["width"] = "输出图片的宽度"
		flagStrings["height"] = "输出图片的高度"
		flagStrings["quality"] = "输出图片的质量,当输出格式为jpg时有效"
		flagStrings["dynamic"] = "是否根据高度自动调整宽度"
		flagStrings["destPath"] = "输出图片的路径,默认为当前目录"
		flagStrings["format"] = "输出图片的格式(可选jpg,png),默认为jpg"
	default:
		errStrings["openfailed"] = "open image file failed"
		errStrings["specifypath"] = "please specify file or directory path"
		flagStrings["fixed"] = "output directory and name keep same as source"
		flagStrings["width"] = "resized width "
		flagStrings["height"] = "resized height "
		flagStrings["quality"] = "jpg image quality "
		flagStrings["dynamic"] = "adjust width and height dynamicly"
		flagStrings["destPath"] = "output directory default current location "
		flagStrings["format"] = "output image format (jpg or png)"
	}
}

func usage(lang string) {
	zh := strings.Contains(lang, "zh")
	switch zh {
	case true:
		fmt.Print(`
使用方法: 

  resize [options] image_file_path | images_directory
	
举例:

  resize -w 640 -o /tmp /home/user/sample.jpg 
	
  调整 sample.jpg 宽度 640 
  保存到 tmp 临时目录,命名为 sample_resized.jpg 
	
选项:

`)
	default:
		fmt.Print(`
Usage: 

  resize [options] image_file_path

Example:

  resize -w 640 -o /tmp /home/user/sample.jpg 
	
  resize sample.jpg to width 640 
  and save to tmp with name sample_resized.jpg 
	
Options: 

`)
	}
}
