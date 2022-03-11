package simplemoduleproject

import (
	"fmt"
	"strings"
)

type Student struct {
	studentName string
	grade       int
}

type Status func(string) bool

func (s Student) Greeting() {
	fmt.Printf("Hello %s from grade %d \n", s.studentName, s.grade)
}

func (s Student) GetNameAt(i int) string {
	return strings.Split(s.studentName, " ")[i-1]
}

func User(userName string, status Status) {
	if status(userName) {
		fmt.Printf("%s adalah siswa berstatus aktif \n", userName)
	} else {
		fmt.Printf("%s adalah siswa berstatus tidak aktif \n", userName)
	}
}
