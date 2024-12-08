package token

import (
	"github.com/golang-jwt/jwe"
	"shop_erp_mono/internal/config"
)

func CreateJWE(payload interface{}, env *config.Database) (string, error) {
	context := []byte(env.Token)
	token, err := jwe.NewJWE(jwe.KeyAlgorithmRSAOAEP, payload, jwe.EncryptionTypeA256GCM, context)
	if err != nil {
		return "", err
	}

	compact, err := token.CompactSerialize()
	if err != nil {
		return "", err
	}

	return compact, nil
}

func DecryptJWE(compact string, env *config.Database) (string, error) {
	context := []byte(env.Token)
	token, err := jwe.ParseEncrypted(compact)
	if err != nil {
		return "", err
	}

	payload, err := token.Decrypt(context)
	if err != nil {
		return "", err
	}

	return string(payload), nil
}
