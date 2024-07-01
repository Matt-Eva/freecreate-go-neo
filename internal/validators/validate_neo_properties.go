package validators

func ValidateWritingProperties(){

}

// to be used for novels, novellas, and novelettes
func ValidateNovelProperties(){

}


type UserProperties struct {
	Username string // unique user identifier
	Name string
	Uid string // ? is uuid a string or int?
	Password string // to be hashed
}
func ValidateUserProperties() UserProperties{
	// validationMap := map[string]bool{
	// 	"username": false,
	// 	"name": false,
	// }

	return UserProperties{}
}

func ValidateCreatorProperties(){

}

func ValidateDonationProperties(){

}

func ValidateFlagProperties(){

}