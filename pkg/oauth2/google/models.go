package google

type OauthToken struct {
	AccessToken string
	IDToken     string
}

type UserResult struct {
	Id              string
	Email           string
	IsVerifiedEmail bool
	Name            string
	GivenName       string
	FamilyName      string
	Picture         string
	Locale          string
}
