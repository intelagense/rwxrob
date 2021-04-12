package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	out, err := os.Create("index.html")
	defer out.Close()
	check(err)

	data := map[string]interface{}{}

	buf, err := ioutil.ReadFile("data.yml")
	check(err)
	err = yaml.Unmarshal(buf, &data)
	check(err)

	t, err := template.ParseGlob("tmpl/*")
	check(err)
	err = t.Execute(out, data)
	check(err)
}
