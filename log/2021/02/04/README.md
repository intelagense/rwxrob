## Thursday, February 4, 2021, 6:18:54PM EST <1612480734>

Yeah, it's pretty clear that other people are lurking on my live streams
and picking up new idea for their more mainstream YouTube stuff. It's
flattering. As long as the momentum to replace the Web is moving on
several levels when people see what KEG can do they'll come around. None
of the alternatives --- including Gemini and Tim's --- have anything like the
semantic combination of Pandoc Markdown, YAML and JSON that I've been
working since [2014](https://github.com/essentials-web) and it now in
[KEG](https://github.com/afkworks). There isn't a moment to lose on this
stuff. All of these fail to realize that none of this has to do with the
Web at all really, it is about creating a *solution* to the unmet needs
of millions of knowledge workers, a need that manifests itself in both
the bloated Web of today and the search engines for it. No solution I
have encountered so far addresses this core need other than KEG.

The really great part is that I know what I have and I don't *really*
need to know that everyone else knows it. I can do my own thing and
astute, intelligent people will naturally gravitate to it because it is
simply superior, just like they did with Linux and Git. I just need to
demonstrate and document that superiority in plain working terms that
show people addressing the need and saying inside, "I want that!"

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
