## Sunday, February 14, 2021, 8:30:53PM EST <1613352653>

All this hacking CTF watching has me realizing how much this entire
industry needs really fast knowledge exchange of the type designed to be
in KEG. I have to get a CTF person involved to provide input into the
entire interface to make `kn` the best possible tool for the secops
knowledge working. It is such a core part of reporting and composing
pen testing efforts of all kinds.

## Sunday, February 14, 2021, 8:23:53PM EST <1613352233>

Need to create a method using Go reflection to get the documentation
provided for a high-level function and insert it derive it during
`init()` to populate subcommand documentation in `cmdtab`.

I suppose I could just create a `cmdtab` subcommand that mines the
content of a function and generates a stub subcommand that just calls it
(once `cmdtab` is not just a package but also a utility command). That
makes the most sense.

## Sunday, February 14, 2021, 8:19:10PM EST <1613351950>

I cannot help but think how irrelevant all this hacking competition
would be if shit like PHP and JavaScript cross-site attacks and SQL
injection didn't exist. That's almost everything. We fucking failed. We
should never have let that shit software off the ground. The entire
world has paid the price. People have died from hospitals being
hit with ransomware for having shitty pen defense and crappy email that
runs attachments. It's infuriating because when something as simple as a
cookie came out, the entire tech community overwhelmingly came out
against it, but the assholes that made it happen didn't care. Now people
are literally dying because of their fucking greed.

## Sunday, February 14, 2021, 7:59:24PM EST <1613350764>

After a bit of watching I can feel the call to rewrite Burp and make it
better using an open source alternative that leverages the running
browser without using it, or better yet, includes something like the
Chromium engine as a proxy when needed (if needed).

## Sunday, February 14, 2021, 7:42:42PM EST <1613349762>

I'm so glad I stayed on a bit longer. I found something called "Hacking
Esports" and am remembering that there is an entire *massive* community
of people *not* like those I described earlier. I really am a hacker in
at heart, even if I cannot stand learning about bad shit even to hack
it. I will eventually be all about finding zero days though. I have to
make my own Web replacement first. But soon.

## Sunday, February 14, 2021, 7:01:05PM EST <1613347265>

Chuckling at how fucking stupid exclusive Twitch group memberships are.

But worse, are how fucking stupid people look when they obviously spend
more time on their online persona, branding, and cliche usernames than
making even the most minimally significant contribution to things that
matter most in our world --- especially if you *claim* to be an expert
in. 

Doesn't matter. 

I'm sure many of those people can't imagine spending a single second
complaining about what I'm complaining about (mostly because it would
take them five times as long to write the words contained in the
complaint). 

Yeah, I'm an asshole.

People are financially incented to play the exclusive thing, plus it
makes them feel included. 

I don't care. As much as I enjoy the energy and fun of the Twitch coding
streams, sometimes I find them so *fucking* juvenile I have to look
away.

I'd rather watch a brilliant digital nomad hiking through a frozen
forest with her stream community.

I feel my get-off-my-lawn guy coming out. Time to watch some Chuck and
go to sleep.

Obviously, I never say anything about any of this to anyone. I just
share it in this  super secret log, using, um, prose. They'll never see
it. Most of them don't read. It's obviously too boring for them. 

But, I'm sure someone out there remembers what prose is, and might even
appreciate this dark prose saturated with masturbatory, confessional
self-indulgence. 

I'm in good company. I'm quite sure this is what Rob Pike and Linus
Torvaldz catch themselves doing in private. I mean, Pike called syntax
color highlighting "juvenile"  and Torvaldz created a "code of
conflict."

Yeah, I'm just fine. 

Plus there a ton more amazing human beings who completely agree:

> Those who speak do not know. Those who know do not speak. (Lao Tzu)

Here's my favorite:

> The level of technical content quality is inversely proportional to
> the prolificness of the author. This is why it is so difficult and
> valuable to uncover those few who actually know what they are talking
> about --- because they do it every day for hours a day --- and who are
> willing and able to write about it. (TJ Halowaychuk)

I aspire to leave this world with a fraction of the contributions these
people have produced. Spoiler alert: it doesn't happen on Twitch.

(That reminds me, maybe I should actually read some of our Discord
messages. I never do. Nah, Imma get the docs and structure for AFK.Works
up instead.)

## Sunday, February 14, 2021, 2:50:52PM EST <1613332252>

I hate that Jekyll (GitHub Pages) decided to insert an `_config.yml`
file without getting confirmation. This reminds me I have to add `doc`
to the semi-reserved list of knowledge nodes.

## Sunday, February 14, 2021, 11:08:57AM EST <1613318937>

If KEG works and people start saving knowledge in Git repos on GitHub,
then the Arctic Project will preserve it for all time. Think about that
for a bit.

## Sunday, February 14, 2021, 10:26:48AM EST <1613316408>

Having smaller files organized in directories is really reminding me of
the Scrivener vibe from the days when I used it for book writing. This
smallest possible knowledge node concept is solid. It's the basis for
the entire
[Zettelkasten](https://duck.com/lite?kd=-1&kp=-1&q=Zettelkasten) that I
only heard of in February of 2020. I really love that so many great
thinkers of burned the midnight oil on this same question and just come
up with similar solutions in different languages.

## Sunday, February 14, 2021, 10:16:59AM EST <1613315819>

So very happy with moving back to pure Markdown files for KEG nodes.
Requiring that the first line be a hashtag makes creating generators
such a since using minimal scripting.

Before when there was a YAML header there it was anything but simple and
intuitive. Best decision in the last three months has been going with
the `data.*` files. It also brought in the entire
File-System-As-A-Database (FADB) stuff that I'd practically written off
from 2015. It's all there and combined into KEG. In fact, so many one
off ideas are seemingly coming together on this, stuff that dates back
even to my IBM days when I used the notion of a node to send knowledge
packages across the wire.

## Sunday, February 14, 2021, 8:34:01AM EST <1613309641>

Sometimes I fucking hate how verbose JSON is. I get that it makes things
flexible, but my God is it wasteful (but so much better than XML).

## Sunday, February 14, 2021, 6:49:33AM EST <1613303373>

It occurred to me that there is no strict limitation requiring root
nodes to be rooted at all. By shifting the focus away from a root node
onto special node types there can be multiple root entry points into a
node tree that might be maintained as a single Git repo.

This has tremendous implications for how knowledge node trees can be
maintained and migrated and address important problems with topical
guides and indexes.

## Sunday, February 14, 2021, 12:04:48AM EST <1613279088>

Important changes to KEG root knowledge nodes:

* DEX (formerly MANIFEST) with one line YAML objects but as a node
  itself (DEX/data.yml)
* KEG (formerly a file) with LAST containing timestamp and KEG version
* META and WORDS considered tool helpers and kept outside node
* Nothing not related to knowledge kept in the node
