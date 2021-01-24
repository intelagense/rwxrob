## Sunday, January 24, 2021, 8:45:53AM EST <1611495953>

Just added this repo as well with log and everything. The cool thing
about that is that I have a portable version of my KEG knowledge base
anywhere I want to go. Say I had to grab exactly one thing with me
before leaving the room. The USB stick would be in, constantly synced
with my most valuable personal content, with or without an Internet
connection. Doing the same thing for dotfiles as well.

Essentially, it's like having two core repos, one private and one
public, and making sure you have them with you wherever you may want to
go.

## Sunday, January 24, 2021, 8:35:09AM EST <1611495309>

Love the new local Git repo strategy for dealing with my private repos.
I documented it in my Git notes but here it is anyway.

```
cd # to get home
mkdir priv
cd priv
git init
git init --bare /media/rwxrob/06CF-482C/priv.git
git remote add bak /media/rwxrob/06CF-482C/priv.git
git push --set-upstream bak master
```

