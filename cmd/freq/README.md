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
% ps -eo user | freq | head
361     calvin
158     root
13      _accessoryupdater
8       _cmiodalassistants
7       _locationd
7       _softwareupdate
7       _rmd
5       _coreaudiod
5       _spotlight
4       _nsurlsessiond
```

```text
$ cat /bin/ls | freq -by byte | head
138812  00
1846    ff
1625    01
1407    03
1157    48
929     5f
899     74
883     20
757     40
720     65
```

```text
$ cat /usr/share/locale/zh_CN/LC_TIME | freq -by rune | head
58      '\n'
25      '月'
21      '%'
16      ' '
7       '星'
7       '期'
6       '一'
6       '二'
6       '十'
5       '1'
```