## Why `lynx` and not `w3m`?

Because `lynx` is simply superior:

* Better `vi` bindings
* Faster since not downloading and positioning images (just stupid)
* Older and better supported
* Has its own mailing list

But the *single* biggest advantage of Lynx is that it *keeps a
consistent text flow on the screen*. I cannot overstate how important
this subtle difference is. You won't notice it until you have logged
thousands of hours using it. But once you have you will find your eyes
jumping and scanning much more efficiently over web pages in text form
from Lynx than you ever could in `w3m`. This scanning ability is the
*entire reason* for using a text based browser in the first place. I can
search for full documentation on anything that an "intellisense" pop-up
would tease you with in less time than it takes for the fucking pop-up
to stop annoying you in your dumb-ass intellisense-enabled IDE. Very few
people on planet Earth (including most inexperienced YouTubers with big
mouths) will appreciate that difference, but I promise you, those that
do *destroy* the research efficiency of the others.

## Do you have a Discord? How can I message you?

[Yes.](https://discord.com/invite/9wydZXY) And you can just DM me there.

## How are you streaming and chatting like that?

It's a combination of things:

* WeeChat IRC client connected to Twitch in a nested TMUX pane so it
  stays even when I create a new window.

* Streaming <https://restream.io> with the bot enabled to messages
  appear to and from Twitch, YouTube, Periscope/Twitter, Discord and
  eventually LinkedIn (I hope). 

* OBS Studio (of course) which installed and worked without any problems
  or driver tweaking whatsoever on my PopOS distro.

* The
  [screenkey](https://gitlab.com/rwxrob/dotfiles/-/blob/master/scripts/screenkey)
  tool positioned down in the corner displaying what I type.

:::co-rant
I had to *write an email* to have my Twitch Affiliate status removed in
order to be compliant with their shitty terms of service that says you
have to wait 24 hours to put content on anything else. Then my subs
couldn't even unsub because the removed the button. Fuck that shit.
Twitch can bite me. They are *lucky* to have my content. Fuckers. It's
clear where their priorities are: selling out for the quickest buck, not
improving the world through education and engineering streams. Yeah,
I've come to really fucking hate Twitch (as a company) even though I've
met a lot of people through it.
:::

* Restream.io (currently,
although I am working on removing dependency on that bot because it is
just ugly).

## How do you get that large figlet font on terminal?

All the fonts that I use are in a [separate
repo](https://gitlab.com/rwxrob/fonts). The figlet fonts I've collected
over the years are all under `figlet`. I just combine the `figlet` and
`lolcat` command to get the result you see in these scripts:

* [figl](https://gitlab.com/rwxrob/dotfiles/-/blob/master/scripts/figl)
* [back](https://gitlab.com/rwxrob/dotfiles/-/blob/master/scripts/back)

The main figlet font I use is `future`.
