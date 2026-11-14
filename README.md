# ledger-logger

A minimal command-line tool for logging timestamped entries to monthly text files. Written in Go, no dependencies.

## Usage

```
ledger-logger [--date YYYY-MM-DD] [--time HH:MM] <log message>
```

- Appends a log entry to the current month's file (e.g., `2026-11.txt`).
- Entries are sorted by date and time after each write.
- Log directory can be set with the `LEDGER_LOGGER_DIR` env variable (defaults to current directory).

## Build

```
go build -o ledger-logger .
```

or

```
./build.sh
```

A macOS Apple Silicon build script is also included: `./build-macos-arm.sh`.

## Examples

```
ledger-logger "Deployed new release"
ledger-logger --date 2026-11-13 --time 09:00 "Backdated entry"
LEDGER_LOGGER_DIR=~/logs ledger-logger "Goes to ~/logs/2026-11.txt"
```

Output format per line:

```
2026-11-14 10:00 Deployed new release
```
