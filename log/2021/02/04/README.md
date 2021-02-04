## Thursday, February 4, 2021, 2:54:52PM EST <1612468492>

Just had a major breakthrough that was the combination of many other
ideas over the last few years:

* Conversational UI Responders
* `cmdtab` Contextual Tab Completion
* `kn` Actions (Reserved, Builtin, Extensions)
* Common Gateway Interface
* `net/http` in Go

In short, I'll be creating what I am going to call the MimTab™ command
line interface.

The concept is simple: treat the entire command line as if it is a query
to be handled by one or more registered Handlers keyed to specific PEGN
or Regular Expression matches every time the user taps tab, essentially
turning every command line into a localized intelligent search fulfilled
by a single command line application (which can optional support other
commands as extensions).

This approach obliterates the ancient, dead, broken idea of a command
line with options and hashes and specific usage and ushers in further
the era of conversational user interfaces that anyone can use, not just
"terminal masters."

The obvious next step to all of this is to create Mim™ Shell (a
variation of a REPL) which is composed entirely of handlers and adds the
concept of temporary cache between command lines allowing the scripting
of 100% natural language.
