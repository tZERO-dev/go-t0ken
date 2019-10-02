# tZERO Go-T0ken Makefile
#

TIME    = $(shell date '+%A %W %Y %X')
COMMIT  = $(shell git rev-parse --short=10 HEAD)
TARGETS = "darwin/amd64,linux/amd64,windows/amd64"

LDFLAG_COMMIT_EXPERIMENTAL = -X "github.com/tzero-dev/go-t0ken/commands.GitCommit=$(COMMIT) - experimental"
LDFLAG_COMMIT              = -X "github.com/tzero-dev/go-t0ken/commands.GitCommit=$(COMMIT)"
LDFLAG_TIME                = -X "github.com/tzero-dev/go-t0ken/commands.BuildTime=$(TIME)"
LDFLAGS_EXPERIMENTAL = '$(LDFLAG_COMMIT_EXPERIMENTAL) $(LDFLAG_TIME)'
LDFLAGS_DIST         = '$(LDFLAG_COMMIT) $(LDFLAG_TIME)'


.PHONY: generate t0ken t0kenbatch dist


print-%: ; @echo '$(subst ','\'',$*=$($*))'

all: generate t0ken t0kenbatch

generate:
	# Generating contracts...
	@go generate ./...

t0ken:
	# Building t0ken...
	@go build -v -ldflags=$(LDFLAGS_EXPERIMENTAL) -o bin/t0ken ./cmd/t0ken

t0kenbatch:
	# Building t0kenbatch...
	@go build -v -ldflags=$(LDFLAGS_EXPERIMENTAL) -o bin/t0kenbatch ./cmd/t0kenbatch

dist: generate
	@xgo -v --dest=dist --targets=$(TARGETS) -ldflags=$(LDFLAGS_DIST) github.com/tzero-dev/go-t0ken/cmd/t0ken

clean:
	@rm -f ./bin/t0ken
	@rm -f ./bin/t0kenbatch
	@rm -f ./build/*.abi
	@rm -f ./build/*.bin
	@find ./contracts -type f -not -name "generate.go" -exec rm {} \;
