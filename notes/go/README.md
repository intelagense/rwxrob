# Notes on Go (Golang)

## Handing Off Process (exec) to Subprocess Executable

```go
s := exec.Command("less", "-r", "-Ps"+status)
less.Stdin = strings.NewReader(buf)
less.Stdout = os.Stdout
less.Run()
```

