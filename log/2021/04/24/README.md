## Saturday, April 24, 2021, 10:45:18AM EDT <1619275518>

Been thinking that the new [rwx.gg] should actually just be a big
Zettelkasten, but I'm afraid of stuff not being baked again. The main
reason is the value of isosec unique identifiers.

It's a toss up between that and putting it into my own main Zettelkasten
and moving it over. 

Then there is the question of how baked to make my own personal ZK site.
Right now it is just a GitHub repo, which is fine. I don't use
`rwxrob.live` for anything significant these days. Also, there's
something to be said for using a `github.io` domain because they are
free and won't fade away after I pass on (which hopefully isn't for a
while, but still). The [rwx.gg] site, on the other hand, would either
need to be maintained by someone else to keep it up to date, or it could
just fade away and no one would mind because all the content might have
been covered elsewhere by then.

* `zet/<isosec>` - rough zettels, including text and video forms
* `lab/<isosec>` - learning labs (formerly "codebook")
* `data/<ident>` - `yq` query-able data in `data.yml`
* `def/<term>` - term definitions, encyclopedia like 
* `cal` - calendar of events in YAML, gen as `README.md/index.html`
* `week` - weekly schedule
* `log/<isosec>` - time-based log entries (not topics, that's zet)
* `work/<isosec>` - work to be completed usually by a certain time
* `pub/<ident>` - publications, articles, blogs, pdfs, videos
* `reg` - main ZK register, always created by hand
* `dex/<word>` - word index (generated)

After repeating that (I've written it before) it occurs to me that
nothing I am doing for [rwx.gg] cannot be contained in one of these
areas.

The KEG types concept is coming back as well. This stuff is organized by
type more than category or content:

* Zettel - a single zettel (slip), both "fleeting" and main
* Lab - exploration and learning content with some structure
* Data - structured data in YAML/JSON for any purpose
* Definition - definition of a term or concept
* Calendar - calendar of upcoming and past events in YAML
* Schedule - weekly general schedule in YAML and generated
* Work - work to be accomplished usually by given time
* Publication - polished content suitable for publishing
* Register - Zettelkasten register, labels and isosec refs
* Index - auto-generated index of words occurring throughout

Yep, that covers everything. I do have a few other KEG types that I've
used but wouldn't need:

* CV - curriculum vitae (resume) in YAML and generated
* Endorsement - professional endorsement (under CV usually)

I think the key here (like they learned with other semantic web efforts)
is to allow people to expand and define their own types. I've already
covered this in other areas by providing YAML/JSON for the JSON-Schema
of the type definition (which is so widely supported now its
ridiculous despite how bad it is).

[rwx.gg]: <https://rwx.gg>
