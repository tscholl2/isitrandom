# Test randomness

```
go run shuffle.go && ent test.dat
```

Compare to

```
head -c 6 /dev/urandom > test.dat && ent test.dat
```

# Todo

- [ ] Determine empirical pattern for riffle shuffle (position of split, size of each alternating packet)
- [ ] Design `shuffle(n int64)` function to emulate empirical data
- [ ] Test various amounts of shuffling to see how many it takes to past tests

# Notes

- [Gilbert-Shannon-Reeds model](https://en.wikipedia.org/wiki/Gilbert–Shannon–Reeds_model) of shuffling
- [Analysis of real shuffling](http://jdc.math.uwo.ca/M9140a-2014-summer/Diaconis-1988.pdf) on page 86, uses frequencies in bridge hands as proxy.
- [7 shuffles is optimal](http://projecteuclid.org/download/pdf_1/euclid.aoap/1177005705)
