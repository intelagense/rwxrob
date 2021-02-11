## Thursday, February 11, 2021, 4:11:46PM EST <1613077906>

Still experimenting with live streaming during "monitored mentoring
sessions" that only require me to monitor the screen next to me and
ensure they don't get stuck during their work. It requires that I mute
the live stream currently because I can't risk having someone ask for
help and having it come through the desktop audio. 

This morning I forgot that I was muted in OBS, so there's that to
remember.

Also, after going through a "stand by" screen need to remember to bring
up the live streamed screen and not the mentored one. It's not really
that hard, just something to remember.

Just noticing as well that my `screenkey` setup needs a different mode
for monitored mentoring since it shows up on the other screen and not
the one that is streamed.

## Thursday, February 11, 2021, 1:27:22PM EST <1613068042>

Fastest method to get the keys from a Golang Map:

```go
// Keys returns a sort list of keys from the Data
func (js Config) Keys() []string {
	keys := make([]string, len(js.Data))
	n := 0
	for k, _ := range js.Data {
		keys[n] = k
		n++
	}
	return keys
}
```

The `make` gets rid of the slower requirement for an `append`.

## Thursday, February 11, 2021, 8:53:59AM EST <1613051639>

Noticed that there is `path.Join()` and `filepath.Join()` and I don't
know the difference other than `filepath` is the one used in every
example I've seen in the last three years. 

I found some of my old code that uses `path` and changed it.

Thanks to @afrology from stream I learned that `path` observes slashes
while `filepath` does the expected UNIX thing and cleans them up along
with dots.

