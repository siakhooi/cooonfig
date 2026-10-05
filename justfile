# List available recipes
default:
    @just --list

# Remove build outputs
clean:
    rm -rf bin dist test-report.json test-coverage.out test-coverage.html

# Run golangci-lint
golangci-lint:
    golangci-lint run

# Run tests and write coverage reports
test:
    scripts/test.sh

# Lint and test
ci:
    scripts/ci.sh

# Cross-compile into bin/ and build snapshot archives into dist/
build:
    scripts/build.sh
    scripts/goreleaser.sh snapshot

# Clean, lint, test, and build
all: clean ci build

# Create a GitHub release from release.env
release:
    scripts/create-release.sh

# Build snapshot archives in dist/ without uploading
go-release:
    scripts/goreleaser.sh snapshot

# Watch the current GitHub Actions run
commit-watch:
    gh run watch

# Create a GitHub release and watch the Actions run
release-watch: release
    gh run watch

binary := "bin/cooonfig-linux-amd64"

# Show CLI help
run-help:
    {{ binary }} --help

# Print the version
run-version:
    {{ binary }} --version

# Run the CLI with the given arguments
run *args:
    {{ binary }} {{ args }}
