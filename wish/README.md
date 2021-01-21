---
Title: I Wish ...
Subtitle: Maybe Someone Will Grant Them
---

Another reason to actually have a web site is simply a place to publicly
proclaim your wishes for yourself and the world. I suppose a wish is not
as weighty as a *goal* or a *todo* list. A wish is something that you
just want, not that you necessarily need to do yourself. In fact,
usually quite the opposite. You *wish* that someone else would do it
instead of forcing you to. Inevitably, I end up granting my own wishes
rather than complaining or sitting around waiting on someone else.

## Lynx Hotkey Binding to Run a Shell Program and Pass URL

Lynx is still the faster browser in town for research, but it is a bit
hard to get URLs and open graphic web browsers when needed. Having a
hotkey binding to run an arbitrary shell command would be so useful.

##  Protobuf Vim Plugin That is Syntax Aware

Currently, the Vim file type of `proto` is supported out of the box, but
the syntax is not using an AST of any kind (it appears) so it can't do
things like line up in-line comments and stuff I've come to love about
`vim-go`. This is probably the type of project I want to work on *after*
finishing PEGN so that I can define the grammar of the syntax in PEGN
and autogenerate a Vim plugin in Go that handles the AST.

## JSON Schema to Protobuf Converter

The world (since 2019) is apparently going all gaga for gRPC and
its Protocol Buffer interface definition language (IDL), and, I'm
thinking, for good reason. I have *really* enjoyed working with it, much
more than I would have imagined having PTSD for all things RPC from
previous decades. In fact, I plan on doing everything in Protobuf for
new projects now because it generates the most important code in a
code-agnostic.

That said, there is a bunch of stuff already specified in JSON-Schema
that would be really nice to just convert over to `.proto` files. Such a
tool does not exist so far that I can find.

## Twitch Font

Someone needs to create a font that has close equivalents of all the
most common Twitch emojis and icons. It would be a lot of work, but it
would be awesome. It could be combined into a variation on one of the
"nerd" fonts that already have the adjustments for powerline and stuff.

## Smarter `back` Utility

