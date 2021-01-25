## Monday, January 25, 2021, 8:26:55AM EST <1611581215>

One of the main cleanup activities I must do is change all my scripts
library to have zero insignificant dependencies on anything else, such
as for dumb shit like colors and usage output. Makes is much harder to
share them. 

This means I need to create tools for adding these useful things more
easily in an independent way, like snippets.

## Monday, January 25, 2021, 8:01:26AM EST <1611579686>

By the way, I forgot to mention that Ez-Lynx is still on (a simple shell
script that configures an installation of Lynx for a specific user). I
just cannot give up how much easier Lynx is to read than W3M because of
indentation of paragraphs and such that seems to be hard coded.

## Monday, January 25, 2021, 7:51:23AM EST <1611579083>

Having that thing where I have so much to get done that I don't know
where to start. Most important thing I have go get done is a full online
profile cleanup so I look respectable to SMEs vetting my job
applications next month.

I'm on it.

## Monday, January 25, 2021, 6:38:14AM EST <1611574694>

Been meaning to set a key binding for `:set paste` for years. Using
`:map` shows all the Vim keybindings, which is a start. I've tried `set
paste` in my `.vimrc` and learned just what a bad idea that is because
it turns off all mappings while in insert mode (which usually isn't an
issue but I've become partial to `zz` and `jj` and the like (in my old
age) for replacing the pinky-stomping alternatives to escape. (Don't
worry I'm still plenty able to use `Cntl-[` or `ESC` when the situation
requires.

Looks like I'll be trying out Tim Pope's [unimparred
plugin](https://github.com/tpope/vim-unimpaired) for a while.

## Monday, January 25, 2021, 6:36:00AM EST <1611574560>

Adding the concept of shortcut to my KEG node schema for bookmarks.

```yaml
- T: Gordon's Art Work
  L: https://www.deviantart.com/zephyrwork
  B: gordon
```

## Monday, January 25, 2021, 6:27:17AM EST <1611574037>

Just discovered [surfraw](https://surfraw.org) and I'm already sold on
their introduction line:

> Surfraw liberateur is capable of navigating speeds that leave GUI
> tainted idolaters agape with fear and wonder.

In fact, the name is apparently an acronym for *Shell Users
Revolutionary Front Rage Against the Web*. 

Vive la revolution!

Now that's my kinda people.

Looks like it is my `?` and `??` and `???` to the nth degree and
includes pretty much everything most terminal users would ever care to
search these days, but it's an interactive shell rather than something
that will get saved in my shell history.

It's also hosted on GitLab for all the right reasons (which
unfortunately are overwhelmingly rendered [irrelevant](https://www.youtube.com/c/rwxrob/search?query=github) to most modern
developers.

## Monday, January 25, 2021, 5:10:30AM EST <1611569430>

Through random experimentation we have concluded a few things about Lynx
and w3m:

* W3M is absolutely the best *default* experience.
* Lynx requires *tons* of customization to be usable .
* W3M's annoying lack of indentation is unreadable.
* Thou shalt lie about user-agent in either one.
* W3M has way better information (`=`).
* W3M has way better internationalization support.
* Browsh is out of the question since requires Firefox.
* W3M loads CSS if it can be found (ex: table border).
* There are a ton of `-o` options with W3M still to consider.

By the way, here's a user-agent that works:

```
"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_0) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.79 Safari/537.1 Lynx"
```

Make sure you put that all on a single line.

## Monday, January 25, 2021, 3:45:47AM EST <1611564347>

Look like `w3m` might actually *finally* be better than Lynx. People
have claimed it was better for years, but the last time I compared it
was no contest. Up until recently my reasons for picking Lynx had more
do to with bad decisions from the `w3m` team:

* It tried to hard to render graphics
* It was just slower for text browsing
* It wanted to be a graphics browser, still

But it turns out that many of advantages of `w3m` are overtaking Lynx:

* Better support for HTML5 elements
* Zero tweaking out of the box
* Vim bindings by default
* Modern concurrency
* Just plain faster
* Image loading no longer automatic

