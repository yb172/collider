# Small local collider - example project to explore test coverage

In your go project you can get test coverage by running

```bash
go test ./... -coverprofile=cover.out
go tool cover -cover=cover.out
```

This will produce following output:

```text
github.com/yb172/coverage-lines/collider/starter.go:8:   Collide           100.0%
github.com/yb172/coverage-lines/collider/starter.go:18:  checkMass         100.0%
github.com/yb172/coverage-lines/loader/yaml.go:12:       LoadFromYaml      100.0%
total:                                                   (statements)      100.0%
```

If you have multiple projects and all of them are writtenn in Go you can also get total coverage by running the same commands on directory where all projects are contained.

However if you have some projects not in Go that won't work.

Instead we need to get total LOC and LOC under test for our Go and non-Go projects and covered/all would be our coverage.

## Straightforward approach

Straightforward approach would be to use LOC counter tool such as [gocloc](https://github.com/hhatto/gocloc). Let's run it on our example program and see what it will produce:

```bash
gocloc --by-file --not-match-d=vendor --include-lang=go .
```

It will produce following result:

```text
----------------------------------------------------------------------------
File                      files          blank        comment           code
----------------------------------------------------------------------------
collider/starter_test.go                     9              0             54
loader/yaml_test.go                          7              0             35
collider/starter.go                          4              1             24
loader/yaml.go                               3              1             18
run-collider.go                              3              0             18
collider/structs.go                          1              1              5
----------------------------------------------------------------------------
TOTAL                         6             27              3            154
----------------------------------------------------------------------------
```

First obvious problem is that `*_test.go` are also included. Let's calculate what is LOC count without tests:

```text
24 (starter.go) + 18 (yaml.go) + 5 (structs.go) = 47
```

I also excluded `run-collider.go` since it wasn't included in cover.out.

Now let's manually count what is the number of lines covered by tests. Html report for go coverage would help us:

```bash
go tool cover -html=cover.out
```

Results are (if count opening curly brace `{` also a line):

```text
18 (starter.go) + 10 (yaml.go) = 28
```

So according to our straightforward approach our test coverage is:

```text
28 / 47 = 59.6%
```

Not great. Why is it so? Because `gocloc` counts lines of code and `import` as well as `type Particle struct {` are also considered as lines of code.

So to correctly calculate coverage we need another approach.

## Use "coverage-lines"

That would be great if there would be a special flag in `go tool cover` that will output how many lines are covered and total lines. Unfortunately I haven't found such flag.

Good news is that all the information we need is already calculated in `go tool cover`: [github.com/golang/tools/blob/master/cmd/cover/func.go#L57](https://github.com/golang/tools/blob/5ebbcd132f1e9c320204496eeeaac47007339543/cmd/cover/func.go#L57). We just need to print `covered` and `total` instead of it's ratio. That is what is done in "[coverage-lines](https://github.com/yb172/coverage-lines)". Let's install and run it:

```bash
coverage-lines cover.out
```

Result is

```text
20 20
```

Interesting: there are actually 20 lines to test according to the profile, not 28 that I got by counting green lines. Apparently `{` as well as `}` are not counted as a line (which makes sense).
