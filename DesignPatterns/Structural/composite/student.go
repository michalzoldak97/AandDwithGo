package composite

import "strings"

type Student struct {
	fullName string
}

func (s *Student) FullName() string {
	return s.fullName
}

func (s *Student) SetFullName(fullName string) {
	s.fullName = fullName
}

func (s *Student) SearchName(phrase string) (bool, string) {
	containsPhrase := strings.Contains(strings.ToLower(s.fullName), strings.ToLower(phrase))
	if containsPhrase {
		return containsPhrase, s.fullName
	}

	return containsPhrase, ""
}
