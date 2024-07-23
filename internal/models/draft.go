package models

// drafts exist only in Mongo, not in Neo4j

// whenever a piece of content is edited, the draft itself is what is first updated
type Draft struct {
	NeoId       string         `bson:"neo_id"`     // optional
	ContentId   string         `bson:"content_id"` // optional
	Content     map[string]any `bson:"content"`    // optional
	Tags        []string       `bson:"tags"`       // optional
	Genres      []string       `bson:"genres"`     // optional
	Description string         `bson:"description"`
	WritingType string         `bson:"writing_type"`
	CreatedAt   int64          `bson:"created_at,omitempty"`
	UpdatedAt   int64          `bson:"updated_at,omitempty`
}
