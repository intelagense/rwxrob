## Sunday, April 18, 2021, 2:48:11PM EDT <1618771691>

Hit another snap using containers with `you` as the account for this Go
stuff. They can't use GitHub without setting up an account to use it.
That's a complete killer of this learning lab container idea.

Looks like back to just plain old Markdown `README.md` pages that
describe what to do for the lab without any dependency on anything being
setup in a container. A derived registry of such labs would be easier to
maintain because it can point to repos or containers that have example
solutions all setup to examine but force the person learning it to set
it all up for themselves.

## Sunday, April 18, 2021, 2:32:58PM EDT <1618770778>

I'm seriously thinking about simplifying the entire lab concept into a
single `kind cluster create -f ...` since I'm headed down the
`docker-compose` path already and `kind` is definitely a better
investment.

I really prefer `docker` containers but the amount of setup, well, I
wonder if at that point it is just better to automate it with a shell
script and `kind`.

One of the reasons is long term. Say I get Gitea up and working in a
single container, what happens next? Eventually, I will have labs that
involve multiple computers, one with a database, another serving a REST
service, etc. In those scenarios I'll want to have the person learning
ssh/exec into a login system and from there do all the stuff. Each lab
at that point is literally a lab of computers abstracted as a `kind`
cluster by name. It's easier to help people learning get their heads
around the idea of connecting to that one login system and getting work
done.

I'm afraid at that pointing having a `lab` command is going to be the
easiest for people with zero skills to use to begin.

Another option is to have them just deal with the GitHub dependencies as
they would in the real world using private repos until they are ready.
That way they are learning the actual skills.

## Sunday, April 18, 2021, 2:21:20PM EDT <1618770080>

Arg, I hate that I have to build in GitHub dependencies on my own repos
into these labs. The whole point is for them to be stand alone. I really
don't feel like standing up my own git hosting service just to handle
all of these, but why the hell not. People might even end up realizing
they don't even need GitHub at all eventually. Imma try to add GitTea
setup to my base container and use that for Go code hosting. It it
works, it could be a good boilerplate for encapsulated labs that are
*truly* local only.

## Sunday, April 18, 2021, 1:56:26PM EDT <1618768586>

So frustrating sometimes that Go is *so* dependent on a Git hosting
service. Now that `go.mod` is pretty much mandatory, a move that is
worth of the Node governance team (what next `go_packages` or some
shit?). 

I'm frustrated because it completely defeats the possibility for
experimentation without *serious* scaffolding to prop it all up. This
makes instructing others in Go programming difficult to say the least
and overcomplicates experimental labs which have to either depend on an
Internet connection and something like GitHub or have to completely mock
that environment locally including a customized git account setup to
resolve the git addresses locally to bare repos or something.

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

