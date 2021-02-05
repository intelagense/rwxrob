## Friday, February 5, 2021, 10:23:30AM EST <1612538610>

It seems like putting a natural language layer that processes a raw
command-line into a generic, traditional function call with parameters
is the best approach and meets the needs to be able to create simple
actions without the extra thought as to how those actions will be
determined and called by the smart layer.

It also occurred to me that this separation applies to the work to
create an application. You have the person or team developing the
actions independently and without thought for how they will be
ultimately invoked from the command line. And you also have a person or
team dedicated to creating the language grammar that translates every
possible command line permutation into those actions. This has immediate
implications for creating fully functional interfaces in *any* foreign
language.

## Friday, February 5, 2021, 9:40:49AM EST <1612536049>

I am curious to know how long something remains in the Go caching
servers before (and if) it ages out.

TIL that `go mod vendor` and `go build -mod=vendor` will pull any
imported dependencies into the vendor folder so that they are
permanently a part of the source code rather than pulled even with the
saved checksums associated with stuff in `go.mod`.

## Friday, February 5, 2021, 8:10:57AM EST <1612530657>

Realized I cannot use case to distinguish context if I want an interface
that ultimately can be compatible with conversational UIs.

