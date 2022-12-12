package composite

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type School struct {
	classes []*Class
}

func readStudentNames() []string {
	f, err := os.Open("Composite/data/names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	students := make([]string, 1)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		students = append(students, scanner.Text())
	}

	return students
}

func (s *School) CreateClasses() {
	students := readStudentNames()
	cCount := CLASS_COUNT * len(GetClassNames())
	currentStudentIdx := 0

	for i := 0; i < CLASS_COUNT; i++ {
		for _, cName := range GetClassNames() {
			c := &Class{}
			c.SetClassName(strconv.Itoa(i) + " " + cName)
			maxIdx := currentStudentIdx + len(students)/cCount
			for j := currentStudentIdx; j < maxIdx; j++ {
				student := &Student{}
				student.SetFullName(students[j])
				c.AddStudent(student)
				currentStudentIdx++
				s.classes = append(s.classes, c)
			}
		}
	}
}

func (s *School) SearchName(phrase string) (bool, string) {
	for _, c := range s.classes {
		hasMatch, name := c.SearchName(phrase)
		if hasMatch {
			return true, name
		}
	}
	return false, ""
}

func (s *School) ClassesLen() int {
	return len(s.classes)
}
