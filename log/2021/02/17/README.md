## Wednesday, February 17, 2021, 7:42:39AM EST <1613565759>

Combining source code and user documentation (that is not about how to
use the API) has always been a dubious decision. But, I realize even
more now that forcing people to follow a KEG root node --- even if the
content is relevant with the source code in the same Git repo --- will
most likely get you added to people's ignored lists.

I'm realizing that a fundamental design decision when deciding how to
organize your KEG content is simply ask the question. "Would I want to
download a zip of all of this?" Normally, you will realize that a bunch
of stuff doesn't belong, including source code that isn't part of the
knowledge itself (or it's generators).

