# renew

Upgrade binaries installed by `go install`.

## Install

```text
go install github.com/clfs/m/cmd/renew@latest
```

## Uninstall

```bash
rm -i $(which renew)
```

## Usage

```text
$ renew -h
Usage of renew:
  -list
        list binaries
  -update string
        update a binary to @latest
```

## Examples

```text
$ renew -list
csvq               github.com/mithrandie/csvq
dlv                github.com/go-delve/delve/cmd/dlv
fzgen              github.com/thepudds/fzgen/cmd/fzgen
go-118-fuzz-build  github.com/AdamKorcz/go-118-fuzz-build
go-mod-upgrade     github.com/oligot/go-mod-upgrade
godoc              golang.org/x/tools/cmd/godoc
gokart             github.com/praetorian-inc/gokart
gomodifytags       github.com/fatih/gomodifytags
goplay             github.com/haya14busa/goplay/cmd/goplay
gopls              golang.org/x/tools/gopls
gosec              github.com/securego/gosec/v2/cmd/gosec
gotests            github.com/cweill/gotests/gotests
gotip              golang.org/dl/gotip
govulncheck        golang.org/x/vuln/cmd/govulncheck
impl               github.com/josharian/impl
protoc-gen-go      google.golang.org/protobuf/cmd/protoc-gen-go
protoc-gen-go-grpc google.golang.org/grpc/cmd/protoc-gen-go-grpc
redress            github.com/goretk/redress
staticcheck        honnef.co/go/tools/cmd/staticcheck
tfsec              github.com/aquasecurity/tfsec/cmd/tfsec
```

```text
$ renew -update staticcheck
go: downloading honnef.co/go/tools v0.4.5
go: downloading golang.org/x/tools v0.9.4-0.20230601214343-86c93e8732cc
```