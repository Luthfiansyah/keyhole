// Copyright 2020 Kuei-chun Chen. All rights reserved.

package mdb

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Logger stores logger info
type Logger struct {
	Collected time.Time `bson:"collected"`
	Warnings  []string  `bson:"warnings"`
	Logs      []string  `bson:"logs"`
	Params    string    `bson:"params"`
	Version   string    `bson:"version"`

	nocolor bool
}

// NewLogger returns Logger
func NewLogger(fullVersion string, params string) *Logger {
	p := Logger{Version: fullVersion, Params: params, Warnings: []string{}}
	p.Collected = time.Now()
	p.Logs = []string{fmt.Sprintf(`%v keyhole begins`, p.Collected.Format(time.RFC3339))}
	return &p
}

// SetNoColor set nocolor flag
func (p *Logger) SetNoColor(nocolor bool) {
	p.nocolor = nocolor
}

// Add adds a message
func (p *Logger) Add(message string) {
	str := fmt.Sprintf(`%v %v`, time.Now().Format(time.RFC3339), message)
	p.Logs = append(p.Logs, str)
}

// Warn adds an warning message
func (p *Logger) Warn(message string) {
	p.Warnings = append(p.Warnings, message)
	fmt.Println(codeRed, "*", message, codeDefault)
}

// Log adds and prints a message
func (p *Logger) Log(message string) {
	str := fmt.Sprintf(`%v %v`, time.Now().Format(time.RFC3339), message)
	p.Logs = append(p.Logs, str)
	log.Println(message)
}

// Print prints keyhole info
func (p *Logger) Print() string {
	if p == nil {
		return ""
	}
	strs := []string{fmt.Sprintf(`{ keyhole: { version: "%v", args: "%v" } }`, p.Version, p.Params)}
	strs = append(strs, p.Logs...)
	if len(p.Warnings) > 0 {
		strs = append(strs, "\nWarnings:")
		for _, warning := range p.Warnings {
			if p.nocolor {
				strs = append(strs, warning)
			} else {
				strs = append(strs, fmt.Sprintf(`%v%v%v`, codeRed, warning, codeDefault))
			}
		}
		strs = append(strs, "")
	}
	return strings.Join(strs, "\n")
}
