# 🧹 Konmari – File Cleanup CLI

Konmari is a command-line tool written in Go to help you clean up old files on macOS and Linux. Featuring both a straightforward CLI and an interactive TUI powered by Bubble Tea. Konmari brings joy to file cleanup.

## Features
- Scan directories for old files
- Delete files based on age
- Dry-run mode to preview deletions safely
- Interactive TUI using Bubble Tea + Lipgloss
- Multi-file selection with visual toggling
- Adaptive styling for light/dark terminals

## Installation
```bash
go install github.com/Viriathus1/konmari@latest
```

## Usage
### CLI mode
```bash
konmari clean --dir ~/Downloads --days 30 --dry-run=false
```
| Option         | Description                           |
| -------------- | ------------------------------------- |
| `--dir`        | Directory to scan                     |
| `--days`       | Files older than X days (default: 30) |
| `--dry-run`    | Preview files without deleting        |

### TUI mode
Launch the interactive file picker
```bash
konmari method
```

## TODO
- [ ] Add verbosity flag
- [ ] Enable file directory deletion
- [ ] Configurable exclusion patterns
- [ ] Cron job setup
- [ ] Windows support

## Licence
MIT Licence