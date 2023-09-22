# stir

Shuffle lines from standard input.

## Install

```text
go install github.com/clfs/m/cmd/stir@latest
```

## Uninstall

```bash
rm -i $(which stir)
```

## Usage

```
$ stir -h             
Usage of stir:
  -seed int
        seed for random number generator
```

## Examples

Shuffle lines:

```text
$ seq 5 | stir -seed 6
1
2
5
3
4
```

The default seed is 0:

```text
$ seq 5 | stir        
5
3
4
1
2
```
