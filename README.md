# Pikchr for Goldmark

This is an extension for [Goldmark](https://github.com/yuin/goldmark) that adds
support for generating figures using [Pikchr](https://pikchr.org). It doesn't
rely on cgo or an external binary, instead using the
[wazero](https://wazero.io/) WebAssembly runtime to embed Pikchr.

## Usage

Register the extension:

``` go
goldmark.New(
	goldmark.WithExtensions(&pikchr.Extender{}),
).Convert(src, dst)
```

Fenced code blocks like below will now be rendered using Pikchr:

~~~markdown
```pikchr
arrow right 200% "Markdown" "Source"
box rad 10px "Markdown" "Formatter" "(markdown.c)" fit
arrow right 200% "HTML+SVG" "Output"
arrow <-> down 70% from last box.s
box same "Pikchr" "Formatter" "(pikchr.c)" fit
```
~~~

Resulting in a figure like below:

![](testdata/basic.svg)

## Performance

Performance will be the same as for the underlying
[go-pikchr](https://github.com/rigtorp/go-pikchr) package. For a small Pikchr
figure, go-pikchr takes about 7.5ms:

```shell
$ go test -test.bench .
goos: linux
goarch: amd64
pkg: github.com/rigtorp/go-pikchr
cpu: Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz
BenchmarkPikchr-8            142           7520289 ns/op
PASS
ok      github.com/rigtorp/go-pikchr    1.471s
```

## Acknowledgements

This extension is based on the
[goldmark-pikchr](https://github.com/jchenry/goldmark-pikchr) extension by
[Colin Henry](https://github.com/jchenry) and the
[goldmark-d2](https://github.com/FurqanSoftware/goldmark-d2) extension by
[Furqan Software](https://github.com/FurqanSoftware).
