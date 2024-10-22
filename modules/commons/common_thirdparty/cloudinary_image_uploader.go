package commonthirdparty

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryImageUploaderInterface interface {
	UploadImage(imagePath *multipart.FileHeader, publicId string) (*uploader.UploadResult, error)
}

type cloudinaryImageUploaderImpl struct {
}

func CloudinaryImageUploader() CloudinaryImageUploaderInterface {
	return &cloudinaryImageUploaderImpl{}
}

func (c *cloudinaryImageUploaderImpl) UploadImage(imagePath *multipart.FileHeader, publicId string) (*uploader.UploadResult, error) {

	// Add your Cloudinary product environment credentials.
	// cloudName := c.cfg.CloudName()
	// apiKey := c.cfg.ApiKey()
	// apiSecret := c.cfg.ApiSecret()

	// println("cloudName: ", cloudName)
	// println("apiKey: ", apiKey)
	// println("apiSecret: ", apiSecret)

	// TODO change this to use the config file
	cloudName := "dgfgbympj"
	apiKey := "941573432367425"
	apiSecret := "tTmG8SZ23_d9XipES33-kN8mh78"

	cld, _ := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)

	// Upload the my_picture.jpg image and set the PublicID to "my_image".

	var ctx = context.Background()
	resp, err := cld.Upload.Upload(ctx, imagePath, uploader.UploadParams{PublicID: publicId})

	if err != nil {
		return nil, err
	}

	// Get details about the image with PublicID "my_image" and log the secure URL.
	getImageDetails(publicId, cld, ctx)

	return resp, nil

}

func getImageDetails(publicId string, cld *cloudinary.Cloudinary, ctx context.Context) {
	// Get details about the image with PublicID "my_image" and log the secure URL.
	uploadResp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicId})
	if err != nil {
		fmt.Println("error")
	}
	log.Println(uploadResp.SecureURL)
}
