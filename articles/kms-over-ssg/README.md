---
Title: Knowledge Management Over Site Generation
Subtitle: Better Than Hugo, VuePress, Gatsby
Created: 2020-06-28T08:25:51-0400
Query: true

Summary:
  Hugo showed the world what Go could do with speed, and how to
  over-engineer your project to death with unnecessary complexity and
  shitty design. In contrast, VuePress showed us what elegant,
  zero-configuration design looks like. Gatsby built everything around
  data sources and the world swooned. But they all share the same flaw,
  they are static *site* generators specifically focus on generating
  HTML above all. But the real need has *never* been about HTML it's 
  *always* been about knowledge. Look around. There are
  very few simple, open knowledge management solutions (KMS).

---

One day we will all look back at how silly and specific the term SSG
really is. "It's not about sites, stupid. It's about knowledge." That's
what people will say in five years as the mainstream suddenly
understands the *real* problem. Even if `kn` just remains a fun pet
project of mine, the dawn of the knowledge management solution is at
hand and it just might be the first *true* KMS ever made.

## What's a KMS?

Use of the term *knowledge management* (KM) has been overrun by buzzword
corporate speak and advertising. Everyone *talks* about knowledge
management but what are they actually *doing* about it (besides trying
to sell you something).

A *knowledge management system* is about knowledge and people, not just
sites, data and information.

* Focused on *individual* knowledge management anyone can use
* Universal knowledge source syntax (Pandoc Light Markdown)
* Social sharing: attribute, subscribe, recommend, and warn
* Dependable semantic information structure
* Search configured and controlled by the *user* 
* Render-neutral content organization 
* Knowledge content component model
* Zero external dependencies
* Required command-line interface
* Optional graphical interface
* Universal protocol support
* Decentralized

### Not About Content Management

Content management systems have been around a while, but they all focus
on the wrong problem, managing content. They don't focus on the problem
of making that content more digestible and sustainable so that it can be
of best use to the humans who ultimately need to convert it into
practical knowledge.

A *true* knowledge management solution focuses primarily on how to get
skills and information into the human brain thus converting the most
sustainable and digestible content possible into practical human
knowledge. A working KMS can be thought of as a framework and interface
for programming the human brain.

### Not Even About Learning Management

Learning management is another term thrown around by people trying to
sell you something. It is usually slapped on crappy software that makes
far too many assumptions and dictates about what learning is and how to
do it. 

Modern, autodidactic learning focuses on the *individual* and their
particular learning style and needs the most important of which is
highly organized, social, and searchable knowledge organized
semantically so that content consumers can use it in ways that work best
*for them.*

Ultimately, however, learning is simply about reading and researching,
writing and working out the meaning and importance of what was read, and
exercising the new knowledge in ways that reinforce it. None of that
requires a specific learning management system of any kind, just a
*really good* knowledge management solution catered to individuals.

## Hugo Has Become Over-Engineered Crap

Hugo has become exactly what you would expect when someone's pet blog
generator that they learned Go on becomes popular with a bunch of
engineers who don't actually create any knowledge content of value. It's
design is a cluster-fuck testament to over-engineering. 

This is *really* obvious when compared to the elegance of Pandoc,
VuePress, and Gatsby. Hugo does things in unnecessarily complicated
ways. Just don't believe their bullshit. There are *very* simpler ways
to do what they claim *has* to be done that way. This approach gives
Hugo the appearance of power but for *knowledge-driven* content it
actually can't touch the others. 

The creators clearly didn't have knowledge management as a priority
going into it and even though lots of cool stuff has been bolted onto it
to try and make up for that fact it continues to miss. This was obvious
from the start when I began working with it some three years ago. I
started *Hugonot* at the time and abandoned it. But I knew then just how
bad Hugo's fundamental design was and is. It's no surprise its project
creator is a narcissistic engineer who doesn't write much of anything
but code and enjoys attacking people in forums more than educating
anyone on any level. You know, someone with values completely the
opposite of mine.

## VuePress Zero-Configuration

I realized this several years ago when looking at VuePress and how
elegant its raw design is. Take away is horrible server-side rendering
to provide hydration and you can see the elegance of its
zero-configuration approach. I'm really glad I experienced VuePress

## Gatsby Data-Driven Sources

Gatsby might have *a lot* of flaws but when I heard there was an SSG
that was built on top of a data-driven design I was ecstatic. In fact,
at the time I was building a tool called Xstatic for just the same
reason. Hugo seemed to completely miss the data-driven boat and was
spinning along making yet another fancy blog making SSG. To this day
Hugo's design is hobbled and centered on its original blog focus.

## Pandoc Knowledge-First

People still don't think of Pandoc as an SSG at all, and that's a good
thing. It means that the priority of the Pandoc team was, is, and will
always be the *knowledge data* above all else. 

You see Pandoc's knowledge focus first and foremost in its ultra simple
template design. Only the most basic logic and iteration exists for a
reason. The focus has *always* been on organizing your data better
*first* and then providing a template for that, not cramming an entire
language into your templates and giving them power to even dynamically
make database and other queries. Doing such things from the template is
just plain stupid.

## Hugo's Silly Templates

I was tempted for about five minutes to add some compatibility to `kn`
with Hugo templates until I looked at the cluster-fucked design they
have become. I could not be happier to be free from that project. The
design decisions over the last three years have been completely abysmal
and ridiculously stupid. It is so fucking obvious the people running
that project have never *actually* worked with any *critical* knowledge
data that they have to maintain, manage, and share themselves. Nothing
worse than a project run by people who don't even use their code for
anything substantial. Yet another silly Hugo site doesn't prove jack
shit.

## Bigger Vision

All the SSGs fail to look at the *real* problem, to back away from the
technology and take in the overall user stories and use cases that
involve knowledge creation, management, searching, and sharing. Every
single one of them is 100% focused on generating *web* content and not
on creating sustainable *knowledge* solutions instead. It's even in the
name, static *site* generator.

What I'm building is much larger in solution scope and much smaller in
complexity. It focuses on *knowledge* not sites, blogs, or even content
alone. Sure it is taking me forever to make, but looking briefly at
Hugo's internals again makes be convinced it is desperately needed.
