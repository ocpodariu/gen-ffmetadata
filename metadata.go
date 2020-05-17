package main

import (
	"fmt"
	"io"
	"time"

	"gopkg.in/yaml.v2"
)

type ChapterMarker struct {
	Title        string        `yaml:"title"`
	StartTimeStr string        `yaml:"start"`
	StartTime    time.Duration `yaml:"-"`
	EndTime      time.Duration `yaml:"-"`
}

type VideoMetadata struct {
	Title       string          `yaml:"title"`
	DurationStr string          `yaml:"duration"`
	Duration    time.Duration   `yaml:"-"`
	Artist      string          `yaml:"artist"`
	Episode     int             `yaml:"episode"`
	DateStr     string          `yaml:"date"`
	Date        time.Time       `yaml:"-"`
	Chapters    []ChapterMarker `yaml:"chapters"`
}

// ParseYAML extracts video metadata from a YAML file
func (vm *VideoMetadata) ParseYAML(r io.Reader) error {
	err := yaml.NewDecoder(r).Decode(vm)
	if err != nil {
		return err
	}

	vm.Date, err = time.Parse("02-Jan-2006", vm.DateStr)
	if err != nil {
		return fmt.Errorf("invalid date (expected format DD-MON-YYYY): %v", err)
	}

	vm.Duration, err = parseStopwatchTime(vm.DurationStr)
	if err != nil {
		return fmt.Errorf("invalid duration (expected format HH:MM:SS): %v", err)
	}

	// Parse chapter times
	for i, ch := range vm.Chapters {
		d, err := parseStopwatchTime(ch.StartTimeStr)
		if err != nil {
			return fmt.Errorf("invalid chapter time '%s' (expected HH:MM:SS): %v", ch.StartTimeStr, err)
		}
		vm.Chapters[i].StartTime = d
	}
	// Each chapter ends right before the next one starts
	for i := 0; i < len(vm.Chapters)-1; i++ {
		vm.Chapters[i].EndTime = vm.Chapters[i+1].StartTime - 1
	}
	// The last chapter ends at the end of the video
	vm.Chapters[len(vm.Chapters)-1].EndTime = vm.Duration

	return nil
}

// parseStopwatchTime parses strings that contain times in stopwatch format "HH:MM:SS"
func parseStopwatchTime(s string) (time.Duration, error) {
	var hh, mm, ss int
	_, err := fmt.Sscanf(s, "%d:%d:%d", &hh, &mm, &ss)
	if err != nil {
		return 0, err
	}

	var d time.Duration
	d += time.Duration(hh) * time.Hour
	d += time.Duration(mm) * time.Minute
	d += time.Duration(ss) * time.Second
	return d, nil
}
