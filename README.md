# sedappeal

Did you know that `sed` substitute command does not have to be in form of
`s/foo/bar/`? It can be `s,foo,bar,` or even `sXfooXbarX`. With that in mind...

```sh
echo cleavage | sed severe
# clearage

echo mała | sed sałatka
# matka
```

## Usage
This simple CLI, given a file with words, will find any pairs (w1, w2) where
`echo w1 | sed w2` yields another word.

Notes:
- w2 will always start with "s"
- while sed supports various flags, e.g. `s/foo/bar/g` for global the CLI
  assumes there are no trailing flags
