## Sunday, January 31, 2021, 1:34:19PM EST <1612118059>

And another thing, Markdown was originally written in Perl.

## Sunday, January 31, 2021, 12:52:35PM EST <1612115555>

I'm actually a little excited to *re-learn* perl. It's absolutely crazy
that I wrote this extremely complicated, low-level
[`perlclasses`](https://github.com/rwxrob/perl-classes-pragma) pragma
and now I barely remember any of it. Looking through it is really
bringing it all back though. It's hard to believe I used to give
presentations on this stuff back in the day.

I am left wondering if, um, this will affect the focus of my imminent
job search. I'm starting to feel like SRE work is still deep in my
bones. I've always seemed to be able to straddle the fence of
development and operations. At this point the main priority is cleaning
up my GitHub repo and that involves porting a lot of this to Perl.
Nothing is significantly difficult.

Another advantage is how powerful having Perl back in my toolbox will be
when it comes to creating the cluster for automated bugbounty
monitoring, not to mention dealing quickly with stuff after getting in.
Perl has always been a preferred hacker language for that reason. Go
compliments it well because it is a single executable payload when the
time comes.

In fact, having a significant amount of Perl in my GitHub will actually
be a good filter for dumb ass potential employers who do not truly
understand when and where Perl is the unquestionable best pick. People
who dismiss Perl immediately out of hand and write shit like "a healthy
disdain for Perl" in their Job descriptions (including that completely
moronic company F-Security) can fuck right off. I don't give a shit. If
you don't understand it --- and especially if you make fun of it without
understanding it --- I want *nothing* to do with you. I'll fucking take
a janitorial job before I work for you.

And to all those people who call me "too emotional," well, you are the
first to understand me. If that makes you uncomfortable hiring or
working with me, you can fuck off as well.

## Sunday, January 31, 2021, 12:21:15PM EST <1612113675>

And, of course, Perl has its own license which is *not* GPLv3!!! Oh my
god, I love Bash, but that means I can keep my stuff free from that
shitty license.

Here's the greatest irony of all. This will make some people fucking
come unglued on my flip-flop. But I have always said the only reason to
learn Zsh was because it was not GPLv3 (which is why Mac uses it). The
simply fact is, that if I keep all my shell in Bourne/Dash/Ash (POSIX-y)
and all the other stuff in Perl that I've avoid GPLv3 completely. That's
something I simply cannot be ignored.

So here's what 2021's experiment will look like:

Language|Uses
|:-:|-
`sh`|Rudimentary Scripts (Bourne/POSIX/Dash/Ash)
`bash`|Interactive Shell (Only)
`perl`|Scripts Beyond Dash
`go`|Anything Compiled
`c`|Stuff That Requires It

And no fucking Python unless I'm writing automations or something, not
even for ML, there is plenty in Go for that now (that is actually
faster).

## Sunday, January 31, 2021, 12:04:50PM EST <1612112690>

I think I finally hit my limit with Bash. For the past year I've been
conducting a personal programming experiment. I decided to allow myself
to code Bash and use all the Bashisms I possibly could in an effort to
see if I could effectively replace Perl and Python with Bash and Go.

Well today I hit two *big* fucking blocks to productivity during this
experiment.

