package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Update function for the Bubble Tea model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.list.SetSize(msg.Width, m.height-4)
		m.textArea.SetWidth(msg.Width - 4)
		m.textArea.SetHeight(msg.Height - 8)
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		switch m.state {
		case mainMenuView:
			return m.updateMainMenu(msg)
		case noteListView:
			return m.updateNoteList(msg)
		case noteEditView:
			return m.updateNoteEdit(msg)
		case searchView:
			return m.updateSearch(msg)
		case folderManageView:
			return m.updateFolderManage(msg)
		case tagManageView:
			return m.updateTagManage(msg)
		case templateView:
			return m.updateTemplate(msg)
		case inputDialogView:
			return m.updateInputDialog(msg)
		}
	}

	return m, cmd
}

// updateMainMenu handles keypresses in the main menu view.
func (m model) updateMainMenu(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		i, ok := m.list.SelectedItem().(item)
		if ok {
			m.message, m.messageType = "", ""
			switch i.title {
			case "📝 View Notes":
				m.state = noteListView
				m = m.loadNoteList()
			case "➕ New Note":
				m.state = inputDialogView
				m.inputMode = "note_title"
				m.previousState = mainMenuView
				m.textInput.SetValue("")
				m.textInput.Placeholder = "Enter note title..."
				m.textInput.Focus()
			case "🔍 Search Notes":
				m.state = searchView
				m.textInput.SetValue("")
				m.textInput.Focus()
			case "📁 Manage Folders":
				m.state = folderManageView
				m = m.loadFolderList()
			case "🏷️  Manage Tags":
				m.state = tagManageView
				m = m.loadTagList()
			case "📋 Templates":
				m.state = templateView
				m = m.loadTemplateList()
			case "❌ Exit":
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// updateNoteList handles keypresses in the note list view.
func (m model) updateNoteList(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = mainMenuView
		m = m.loadMainMenu()
		m.message, m.messageType = "", ""
	case "enter":
		if i, ok := m.list.SelectedItem().(item); ok {
			for _, note := range m.data.Notes {
				if note.ID == i.id {
					m.currentNote = &note
					m.state = noteEditView
					m.textArea.SetValue(note.Content)
					m.textArea.Focus()
					break
				}
			}
		}
	case "d":
		if i, ok := m.list.SelectedItem().(item); ok {
			m.data.Notes = removeNoteByID(m.data.Notes, i.id)
			saveData(m.data)
			m = m.loadNoteList()
			m.message, m.messageType = "Note deleted successfully!", "success"
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// updateNoteEdit handles keypresses in the note editing view.
func (m model) updateNoteEdit(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+s":
		m.currentNote.Content = m.textArea.Value()
		m.currentNote.UpdatedAt = time.Now()

		found := false
		for i, note := range m.data.Notes {
			if note.ID == m.currentNote.ID {
				m.data.Notes[i] = *m.currentNote
				found = true
				break
			}
		}
		if !found {
			m.data.Notes = append(m.data.Notes, *m.currentNote)
		}

		saveData(m.data)
		m.state = noteListView
		m = m.loadNoteList()
		m.message, m.messageType = "Note saved successfully!", "success"
	case "esc":
		m.state = noteListView
		m = m.loadNoteList()
	}

	var cmd tea.Cmd
	m.textArea, cmd = m.textArea.Update(msg)
	return m, cmd
}

// updateSearch handles keypresses in the search view.
func (m model) updateSearch(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.state = mainMenuView
		m = m.loadMainMenu()
	case "enter":
		query := m.textInput.Value()
		results := m.searchNotes(query)

		items := []list.Item{}
		for _, note := range results {
			desc := fmt.Sprintf("📁 %s | 🏷️ %s | %s", note.Folder, strings.Join(note.Tags, ", "), note.CreatedAt.Format("2006-01-02"))
			items = append(items, item{title: note.Title, desc: desc, id: note.ID})
		}

		m.list = m.createList()
		m.list.Title = fmt.Sprintf("Search Results for: '%s'", query)
		m.list.SetItems(items)
		m.state = noteListView
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// updateFolderManage handles keypresses in the folder management view.
func (m model) updateFolderManage(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = mainMenuView
		m = m.loadMainMenu()
	case "enter":
		if i, ok := m.list.SelectedItem().(item); ok && i.id == -1 {
			m.state = inputDialogView
			m.inputMode = "folder_name"
			m.previousState = folderManageView
			m.textInput.SetValue("")
			m.textInput.Placeholder = "Enter folder name..."
			m.textInput.Focus()
		}
	case "d":
		if i, ok := m.list.SelectedItem().(item); ok && i.id >= 0 && i.title != "General" && i.title != "Work" && i.title != "Personal" {
			folderNameToDelete := i.title
			var newFolders []string
			for _, folder := range m.data.Folders {
				if folder != folderNameToDelete {
					newFolders = append(newFolders, folder)
				}
			}
			m.data.Folders = newFolders
			saveData(m.data)
			m = m.loadFolderList()
			m.message, m.messageType = "Folder deleted!", "success"
		} else {
			_, ok := m.list.SelectedItem().(item)
			if ok {
				m.message, m.messageType = "Cannot delete default folders!", "error"
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// updateTagManage handles keypresses in the tag management view.
func (m model) updateTagManage(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = mainMenuView
		m = m.loadMainMenu()
	case "enter":
		if i, ok := m.list.SelectedItem().(item); ok && i.id == -1 {
			m.state = inputDialogView
			m.inputMode = "tag_name"
			m.previousState = tagManageView
			m.textInput.SetValue("")
			m.textInput.Placeholder = "Enter tag name..."
			m.textInput.Focus()
		}
	case "d":
		if i, ok := m.list.SelectedItem().(item); ok && i.id >= 0 {
			tagToDelete := i.title
			var newTags []string
			for _, tag := range m.data.Tags {
				if tag != tagToDelete {
					newTags = append(newTags, tag)
				}
			}
			m.data.Tags = newTags
			saveData(m.data)
			m = m.loadTagList()
			m.message = "Tag deleted!"
			m.messageType = "success"
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// updateTemplate handles keypresses in the template view.
func (m model) updateTemplate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "esc":
		m.state = mainMenuView
		m = m.loadMainMenu()
	case "enter":
		if i, ok := m.list.SelectedItem().(item); ok && i.id < len(m.data.Templates) {
			template := m.data.Templates[i.id]
			m.currentNote = &Note{
				ID:        m.data.NextID,
				Title:     template.Name,
				Content:   template.Content,
				Tags:      append([]string{}, template.Tags...),
				Folder:    "General",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			m.data.NextID++
			m.state = noteEditView
			m.textArea.SetValue(template.Content)
			m.textArea.Focus()
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// updateInputDialog handles keypresses in the input dialog view.
func (m model) updateInputDialog(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.state = m.previousState
		switch m.previousState {
		case folderManageView:
			m = m.loadFolderList()
		case tagManageView:
			m = m.loadTagList()
		default:
			m = m.loadMainMenu()
		}
	case "enter":
		input := strings.TrimSpace(m.textInput.Value())
		if input == "" {
			m.message, m.messageType = "Name cannot be empty!", "error"
			return m, nil
		}
		switch m.inputMode {
		case "note_title":
			m.currentNote = &Note{
				ID:        m.data.NextID,
				Title:     input,
				Folder:    "General",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			m.data.NextID++
			m.state = noteEditView
			m.textArea.SetValue("")
			m.textArea.Focus()
		case "folder_name":
			m.data.Folders = append(m.data.Folders, input)
			saveData(m.data)
			m.state = folderManageView
			m = m.loadFolderList()
			m.message, m.messageType = fmt.Sprintf("Folder '%s' created!", input), "success"
		case "tag_name":
			m.data.Tags = append(m.data.Tags, input)
			saveData(m.data)
			m.state = tagManageView
			m = m.loadTagList()
			m.message, m.messageType = fmt.Sprintf("Tag '%s' created!", input), "success"
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}