---
Title: Evils of Long Comments
Subtitle: Unreadable, Hard to Update, Inconsistent
Query: true

Summary:
  After trying an alternative long comment format for a day I am
  convinced there is actually a right way, and an evil way to do long
  comments in large code bases.

---

First of all, I'm not saying that long comments are evil. I'm saying
there is an *evil way to do them* and a good way. 

This doesn't have to be that long of an explanation. Most of it is
common sense that is becomes apparent after staring at code with syntax
highlighting off for more than an hour.

## I Need the LHS Border

I don't think I realized how important having something consistently on
the left hand side of the screen and code is to allow better scanning of
code. I'm sure this is why the C people have been putting that `*` over
there for ages. It also supports the Go `//` approach because the
distinction is so pronounced in comparison even to the C approach.

## I Need Wrapping

All good editors will wrap pretty much any comment section properly. But
long comments with ` /* ... */` are a bit tougher to find support for
mostly because they involve more than just simple folding. Blocks of
single-line comments resolve better and can even be created with a
simple vi/bash filter script:

```bash

#!/bin/bash

while read -r line; do 
  echo "${1-##} $line"
done

```

## I Need Whitespace Before First Line of Code

I thought I really liked that the ending token `*/` of the long comment
on its own line would give me a blank line before the code it covers,
but then I realized when the comment isn't long enough enough for two
lines things get really complicated and ugly really fast.

For a while I had single line comments if only one line and long
comments for the rest. But having to context switch between the two is
just fucking insane. I don't know what I was thinking. With a consistent
comment approach I don't need to spend any cognitive overhead on it.

I did start putting a extra blank before the code if there is more than
one line of comments for better visual separation as I have seen other
Go projects do.

## I Need to Be Able to Comment Out Code  

By far the most annoying realization I had when using the ` \* ... *\ `
form of long comments was that I no longer could comment out any code
that had such long comments in them. The conflict between the comment
tokens forced it not to work. This is never an issue if you use
single-line comments --- as is common practice in the Go community.

```go{.wrong}

/*
/* I'm a comment */

func main() {
  fmt.Println("Hello World!"
}
*/

```

This one works:

```go

/*
// I'm a comment

func main() {
  fmt.Println("Hello World!"
}
*/

```

## Still Makes Sense Sometimes

There are cases when the stuff on the left-hand side is just not needed.
For example, when writing a rather long `doc.go` file it is pretty
standard practice to use the C long-comment form. This is all over in
the Go standard library, but we are talking a *lot* of comments and
*zero* code (well, there's one line).
