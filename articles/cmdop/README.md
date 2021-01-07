---
Title: Command Oriented Programming / CMDOP
Subtitle: Recovering from OOP Thinking
Created: 2020-06-23T07:58:39-0400

Summary: 
  For me the worst part about having been raised to code and think like
  an object-oriented programmer is how uncomfortable I feel when methods
  don't have an object. The solution seems to be thinking of everything
  as commands to be executed that require data and then working out
  where the action is, what the parameters are.

---

## But Who *Owns* That Action?!

In my mind, I fight with myself about what things do and who's *job* it
is to do what even though that isn't even the right thing to think
about. I'm ruined from people telling me I should be writing everything
down that a thing does on an index card. I *have* to have a virtual
index card to contain all the actions. 

I'm so pathetic.

I'm uneasy with the idea that a method can exist all by itself and not
be *owned* by anything except the program or package in which it is
contained. 

The truth is our reality has a *lot* of stuff that is just stuff, that
doesn't *do* anything, that is acted upon by other things.

## Factory Hell

One sure sign OOP has fallen on its fucking face was when suddenly
everything became a "factory" in Java and CPP. How stupid did we all
have to be to allow ourselves to make nouns for things that were clearly
always just verbs. Rather than allow the calling program or package to
be the owner of such actions we made up all kinds of stupid classes that
could receive our nice pretty, little, concrete blueprints for how to do
everything, never to be changed again. 

## Unix Philosophy in Code

* Create things that do one thing and do it well.  
* Create things to work together. 
* Use data as the universal interface.

One of the practices that clearly killed OOP was building monolithic
classes sort of out of necessity. Inheritance was the same thing because
the monolith depended on everything above it. As usually the solution is
to break big things down into agile, composable components. It works for
*everything* from social structures, to companies, to operating systems
and software.

The more I think of each bit of code as one procedure, function, or
method that does one thing extremely well the better my code base.

## Command Structs to the Rescue {#command-structs}

A *command struct* is my word for a light-weight struct that has *one
action method* that is an implementation of a one-method interface. Such
will almost always be named something related to the interface with the
interface name also matching the suffix of the command struct's name:

```go

type interface Renderer {
      Render(node string) error
}

```

```go

// here's a minimal command struct

type PandocHTMLRenderer struct {}

func (r *PandocHTMLRenderer) Render(node string) error {
      // ...
}

```

```go

func Render(node string, base *Base, r Renderer) {
      // ...
}


```

Go has figured out a big part of the solution: one method interfaces
combined with light-weight structs that fulfill that method. The struct
mostly exists to provide context to method simplifying its argument
signature sometimes to nothing.

This is *far* more effective and sustainable than creating complicated
parameter signatures --- or worse --- non-specific signatures (like
Python) that allow *anything* to be passed in forcing the parsing of
them every single time it is called preventing calling again with the
same or similar parameters without parsing them all over again.

Go even allows for drop-dead simple marshaling and unmarshalling of the
data in the struct making it even better suited for modern software
development. Rather than passing arguments on every call, the
light-weight struct can be inflated from JSON or YAML and then only
changed as needed between each call.

The concept has been around for a while. It's sort of another take on
mixins. You take some functionality, encapsulate it, and make it
available to anything that needs it. Each light-weight structure
wrapping one main method becomes the equivalent of a command in a Unix
system.

## Think In Terms of Commands 

I don't have to go all the way into a strictly functional programming
paradigm to benefit from not creating dependencies on the encapsulating
context and state of the thing, at least on a large scale. The idea of
keeping small command-like structs is that the struct data is more of a
set of highly-efficient argument parameters than an environment.

The point is, the less dependent everything I write is on other stuff
that was not fed into it the better. It's easier to test, debug, and
maintain later. Those are the big wins of the functional paradigm.

Here's an example I recently had to decide. I'm creating a knowledge
management utility that uses Pandoc currently for rendering Markdown
into HTML. The action could be easily stated in English as the
following:

```

Render a node from a knowledge base using pandoc as html. 
^^^^^^   ^^^^        ^^^^^^^^^^^^^^       ^^^^^^^^^^^^^^
action   what            where                 how

```

In code this would look like this:

```go

kn.Render(node,base,renderer)

```

For a time I was tempted to make an implicit renderer with a strong
dependency on something in the `kn` package itself, say a default
renderer.

```go

kn.Render(node,base)

```

Before that I was thinking of putting all that action into the knowledge
base *class* itself:

```go

base.Render(node)

```

It's obvious to see how different that is from the command approach. 

The OOP looks downright stupid. The focus is on the *thing* and not the
*command* (action).

The first is just a translation of most spoken languages and could be
easily thought of as a shell command:

```sh

kn render <node> [<base>] [<renderer>]

```

In fact, that is *exactly* what the command is. The place for implicit
defaults is at the highest level --- the user level --- not in the code
itself.

That command is *so* clear, in fact, that it could be spoken to a
conversational voice interface and executed. Think of `kn` as *Alexa* or
*Ciri*.

Thankfully the world has *really* woken up from the absolute idiocy of
thinking of everything in terms of objects instead of actions. I know it
isn't a thing yet, but I am going to start calling it *Command-Oriented
Programming*. 

## Parallels to Well-Organized Militaries

Creating a military force --- as with some war games --- requires
different units that perform one main task really well. The pikeman
holds off the initial attack. Archers hold back and launch from the back
ranks. Cave trolls haul heavy war machines. The more organized and
specific the roles the better. Commanders then give the command with a
few parameters and stuff just works. The commands are not complicated.
Only the minimum amount of information is provided as parameters to keep
the communication concise, clear, and quick. The units have been *coded*
with all the other instructions and data they need to perform their main
tasks when commanded.

When a military is well trained and prepared and able to adjust to the
environment of the battlefield the only command needed --- for the
entire collective of battle-ready fighters --- is "fight!"

We can learn a lot from this comparison while creating software designs
dynamically.

## Learning from the `exec()` System Call

Perhaps it is all my love for the Unix command line that has driven this
discovery. The secret to what might be the most successful software
design paradigm in history might *already* have been captured in the
most successful human-computer interface, which is captured in the
simple but powerful Unix `exec()` system call. The entire Unix
philosophy can be entirely represented in its humble function signature:

```go

func Exec(argv0 string, argv []string, envv []string) (err error)

```

The `argv0` is just the name of the command. Then all the information
needed by the command are passed as `argv`. But the secret just might be
that the environment, the context, the surrounding conditions of that
command, are passed as `envv`. 

The not-so-amazing reminder here is that the `envv` can be passed to
*multiple different commands*. Rather than encapsulating each, the
commands are given the information they need to do their jobs and the
larger environment that they *share with every other command.*.

Obviously, using the environment can create problems, but it is an
essential part of Unix's success. Considering that *every other major
operating system* also has environment variables in some form it is
worth taking note.

In terms of our fantasy army metaphor a number of soldiers are given a
single command with parameters, the tools and munitions to execute the
command, and have an implicit environment --- the battlefield --- in
which to carry them out. 

Say its raining or windy. The archers are expected to check the
environment and adapt, not be commanded specifically about how to
execute the task they've been commanded. The only information besides
"fire" in the command given is where, not how.

## No It's Not Functional Programming

Some might be tempted to think of this as functional programming. It's
not. It has a lot of inspiration from it, but strictly speaking command
oriented programming is more broad. It is okay to violate the sacred
functional commandments in order to code and execute a command in the
easiest and clearest way possible. 
