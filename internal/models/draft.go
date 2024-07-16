package models

// drafts exist only in Mongo, not in Neo4j

type Draft struct {
	NeoId     string         `bson:"neo_id,omitempty"`
	Content   map[string]any `bson:"content"`
	CreatedAt int64          `bson:"created_at,omitempty"`
}
