package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/guranytou/imageconvert"
)

var (
	inputFileType  string
	outputFileType string
)

func main() {
	flag.StringVar(&inputFileType, "e", "jpeg", "入力ファイルの形式を指定する")
	flag.StringVar(&outputFileType, "o", "png", "出力ファイルの形式を指定する")
	flag.Parse()

	fileTypeCheck(inputFileType)
	fileTypeCheck(outputFileType)
	dirCheck(flag.Arg(0))
	dirCheck(flag.Arg(1))

	err := imageconvert.Convert(inputFileType, outputFileType, flag.Arg(0), flag.Arg(1))
	if err != nil {
		fmt.Println("Err:", err)
	}

}

func fileTypeCheck(filetype string) error {
	err := imageconvert.FileTypeCheck(filetype)
	if err != nil {
		fmt.Println("Err:", err)
		os.Exit(1)
	}
	return nil
}

func dirCheck(dir string) error {
	err := imageconvert.DirCheck(dir)
	if err != nil {
		fmt.Println("Err: ", err)
		os.Exit(1)
	}
	return nil
}
