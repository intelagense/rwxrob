---
Title: Golang JSON/YAML Tags Actually Suck
Subtitle: At Least for Some Things
---

> You know those things that you thought were cute? Yeah, don't included
them. They always end up to bite you later. (Ryan Dahl, 10 Things I
Regret About NodeJS)

As cool as I find Golang automatic JSON and YAML marshalling and
unmarshalling directly into tagged structs I'm kinda starting to thing
using that technique for anything significant is a really bad idea. Why?

## The Problem

**Unmarshalling often quietly leaves you with an empty structure.**
This can be disastrously bad if not caught before depending on that
struct for anything later, like the *ultimate* null pointer error in a
way.

**Multiple tags often needed.** For YAML, for example, there are nifty
packages to first convert the YAML to JSON and then unmarshal that
allowing you to use no tags or either `json` or `yaml` tags but these
solutions are *horribly* inefficient. Reading a config file once in a
while might be okay, but for parsing hundreds if not thousands of files
these methods are simply unsustainable and unacceptable.

**Requires dependency on structs with public attributes.** Tagging
*requires* public struct attributes. If you do not know why that is a
bad thing you haven't been coding Go long enough or on a big enough
project yet. For most uses structs are actually a bad pick. You should
probably be building things that have proper interfaces instead.
Favoring interfaces --- even though it about doubles the code you need to
use them --- simplifies the code that uses them and allows your things
to be used in many different contexts --- especially if you have
followed [SOLID methodology](https://duck.com/lite?kae=t&q=SOLID methodology) and kept most of your interfaces with one
or two methods.

## One Solution

The solution for me is to create a proper interface for all my things
and then to implement my own `Unmarshal()` method that uses a ubiquitous
`map[string]interface{}`. Then I can decide if I want to cherry pick
values out of it to populate an internal, private struct that fulfills
the interface or simple pull them directly from it and cast the
`interface{}` to the property type. This indirection is
not as heavy as it might seem. Sure it requires a bit more code, but the
result provides a solution to all the above problems.

* I can validate in one place that all the required data was present in
  the YAML/JSON.

* I don't need any cumbersome tags *at all.*

* I can build dependencies on my public interface instead of exposing
  the internals of a struct implementing it that I would rather not.

