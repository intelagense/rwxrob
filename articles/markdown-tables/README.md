---
Title: GFM Markdown Tables FTW!
Subtitle: Fastest, Safest, Most Supported
Created: 2020-07-06T19:02:42-0400
---

I've come to realize that even my beloved Pandoc tables are just not
enough to keep me using Pandoc as a primary rendering engine over
Goldmark's builtin support for GitHub-Flavored-Markdown. 

|Column One|Column Two|Column Three|
|:-|-|-:|
|Here is|some|thing|
|And|another|thing|  

```markdown
|Column One|Column Two|Column Three|
|:-|-|-:|
|Here is|some|thing|
|And|another|thing|  
```

Here's the pandoc version:

```
Column One      Column Two    Column Three
------------   ------------  -------------
Here is             some       thing 
And              another       thing
```

I know. I know. I've railed on GFM for a *very* long time and there is a
lot of crap in it, but if using the above simplified table format is
required over Pandoc's better table format so be it.

## Better for All Fonts

When I read that there are actually a decent enough number of people
writing Markdown using non-proportional fonts it really made me stop and
realize that making your tables depend on a mono-spaced font is actually
a pretty bad idea. That killed my love for Pandoc tables more than the
rest. In fact, it made GFM tables look pretty damn intelligent.

GFM tables require a bar to separate the columns. There is nothing
needed with the whitespace at all. This might make them look rather ugly
when in raw text but who gives a shit. They are *way* faster to write
than all the other Pandoc tables. I seriously was not expecting to
discover that when I gave GFM another chance.

## Faster to Write

To me writing speed trumps source appearance by *a lot*. Because GFM
tables don't depend on lining anything up *at all* you can type them way
faster than anything else. Sure the source will look a bit mangled but
you don't really need to care. It is one of those trade-offs from the
"source should look good as render" principle from original Markdown.
You know, the same silliness that gave us literally an infinite number
of ways to express a fucking horizontal rule.

## Things I've Come to Hate About Pandoc Tables

**I'm always messing up the alignment.** They might look pretty but damn
they are hard to line up and *maintain*. If not exactly under the right
`-` you lose all your alignment. GFM tables have an explicit `:` to
visual show where the alignment is and it has *nothing* to do with the
alignment. You can just add that if you feel like futzing around with
how the source text should look. Frankly I usually don't care that much
about that.

**Long table cell content actually sucks.** I thought I would love
having long table cells full of paragraphs and all kinds of multi-line
goodness, but keeping all that shit wrapped properly and lined up is a
fucking nightmare, if with all the greatest settings. Besides, tables
are *not* for long content. Every time you need a cell with more than
two lines in it you are probably in need of another way to express that
information.
