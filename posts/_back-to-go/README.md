---
Title: Moving Back to Golang for Most Stuff
Subtitle: For README.World Knowledge Management Tool and More
Query: true
Created: 2020-06-21T16:01:29-0400

Summary: 
  Writing everything that has to be compiled in Golang again now that I
  have gotten over my latest Rust fixation. I'm happy I survived. Future
  Rob would have been pissed at taking on that much unnecessary
  complexity and committment to a politically-motivated community that I
  really just cannot force myself to like. Turns out that Go *is* the
  best job for the `kn` tool. Time to get busy and port it from Bash
  even though I am feeling the sunk-cost fallacy of all that work I did
  to make it so pretty and compliant with Google's Shell Scripting
  Guide. Ironically, by dropping shell I'm complying with the part that
  says, "use another language if you script gets over 500 lines long".

---

Recent experimentation with Rust has once again reminded me that it is
*not* the right tool for most jobs that *I need* a compiled language for
(nor is it the most likely pick for most other jobs out there). Rust is
*highly* specialized and overly and unnecessarily complicated. But it is
still the right tool for *some* jobs. With that in mind I've decided to
move *back* to porting the `kn` tool to Go instead of Rust.

## YAML Knowledge Node Types with Go

It works out nicely that the new focus on YAML knowledge node types goes
better with Go than Rust or anything else. I can create a Go struct for
each type and just unmarshal it --- especially since I kept the much
more readable initial cap names for everything. In fact, that should
*really* speed up that part of the development that easy equivalent in
Bash.

