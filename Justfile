default: fmt mod-tidy

docgen: docgen-man docgen-markdown docgen-yaml

docgen-markdown:
  command rm --force --recursive "docs/markdown/"
  go run "cmd/tidy/main.go" docgen markdown --output "docs/markdown/"
  prettier --write "docs/markdown/"

docgen-man:
  command rm --force --recursive "docs/man/"
  go run "cmd/tidy/main.go" docgen man --output "docs/man/"

docgen-yaml:
  command rm --force --recursive "docs/yaml/"
  go run "cmd/tidy/main.go" docgen yaml --output "docs/yaml/"
  prettier --write "docs/yaml/"

fmt:
  go fmt ./...

mod-tidy:
  go mod tidy

upgrade:
  go get -u ./...
  go mod tidy
