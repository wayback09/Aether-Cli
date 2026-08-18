# Aether CLI

The `aether` CLI is the official developer toolkit for creating, building, and packaging extensions for the Aether Minecraft Launcher. It allows developers to quickly scaffold new projects and package them securely into `.aex` format.


## Installation

Install the CLI with Go:

```bash
go install github.com/wayback09/aether-cli@latest
```

Go installs the binary as `aether-cli` (named after the module). To use the `aether` command shown throughout this README, rename it and make sure your Go bin directory is on `PATH`:

```bash
# macOS / Linux
mv ~/go/bin/aether-cli ~/go/bin/aether
export PATH="$PATH:$(go env GOPATH)/bin"   # add to ~/.zshrc or ~/.bashrc

# Windows (PowerShell)
ren "$env:USERPROFILE\go\bin\aether-cli.exe" aether.exe
[Environment]::SetEnvironmentVariable("Path", $env:Path + ";$env:USERPROFILE\go\bin", "User")
```

Then open a new terminal and run `aether help`.

> **Note:** If you run `go install` again, the `aether-cli` binary will be recreated — re-apply the rename (or skip it and call the CLI as `aether-cli`).

## Commands

- `aether init <name> <id>`: Scaffolds a new Aether extension with a `manifest.json`, `main.js`, and `ui/index.html` frontend. Automatically integrates the `@aethermc/sdk`.
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

## License

Licensed under the GNU General Public License v3.0 only. See [LICENSE](LICENSE).
