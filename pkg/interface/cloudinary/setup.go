package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"shop_erp_mono/internal/config"
)

func SetupCloudinary(env *config.Database) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromParams(env.CloudinaryCloudName, env.CloudinaryAPIKey, env.CloudinaryAPISecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
