---
Title: Goldmark vs Pandoc Markdown
Subtitle: Speed Over Powerful Syntax
Query: true

---

The more I get into the guts of developing the `kn` Component-Oriented
Knowledge Management tool the more I realize that many of Pandoc's
extras are actually not needed at all --- the biggest of which is Pandoc
tables. Spans and divs are so incredibly important when creating Pandoc
documents, but spans with attributes are being discussed in the
CommonMark forum and will likely be adopted while most Pandoc divs are
better managed as a separate knowledge component that is compiled into a
dynamic type so that the content of the div can also stand on its own
and be included elsewhere.

## Callouts Kinda Suck

I wasn't expecting to realize that the entire idea of callouts and
asides in publishing is based on the limitation books place on content
in general. Since you cannot link to that other content it gets put in a
big rectangle with a different background and perhaps a fun icon
underlay. But that content could be just as easily compiled into another
document for its parenthetical relevance.

Being able to cover divs without using Pandoc divs by providing for it
at the meta-content component model level frees me entirely from Pandoc
since there is no other compelling reason to use it.

## Other Stuff is Covered

* LaTeX is covered with MathJax
* Spans with attributes are not critical
* References, bibliography, and the rest kinda suck

Seriously, there just isn't anything else that most people would care
about.

## Speed Increase

Removing the dependency on Pandoc means *zero* loss of performance in
order to create a `pandoc` subprocess every time. This is not a small
increase.

It also means that I can use Go templates instead of the sufficient, but
very weak Pandoc templates. Since the very idea of a knowledge component
is itself similar to the new Pandoc partials I get all the additional
speed of *precompiled* Go templates with the solid association from the
Pandoc model.

## Goldmark Extensibility

Goldmark is 100% pure GO and is rock-fucking-solid. It is the fastest
CommonMark parser in the world right now second only to the C reference
implementation. *No* other parsing engine can claim that. This is why
the Hugo project immediately switched to it.

## Cross Training from Hugo

Hugo is quickly becoming the number one SSG in the world and deserves
it. It has been kinda fun to watch it kick all the others asses slowly
but surely on every metric --- finally adoption. Being able to leverage
all the learning everyone has done about Go templates is an obvious win.
In fact, Pandoc template syntax has always been rather quirky and weird.
It is definitely Pandoc's biggest liability next to being written in
Haskell, which practically no one else codes in these days.

## Knowledge Component Registry

Using a template syntax from Go also brings a solid stability that could
easily lend itself to a full knowledge component ecosystem with a
registry and everything. This is perhaps the most compelling reason of
all.
