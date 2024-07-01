package queries

func BuildNeoWritReturnQuery() string {
	return "RETURN w.title AS title, w.description AS description, c.name AS author, u.username AS username"
}
