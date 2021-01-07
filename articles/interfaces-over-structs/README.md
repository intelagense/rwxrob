---
Title: Choose Interfaces Over Structs in Golang
Created: 2020-07-04T11:04:50-0400
Query: true
---

These days it's pretty common knowledge that creating anything in Golang
as a struct is a bad idea unless you *really* have a good reason. That's
because interfaces are so much more flexible and surprisingly often
faster than structs. Here are a few examples.

## Easier to Document

It's a small but significant advantage to be able to document each
attribute of thing individually instead of jamming a ton of notes that
make a struct definition rather hard to read. Let's take the example of
a thing we'll call `Entry`, as in an entry to a log or list. As a struct
it might look like this:

```go
type Entry struct {
  Tstamp time.Time
  Amount float64
  Note string
}
```

To document that you have three choices:

1. Add a bunch of trailing comments
1. Add preceding line comments
1. Add some sections to the Entry comment

I'll just let you imagine what that looks like, but believe me, none of
them are really good options. The trailing comments get in the way of
`json` and `yaml` tags. The preceding line comments break the structure
up so seeing the overall structure of the structure is difficult and
changing to exported is problematic. The section in the Entry comments
is removed from the `struct` field so that you forget to remove or
update it when you change something. No one is going to die if you use
one of these methods, and quite frankly you *have* to pick one when
documenting a `struct` at all. The standard library is full of examples.

But let's consider what happens when you use an `interface` instead:

```go

type Tstamp interface {
  Tstamp time.Time
}

type Amount interface {
  Amount float64
}

type Note interface {
  Note string
}

type Entry interface {
  Tstamp
  Amount
  Note
}
```

Sure the code gets *way* longer, but that is not a bad thing in the long
run. It provides an obvious way to document each interface fully. Not
only can each interface be documented by itself, but the inline comments
on the interface methods will never have to contend with tags for room
at the end of a line.

Of course, this also allows other combinations of these fields to be
made into other `interfaces`. This is composition at its finest. 

## Pass by Pointer or Value? Neither It's an Interface

When dealing with interfaces you also don't have to make that painful
decision about accepting a pointer or copy of a thing into your
functions and methods. Anything that fulfills an interface can simply
use the name of that interface.

## Flexibility

The above example with `Entry` works for any of the following function
calls:

```go
func PrintTstamp(t Tstamp){fmt.Println(t)}
func PrintAmount(a Amount){fmt.Println(a)}
func PrintNote(n Note){fmt.Println(n)}

type entry struct {
  t time.Time
  a float64
  n string
}

func (e *entry) Tstamp() time.Time {return e.t}
func (e *entry) Amount() float64 {return e.a}
func (e *entry) Note() string {return e.n}

func main() {
  ent := entry{time.Now(),23.3,"Just a test"}
  PrintTstamp(ent)
  PrintAmount(ent)
  PrintNote(ent)
}
```

The time spent writing out this out using the interfaces can easily be
made back later by allowing your code to accept the same argument for
multiple different functions and methods that would otherwise force you
to only allow a single structure to be passed. When those conflicts
happen you have to do gymnastics to extract out the data needed and
create an entirely new struct just to comply with a poorly designed
function that is not written to accept arguments that are interfaces
instead. In fact, *thou shalt not write functions and methods that
accept structs* unless you *really* want to hate yourself later.  Don't
believe me? Spend an afternoon looking at all the ways this has been
done very well in the Go standard library.

## When Data Has Already Been Parsed or Loaded

Very often there are already low-level libraries for parsing everything
you need up into structs of data. Such is the case, for example, with
`goldmark`. It parses rather complicated CommonMark --- including the
YAML data --- into a full abstract syntax tree. The YAML data --- even
if you have multiple YAML data blocks throughout your CommonMark --- is
all combined into a single, wonderful `map[string]interface{}`.

Now consider the problems and inefficiencies of defining your data model
using `struct` instead of `interface`. Here's a `Person` `struct`:

```go
type Person struct {
  Name string
  Age int
}
```

