## Saturday, April 3, 2021, 12:38:43PM EDT <1617467923>

TIL about `irc_raw_in_*` within `weechat` to help with parsing and
responding to messages.

> In python and those functions have to exist in the base file  so
what i did was write `api.py` which exports a callback function, and you just
import that callback + call a register function then use API's own
functions to register your own events.
