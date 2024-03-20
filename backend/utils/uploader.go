package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Uploader(w http.ResponseWriter, r *http.Request, size int, formFileName string, imageName string) (string, error) {
	//*checking the file 's size
	if r.Method == "POST" {
		maxsize := size * 1024 * 1024
		err := r.ParseMultipartForm(int64(maxsize))
		if err != nil {
			return "", errors.New("❌ could not allocted memory due to empty file in form")
		}
		file, header, err := r.FormFile(formFileName)
		if err != nil { //!empty value sent wwhile submitting form
			fmt.Println("🚫 empty image")
			return "", nil
		}
		defer file.Close()

		if header.Size > int64(maxsize) { // Check if file size is greater than 5 MB
			fmt.Printf("⚠ Image exceeds %vMB", size)
			return "", errors.New("file size exceeds limit")
		}
		fmt.Println("✅ image size checked")

		//*creating a copy of the uploaded in the server
		//!--checking extension validity
		if !IsValidImageType(header.Filename) {
			fmt.Println("⚠ Wrong image extension")
			return "", errors.New("invalid extension")
		}

		if imageName == "" {
			ImgName, errImg := GenImageName(header.Filename)
			imageName = ImgName
			if errImg != nil {
				fmt.Println("🚫 empty image")
				return "", errImg
			}
		}

		uploaded, err := os.Create("uploads/" + imageName)
		if err != nil {
			fmt.Println("⚠ wrong image path")
			return "", err
		}
		defer uploaded.Close()

		//*Copying the uploaded file's content in the local one
		if _, err := io.Copy(uploaded, file); err != nil {
			fmt.Println("⚠ couldn't copy image in local")
			return "", err
		}
		return imageName, nil
	}

	return "", nil
}
