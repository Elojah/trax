package pbtypes

func GetString(s *String) *string {
	if s == nil {
		return nil
	}

	return &s.Value
}

func GetStringEmpty(s *String) string {
	if s == nil {
		return ""
	}

	return s.Value
}
