package models

// chapters exist only in Mongo, not in Neo4j

type Chapter struct {
	NeoId     string         `bson:"neo_id,omitempty"`
	Content   map[string]any `bson:"content"`
	CreatedAt int64          `bson:"created_at,omitempty"`
	UpdatedAt int64 `bson:"updated_at,omitempty`
}
