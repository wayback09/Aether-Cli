# Aether CLI

The `aether` CLI is the official developer toolkit for creating, building, and packaging extensions for the Aether Minecraft Launcher. It allows developers to quickly scaffold new projects and package them securely into `.aex` format.


## Installation

Install the CLI with Go:

```bash
go install github.com/wayback09/aether-cli@latest
```

Make sure your Go bin directory is on `PATH`, then run `aether help`.

## Commands

- `aether init <name> <id>`: Scaffolds a new Aether extension with a `manifest.json`, `main.js`, and `ui/index.html` frontend. Automatically integrates the `@aether/sdk`.
- `aether validate`: Parses the `manifest.json` in the current directory and checks for missing required fields (`id`, `name`, `version`, `main`, `api`) to ensure compatibility.
- `aether build`: Validates the extension and then securely packages it into an Aether `.aex` archive while stripping out development files like `.git` and `node_modules`.

## Development

The CLI is written in Go and uses standard libraries to keep the binary small and dependency-free.

```bash
# Build the CLI
go build -o aether.exe

# Run the CLI
./aether.exe help
```
