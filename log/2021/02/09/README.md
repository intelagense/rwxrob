## Tuesday, February 9, 2021, 3:13:45AM EST <1612858425>

Realized that dealing with Golang templates is so genericised now that
you don't really even need something like Hugo. In fact, the following
could be easily worked into a `kn` action that follows the KEG standard
for Knowledge Nodes:

```golang
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"html/template"

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
```


