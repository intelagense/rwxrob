## Sunday, April 25, 2021, 9:39:53AM EDT <1619357993>

Here's a trick for testing anything in Go  with `os.Exit()` (or
something like it):

```go
// go coverage detection is fucked for this sort of stuff, oh well, we
// did the test even if coverage falsely reports 50%
func TestExec(t *testing.T) {
  if os.Getenv("TESTING_EXEC") == "1" {
    err := Exec("go", "version")
    if err != nil {
      t.Fatal(err)
    }
    return
  }
  cmd := exec.Command(os.Args[0], "-test.run=TestExec")
  cmd.Env = append(os.Environ(), "TESTING_EXEC=1")
  err := cmd.Run()
  if err != nil {
    t.Fatalf("process exited with %v", err)
  }
}
```

It's an amazing little trick since it actually calls its own temporary test binary (`os.Args[0]`) and has it just run that specific test. The environment variable is to prevent an infinite (but interesting) function recursion situation.

One thing that took me a while to notice is that there is *no* output whatsoever. All you get is the return value of the command that was executed. To get any kind of data back from the test of an executable like this you would need to create a temporary file and write to it. But still, this is pretty damn awesome for testing *any* executable.

I'm considering this for many of the highest level commands that I create that are difficult to test without actually calling them. Hell, you could even test things like tab-completion shell code generation that's designed to be evaluated by an `bashrc` file or something.
