package composite

type Class struct {
	className string
	students  []Composite
}

func (c *Class) ClassName() string {
	return c.className
}

func (c *Class) SetClassName(className string) {
	c.className = className
}

func (c *Class) AddStudent(s Composite) {
	c.students = append(c.students, s)
}

func (c *Class) SearchName(phrase string) (bool, string) {
	for _, m := range c.students {
		containsPhrase, name := m.SearchName(phrase)
		if containsPhrase {
			return containsPhrase, name
		}
	}
	return false, ""
}
