package simplemoduleproject

import (
	"fmt"
	"strings"
)

type Student struct {
	name  string
	grade int
}

type Status func(string) bool

func (s Student) Greeting() {
	fmt.Printf("Hello %s from grade %d \n", s.name, s.grade)
}

func (s Student) GetNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func User(name string, status Status) {
	if status(name) {
		fmt.Printf("%s adalah siswa berstatus aktif \n", name)
	} else {
		fmt.Printf("%s adalah siswa berstatus tidak aktif \n", name)
	}
}
