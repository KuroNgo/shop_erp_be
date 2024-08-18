package cloudinary

type UploadImage struct {
	ImageURL string `bson:"image_url" json:"image_url"`
	AssetID  string `bson:"asset_id" json:"asset_id"`
}

type Update struct {
	URL       string `bson:"url" json:"url"`
	SecureURL string `bson:"secure_url" json:"secure_url"`
	CreateAt  string `bson:"create_at" json:"create_at"`
}

type RemainingDataAccount struct {
	RemainingData int16 `bson:"remaining_data" json:"remaining_data"`
	RemainingFile int16 `bson:"remaining_file" json:"remaining_file"`
}

type UploadAudio struct {
	AudioURL string `bson:"image_url" json:"image_url"`
	AssetID  string `bson:"asset_id" json:"asset_id"`
}
