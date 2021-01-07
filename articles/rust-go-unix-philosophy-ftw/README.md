---
Title: Rust, Go, and the Unix Philosphy FTW!
Subtitle: More Deving, Less Dissing
Query: true

Summary:
  Rust and Go are *not* competitors and never have been. Go is made for
  everyone. Rust is made for people who like a bit more speed and
  low-level development possibilities at the cost of developer
  productivity and unnecessary complexity. Rust is denser than Perl.
  They have *very* different goals and *massively* different
  communities. But all of that is fine because you are following the
  Unix philosophy of development, right? Use both, or not, It's your choice.

Created: 
---

## Yeah, Rust Actually Sucks More Than I Realized

I'll admit after spending the last two days messing around with Rust and
reading about the community and core team (one of whom was the directory
of NodeJS) that I'm no longer impressed with Rust, at all.

I'm kinda glad that the weird Rust excitement I felt --- prompted by
<https://pest.rs> --- has waned and brought me back to a more practical
thinking. You know, the kind of thinking that actually *gets shit done*
instead of wasting time on overly complex cool stuff. 

No wonder businesses love Python and Go. They want "safe enough"
*development speed* even if it costs a bit of run-time performance. The
interpreted language mania has completely confirmed that.

## Safe? Mmmm, Okay.

Sure Rust has built in safety, sort of, and yes it has *very* secure
concurrency *because it doesn't really have any*. The `async` and
`await` stuff was recently added and is yet another example of stupid
promise mania. 

Promises are butt-ugly horrible. Why can't people see that? It is as if
people *want* unnecessary complexity. It makes Go's goroutines and pipes
look genius in comparison.

I blogged earlier about how I felt that Go might not be as safe for
beginners because of the ease with which a beginner can write highly
performant, truly parallel code in Go. The logic was that such beginners
would be more likely to write race conditions with code not safe for
concurrency. But the fact is the amount of learning and unnecessary
complexity to do *any* concurrency in Rust is so high that *other*
errors are very likely to creep in as well as increase developer time by
a factor of 10.

In other words, learning how to do `go build -race` instead of just `go
build` is *way* more worth it in the long run.

## Go's Unsexiness Makes It Sexy

People who appreciate how plain Go is are those who have decades of
development experience on large applications. That is why a unit test
framework is *included*. It is why code documentation is *included*. It
is why things are forced to be *explicit*. Even being left-justified
makes is so much quicker to scan.

Go design and ongoing development represents a core commitment to the
software development practices *that actually matter*: sustainability,
maintainability, portability, testability, and documentation. Rust fails
on every point in comparison. 

## What I Think Doesn't Matter

But it doesn't matter what I think. There are a lot of people picking up
on Rust and the [Gartner Hype Cycle](https://duck.com/lite?kae=t&q=Gartner Hype Cycle) is about to kick into high gear.
Rust *will* be all over the place. The *promise* of safe code is enough
to get *anyone* to look past the absolutely disaster that is Rust
syntax. 

## Go Will Win *Most* Projects

If only Go would have *at least* added `<>` for the new generics design.
But that would have made compilation slower and Go's priorities have
always been clear.

Still the fact that Go will have generics even though it took its own
sweet time to add them (which is a methodical good thing) means it will
continue to destroy Rust for mainstream adoption. Rust is enjoying a bit
of interest because it can replace C++, but as more people who care
about actually producing and maintaining software make their choices
about the best modern compiled language to the job it is very clear they
will continue to pick the unsexy option. Those companies swayed by the
sexy promises (no pun intended) of Rust will ultimately pay the price in
organizational productivity.

## Unix Philosophy Saves Us All

Rather than building monolithic applications building one command that
does something really well will continue to be of value for most things.
For example, I may implement the PEG parser in <https://pest.rs> while
doing the rest in Go. There's no reason we can't have the best of both
worlds.
