# Contributing to GoContainer

We love your input! We want to make contributing to GoContainer as easy and transparent as possible, whether it's:

- Reporting a bug
- Discussing the current state of the code
- Submitting a fix
- Proposing new features
- Becoming a maintainer

## Development Process

We use GitHub to host code, to track issues and feature requests, as well as accept pull requests.

## Pull Requests

1. Fork the repo and create your branch from `main`.
2. If you've added code that should be tested, add tests.
3. If you've changed APIs, update the documentation.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Development Setup

### Prerequisites

- Go 1.21 or later
- Make (optional, for running common tasks)

### Getting Started

1. Clone your fork:
   ```bash
   git clone https://github.com/yourusername/gocontainer.git
   cd gocontainer
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build the project:
   ```bash
   go build ./cmd/gocontainer
   ```

4. Run tests:
   ```bash
   go test ./...
   ```

## Testing

We use Go's built-in testing framework along with [testify](https://github.com/stretchr/testify) for assertions.

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -coverprofile=coverage.out ./...

# View coverage report
go tool cover -html=coverage.out
```

### Writing Tests

- Write tests for all new functionality
- Follow Go testing conventions
- Use table-driven tests where appropriate
- Mock external dependencies

Example test:

```go
func TestNewContainer(t *testing.T) {
    container, err := New()
    require.NoError(t, err)
    require.NotNil(t, container)
    require.Equal(t, ocischemav1.MediaTypeImageIndex, container.Index.MediaType)
}
```

## Code Style

We follow standard Go conventions:

- Use `gofmt` to format your code
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- Write clear, self-documenting code
- Add comments for exported functions and complex logic

### Linting

We use [golangci-lint](https://golangci-lint.run/) for linting:

```bash
golangci-lint run
```

## Project Structure

```
.
├── cmd/gocontainer/     # CLI application
├── pkg/
│   ├── binary/          # Binary analysis and platform detection
│   └── container/       # OCI container creation and management
├── examples/            # Usage examples
├── .github/workflows/   # CI/CD pipelines
└── docs/               # Additional documentation
```

## Commit Messages

We follow [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `style:` formatting changes
- `refactor:` code refactoring
- `test:` adding tests
- `chore:` maintenance tasks

Examples:
- `feat: add support for Windows containers`
- `fix: handle empty binary paths correctly`
- `docs: update installation instructions`

## Issue and Bug Reports

We use GitHub issues to track public bugs. Report a bug by [opening a new issue](https://github.com/sntns/gocontainer/issues/new).

**Great Bug Reports** tend to have:

- A quick summary and/or background
- Steps to reproduce
  - Be specific!
  - Give sample code if you can
- What you expected would happen
- What actually happens
- Notes (possibly including why you think this might be happening, or stuff you tried that didn't work)

## Feature Requests

We welcome feature requests! Please:

1. Check if the feature has already been requested
2. Provide a clear description of the problem you're trying to solve
3. Explain why this feature would be useful to other users
4. Consider providing a rough implementation plan

## Release Process

Releases are automated through GitHub Actions when tags are pushed:

1. Update version in relevant files
2. Create a new tag: `git tag v1.0.0`
3. Push the tag: `git push origin v1.0.0`
4. GitHub Actions will build and create the release

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

## Questions?

Feel free to [open an issue](https://github.com/sntns/gocontainer/issues/new) with the `question` label if you have any questions about contributing.