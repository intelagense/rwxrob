# Notes About PicoCTF 2019

## Let's Warm Up (50)

If I told you a word started with 0x70 in hexadecimal, what would it start with in ASCII?

`picoCTF{p}`

## Warmed Up (50)

What is 0x3D (base 16) in decimal (base 10).

`picoCTF{61}`

## 2Warm (50)

Can you convert the number 42 (base 10) to binary (base 2)?

`picoCTF{101010}`

## Bases (100)

What does this `bDNhcm5fdGgzX3IwcDM1` mean? I think it has something to do with bases.

`picoCTF{l3arn_th3_r0p35}`

## First Grep (100)

Can you find the flag in file? This would be really tedious to look through manually, something tells me there is a better way. You can also find the file in `/problems/first-grep_0_93be1631acf1a93b98cdcc3e7b9fdc52` on the shell server.

`picoCTF{grep_is_good_to_find_things_4b2451ea}`

## First Grep: Part II (200)

Can you find the flag in `/problems/first-grep--part-ii_4_ca16fbcd16c92f0cb1e376a6c188d58f/files` on the shell server? Remember to use grep.

```sh
grep pico /problems/first-grep--part-ii_4_ca16fbcd16c92f0cb1e376a6c188d58f/files/**/*
```

`picoCTF{grep_r_to_find_this_0e28f3ee}`

## whats-the-difference

Can you spot the difference? `kitters` `cattos`. They are also available at `/problems/whats-the-difference_0_00862749a2aeb45993f36cc9cf98a47a` on the shell server

```go
package main

// go run flag.go | xclip

import (
	"fmt"
	"io/ioutil"
)

func main() {
	cat, _ := ioutil.ReadFile("cattos.jpg")
	kit, _ := ioutil.ReadFile("kitters.jpg")
	for i := 0; i < len(cat); i++ {
		k := kit[i]
		c := cat[i]
		if k != c {
			fmt.Printf("%c", c)
		}
	}
	fmt.Println()
}
```

My god that solution is *so* much better than the other crappy ones in Python out there. Go is definitely going to be the overwhelming leader over Python for this sort of coding.

`picoCTF{th3yr3_a5_d1ff3r3nt_4s_bu773r_4nd_j311y_aslkjfdsalkfslkflkjdsfdszmz10548}`

## Resources (100)

We put together a bunch of resources to help you out on our website! If you go over there, you might even find a flag! <https://picoctf.com/resources>

`picoCTF{r3source_pag3_f1ag}`

## strings is (100)

Can you find the flag in file without running it? You can also find the file in `/problems/strings-it_1_7a67382a38fc00751a6b9b29b0872813` on the shell server.

`picoCTF{5tRIng5_1T_0690b2a5}`

## what's a netcat (100)

Using `netcat` (`nc`) is going to be pretty important. Can you connect to `2019shell1.picoctf.com` at port 4158 to get the flag?
