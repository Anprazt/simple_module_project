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

type TeacherProfile struct {
	ID      int
	Name    string
	Age     int
	Address string
	Telp    string
}

func About(asalSekolah string) {
	fmt.Println("Ini adalah data-data siswa dari sekolah ", asalSekolah)
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
func ProfileTeacher() {
	var profile TeacherProfile
	profile.ID = 1
	profile.Name = "Prazt Andi Wicaksana"
	profile.Age = 25
	profile.Address = "Cikoneng"
	profile.Telp = "081284846363"

	fmt.Println(profile)
	fmt.Printf("Id Teacher \t: %d \n", profile.ID)
	fmt.Printf("Name \t\t: %s \n", profile.Name)
	fmt.Printf("Age \t\t: %d tahun \n", profile.Age)
	fmt.Printf("Address \t\t: %s \n", profile.Address)
	fmt.Printf("Telephone Number: %s \n", profile.Telp)
}
