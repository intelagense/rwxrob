## Tuesday, February 16, 2021, 12:42:49AM EST <1613454169>

I really don't know why it took me so long to turn back on that window
titles and status bar in tmux now that I've switch back to windows
instead of panes. So glad I did though. It will really help everyone
watching as well because they'll be able to see what is actually
running. 

By the way, that is yet another reason to use `exec` in scripts as the
last thing you do because it updates to the correct name in the tmux
status bar rather than just putting `sh` for everything.

By the same token, a compiled Go program will *always* be better than
any script because the binary of a script is always just the name of the
interpreter and not the actual name of the code that it is running. 

I especially love that people will be able to see more and more `perl`
appearing in the status bar as well.
