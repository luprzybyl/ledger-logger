package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func main() {
	// Define flags
	dateFlag := flag.String("date", "", "Date for the log entry in YYYY-MM-DD format (optional)")
	timeFlag := flag.String("time", "", "Time for the log entry in HH:MM format (optional)")
	flag.Parse()

	// Remaining args are the message
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: ledger-logger [--date YYYY-MM-DD] [--time HH:MM] <log message>")
		os.Exit(1)
	}
	message := strings.Join(args, " ")

	now := time.Now()
	dateStr := now.Format("2006-01-02")
	timeStr := now.Format("15:04")

	if *dateFlag != "" {
		dateStr = *dateFlag
	}
	if *timeFlag != "" {
		timeStr = *timeFlag
	}

	fileName := now.Format("2006-01") + ".txt"
	logLine := fmt.Sprintf("%s %s %s\n", dateStr, timeStr, message)

	// Get log directory from env variable, default to current directory
	dir := os.Getenv("LEDGER_LOGGER_DIR")
	if strings.TrimSpace(dir) == "" {
		dir = "."
	}

	logPath := filepath.Join(dir, fileName)
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if _, err := f.WriteString(logLine); err != nil {
		fmt.Printf("Failed to write log: %v\n", err)
		os.Exit(1)
	}

	// Sort the file by date and time after writing
	sortLogFile(logPath)
}

// sortLogFile reads all lines, sorts them by date and time, and writes them back
func sortLogFile(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Failed to read log file for sorting: %v\n", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) < 2 {
		return // nothing to sort
	}
	type logEntry struct {
		dt  time.Time
		raw string
	}
	var entries []logEntry
	for _, l := range lines {
		parts := strings.Fields(l)
		if len(parts) < 3 {
			continue // skip malformed
		}
		datetime := parts[0] + " " + parts[1]
		t, err := time.Parse("2006-01-02 15:04", datetime)
		if err != nil {
			continue // skip malformed
		}
		entries = append(entries, logEntry{dt: t, raw: l})
	}
	// Sort by dt ascending
	sort.Slice(entries, func(i, j int) bool { return entries[i].dt.Before(entries[j].dt) })
	var sortedLines []string
	for _, e := range entries {
		sortedLines = append(sortedLines, e.raw)
	}
	os.WriteFile(path, []byte(strings.Join(sortedLines, "\n")+"\n"), 0644)
}
