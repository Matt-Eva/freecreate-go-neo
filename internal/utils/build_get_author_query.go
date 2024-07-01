package utils

func BuildGetAuthorQuery() string {
	return "WITH w MATCH (w) <- [:CREATED] - (c:Creator) <- [:IS_CREATOR] - (u:User)"
}