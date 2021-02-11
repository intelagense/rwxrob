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

