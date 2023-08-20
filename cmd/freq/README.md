# freq

Compute frequency distributions from standard input.

## Install

```text
go install github.com/clfs/m/cmd/freq@latest
```

## Uninstall

```bash
rm -i $(which freq)
```

## Usage

```text
$ freq -h
Usage of freq:
  -by string
        line, byte, rune, or word (default "line")
```

## Examples

```text
$ ps -eo user | freq | head
356     calvin
158     root
13      _accessoryupdater
8       _cmiodalassistants
7       _softwareupdate
7       _rmd
7       _locationd
5       _spotlight
5       _coreaudiod
4       _nsurlsessiond
```

```text
$ cat /bin/ls | freq -by byte | head
138812  0x00
1846    0xff
1625    0x01
1407    0x03
1157    0x48
929     0x5f
899     0x74
883     0x20
757     0x40
720     0x65
```

```text
$ cat /usr/share/locale/zh_CN/LC_TIME | freq -by rune | head
58      "\n"
25      "月"
21      "%"
16      " "
7       "期"
7       "星"
6       "十"
6       "二"
6       "一"
5       "1"
```

```text
$ man tar | freq -by word | head
299     the
128     and
113     to
113     is
81      or
76      tar
76      of
65      archive
60      in
60      be
```