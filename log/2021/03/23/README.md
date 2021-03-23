## Tuesday, March 23, 2021, 10:33:35AM EDT <1616510015>

I can't fucking believe they fucked with `alt-tab` (which not does not
switch to the latest and toggle, but forces you to use your mouse to
select something). It's so FUCKING stupid! Hey, let's break everyone's
muscle memory burned in over two fucking decades for no fucking reason!

Yeah, I want NOTHING to do with that fucking idiotic company. That's
more moronic that removing the `esc` key. These type of shitty decisions
that make it onto millions of computers are evidence of HOW FUCKING
STUPID MICROSOFT IS! (and why everyone should use Linux).

## Tuesday, March 23, 2021, 8:54:20AM EDT <1616504060>

Here's some ideas on the components and tools of successful
infrastructure engineer that I want to work up to having:

* A primary Linux workstation for streaming and `ssh`-ing into from your
  laptop when you are away from it (like in the backyard)
* A primary cloud server running your `dotfiles` container
* A small `k8s` cluster
* A good hard-wired network throughout the home
* A reasonable wifi network mostly for devices and family
* A great zero-latency, far-reaching wireless headset (not bluetooth)
* A kick-ass KVM with keyboard hotkeys (instead of just a switch)
* A good coffee machine

I would add the following for personal reasons:

* An iPad for watching stuff and testing iOS views of the world
* 

## Tuesday, March 23, 2021, 8:50:39AM EDT <1616503839>

Another thing I've realized is that `dotfiles` is so 1990s. Today having
your own image on <https://hub.docker.com> is so much cooler, and more
flexible. By maintaining one, you can get on any system with Docker and
just `docker pull <your>/image` and get to work.

Say you forgot your computer for some reason and are staying with
family. You can just ask to install Docker and borrow your cousins
computer (with a new account) and do stuff.

Still, if you have Internet, which you would need to install Docker, you
could just `ssh` into your cloud system. That really is the best
solution for most people still. 

## Tuesday, March 23, 2021, 8:37:32AM EDT <1616503052>

Woke with the idea of creating `skilstak/bboost-core` to get people
started quickly and correctly. This way SKILSTAK onboarding will get
pretty simple:

1. Get a computer
2. Get a terminal
3. Learn terminal basics
4. Install Docker
5. `docker pull skilstak/bboost`

I could have container images that build on one another so that people
can practice installing and configuring the parts if they wish. If not,
they can just get doing with exactly the same environment I use every
day --- including the `kn` commands (`pomo`, etc.).

The reason I think this is fine now, is because learning to *use* Docker
is as fundamental as learning to use `vi` today. Might as well start out
right.

This approach removes a lot of the other paths that I really do not feel
like documenting since they are wasteful, problematic, and frankly
archaic. I now believe even having Linux on bare metal might become a
thing of the past (unless Windows shifts to Linux, which they
practically have).

I've felt for a long time that when people ask which "window manager" I
use that they are coming at the problem entirely wrong. No one *today*
should give a shit about your computer, only whether or not you can
effectively use `docker`. Docker is the ultimate equalizer ---
particularly for anything and everything from the terminal.

*Real* work these days doesn't give a shit about your operating system
at all. People who don't understand this ask dumb shit like the
following:

* Which window manager do you use?
* Which distro are you on?

These are all symptoms of a new term I will coin today: *container
cluelessness*. When someone understands containers they suddenly realize
how foolish these questions really are.

I've had these feelings for while since I'm an
everything-in-the-terminal guy already, but the more *I* learn about
containers, the more its not just a feeling, but an objective,
career-limiting reality that *not* understand `docker` is a huge deal.

