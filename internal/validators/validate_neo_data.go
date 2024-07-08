package validators

// WRITING
// to be used for all fiction formats, including collections, series, and fictional universes
type WritingProperties struct {
	Title        string
	Description  string
	Uid          string
	Rank         int64
	RelRank      int64
	Likes        int64
	Libraries    int64
	ReadingLists int64
	Donations    int64
	Reads        int64
}

func ValidateWritingProperties() {

}

type RelWritHasTag struct {
}

type RelInCollection struct {
}

type RelInSeries struct {
}

type RelInUniverse struct {
}

// Users
type UserProperties struct {
	Username string // unique user identifier
	DisplayName     string
	Uid      string // ? is uuid a string or int?
	Password string // to be hashed
}

func ValidateUserProperties() {

}

type BookshelfProperties struct {
	Name string
	Uid  string
}

// User -> Creator
type RelFollows struct {
	Uid string
}

// User -> Creator
type RelSubscribed struct {
	Uid string
}

// User -> Writing
type RelRead struct {
	Uid string
}

// User -> Writing
type RelLikes struct {
	Uid string
}

// User -> Writing
type RelWritInLibrary struct {
	Uid string
}

// User -> Writing
type RelWritOnList struct {
	Uid string
}

// Creators
type CreatorProperties struct {
	Name string
	Uid  string
}

func ValidateCreatorProperties() {

}

func ValidateDonationProperties() {

}

func ValidateFlagProperties() {

}

func ValidateGenreProperties() {

}

func ValidateTagProperties() {

}
