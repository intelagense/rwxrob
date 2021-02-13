# Git Notes

Recipes and such for Git source management system.

## Delete Tags

```sh
git push origin -d <tag>
```

## List All Remote Tags

```sh
git ls-remote --tags | perl -pe 's,.+/,,'
```

## Create a "Bare" Repo for Local Backups

Some repos have no business being on GitHub or GitLab, but using the
same workflow is nice. I prefer to not even keep this anywhere near my
`$REPOS` directory with everything else.

```
cd # to get home
mkdir priv
cd priv
git init
git init --bare /media/rwxrob/06CF-482C/priv.git
git remote add bak /media/rwxrob/06CF-482C/priv.git
git push --set-upstream bak master
```

And can clone it with the pull path.

```
git clone /media/rwxrob/06CF-482C/priv.git
```
