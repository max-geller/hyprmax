# Development Notes

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
