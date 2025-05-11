# Development Notes

## Core Functionality

### Configuration Management
- Uses structured config types for type safety
- Automatic validation of all settings
- Backup system with timestamped files
- Test mode for safe development

### UI Architecture
- Main menu → Section views → Editors
- Each section has its own model and view
- Context-sensitive help available with `?`
- Error display system with multiple levels

### Validation System
- Type-specific validators (int, float, bool, string)
- Special validators for:
  - Window rules
  - Keybindings
  - Input settings
  - Animation parameters

### Editor Interfaces
- Generic editor model for rules and bindings
- Field-specific validation
- Required field marking
- Cancel/Save functionality

## Important Considerations

### Safety Measures
1. Always run in test mode during development
2. Backup creation before any real config changes
3. Validation before saving
4. Error checking during parsing

### User Experience
1. Context help available in all views
2. Clear error messages
3. Visual feedback for actions
4. Consistent navigation

### Configuration Handling
1. Parse → Validate → Modify → Validate → Save
2. Automatic type conversion
3. Range checking for numeric values
4. Format validation for complex settings

## Testing Guidelines
1. Use test configuration file
2. Validate all user inputs
3. Check edge cases
4. Verify backup system

## Future Enhancements
1. Live preview of changes
2. Configuration profiles
3. Import/Export functionality
4. Plugin system

## Known Limitations
1. Some settings require Hyprland restart
2. Complex animations need careful validation
3. Some settings are interdependent

## Safety Measures

### Test Mode

The application runs in test mode during development to prevent accidental modifications to the user's active Hyprland configuration.

```go
// Usage in main.go
cfg, err := config.LoadConfig("", true) // true enables test mode
```

Test mode uses a separate configuration file located at:

- `config/testdata/hyprland.conf`

### Backup System

Before any modifications to the real configuration file, the system automatically creates timestamped backups:

```
/path/to/hyprland.conf.20240311_153000.backup
```

Backup files are created:

- Before loading the real config file
- Before saving any changes
- With timestamp format: YYYYMMDD_HHMMSS

## Configuration Handling

### File Locations

- Default config: `~/.config/hypr/hyprland.conf`
- Test config: `config/testdata/hyprland.conf`
- Backup configs: `~/.config/hypr/hyprland.conf.[timestamp].backup`

### Config Loading Process

1. Check if test mode is enabled
2. Determine correct config path
3. Create backup if using real config
4. Parse configuration file
5. Load into structured format

## Development Workflow

### Testing New Features

1. Modify test configuration in `config/testdata/hyprland.conf`
2. Run application in test mode
3. Verify changes work as expected
4. Implement safety checks
5. Test with real configuration

### Before Production Use

- [ ] Verify backup system works
- [ ] Add user confirmations for changes
- [ ] Test backup restoration
- [ ] Add error recovery
- [ ] Document recovery procedures

## Safety Checklist

Before modifying real config:

- [ ] Backup exists
- [ ] Changes are validated
- [ ] User confirmed action
- [ ] Recovery path exists
- [ ] Changes can be reverted
