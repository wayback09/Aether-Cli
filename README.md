# Aether CLI

<p align="center">
  <a href="https://discord.gg/pQc9NnGhpG">
    <img src="https://img.shields.io/badge/discord-Join%20our%20Discord-5865F2?logo=discord&logoColor=white&style=for-the-badge" alt="Discord">
  </a>
</p>

The `aether-cli` is the official developer toolkit for creating, building, and packaging **extensions** and **themes** for the Aether Minecraft Launcher. It lets developers scaffold new projects and package them securely into `.aex` (extension) or `.theme` (appearance pack) format.

---

## Installation

Install the CLI with Go:

```bash
go install github.com/wayback09/aether-cli@latest
```

The binary is installed as `aether-cli`. Make sure your Go bin directory is on `PATH` (`$(go env GOPATH)/bin` on macOS/Linux, `%USERPROFILE%\go\bin` on Windows), then open a new terminal and run:

```bash
aether-cli help
```

---

## Commands

### `init` — Scaffold a new project

```bash
# Scaffold a new extension
aether-cli init <name> <id>

# Scaffold a new theme
aether-cli init <name> <id> --theme
```

| Flag | Alias | Description |
|:---|:---|:---|
| `--theme` | `-t` | Scaffold a theme instead of an extension |

**Extension output:**
```
<name>/
├── manifest.json     # Extension metadata (id, name, version, permissions, api)
├── package.json      # npm typings dep (@aethermc/sdk)
├── main.js           # Backend sandbox entry point (Goja)
└── ui/
    └── index.html    # Frontend rendered in a sandboxed iframe
```

**Theme output:**
```
<name>/
├── package.json      # Theme metadata (id, name, version, author, css, overwrite)
├── theme.css         # CSS overrides (Aether :root design tokens pre-filled)
├── overwrite.json    # Optional asset overrides (sidebar-logo, background, etc.)
└── README.md
```

---

### `validate` — Validate an extension or theme

```bash
# Validate extension (reads manifest.json)
aether-cli validate

# Validate theme (reads package.json)
aether-cli validate --theme
```

**Auto-detection:** if `manifest.json` is absent but `package.json` is present, the project is automatically treated as a theme.

| Flag | Alias | Description |
|:---|:---|:---|
| `--theme` | `-t` | Force theme validation mode |

**Extension checks:** `manifest.json` exists, valid JSON, required fields: `id`, `name`, `version`, `main`, `api`.

**Theme checks:** `package.json` exists, valid JSON, required fields: `id`, `name`, `version`, CSS file referenced by `css` field actually exists on disk.

---

### `build` — Package into .aex or .theme

```bash
# Build extension → <id>-<version>.aex
aether-cli build

# Build theme → <id>-<version>.theme
aether-cli build --theme
```

**Auto-detection:** same logic as `validate` — detects project type automatically.

| Flag | Alias | Description |
|:---|:---|:---|
| `--theme` | `-t` | Force theme build mode |

Runs validation first, then packages all project files into a zip-format archive, automatically excluding `.git/`, `node_modules/`, and existing archive files.

---

### `help` — Show help

```bash
aether-cli help
aether-cli --help
aether-cli -h
```

---

## Quick Start — Extension

```bash
aether-cli init my-extension com.example.myextension
cd my-extension
aether-cli validate
aether-cli build
# → com.example.myextension-1.0.0.aex
```

## Quick Start — Theme

```bash
aether-cli init my-theme com.example.mytheme --theme
cd my-theme
# Edit theme.css to customise colours, radii, spacing...
aether-cli validate
aether-cli build
# → com.example.mytheme-1.0.0.theme
```

---

## Development

The CLI is written in Go and uses only the standard library to keep the binary small and dependency-free.

```bash
# Build locally
go build -o aether-cli.exe   # Windows
go build -o aether-cli       # macOS / Linux

./aether-cli help
```

---

## License

Licensed under the GNU General Public License v3.0 only. See [LICENSE](LICENSE).
