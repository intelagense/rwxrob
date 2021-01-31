# Perl Notes

Yes, Perl is still a *very* valid shell scripting language ---
particularly when dealing with lots of text. Did you know *no* other
language supports `\p{Cc}` Unicode regular expressions and that every
other language on planet Earth uses PCRE because Perl set the standard
for regular expressions? So, um, yeah.

## Bash Tab Completion in Perl

```perl
if ($ENV{'COMP_LINE'}) {
  my $partial = $ARGV[1];
  my [\@actions](https://twitch.tv/actions) = map {s,^.*/,, and $_} glob "kn.d/*";
  map {/^$partial/ and say $_} [\@actions](https://twitch.tv/actions);
  exit 0;
}
```
