package cloudinary

import "github.com/cloudinary/cloudinary-go/v2"

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := "vbGeG7reX1mtT7isZPRU7ZWEHxM"
	cldName := "df4zm1xjy"
	cldKey := "138627467233122"

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
