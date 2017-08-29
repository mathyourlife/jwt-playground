package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

type Settings struct {
	DB struct {
		Name  string `yaml:"name"`
		Admin struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
		} `yaml:"admin"`
		AppUser struct {
			Username string `yaml:"username"`
		} `yaml:"appuser"`
	} `yaml:"db"`
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)

	b, err := ioutil.ReadFile("/home/dcouture/gopath/src/github.com/mathyourlife/jwt-playground/sql/settings.yml")
	if err != nil {
		log.Fatal(err)
	}
	s := Settings{}
	err = yaml.Unmarshal(b, &s)
	tmpl, err := template.ParseFiles("create-db.sql")
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(os.Stdout, "create-db.sql", s)
	if err != nil {
		panic(err)
	}

}
