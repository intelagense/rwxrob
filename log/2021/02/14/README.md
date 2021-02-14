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
