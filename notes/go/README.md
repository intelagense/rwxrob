# Notes on Go (Golang)

## Stuff That Annoys Me About Go

* Goroutines leak if not dealt with properly
* Stuttering and complexity of `context.Context`
* Utter lack of beginner focus and documentation

## Handing Off Process (exec) to Subprocess Executable

```go
s := exec.Command("less", "-r", "-Ps"+status)
//less.Stdin = strings.NewReader(buf)
less.Stdin = os.Stdin
less.Stdout = os.Stdout
less.Run()
```

## Golang Template Joins

```go
{{ range $index, $element := .Equipment}}
	{{if $index}},{{end}}
	{{$element.Name}}
{{end}}
```

