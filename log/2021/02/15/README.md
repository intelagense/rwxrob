## Monday, February 15, 2021, 9:07:27AM EST <1613398047>

You know, I'm a hacker. I always have been. I don't know why I don't
just embrace that as a career. Thinking about what I was watching
on EsportsHacking and I was almost yelling at the screen about all the
better ways to do what everyone was doing.

I have the core skills and
the experience, I just fucking hate working with shitty tech, even to
own it. For example, I'd have to (re)learn PHP all over again.

Today I'm really feeling the pull of what I'm naturally good at. Sure
I can do all the other stuff. Sure I can be a software engineer, or an
SRE, or a DevOps dude. But when I think about NahamSec and other
CTFers success I'm reminded that the coding that makes these hackers
successful is *exactly* the coding I'm really good at and drawn
too, quick, dirty, effective, and powerful, which is why I *still* love
Perl so much. It blows the fuck out of Python for everything in
pentesting *except* the shit that scriddies use. 

Go could replace Burp suite with something so much better. Hell, I could
sell that shit after releasing as open source.

When asked about "note taking" --- something that comes up all the time
during conversations between hackers --- the methods they use are arcane
and silly compared to `kn`. They don't even know it exists. No wonder
NahamSec dropped by during my knowledge management streams. Dude knows
what's up, which is how he manages to *pay* \$7k/month just to keep his
k8s rig up at home doing all the scanning. Hell, that's easier than
setting up a Bitcoin mining operation for someone with a lot of coding
experience.

I'm realizing that one of the core differences I have from *most*
hackers is simply the ability to code good software. Most hardly know
how to code, and when they can, well, they look like kindergarteners to
*actual* software developers. I have the unique experience of doing both
systems operations and systems engineering and software development. I
just have to fucking use them both, and frankly, not even the modern
idea of "admins who code" (SREs) touches what is possible for a someone
with such skills in the redteam world. 

Why the fuck am I not going for this? Is it just my pristine arrogance
not to be exposed to shitty software? Probably. But let's face it.
Accepting a corporate gig of *any* kind means seeing that shit
*everywhere* so might as well take advantage of it and fucking own it
for greatest satisfaction rather than trying to sustain it until it
blows up in our fucking faces and everyone points the finger at me. Imma
be the blowing up that shit, not taking the fall.

All of this has caused me to really rethink my long term goals. After
all, if was hackers who created all the cool stuff we use every day on
the Internet. And it will be hackers who save the world from the
Idiocracy taking control of it right now. 

Kubernetes and such are still a core part of the entire thing. They all
overlap. So learning that shit ain't taking a back seat at all, right
after I redo the Web, I mean, finish the first iteration of KEG.

I really just want to be Gilfoyle. Is that asking too much?

## Monday, February 15, 2021, 8:59:55AM EST <1613397595>

Tough getting my Perl chops back after working in Bash for so long. Bash
is decidedly lame at doing things quickly and tersely, the very thing
Perl gets a bad rap for is the thing that makes it dominate so
completely as a shell scripting language. Again, why the fuck did I ever
leave it?

## Monday, February 15, 2021, 8:44:55AM EST <1613396695>

I'm wondering if I have been too harsh on the simple `#!/usr/bin/env
perl` in the past. I've always hated how insecure and unnecessarily slow
it is, but it is damn convenient and great for writing scripts that can
easily and quickly be executed and shared.

I suppose my biggest complaint is that people (once again) take a simple
convenience and deploy it into production without doing the minimal
effort to correct the shebang line in their installers. This is why I
hate it more than anything. I'm afraid that if I start doing it on
stream people will think they can (and should) do that in production.

After all, scripts should only be for hacking CTFs and utilities that
don't really matter. Everything else should be in a more substantial
language. And the `env` trick is in-line with the goals of that usage.
So I'm going to start using it again with that in mind. I can't prevent
people from doing stupid things either because they don't know better or
willingly choose to deploy that shit into production. Why should my
effectiveness and ability to share be affected by other people's
ignorance and/or carelessness?

