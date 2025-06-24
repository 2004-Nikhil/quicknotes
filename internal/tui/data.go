package main

import (
    "encoding/json"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"
)

// Data structures
type Note struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    Tags      []string  `json:"tags"`
    Folder    string    `json:"folder"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type AppData struct {
    Notes     []Note     `json:"notes"`
    Folders   []string   `json:"folders"`
    Tags      []string   `json:"tags"`
    Templates []Template `json:"templates"`
    NextID    int        `json:"next_id"`
}

type Template struct {
    Name    string   `json:"name"`
    Content string   `json:"content"`
    Tags    []string `json:"tags"`
}

// Load data from JSON file
func loadData() *AppData {
    dataFile := getDataFilePath()

    data := &AppData{
        Notes:     []Note{},
        Folders:   []string{"General", "Work", "Personal"},
        Tags:      []string{"important", "todo", "idea"},
        Templates: getDefaultTemplates(),
        NextID:    1,
    }

    if _, err := os.Stat(dataFile); os.IsNotExist(err) {
        saveData(data)
        return data
    }

    content, err := ioutil.ReadFile(dataFile)
    if err != nil {
        return data
    }

    json.Unmarshal(content, data)
    return data
}

// Save data to JSON file
func saveData(data *AppData) error {
    dataFile := getDataFilePath()
    dir := filepath.Dir(dataFile)
    os.MkdirAll(dir, 0755)

    content, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(dataFile, content, 0644)
}