In order to use that you have to write some sort of code that would
*copy over the data* from the `map[string]interface{}`. Because you are
using primitive data you can't really use pointers and if you did it
would negate the other advantages of having a struct for marshaling and
such. But you don't need built in Go unmarshaling at all and probably
don't want it because of how finicky it can get and silently just fail
to unmarshal properly because some data didn't match. No, we are going
to use `goldmark` to do all that heavy lifting for us.

Besides, we are talking about *reading* data that is maintained in
CommonMark and YAML. That means that any requirement to change any of it
through any coding interface is just a bad idea. It would be better to
leverage a good raw CommonMark editor than try to build a GUI on top of
that. Because we are just reading from parsed data all we really need
are methods that return the data that we want. Here's the same `Person`
as an `interface` instead:

```go
type Person interface {
  Name() string
  Age() int
}
```

It might seem counter-intuitive, but `interface` method calls are actually
about the same speed as returning the raw struct data. This is because
only a pointer to that data is required to return, such as with
many of the `FileInfo` fields:

```go
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)
}
```

Look familiar? There's a reason this is an `interface` and not a `struct`. You
guessed it. It's the same reason you just saw. If an `interface` is good enough
for something as frequently accessed as `FileInfo` you can be sure it is plenty
fast enough for your applications.

## Full Benchmark Test Comparison

Thanks to Go's great built-in benchmark testing we can see all of this
in action rather easily. We'll create three alternatives:

1. Using just a `struct`
1. Using an interface where the values are returned directly
1. Using an interface where the values are indirectly returned 

*Or you can [download it.](./lib.go)*

```go
package lib

// -------------------------------------------------

type Person struct {
	Name string
	Age  int
}

func NewPerson() {
	person := Person{Name: "Rob", Age: 52}
	if person.Name == "Rob" {
	}
	if person.Age == 42 {
	}
}

// -------------------------------------------------

type PersonI interface {
	Name() string
	Age() int
}

type aperson struct {
	// fair to leave out since usually pointing to other data
}

func (p *aperson) Name() string {
	return "Rob"
}

func (p *aperson) Age() int {
	return 42
}

func NewPersonI() {
	person := new(aperson)
	if person.Name() == "Rob" {
	}
	if person.Age() == 42 {
	}
}

// -------------------------------------------------

type PersonII interface {
	Name() string
	Age() int
}

type aaperson struct {
	name string
	age  int
}

func (p *aaperson) Name() string {
	return p.name
}

func (p *aaperson) Age() int {
	return p.age
}

func NewPersonII() {
	person := aaperson{name: "Rob", age: 42}
	if person.Name() == "Rob" {
	}
	if person.Age() == 42 {
	}
}
```

Here's what the benchmark ([`lib_test.go`](./lib_test.go)) code looks
like:

```go
package lib

import "testing"

func BenchmarkNewPerson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPerson()
	}
}

func BenchmarkNewPersonI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonI()
	}
}

func BenchmarkNewPersonII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPersonII()
	}
}
```

And finally the results:

```
goos: linux goarch: amd64
BenchmarkNewPerson-8      853628504   1.41 ns/op
BenchmarkNewPersonI-8     1000000000  0.412 ns/op
BenchmarkNewPersonII-8    873169317   1.39 ns/op
```

As you can see the difference to call via `interface` methods is 0.02
ns/op. That's pretty much the dictionary definition of "negligible". If
that kind of time difference is a problem then you are probably coding
in the wrong language to begin with (try C or Assembly instead).

What is even more interesting is the ridiculously faster times when the
value simply has to be returned and not looked up. Such is often the
case when the thing is *already* there and just needs a pointer
returned. This is why `FileInfo` using `interface` methods instead of
having to copy over that same data into a `struct` and then return that,
or provide a `struct` that contains pointers to rather primitive data.
This means that, depending on how much data is just being pointed at,
you could end up with even faster times matching or potentially beating
that of `structs`.

## Moral of the Story

The difference in execution speed between a thing designed as a `struct`
and another as an `interface` with accessor methods is negligible. While
an `interface` requires a bit more code to write it buys you *a lot* of
flexibility later. We're talking the kind of flexibility that made
`goldmark` dominate so completely over the `blackfriday` Markdown
engine. But that's not the only example, do your own research. You'll
find so many more. Now its time to start your own. I know I am.


