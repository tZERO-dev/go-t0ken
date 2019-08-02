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


.PHONY: t0ken t0kenbatch dist


print-%: ; @echo '$(subst ','\'',$*=$($*))'

all: t0ken t0kenbatch

t0ken:
	@go build -ldflags=$(LDFLAGS_EXPERIMENTAL) -o bin/t0ken ./cmd/t0ken

t0kenbatch:
	@go build -ldflags=$(LDFLAGS_EXPERIMENTAL) -o bin/t0kenbatch ./cmd/t0kenbatch

dist:
	@xgo -v --dest=dist --targets=$(TARGETS) -ldflags=$(LDFLAGS_DIST) github.com/tzero-dev/go-t0ken/cmd/t0ken

clean:
	@rm -f ./bin/*
