## Wednesday, February 10, 2021, 5:36:39PM EST <1612996599>

One of the most annoying things about the otherwise amazing `gh` command
line tool from GitHub is the lack of tab completion for the aliases I
have dutifully added. Had I added them using my shell script approach
they would be automatically detected. The pile of reasons to create a
`repo` script (probably Perl) that encapsulates `gh` and `gl` and
something for other local bare Git repos continues to grow.

## Wednesday, February 10, 2021, 5:33:30PM EST <1612996410>

More reasons to hate on GitLab:

* I have to be logged in to search it
* I have to wait through the browser Cloudflare confirmation before I
  can even open anything.

It is so sad to see such a great product fail to completely.

## Wednesday, February 10, 2021, 3:05:14PM EST <1612987514>

TIL that `readelf` is a thing that I did not know about:

```sh
readelf --all $(which ls)
```

## Wednesday, February 10, 2021, 1:07:08PM EST <1612980428>

I'm done with `exa` which doesn't even show how many hard links there
are. How fucking brain dead can an `ls` be? Fuck the colors and emojis
if it doesn't have the *base* functionality of the `ls` command.

