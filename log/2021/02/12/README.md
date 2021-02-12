## Friday, February 12, 2021, 5:38:14PM EST <1613169494>

Blows me away that the installation instructions for `gh` have the user
use `curl -o` instead of `curl -O` to get the name from the URL. Just
goes to show there is a lot *everyone* has to learn and the learning
never stops.

## Friday, February 12, 2021, 5:12:05PM EST <1613167925>

After hearing that some people recommend using Control-C instead of ESC
for Vim I cannot believe how clueless otherwise really informed people
are. They have no idea that doing this not only encourages using
Control-C for mode changes in other contexts (which will usually
abruptly stop the program from working without saving anything) but they
also have no idea they are bypassing some of the most important Vim
events that Plugins use. But the worst thing of all is completely
destroying their ability to use `set -o vi` on their shell history.

Here's a recap of all the ways people deal with requirement for ESC in
Vi/m:

|::|::||
-|-
👍| CapsLock   | Original ESC position. Requires OS remapping. Universal.
👍| Alt-[hjkl] | Same as ESC Key. Easy to type. Supported everywhere.
👍| Control-[  | Same as ESC Key. Easy to type. Not supported in Europe.
👍| ESC Key    | Default but annoying and bad on pinkies.
👎| ff/jk/kj   | Requires remapping *every* Vim config. Vim specific.
👎| Control-C  | Builds really dangerous habits. Hack for Vi/m only.

I've come to conclude that remapping Caps Lock is objectively the best
solution. It has the added advantage of ensuring you never accidentally
hit it during a Vi/m session and *really* mess yourself up (you know you
have done it at least once, admit it).

Here's some [facts about why Caps Lock became king](https://simplyian.com/2015/01/08/The-real-reason-why-Caps-Lock-and-Escape-are-in-terrible-positions/).

Fuck you, Caps Lock. You always sucked anyway. Bubye. Take your useless
ass off of my keyboard. You are my ESC
once more. 

## Friday, February 12, 2021, 4:41:07PM EST <1613166067>

I cannot say how grateful I am for the `gh` command line tool provided
by GitHub. I would abandon GitLab just because of it.

After being notified that SkilStak no longer qualifies for "educational"
status (according to the new Microsoft owned GitHub) --- and ignoring
their suggestion I pay \$32/month just for the privilege of having what
I have kept for *over seven fucking years* I now have all my repos
locally with the following command:

```sh
for i in $(gh repos skilstak); do gh repo clone $i; done
```

This command cloned down all 144 repos from `skilstak` in under 20
minutes allowing me to sleep better knowing they won't suddenly
disappear because MS suddenly changed their mind about what
"educational" actually means (spoiler alert: it's only those paying into
the bullshit status quo system).

God, I cannot *wait* to be rid of both GitHub and GitLab dependencies.
Even though it is slightly more work and will result in some minor
personal monthly costs, I'm setting up *everything* on Gitea instances
for each domain (and will mirror the content to GitHub for visibility).
I'm so fucking done with bait-n-switch bullshit and the people who are
destroyed because of their whims. It's free, so it's hard to complain,
but the great lie they don't tell you is that you never really mattered
at all.

