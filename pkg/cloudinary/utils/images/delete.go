package images_cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"shop_erp_mono/pkg/cloudinary"
)

func DeleteToCloudinary(assetID string) (string, error) {
	ctx := context.Background()
	cld, err := cloudinary.SetupCloudinary()
	if err != nil {
		return "", err
	}

	deleteParams := uploader.DestroyParams{
		PublicID: assetID,
	}

	result, err := cld.Upload.Destroy(ctx, deleteParams)
	if err != nil {
		return "", err
	}

	return result.Result, nil
}
