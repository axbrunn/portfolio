package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// ANSI color codes
const (
	colorReset  = "\033[0m"
	colorGray   = "\033[90m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorRed    = "\033[31m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

func levelColor(level string) string {
	switch strings.ToUpper(level) {
	case "DEBUG":
		return colorGray
	case "INFO":
		return colorGreen
	case "WARN", "WARNING":
		return colorYellow
	case "ERROR":
		return colorRed
	default:
		return colorCyan
	}
}

func main() {
	service := flag.String("service", "", "Filter by service name")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		// Try to parse as JSON
		var entry map[string]any
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			// Not JSON, print as-is
			fmt.Println(line)
			continue
		}

		// Filter by service if flag is set
		svc, _ := entry["service"].(string)
		if *service != "" && svc != *service {
			continue
		}

		// Extract common fields
		t, _ := entry["time"].(string)
		level, _ := entry["level"].(string)
		msg, _ := entry["msg"].(string)

		// Shorten timestamp
		if len(t) > 19 {
			t = t[:19]
		}

		// Build extra fields (skip the common ones)
		skip := map[string]bool{"time": true, "level": true, "msg": true, "service": true}
		var extras []string
		for k, v := range entry {
			if !skip[k] {
				extras = append(extras, fmt.Sprintf("%s=%v", k, v))
			}
		}

		lc := levelColor(level)
		svcPart := ""
		if svc != "" {
			svcPart = fmt.Sprintf("%s%-8s%s | ", colorCyan, svc, colorReset)
		}

		extPart := ""
		if len(extras) > 0 {
			extPart = fmt.Sprintf("  %s%s%s", colorGray, strings.Join(extras, " "), colorReset)
		}

		fmt.Printf("%s%s%s  |  %s%s%-5s%s  |  %s%s%s\n",
			colorGray, t, colorReset,
			svcPart,
			lc, strings.ToUpper(level), colorReset,
			msg, colorReset,
			extPart,
		)
	}
}