The first is support for `\p{Cc}` and other Unicode class regular
expression support. As usual, since Perl is literally the best language
on planet Earth for regular expressions --- having defined the modern
standard that is now included *in every other language* --- it is
naturally the first to invent and deploy support for full Unicode
classes. (Oh, and by the way, Perl supported Unicode *way* before Python
and didn't need a breaking change like Python3 to fix it.)

You simply cannot code in any other scripting language with that level
of regular expression and parsing support. Just think of `pack` all by
itself?!

Given the amount of language and knowledge manipulation that I do
regularly the primary criteria of the `kn` implementation language is
rock solid support for parsing and regular expressions. For that, Perl
is simply the best.

The second is small but oh so important: POD. I went to create an
efficient way to document my Actions for `kn` only to realize there is
simply no efficient way to do this in Bash. Go has it but they I have a
bunch of compilation steps for different target OSes, perhaps in the
long run, but not for throwing together stuff quickly. Perl POD comes
*after* the code (`__END__`) so the interpreter never gets bogged down
by parsing over it. This is a level of simple genius that most people
today just don't fucking appreciate. Python does *not* have this. It's
compiler is fucking brain dead, which is Guido is pushing so hard for a
new one in PEG. Again, Perl is king of *solid* script documentation.

I have to sit on this conclusion a bit, but these are *huge* discoveries
made in the fire of scientific personal research that I simply cannot
ignore. Perl is *literally* the best tool for *this* job, despite its
moronically uninformed unpopularity today. Given the amount of natural
language manipulation that I have been doing I have to consider the
possibility of prototyping `keg` and `kn` in Perl. 

I kind of like the prospect of building the next Internet Knowledge
Exchange Grid primarily on Perl to get started. Intelligent engineers
will make their considerations and discover why this is the right thing
to do.

## Sunday, January 31, 2021, 11:32:01AM EST <1612110721>

Struggling with the potential need for a convention for documenting
Actions. 

* Should I have the documentation be in a separate YAML file?
* Should embed it?
* Is YAML overkill?

People from the MAN page generation would say to keep is separate since
every execution of a Bash script is slowed by parsing the doc lines
unnecessarily (even though it is nanoseconds to do it).

This is where I *really* miss Perl where the convention was to put it in
the file after a certain point that the interpreter would always skip
and therefore didn't affect compilation. There's really no reason not to
use Perl for this actually, it's on everything by default as much as
Bash is. But people will really frown about it. Eventually Perl Actions
could be replaced with compiled Go code.

The thing about Perl is that it is absolutely the right pick to
prototype this stuff *mostly* because the `perl` executable is on
*everything* where a consistent version of Python is not. Besides,
Python start up times are fucking horrible and always have been. The
compiler is absolute shit.

But if I do this in Perl everyone will immediately write it off for all
the wrong fucking reasons. You know what *no* other scripting language has
right now? `\p{Cc}` Unicode notation. *Only* Perl supports that shit.
And yet assholes have the gall to suggest Perl is a "boomer" language.

I'm in a fucking bad mood as it is today, thinking about lesser
programmers throw shade at a language they know fucking nothing about
just pisses me off more. So Imma remind myself how to code in Perl
(after having set it down and coded nothing in it for over 15 years)
just to fucking piss them off.

## Sunday, January 31, 2021, 11:23:26AM EST <1612110206>

OMG I turned off the lights in my room and am coding like I used to all
the time and I cannot believe how much better it is than when I have
them on. Also my fucking fingernails are fucking with my typing. 

## Sunday, January 31, 2021, 11:15:41AM EST <1612109741>

I'm realizing that having the Action documentation for `kn` cannot be a
comment in a script if I'm going to keep the UNIX way of not caring what
languages the actions are written in. So it will have to be a command
that all of them support, probably `help`.

## Sunday, January 31, 2021, 10:44:41AM EST <1612107881>

Only chat streaming for me today. Last night knocked me out. Had to
cancel Beginner Boost. I need some recovery time. Ironic that coding
quietly is one of my best methods to do that.

And another thing, decided given the horrendously bad live chat API at
YouTube that I'll no longer be supporting live chat at all. There's
another *big* reason. After watching [this John Hammond
video](https://duck.com/lite?kd=-1&kp=-1&q=this+John+Hammond+SUDO+video) I
noticed something odd. He's chapters appeared on the side where live
chat normally would. Looks like when you incorporate live chat you give
up that ability. So I am just going to make sure that people know about
<https://chat.rwx.gg> and the options with Twitch and Discord. I know
that will be hard because it will still be two windows for some people,
but it is just not worth it to throw away chapter playlists and deal
with the stupid polling chat api for YouTube live chat. They won't be
able to keep that.

## Sunday, January 31, 2021, 1:06:42AM EST <1612073202>

I wish I could remember the name of the currency that is in control and
can just be printed more when the country controlling it wants to. It's
not "fiat."

