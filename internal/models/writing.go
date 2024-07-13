package models

type Writing struct {
	Uid          string
	Title        string
	Descpription string
	Thumbnail    string
	WritingType  string
	CreatedAt    int64
	UpdatedAt    int64
	Libraries    int64
	Likes        int64
	Views        int64
	Donations    int64
	Rank         int64
	RelRank      int64
}

type ShortStory struct {
	Writing
}

type Novel struct {
	Writing
}
