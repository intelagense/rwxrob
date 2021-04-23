## Friday, April 23, 2021, 4:31:02PM EDT <1619209862>

Open credentials == managing credentials as software.

## Friday, April 23, 2021, 3:34:15PM EDT <1619206455>

Toying with the idea of adding back the Open Credential thing as a part
of SKILSTAK this time. Basically, I'd be directly ripping off the method
that works for scouting with merit badges and rank advancement, but
making it more real and related to technology. Scouting, after all, has
always been about self-directed learning.

What would the terminology be?

* *rank* - names for different levels that aren't necessarily
         vertically linear
* *boost* - an accelerated review of what learning is needed 

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

