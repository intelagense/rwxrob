# Perl Notes

[The Universe was Actually Hacked Together with
Perl](https://xkcd.com/224/)

Yes, Perl is still a *very* valid shell scripting language ---
particularly when dealing with lots of text.

## Unicode Regular Expression Support from 5.14

Did you know *no* other language supports `\p{Cc}` Unicode regular
expressions and that every other language on planet Earth uses PCRE
because Perl set the standard for regular expressions? So, um, yeah.

## Safe Path Handling

```perl
my $abs = File::Spec->abs_path($0)
my $new = File::Spec->join($some,"README.md")
```

## Bash Tab Completion in Perl

```perl
if ($ENV{'COMP_LINE'}) {
  my $partial = $ARGV[1];
  my [\@actions](https://twitch.tv/actions) = map {s,^.*/,, and $_} glob "kn.d/*";
  map {/^$partial/ and say $_} [\@actions](https://twitch.tv/actions);
  exit 0;
}
```
