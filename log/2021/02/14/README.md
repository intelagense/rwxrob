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
