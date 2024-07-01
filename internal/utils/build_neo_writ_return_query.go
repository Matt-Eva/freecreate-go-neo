package utils


func BuildNeoWritReturnQuery() string {
	return "RETURN w.name AS title, w.description AS description, c.name AS author, u.username AS username"
}