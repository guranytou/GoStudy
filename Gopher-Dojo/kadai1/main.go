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

type ConvImage struct {
	InputFileType  string
	OutputFileType string
	SrcPath        string
	DirPath        string
}

func main() {
	cmd := new(ConvImage)
	flag.StringVar(&cmd.InputFileType, "e", "jpeg", "入力ファイルの形式を指定する")
	flag.StringVar(&cmd.OutputFileType, "o", "png", "出力ファイルの形式を指定する")
	flag.Parse()

	cmd.SrcPath = flag.Arg(0)
	cmd.DirPath = flag.Arg(1)

	fileTypeCheck(cmd.InputFileType)
	fileTypeCheck(cmd.OutputFileType)
	dirCheck(cmd.SrcPath)
	dirCheck(cmd.DirPath)

	err := imageconvert.Convert(cmd.InputFileType, cmd.OutputFileType, cmd.SrcPath, cmd.DirPath)
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
