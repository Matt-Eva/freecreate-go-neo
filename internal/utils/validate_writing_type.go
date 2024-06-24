package utils

func ValidateWritingType(writingType string) string{
	types := GetWritingTypes()
	for _, t := range types {
		if writingType == t {
			return t
		}
	}

	return ""
}