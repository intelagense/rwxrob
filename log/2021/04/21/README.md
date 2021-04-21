## Wednesday, April 21, 2021, 1:56:33PM EDT <1619027793>

I'm realizing that the standard (but slightly annoying) use of
`/usr/bin/env <whateverthefuck>` is really more relevant in the
container world than in other places because people might be using the
same thing in a vastly different location depending on which container
or host OS. That has always been the reason, but working with containers
lately has really brought that home.

However, people should write most things for containers in POSIX shell
which is *always* at `/bin/sh`.

Python is the real stinker in this bunch since could be any number of
places, and at any particular version. Thankfully, that is overcome by
containerization. And if you are creating a container then you can be
sure the script has the explicit path rather than `/usr/bin/env` before
you `COPY` it into the container.

So I guess, no, `/usr/bin/env` still sucks as bad as it always did.
Maybe people find comfort using it for installations and such, but if
you are using it for that you can take a moment and read through it, and
change the first line. After all, if you are not reading random
installation shell scripts before running the you will end up fucked
(like I was after running the Anaconda installer, what a piece of shit).

## Wednesday, April 21, 2021, 12:19:55PM EDT <1619021995>

TIL that `goodfirstissue` is a GitHub issue label/tag that indicates an
issue that is good *for contributors to get started with on the
project.* This is a very critical thing to get out to the community
because this is one way to get involved with Open Source projects, which
are also some of the best ways to improve skills and even find jobs. I'm
going to start making sure to tag issues that are particularly open with
this, as well as `helpwanted`.

## Wednesday, April 21, 2021, 10:55:54AM EDT <1619016954>

Noticed that the optimal settings to get a Dell Latitude 5411 to display
well from TMUX connected into my main streaming computer are the
following:

* Ubuntu Mono Nerd Font (there's a Windows variation as well)
* Laptop Font Size: 21
* Desktop Font Size: 30

Laptop ends up being 97 width while desktop is 95 but there are no
visible lines or anything if you use my borderless TMUX binary.

## Wednesday, April 21, 2021, 2:50:58AM EDT <1618987858>

66 is the number of thy counting, no more, no less. That is the biggest
a directory name can be without pushing my prompt beyond even being able
to display it on two lines. Lately, I've been prone to create rather
large directory names that match the titles word for word. 66 is a nice
round, evil number.

