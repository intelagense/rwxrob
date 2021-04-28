## Wednesday, April 28, 2021, 10:34:10AM EDT <1619620450>

Don't know if it is an official convention, but in Go code the docs
usually to not include parentheses when referring to functions and
methods, which makes sense because doing so implies that the arguments
list needs to be included as well. Plus, it makes for cleaner reading
and easier searching. For example, when searching through code quickly
to find the beginning of a function or method adding the initial
parenthesis isolates what I want and doesn't match all the stuff in the
documents. I know there are cleaner ways to do this that involve
language-specific plugins, but this works for any language, anytime.

Moral of the story: just use names of functions and method when
referring to them within in-code documentation.

