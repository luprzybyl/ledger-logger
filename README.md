# ledger-logger

A concise command-line tool for logging dated entries to monthly text files, written in Go.

## Usage

```
ledger-logger [--date YYYY-MM-DD] [--time HH:MM] <log message>
```
- Appends a log entry to the current month's file (e.g., `2025-09.txt`).
- Entries are sorted by date and time after each write.
- Optionally set log directory with `LEDGER_=LOGGER_DIR` env variable (defaults to current directory).

## Build

```
./build.sh
```

## Example

```
ledger-logger --date 2025-09-08 --time 10:30 "Paid electricity bill"
```
