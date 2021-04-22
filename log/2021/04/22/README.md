## Thursday, April 22, 2021, 2:33:03PM EDT <1619116383>

TIL that you should almost never use a Pod during development and use a
Deployment instead. This is because you are going to change things up
and you cannot do that directly with a Pod. I wish I had heard about
this a *long* time ago.

## Thursday, April 22, 2021, 1:49:07PM EDT <1619113747>

Alacritty finally crapped out. It's just not ready for anything serious
yet. I can't even explain all the subtle things wrong with it. It simply
stopped responding. It also does not have full ANSI terminal escape
support, which I think is a big part of the problem. 

One specific thing is how it gets all fucked up when in full screen and
a Windows notification fires off in the upper right. For some reason it
completely resets the position of the content in the terminal. It's bad
enough getting annoyed by the notification. It's worst when it blows
away your position and cursor location. That bug alone is worth dropping
Alacritty for Windows (and therefore anything, since why do it if it
doesn't work anywhere). 

Thankfully, the Windows Terminal is much better. I threw it out because
I was so pissed about how hard it is to manage the configuration file.
But I'm over that now. So I spent the time to configure my Windows
Terminal and it was worth every second. It even does my Ubuntu Mono Nerd
Font without a problem.

I now prefer to do all my work from a laptop again, which is bad for my
stream because I'm generally not sitting at my desk much anymore. Even
though that is nice for a while, there's something about lounging in a
supine position that provides for more calm and less physical
distraction. Whatever the reason, I've always coded better from a chair
or bed. I think a lot has to do with keeping my legs warm and core
*disengaged* so that the extra body stress is out of the way. I also
find myself getting into a zen state where I control my breathing (and
sometimes forget to breath).

In short, sitting in *any* chair or standing is bad for getting into
flow state. There are too many physical distractions. The other night I
coded from bed for five hours straight without even the slightest
fatigue.

## Thursday, April 22, 2021, 10:37:07AM EDT <1619102227>

Wanna job in cloud-native? Learn go!

Another reason learning Go is really mandatory for any cloud-native
engineer is because documentation of important tools such as `kind`
points *directly* to the Go source code to document structures. This is
even more significant than learning Go `text/template` (which is used in
several essential tools --- particularly that involve filtered output).

## Thursday, April 22, 2021, 10:29:00AM EDT <1619101740>

I need to add `/todo/<isosec>` to my KEG content list.

