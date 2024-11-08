package images_cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"shop_erp_mono/pkg/interface/cloud/cloudinary"
	"shop_erp_mono/pkg/interface/cloud/cloudinary/models"
)

func UpdateToCloud(publicID string, filename string) (interface{}, error) {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()

	renameParams := uploader.RenameParams{
		FromPublicID: publicID,
		ToPublicID:   filename,
	}

	// Access the filename using a desired filename
	result, err := cld.Upload.Rename(ctx, renameParams)
	if err != nil {
		return "", err
	}

	resultRes := models_cloudinary.models_cloudinary{
		URL:       result.URL,
		SecureURL: result.SecureURL,
		CreateAt:  result.CreatedAt.String(),
	}

	return resultRes, err
}
