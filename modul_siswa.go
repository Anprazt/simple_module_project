package simplemoduleproject

import (
	"fmt"
	"strings"
)

type Student struct {
	StudentName string
	Grade       int
}

type Status func(string) bool

func About() {
	fmt.Println("Ini adalah data-data siswa")
}

func (s Student) Greeting() {
	fmt.Printf("Hello %s from grade %d \n", s.StudentName, s.Grade)
}

func (s Student) GetNameAt(i int) string {
	return strings.Split(s.StudentName, " ")[i-1]
}

func User(userName string, status Status) {
	if status(userName) {
		fmt.Printf("%s adalah siswa berstatus aktif \n", userName)
	} else {
		fmt.Printf("%s adalah siswa berstatus tidak aktif \n", userName)
	}
}

func AnonymousOrtu() {
	OrTu := struct {
		OrtuName string
		Alamat   string
		Usia     int
		Notelp   string
	}{
		OrtuName: "Arifin",
		Alamat:   "Cikoneng",
		Usia:     50,
		Notelp:   "081284846363",
	}
	fmt.Println(OrTu)
}
