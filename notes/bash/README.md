# Random Discoveries and Tricks in Bash Shell

## Use `printf` Instead of Date

Turns out all of the [strftime
library](https://duck.com/lite?kd=-1&kp=-1&q=strftime+library) stuff is
available to modern (4.2+) Bash `printf`.

### Get Epoch Seconds

Using `printf -v foo '%(...)T'` is identical to `foo=$(date +'...')` and saves a fork for the call to the external program date.

```bash
printf '%(%s)T\n'
```

```{.out}
1611057884
```

### Convert and Adjust an ISO Timestamp

For this still have to have `date` to parse the date, but can use
`printf` to print the adjusted time.

```bash
#!/bin/bash

secs=$(date -d 2021-01-19T05:52:43-05:00 +%s)
echo $secs
((secs+=60))
printf "%(%s)T\n" $secs

# Output:
# 1611053563
# 1611053623

```

This could, of course, all be done in a single `date` line:

```
secs=$(date -d '2021-01-19T05:52:43-05:00 +1 minute' +%s)
echo $secs
```

