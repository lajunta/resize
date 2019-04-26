# Intro

Resize is a simple image tool specially for resize image 

## Usage

Just type `resize` to see usage.

Or check the following:

```
Usage: 

  resize [options] image_file_path

Example:

  resize -w 640 -o /tmp /home/user/sample.jpg 
	
  resize sample.jpg to width 640 
  and save to tmp with name sample_resized.jpg 
	
Options: 

  -d	adjust width and height dynamicly (default true)
  -f string
    	output image format (jpg or png) (default "jpg")
  -fixed
    	output directory and name keep same as source
  -h int
    	resized height  (default 640)
  -o string
    	output directory default current location  (default "/home/zxy/goapps/src/resize")
  -q int
    	jpg image quality  (default 95)
  -w int
    	resized width  (default 960)

```

## Build

First install MinGW-w64 into your linux system.

Second open build.sh and check MinGW-w64 path , modify it if it is wrong

Third run build.sh , and copy binary file in dist directory to your system path

## License

This project is licensed under the MIT license. Please read the LICENSE file.