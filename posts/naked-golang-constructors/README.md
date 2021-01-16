---
Title: Naked Golang Constructors
Subtitle: Less is More, Explicit is Better Than Implicit, Remember?
Query: true

Summary:
  I've never been a fan of complicated function and method signatures
  --- especially like what are common with Python and Java. The great
  thing about Go is that you can just leave them naked and empty and set
  the struct values *directly and explicitly* immediately after. 

---

In fact, this is one of the dumbest things Python decided to go with,
the most non-specific, wildly open and large constructor signature
possible. It's like th `*arsgs` and `**kwargs` were added to
deliberately tempt developers into making the most obtuse and unclear
constructors possible. Quite ironic give the "explicit" is better mantra
from the [Zen of Python](https://duck.com/lite?kae=t&q=Zen of Python).

Consider the following, which is just a *really bad example* to make a
point:

```go

func (b *Base) NewBase(path string, conf string) *Base {
  b := &Base{
    ConfDir: conf,
    Path: path,
  }
  return b
}

var b *Base = NewBase("./","./.kn")
var a *Base = NewBase("./","")


```

This is simply not *idiomatic* Go programming style.

This could just as easily and more clearly have been done *without* the
monstrous method parameter signature. 

```go

func (b *Base) NewBase() *Base {
  b := new(Base)
  // some derived initialization here
  return b
}

var b *Base = NewBase()
b.Path = "./"
b.ConfDir = filepath.Join(b.Path,".kn")

// ...

``` 

Of course this doesn't work if you intend to do something magic with an
incoming argument during your initialization. But do you really need
that magic? Does that implicit magic belong in your constructor at all?
Are you sure you don't want just simpler [command
structs](/cmdop/#command-structs) in general that do less?

