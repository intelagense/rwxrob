## Friday, April 23, 2021, 7:12:54AM EDT <1619176374>

I have noticed the convention for shell scripting to make some
variables uppercase. I believe the reason for this is the universal
global nature of these variables and the lack of distinction from
aliases and commands, which are always lower case.

I also noticed that POSIX shell scripts have no problem calling
functions that will set variables within the scope of the caller and use
them. This means that such functions also initialize these variables
within the function to ensure their initial values are what are
expected.

The best example of really solid POSIX shell scripting that I have ever
encountered is the Docker install script. (I have a copy in my [dotfiles]
under `install`.) It is an absolute work of art and I'll be using it
regularly to show what amazing POSIX shell code should look like. In
fact, I plan on extracting a style guideline from it (unless I can find
a better one).

[dotfiles]: <https://github.com/rwxrob/dotfiles>

