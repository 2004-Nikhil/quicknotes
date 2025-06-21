package main

import (
	"os"
	"path/filepath"
	"strings"
)

// Get data file path (cross-platform)
func getDataFilePath() string {
    homeDir, _ := os.UserHomeDir()
    return filepath.Join(homeDir, ".quicknotes", "data.json")
}

// Get default templates
func getDefaultTemplates() []Template {
    return []Template{
        {Name: "Meeting Notes", Content: "# Meeting Notes...", Tags: []string{"meeting", "work"}},
        {Name: "Daily Journal", Content: "# Daily Journal...", Tags: []string{"journal", "personal"}},
        {Name: "Project Planning", Content: "# Project Plan...", Tags: []string{"project", "planning", "work"}},
        {Name: "Quick Idea", Content: "# Idea...", Tags: []string{"idea", "brainstorm"}},
    }
}

// removeNoteByID removes a note from a slice by its ID.
func removeNoteByID(notes []Note, id int) []Note {
    for i, note := range notes {
        if note.ID == id {
            return append(notes[:i], notes[i+1:]...)
        }
    }
    return notes
}

// countNotesInFolder counts notes within a specific folder.
func countNotesInFolder(notes []Note, folder string) int {
    count := 0
    for _, note := range notes {
        if note.Folder == folder {
            count++
        }
    }
    return count
}

// countNotesWithTag counts notes that have a specific tag.
func countNotesWithTag(notes []Note, tag string) int {
    count := 0
    for _, note := range notes {
        for _, noteTag := range note.Tags {
            if noteTag == tag {
                count++
                break
            }
        }
    }
    return count
}

// containsTag checks if a slice of tags contains a specific query.
func containsTag(tags []string, query string) bool {
    for _, tag := range tags {
        if strings.Contains(strings.ToLower(tag), query) {
            return true
        }
    }
    return false
}