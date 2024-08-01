package models

// chapters exist only in Mongo, not in Neo4j
// shard key is creatorid
// query for chapter is creatorid/neoid/uid

type Chapter struct {
	Uid       string         `bson:"uid,omitempty"`
	CreatorId string         `bson:"creator_id,omitempty`
	NeoId     string         `bson:"neo_id,omitempty"`
	Title     string         `bson:"title,omitempty"`
	Published bool           `bson:"published,omitempty"`
	CreatedAt int64          `bson:"created_at,omitempty"`
	UpdatedAt int64          `bson:"updated_at,omitempty`
	Content   map[string]any `bson:"content"`
}
