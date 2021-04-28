## Wednesday, April 28, 2021, 11:59:06AM EDT <1619625546>

Finding out that I should use Helm for delivering my modified version of
JupyterHub. I wanted to do that anyway, but misunderstood that "it was
overkill". I've been looking at Terraform scripts as well (for the
future). But they are *definitely* overkill and not really alternatives.
What I'm looking for specifically is an easy way to bundle and track
versioning on these changes and everything combined together.

Now I really need to see how Helm behaves in Kind and how much it
integrates with stuff from the registry. The bottom line, when it comes
to CN applications development identifying what stuff should
go into the container image (Dockerfile) what what goes into the
configuration (effectively though arguments) is a tough decision. Here's
a sample list:

* What goes into the Dockerfile?
* What registry should host the image?
* How configurable should the image be?
* How should the image configuration be done? (env, args, YAML, etc.)
* How much configuration should be Docker-ish?
* How much configuration should be Kubernetes-ish?

These are entirely new considerations for most applications developers,
who generally just think about how to package or containerize their app.
This is not just containerizing but anticipating all the different
configuration possibilities as deployed to Kubernetes, including things
like concurrent Pods and network services. Come to think of it, it is
the same stuff application developers thing about with a single
application, but scaled up to multiple applications combined together,
things like concurrency and consistent, persistent storage.

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

