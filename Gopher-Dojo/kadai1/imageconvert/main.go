package imageconvert

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

func FileTypeCheck(filetype string) error {
	switch filetype {
	case "jpeg", "png", "jpg", "gif":
		return nil
	default:
		return errors.New("ファイルタイプが不正です")
	}
}

func DirCheck(dir string) error {
	if _, err := os.Stat(dir); err != nil {
		return errors.New("指定されたディレクトリは存在しません")
	}
	return nil
}

func getFilePathList(path string, filetype string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		ext := strings.Replace(filepath.Ext(path), ".", "", -1)

		if ext == "" {
			return nil
		} else if ext == filetype {
			fileList = append(fileList, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileList, nil
}

func Convert(inputFileType string, outputFileType string, srcPath string, dirPath string) error {
	fileList, err := getFilePathList(srcPath, inputFileType)
	if err != nil {
		fmt.Println("Err", err)
	}
	if len(fileList) < 1 {
		return errors.New("該当するファイルはありませんでした")
	}

	for _, fileName := range fileList {
		file, err := os.Open(fileName)
		defer file.Close()
		if err != nil {
			return errors.New("ファイルが開けませんでした")
		}

		basename := filepath.Base(fileName[:len(fileName)-len(filepath.Ext(fileName))]) + "." + outputFileType
		filePathName := filepath.Join(dirPath, basename)

		output, err := os.Create(filePathName)
		if err != nil {
			return errors.New("ファイル作成に失敗しました")
		}
		defer output.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return errors.New("画像のデコードに失敗しました")
		}

		switch outputFileType {
		case "jpeg", "jpg":
			jpeg.Encode(output, img, nil)
		case "gif":
			gif.Encode(output, img, nil)
		default:
			png.Encode(output, img)
		}
		fmt.Println("ファイルの変換に成功しました")
	}
	return nil
}