It's just a small thing, but I want to make my `back in 10 minutes`
detect the width of the script and change the figlet font according to
how much room it has. I'm already doing something like that in the
[Just say No to `nano` project](https://duck.com/lite?kd=-1&kp=-1&q=Just say No to `nano` project) but wonder if it could just be done with a
shell script just as easily.

## Go Template Utility: `gotmpl`

[*Here look I even started the repo for you.*](https://gitlab.com/rwxrob/gotmpl)

It occurred to me that Go is actually the perfect language to write a
little `gotmpl` utility that reads from a collection of cached
templates (both `text` and `html`). Sometimes all I want is the lower of
Go's templating system for random things and usually just want to
combine them with some YAML or JSON or just a list of stuff on the
command line.

I could use it to manage snippets from within a `vi` session. Some of my my favorite methods to add to a simple data struct:

* `String() string` - the `Stringer` interface as JSON
* `JSON() []byte` - compressed JSON buffer
* `Print()` - `fmt.Println()` of the string JSON
* `Parse(jsn64 []byte) error` - JSON (optionally base64 encoded)
* `Load(path string) error` - `ioutil.ReadFile()` then `Parse()`
* `Base64() []byte` - `JSON()` in base64 encoding

I should probably also make something that creates the `Example*` test
cases as well. Actually, I should just do one `ExampleData()` (or
whatever) and then just include a bunch of examples of using all the
related methods.

## MainMaker, Generator Replacement to Go Cobra Commander

I fucking hate the [Cobra Golang commander
module](https://github.com/spf13/cobra). I always have. It is far too
heavy and invokes a external dependency for no good reason. I *know* it
is used in Hugo and Docker. I don't need the reminder. It's one reason I
hate it. It sucks ass. The people who created it and maintain it ---
especially Steve --- are absolutely amazing people. This isn't about
them or how great their programming skills are. It's about the shitty
design of a shitty module that squandering an excellent opportunity to
be what the modern, conversational, sub-commander CLI world *really*
needs. Instead, it perpetuates the Boomer-era bullshit we've come to
tolerate and not question because we think we just have to accept it. We
don't.

Perhaps the most annoyingly brain-dead thing Cobra does is generate a tab
completion shell script that actually encourages people to put `eval` in
their configuration files. It is as if no one on the Cobra team ever
even RTFMed the Bash `man` page about *Programmable Completion* that
clearly says you could have used the option to run a program instead
(`-C`). So instead of that big fucking, unnecessary shell script you
could have written all the completion code *in Go* in your program
instead. Duh. What's worse is encouraging people to use *the far more
secure practice* of adding `complete -C foo foo` to their shell
configuration files instead. No need to have root access to put files in
`/etc/bash...<whatever the fuck>`.

Stupid, right?

Rather than just complain I actually tried to follow my "don't get mad,
get busy" mantra and fixed these glaring problems in my
[cmdtab](https://gitlab.com/rwxrob/cmdtab) package. But I'm even
realizing *that* we can still do better. The unnecessary methods always
bothered me. We know enough about the program when we are building the
thing so why spent all that wasted run-time just to figure out the
arguments and get them to the right place? My commander --- as
*beautiful* as it is when it comes to combining completion with amazing
documentation that travels *with* your command instead of in addition to
it requiring it to be installed --- missed one *major* advantage that
*no one* has created: zero dependency with code generation.

Yep. The last great frontier for the world's best commander is code
generation. By moving this work into a tool instead of a module we drop
the dependency entirely and have sustainable code that we can even tweak
to our hearts content. *That* is the Holy Grail. *That* is something Rob
Pike would likely actually approve of. (You'll never get his honest
opinion of Cobra, I promise.) The start up speed impact would be even
more negligible because the generated code has intimate knowledge and
far less indirection through flow control and unnecessary function
calls. There's simply nothing faster and more elegant. It *has* to be
made. Someone make it before I have too.

## Terminal Typing Training

All the web tools are lame. Even those designed for coding suck pretty
badly. There needs to be an exciting terminal typing training tool that
allows all the keys a technologist needs to be learned. By creating
snippets in a given language it can double as a method of familiarizing
oneself with a particular language syntax while at it.

I'd like to write it in Go using the `cview` GUI library.

## Namecheap Command Line Interface

I use Namecheap a lot when looking up domains. It would be nice to have
a better search capability from the command line than the horrible
graphic user interface.

## Restream.io Command Line Tool with Chat

I *really* need finish the <https://restream.io> API command
integration. I'm hopeful that the Websockets API it provides does what
is necessary to replace my WeeChat TMUX pane on my live stream. Please
someone do it for me so I don't have to make it. Who am I kidding, I'll
have to make it myself or I'll never be satisfied. At least I could fork
it, I suppose. It just doesn't exist right now because no one cares
about the terminal anymore, too many low-brow graphics interface users
overwhelm the minority interests of we terminal people.

## Replacement for Find with a Good Query Language

The `find` command is an essential tool for system administration ---
particularly for disk usage and security audits --- but the syntax of
the arguments has always been ridiculously difficult to remember
requiring multiple references to the man page. 

Wouldn't it be cool if we came up with `find` replacement that has a
drop-dead simple query language (maybe even just SQL) that it read
either as a single string or passed into standard output, or just placed
in a file with a she-bang line pointing to it. It would be an *actual*
language for finding things on your system (no not like `awk` or `sed`).

One approach would be simply to create SQLite flat-database files and a
light-weight interface to them so that existing SQL tools could be used.
This is something that could be done rather quickly. By having different
SQLite files created queries could be made across databases to determine
changes in the file system.

## Nark, An Open-Source "Tripwire" Replacement Idea 

The original
[`tripwire`](https://github.com/Tripwire/tripwire-open-source) seems to
be alive and well. However, it looks like it could use an update and
perhaps a rewrite in a modern language (like Go or Rust). I'd also like
to focus on the reporting aspect above all and add both inbound and
outbound ports to the monitor to "nark" on (get it) stuff that is not
supposed to be on my system.

Another consideration would be the license change from GPLv2 to Apache.

What a great project this is for learning a specific language as well?
It's practical, uses tech and techniques that are relevant, and it's
relatively simple to implement.

## Watch With Me Application / Service

In order to comply with copyright law live streamers cannot do what they
would normally do, listen to copyrighted music and watch their favorite
videos and movies. This sucks because in order to live stream currently
you have to give all that up, but that is part of the fun of being a
coder, being able to jam out to Cake while banging on the keyboard.
Something needs to be created that keeps everyone watching the stream
informed about the URL and the location within the content enabling
stream viewers to be at exactly your same location so that people can
chat about it and comment on it like Mystery Science 3000. It might be
something as simple as a browser plugin and a service that can be
periodically polled by the plugin.

And another thing, I don't necessarily want to force those watching my
stream to listen to my music, or to watch the movies I'm watching. They
might now want to be distracted or perhaps they have their own music.
I've heard people say they actually *don't like* streams that have music
for that reason.

To date, most streamers will jam out with headphones without their
viewers knowing what they are listening to. That's fine, but it sort of
misses the point --- especially when wanting to view documentaries
together and such.

## Coder Coffee Shop Desktop App

I'd love to just be able to pull up an app on a secondary screen and see
everyone's screens while they are coding. The app could rotate through
random people and give them a 4x size for a bit just to highlight what
different people are working on. Those screens that don't change could
auto-minimize themselves so only people doing stuff would be on the
screen. Then we could add like a Discord chat room with it all. The one
solid requirement would be that it not crash constantly.

Even something that is just terminal based would be amazing, perhaps
something like what TMUX or Screen shared session do right now.

## Reformed Educational Systems Focused on Learning

Most public education systems (and most private ones as well) are not
[focused on learning](https://youtu.be/GFD37HvGx6k?t=181). They are focused on meeting quotas, perception
management, making money, pushing people through, and glorified
child-care. The most motivated teachers in such a system simply cannot
accomplish true learning in their students. When learning happens it
comes at the cost of dependency on the teacher and system instead of
producing true autodidactic individuals ready to face the world.

Teachers who have gone through the formal education to become one either
love or hate that when I say this. I'm convinced the science surrounding
education is seriously flawed and focused on the wrong outcomes. Ken
Robinson agrees. This is why I want nothing to do with traditional
academia and teacher certification. It turns well-intentioned teachers
into drones with their aspirations burned out by an objectively failing
system with upside-down priorities.

The path to reform means significantly changing the culture of
education. It means turning teachers into mentors and learning
facilitators, not subject-matter experts who have *actual* experience.
It means obliterating ego and promoting empathy in the learning
environment. It means having SMEs and facilitators working together. The
SMEs produce the content, challenges, and activities. The facilitators
organize learning lab environments where *the learner is in charge of their
own learning*. Responsibility for learning is placed on the learner from
the very beginning rather than creating an antagonistic teacher versus
student approach we have today. We need the courage to work
toward motivational skills and let those unwilling to learn simply
not learn, to hold them back for as long as it takes until they decide
on their own terms to begin and own their learning instead of trying to
*force* them to learn.

These techniques have been proven successful over and over again in
environments where the traditional systems gets completely out of the
way, from the groups of "lost learners" who are essentially thrown into
detention during school all day, to the "gifted and talented" deemed to
smart for the current system. Both of *those* systems flourish because
the teacher is not a teacher, but a motivator, a facilitator, and yes,
even a responsible friend. 

All of the plights of 2020 can be traced back to a lack of good
learning. Those who learn and love to learn discover empathy and
facilitators working in this way *must* be above all experts in
emotional intelligence and empathy, councilors more than teachers. Most
people who remember their "best" teacher will recall all of these
qualities in that person who operated this way within a seriously
broken system.

## Better PEG Grammar Notation and Tooling

*Update: I'm actually [making this](https://pegn.dev)
if you want to help.*

PEG ([Parser Expression Grammar](https://duck.com/lite?kae=t&q=Parser Expression Grammar)) is by far the most important advance
in language grammar development in the last 20 years (even though many
CS departments still don't even talk about it in their curriculum). The ecosystem
around PEG is fractured and full of over-engineered code generators that
are implementation language specific. What we need is a standardized
notation based on Bryan Ford's original (but incomplete) "example" in
his paper that started it all. Then we need to show that tooling can
exist --- including code generators --- that simply use the grammar and
don't intermix PEG with additional syntax to render in a specific
language. That way the same grammar files can be used to generate
parsers of different kinds in different languages. The same grammar file
could be used to generate a highly optimized parser within a single file
or another that uses very flexible AST approach and parse function
library.

PEG is particularly useful in computer science education not only to
create grammars, but to document existing ones and get the mind to
understand what syntax is as well as any specific syntax. Capturing an
existing language in PEG is a great educational exercise.

## Fast, One-Pass Parsable Markdown (Ezmark)

Markdown is amazing, but not even the CommonMark initiative has
recognized the value in creating a version of Markdown that can be
parsed in one pass instead of two. This is because of reference links
and other stuff in blocks at the bottom of a Markdown document that must
be parsed before anything above can be rendered that depends on them.

All the advances in Markdown --- mostly in Pandoc --- have been toward
become more flexible with adding extensions and providing more and more
additions to the syntax. This is entirely against the original goal of a
standard, simple language that is better than HTML than anyone can learn
in 20 minutes. We need a Markdown that enough people could use without a
lot of training not just because it is simple, but because HTML needs to
be replaced with some version of Markdown.

## A Decentralized Knowledge Network

This is one I have been working on for several years, something to
replace or supplement the world wide web that maintains the motivations
of the original web: knowledge exchange. It will leverage the
infrastructure of the web without depending on it. It will provide
localized content searching through semantic structuring, parsing, and
local caching. It won't even need a network. It will provide
digitally-signed attribution of authorship. It will use trust to improve
content discovery.

*Development is moving along slowly [here](https://gitlab.com/rwx.gg/kn).*

## Bash Shell Rewritten Under Better License

Seriously, Bash is so bloated and the code base so old and crusty that
the whole thing needs a serious rewrite under a license that will allow
it to ship on embedded hardware and such instead of Bash because of its
completely stupid [GPLv3 license](https://rwx.gg/views/gnuisdead).
(No [Zsh is *not*](https://rwx.gg/advice/dont/zsh) the answer.)

## Kick-Ass Educational Chemistry 3D Shooter

I've long wanted someone to take me up on the idea of creating a really
rated M educational chemistry shooter that does *smell* like a "for
education" game. It could be a full triple-A game that requires you to
prepare different ammo and weapons to fight an wildly different set of
monsters and mobs using the right stuff. There could be a chem lab where
you see the periodic table in 3D, play around with simulated
combinations of all the chemicals to learn about their properties and
more.

Hell just a realistic chem lab that taught about what happens when you
combine different chemicals would be amazing.

## YAML "Lite"

Let's face it. YAML kinda sucks. It allows for execution of code and
does far more than it ever should. There needs to be a subset of YAML
that is strongly specified, a YAML *lite* so to speak. 

The tough part would be deciding what to throw out and what to keep. For
example, some people despise linking, but I love it. It saves on so much
redundancy when using YAML as a replacement for a database --- one of
its stronger use cases.

~~If~~ When I finish [PEGN](https://pegn.dev) something like this will
be a lot easier to consider.
