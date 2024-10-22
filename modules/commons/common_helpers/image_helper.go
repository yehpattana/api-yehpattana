package commonhelpers

import (
	"fmt"
	"mime/multipart"
	"time"

	commonthirdparty "github.com/natersland/b2b-e-commerce-api/modules/commons/common_thirdparty"
)

func UploadImageOrUseDefaultImage(imageFile *multipart.FileHeader, defaultImageUrl string, folderName string) (string, error) {
	if imageFile == nil {
		return defaultImageUrl, nil
	} else {

		currentTime := time.Now()
		formattedTime := currentTime.Format("20060102_150405")
		filename := fmt.Sprintf("%s/image_%s.jpg", folderName, formattedTime)

		uploadImageResult, err := commonthirdparty.CloudinaryImageUploader().UploadImage(imageFile, filename)

		println("uploadImageResult: ", uploadImageResult)

		if err != nil {
			return "", err
		}

		return uploadImageResult.SecureURL, nil
	}

}
func UploadAttachment(file *multipart.FileHeader, folderName string) (string, error) {

	currentTime := time.Now()
	formattedTime := currentTime.Format("20060102_150405")
	filename := fmt.Sprintf("%s/attachment_%s", folderName, formattedTime)

	uploadResult, err := commonthirdparty.CloudinaryImageUploader().UploadImage(file, filename)

	println("uploadResult: ", uploadResult)

	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil

}
