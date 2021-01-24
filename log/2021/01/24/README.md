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

