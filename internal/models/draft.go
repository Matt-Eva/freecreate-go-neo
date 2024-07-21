package models

// drafts exist only in Mongo, not in Neo4j

type Draft struct {
	NeoId       string         `bson:"neo_id"`
	ContentId   string         `bson:"content_id"`
	Content     map[string]any `bson:"content"`
	Tags        []string       `bson:"tags"`
	Genres      []string       `bson:"genres"`
	Description string         `bson:"description"`
	WritingType string         `bson:"writing_type"`
	CreatedAt   int64          `bson:"created_at,omitempty"`
	UpdatedAt   int64          `bson:"updated_at,omitempty`
}
