# Notes About KEG, the Knowledge Exchange Grid

## Random Notes

* `gen`

## Organization Needed

Need to organize all the content and code. Here's some initial thoughts
and questions:

* Do we host our own and mirror to GitHub?
* What are the top-level domains and where do they go?
* Do we go mono-repo or submodules?
* Need a root knowledge node for notes from meetings, etc.
* Where does the specification go?
* What about the potential for a KEG/ node? (KEG/LAST)
* Should the DEX node be local or KEG only? Optional? Local and KEG.
* Rethinking DEX as one-liner YAML data for humans? Yes with redirects.
* Do the WORDS and META files even belong? This is a standard that tools
  agree on the consumer side? In or out of the actual node itself?
* Where to we create searchable cache data? On the consumer side, duh.

## Domains

Domain|Use
|:-:|-
afk.works|Association for Federated Knowledge Work
keg.network|Main KEG node registry (like Go docs)
keg.sh|Query the KEG registry directly from the shell
know.sh|Unknown
pegn.dev|Parsing Expression Grammar Notion
readme.world|Unknown
rwx.gg|KEG Node for RWX learning content
skilz.dev|Unknown
mim.sh|Unknown
mim.network|Unknown
skilstak.sh|SkilStak remote shell login
skilstak.io|SkilStak main site (Easter egg learning game)

Hosting of Specifications:

```

```

Potential Import Dependencies:

```go 
import "code.afk.works/keg"
```

Feels like a single reference implementation in Go is the best way to
keep API development rapid and simple.

Potential Apps:

All the apps will use `keg` API but are separate for all the right
separation of concerns reasons, the biggest being that `kn` is for
working with local knowledge nodes, and `keg` is for the larger
Knowledge Exchange Grid.

```
go get -u apps.afk.works/keg
go get -u apps.afk.works/kn
```
