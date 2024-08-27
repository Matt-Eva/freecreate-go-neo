package users


type ReturnUser struct {
	Uid        string `json:"uid"`
	UniqueName string `json:"uniqueName"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	BirthDay   int    `json:"birthDay"`
	BirthYear  int    `json:"birthYear"`
	BirthMonth int    `json:"birthMonth"`
	ProfilePic string `json:"profilePic"`
}