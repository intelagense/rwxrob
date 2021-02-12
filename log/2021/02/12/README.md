## Friday, February 12, 2021, 4:46:09PM EST <1613166369>



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

