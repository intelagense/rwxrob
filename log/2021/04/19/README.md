## Monday, April 19, 2021, 1:46:30PM EDT <1618854390>

Just a reminder that the best documentation I can find on the format for
the comments in code in Go is still <https://blog.golang.org/godoc> from
2011. I'm quite sure something has changed since then.

## Monday, April 19, 2021, 12:30:07PM EDT <1618849807>

I have to stop forgetting to omit the `*` from both the `Stringer` and
`MarshalJSON` methods. It keeps making them miss their match as if they
didn't exit.

## Monday, April 19, 2021, 12:04:53PM EDT <1618848293>

That's it. I'm using `x` for my receivers in everything. This variation
based on the name is just plain fucking stupid and impossible to
remember. Using `x` for all my receivers makes everything consistent and
easy to find. Go might not like `this` or `self` but frankly not
standardizing the name of the receiver, not even by convention, is just
fucking brain dead. I don't have to comply.

