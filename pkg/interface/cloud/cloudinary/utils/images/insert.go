package images_cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"mime/multipart"
	"shop_erp_mono/pkg/interface/cloud/cloudinary"
	"shop_erp_mono/pkg/interface/cloud/cloudinary/models"
)

func UploadImageToCloudinary(file multipart.File, filePath string, folder string) (models_cloudinary.UploadImage, error) {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()
	if err != nil {
		return models_cloudinary.UploadImage{}, err
	}

	uploadParams := uploader.UploadParams{
		PublicID: filePath,
		Folder:   folder,
	}

	result, err := cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return models_cloudinary.UploadImage{}, err
	}

	resultRes := models_cloudinary.UploadImage{
		ImageURL: result.SecureURL,
		AssetID:  result.AssetID,
	}
	return resultRes, nil
}
