# Notes on Go (Golang)

## Stuff That Annoys Me About Go

* Goroutines leak if not dealt with properly
* Stuttering and complexity of `context.Context`
* Utter lack of beginner focus and documentation

## Handing Off Process (exec) to Subprocess Executable

When creating command line utilities it is helpful to know some of the
conventional equivalents to the same tasks used in shell scripting. An
important one is to execute another program and, if possible, replace
the current process with a new one running another program but
inheriting the arguments, jobs, and environment of the parent.

The following probably does what you want without the operating system
specific `syscall` dependency:

```go
s := exec.Command("less", "-r", "-Ps"+status)
//less.Stdin = strings.NewReader(buf)
less.Stdin = os.Stdin
less.Stdout = os.Stdout
less.Run()
```

But if you want exactly the same thing as the shell `exec` command
(`execve(3)`) you need something like one of these:

To hand off to `vim`:

```go
myVimArgs := []string{"README.md"}
vim, err := exec.LookPath("vim")
if err != nil {
	return fmt.Errorf("Vim not found")
}
args := []string{vim, myVimArgs...}
syscall.Exec(vim, args, os.Environ())
```

To change directories as if done from the shell directly. (Note that
this version assumes (and trusts) a UNIX/Linux environment variable):

```go
shell = os.Getenv("SHELL")
if shell == "" {
	return fmt.Errorf("SHELL env variable not set.")
}
os.Chdir("/tmp")
return syscall.Exec(shell, []string{shell}, os.Environ())
```

## Testing for Panic

```go
func assertPanic(t *testing.T, f func()) {
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("The code did not panic")
        }
    }()
    f()
}
```

## Golang Template Joins

```go
{{ range $index, $element := .Equipment}}
	{{if $index}},{{end}}
	{{$element.Name}}
{{end}}
```

## Fastest Way to Get Keys from Map 

```go
// Keys returns a sort list of keys from the Data
func (js Config) Keys() []string {
	keys := make([]string, len(js.Data))
	n := 0
	for k, _ := range js.Data {
		keys[n] = k
		n++
	}
	return keys
}
```
