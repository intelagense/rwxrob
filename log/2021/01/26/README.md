## Tuesday, January 26, 2021, 11:17:00AM EST <1611677820>

The more I get into creating the `kn` tool the more shell specific
functionality I realize will eventually need to be hard-coded into Go
(or whatever target language). Here are some specific examples:

* `date -d` notation (ex: yesterday, 6pm tomorrow, etc.)
* `yq`/`jq` query language

These two grammars are fundamental to everything I'm doing so far.

I'm convinced the `jq` query language will become a de-facto standard
for parsing and outputting JSON and YAML data (not unlike `xpath` for
XML). This means compiled applications are going to need to hard code 
query strings. The question is, are there libraries for such things in
Go or Rust or whatever. There's a good chance that `yq` supports it
because it is already written in Go and even if not, the syntax parsing
code could be extracted into a generic library for such things. But if
memory serves, it uses some horrible form of `lex`. Yet another reason
to get PEGN grammars *really* off the ground. We could capture the
entire grammar for `jq` and generate automatic AST parsers for them that
could be embedded in anything.

## Tuesday, January 26, 2021, 5:51:19AM EST <1611658279>

Ya know, it just makes the most sense to standardize KEG log file
structure to be by day. Perhaps calling it KEG daily log or KEG journal
is better. I'll keep log for now, as in "captain's log" or "scientific
log". 

Forcing the standard default to be by day just simplifies everything
when creating tools that use it, which I recently encountered while
integrating recording logs.

Have I mentioned how much I *hate* the term "blog"? No one uses it
properly. Originally, it meant "web log" and was chronological with
entries being short. A "post" to a blog doesn't semantically work. A
"post" is something you hang up more akin to an article or something you
put on a bulletin board. People are stupid. They will find a way to
destroy perfectly good words.
