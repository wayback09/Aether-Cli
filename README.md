# Aether CLI

The `aether-cli` is the official developer toolkit for creating, building, and packaging extensions for the Aether Minecraft Launcher. It allows developers to quickly scaffold new projects and package them securely into `.aex` format.


## Installation

Install the CLI with Go:

```bash
go install github.com/wayback09/aether-cli@latest
```

The binary is installed as `aether-cli` (named after the module). Make sure your Go bin directory is on `PATH` (`$(go env GOPATH)/bin` on macOS/Linux, `%USERPROFILE%\go\bin` on Windows), then open a new terminal and run `aether-cli help`.

## Commands

- `aether-cli init <name> <id>`: Scaffolds a new Aether extension with a `manifest.json`, `main.js`, and `ui/index.html` frontend. Automatically integrates the `@aethermc/sdk`.
- `aether-cli validate`: Parses the `manifest.json` in the current directory and checks for missing required fields (`id`, `name`, `version`, `main`, `api`) to ensure compatibility.
- `aether-cli build`: Validates the extension and then securely packages it into an Aether `.aex` archive while stripping out development files like `.git` and `node_modules`.

## Development

The CLI is written in Go and uses standard libraries to keep the binary small and dependency-free.

```bash
# Build the CLI
go build -o aether-cli.exe

# Run the CLI
./aether-cli.exe help
```

## License

Licensed under the GNU General Public License v3.0 only. See [LICENSE](LICENSE).
