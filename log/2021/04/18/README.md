## Sunday, April 18, 2021, 12:30:20PM EDT <1618763420>

I was fretting over naming conflicts and realized that they are
completely addressed by "registry" (side effects) package imports
because all the subcommands included for that CmdBox module are in that
modules package namespace. Each module gets its own package namespace
that is protected from all others. This resolves a problem I had with
the other method I was struggling with that had to come up with the same
sort of recursive tree within a single commands registry map. In other
words, `_` fixed *everything* without my really planning it that way. It
was just the right thing to do and turns out that works. The key is to
remember that every CmdTab command can only have one set of subcommands.

Renaming when there is a conflict can easily be resolved in the
`Add("mycmd","orig->another")` since the Add happens *after* the
`init()` of the  packages upon which that module depends.

