# Hyprmax

A modern, terminal-based settings manager for Hyprland written in Go.

## Features

- 🎨 Modern, intuitive TUI interface
- ⚡ Live configuration editing
- 🔒 Safe configuration handling with automatic backups
- ✨ Comprehensive settings validation
- 🎯 Context-sensitive help
- 🔑 Advanced keybinding management
- 📋 Window rules editor

## Installation

```bash
go install github.com/max-geller/hyprmax@latest
```

## Usage

### Basic Navigation
- Arrow keys or `j`/`k` for movement
- `Enter` to select/edit
- `Esc` to go back/cancel
- `q` to quit

### Special Commands
- `?` - Show context help
- `n` - Create new entry (in rules/bindings)
- `ctrl+s` - Save changes

### Safety Features
- Test mode for safe development
- Automatic backups before changes
- Validation before saving

## Development

Requirements:
- Go 1.21 or later
- Hyprland

```bash
# Clone repository
git clone https://github.com/max-geller/hyprmax.git
cd hyprmax

# Build
go build

# Run tests
go test ./...
```

## Contributing

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Run the tests
5. Submit a pull request

## License

MIT License
