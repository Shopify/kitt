# kitt: A Golang 1.18+ Generics Toolkit

`kitt` provides some handy data structures and "building blocks" for use with Go 1.18+ generics.

## Available stuff:

- Set
- Queue/BoundedQueue

## Contributors

- [Chris Pappas](https://github.com/chrispappas)

## Quick Start

[main.go](./main.go) contains a comprehensive example of how the `set` and `queue` packages work.

Compile and run the main.go file in the project root in one easy command:

```bash
go run ./main.go
```

## Requirements

- Go 1.18 or newer (must support Generics)

## Running tests/benchmarks

Run the tests and benchmarks for the project using this command:

```bash
go test -v -race -bench=. ./...
```

## CI/CD and Github Actions

This project is configured to use GH Actions to automatically test/benchmark the project whenever pushes occur.
See the [.github/workflows](./.github/workflows) folder for all the details.

## License

kitt is released under the [MIT License](https://opensource.org/licenses/MIT).
