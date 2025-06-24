package main

import (
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/bubbles/textarea"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
)

// Application states
type viewState int

const (
    mainMenuView viewState = iota
    noteListView
    noteEditView
    searchView
    folderManageView
    tagManageView
    templateView
    inputDialogView
)

// List item for Charm's list component
type item struct {
    title, desc string
    id          int
}

func (i item) FilterValue() string { return i.title }
func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }

// Model for the application
type model struct {
    state         viewState
    data          *AppData
    list          list.Model
    textInput     textinput.Model
    textArea      textarea.Model
    currentNote   *Note
    message       string
    messageType   string // "success", "error", "warning"
    width, height int
    inputMode     string // "note_title", "folder_name", "tag_name"
    previousState viewState
}

// Initialize the application
func initialModel() model {
    data := loadData()
    ti := textinput.New()
    ti.Placeholder = "Enter text..."
    ti.Focus()
    ta := textarea.New()
    ta.SetWidth(80)
    ta.SetHeight(20)
    ta.Focus()
    m := model{
        state:     mainMenuView,
        data:      data,
        textInput: ti,
        textArea:  ta,
    }
    m = m.loadMainMenu() // Load initial menu
    return m
}

// Init is the first command that will be executed.
func (m model) Init() tea.Cmd {
    return nil
}