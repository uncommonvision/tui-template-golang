# TUI Template - Go

A comprehensive starter template for building Terminal User Interfaces (TUIs) in Go using [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Cobra CLI](https://github.com/spf13/cobra).

## Features

- 🚀 **Full CLI Structure** - Professional command-line interface with subcommands
- 🎨 **Pre-configured Styling** - Beautiful themes using Lipgloss
- 📱 **Multi-screen Navigation** - Home, Settings, and Help screens
- 🧩 **Reusable Components** - Modular architecture for easy extension
- 📚 **Example Applications** - List navigation, forms, and more
- ⚙️ **Configuration Management** - JSON-based user configuration
- ⌨️ **Keyboard Shortcuts** - Intuitive navigation and controls
- 🔧 **Development Ready** - Makefile and build tools included

## Quick Start

### Installation

```bash
git clone <this-repo>
cd tui-template-golang
go mod tidy
go build -o mytui
```

### Usage

See all available commands:
```bash
./mytui
```

## Commands

| Command | Description |
|---------|-------------|
| `mytui` | Launch the main TUI application (default) |
| `mytui examples list` | Run interactive list example |
| `mytui examples form` | Run form input example |
| `mytui config init` | Create default configuration file |
| `mytui config show` | Display current configuration |
| `mytui version` | Show version information |

## Project Structure

```
tui-template-golang/
├── main.go                 # Entry point
├── cmd/                    # Cobra CLI commands
│   ├── root.go            # Root command setup
│   ├── start.go           # Main TUI launcher
│   ├── examples.go        # Example commands
│   ├── config.go          # Configuration management
│   └── version.go         # Version command
├── internal/              # Private application code
│   ├── models/           # Bubble Tea models
│   │   └── main.go      # Main application model
│   ├── components/       # Reusable UI components
│   ├── styles/           # Lipgloss styling definitions
│   │   └── styles.go
│   └── config/           # Configuration management
│       └── config.go
├── examples/             # Example applications
│   ├── list/            # List navigation example
│   ├── form/            # Form input example
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
├── Makefile            # Build automation
└── README.md           # This file
```

## Keyboard Shortcuts

### Main Application
- `1` or `h` - Go to Home screen
- `2` or `s` - Go to Settings screen
- `3` or `?` - Show Help screen
- `q` or `Ctrl+C` - Quit application

### List Example
- `↑`/`↓` - Navigate list items
- `Enter` - Select item
- `/` - Filter/search
- `q` - Quit

### Form Example
- `Tab` - Next field
- `Shift+Tab` - Previous field
- `Enter` - Submit form (on last field)
- `Esc` or `q` - Quit

## Styling System

The template includes a comprehensive styling system using Lipgloss:

- **Colors**: Primary, Secondary, Success, Warning, Error, Info
- **Styles**: Title, Header, Border, Button, Help, List items
- **Layout**: Consistent spacing and margins

Styles are defined in `internal/styles/styles.go` and can be easily customized.

## Configuration

The application supports JSON-based configuration stored in `~/.config/mytui/config.json`:

```json
{
  "theme": "default",
  "debug": false,
  "auto_save": true,
  "log_level": "info"
}
```

Initialize configuration:
```bash
./mytui config init
```

View current configuration:
```bash
./mytui config show
```

## Development

### Building

```bash
# View available Makefile commands
make

# Build for current platform
make build

# Build for all platforms
make build-all

# Clean build artifacts
make clean
```

### Adding New Commands

1. Create a new file in `cmd/` directory
2. Define your cobra command
3. Add it to the root command in `init()`
4. Implement your TUI model if needed

### Adding New Examples

1. Create a directory in `examples/`
2. Implement your Bubble Tea model
3. Add the command to `cmd/examples.go`

## Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - Pre-built components
- [Cobra](https://github.com/spf13/cobra) - CLI framework

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This template is available under the MIT license. See the [LICENSE](LICENSE) file for more info.

## Inspiration

This template is inspired by the excellent work of the [Charm](https://charm.sh/) team and their ecosystem of TUI tools.

---

⚡ Start building amazing terminal interfaces with Go and Bubble Tea!
