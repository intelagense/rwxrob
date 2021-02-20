## Saturday, February 20, 2021, 9:43:21AM EST <1613832201>

Realized that I have the beginnings of a system-independent, UNIX-like
portable environment contained within a single executable. I was noting
the value of having even simple actions like `date +%s` within `kn`
because it makes them available if I push `kn` to another system even if
it is Windows with no UNIX/Linux shell in sight. Combining a few other
things into it would result in a *single* executable that has all the
wicked fast tools I use to do *everything* making my workspace 100%
portable to any computer, period.

I must be made.

## Saturday, February 20, 2021, 9:32:16AM EST <1613831536>

Here's why `kn tstamp` matters:

```json
{
  "tstamp": "2021-02-20T09:38:03-05:00"
}
```
