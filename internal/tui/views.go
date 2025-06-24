package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

// View function for rendering the UI
func (m model) View() string {
	var content string
	var finalContent strings.Builder

	finalContent.WriteString(titleStyle.Render("🌈 QuickNotes") + "\n\n")

	// Render message if present
	if m.message != "" {
		var msgStyle lipgloss.Style
		switch m.messageType {
		case "success":
			msgStyle = lipgloss.NewStyle().Foreground(successColor).Bold(true)
		case "error":
			msgStyle = lipgloss.NewStyle().Foreground(errorColor).Bold(true)
		case "warning":
			msgStyle = lipgloss.NewStyle().Foreground(warningColor).Bold(true)
		default:
			msgStyle = lipgloss.NewStyle().Foreground(accentColor)
		}
		finalContent.WriteString(msgStyle.Render(m.message) + "\n\n")
		// Clear message after preparing it for display
		m.message = ""
		m.messageType = ""
	}

	switch m.state {
	case mainMenuView, noteListView, folderManageView, tagManageView, templateView:
		content = m.list.View()
		// Add contextual help text
		switch m.state {
		case noteListView:
			content += "\n" + helpStyle.Render("Enter: edit, d: delete, q: back to menu")
		case folderManageView:
			content += "\n" + helpStyle.Render("Enter: select/create, d: delete, q: back to menu")
		case tagManageView:
			content += "\n" + helpStyle.Render("Enter: select/create, d: delete, q: back to menu")
		case templateView:
			content += "\n" + helpStyle.Render("Enter: use template, q: back to menu")
		default: // mainMenuView
			content += "\n" + helpStyle.Render("Use ↑/↓ to navigate, Enter to select, Ctrl+C to quit")
		}
	case noteEditView:
		title := "Editing Note"
		if m.currentNote != nil && m.currentNote.Title != "" {
			title = fmt.Sprintf("Editing: %s", m.currentNote.Title)
		}
		content = headerStyle.Render(title) + "\n\n"
		content += m.textArea.View()
		content += "\n" + helpStyle.Render("Ctrl+S: save, Esc: cancel")
	case searchView:
		content = headerStyle.Render("Search Notes") + "\n\n"
		content += "Enter search query:\n"
		content += m.textInput.View()
		content += "\n\n" + helpStyle.Render("Enter: search, Esc: back to menu")
	case inputDialogView:
		var title string
		switch m.inputMode {
		case "note_title":
			title = "New Note"
		case "folder_name":
			title = "New Folder"
		case "tag_name":
			title = "New Tag"
		}
		content = headerStyle.Render(title) + "\n\n"
		content += m.textInput.View()
		content += "\n\n" + helpStyle.Render("Enter: create, Esc: cancel")
	}

	finalContent.WriteString(content)
	return finalContent.String()
}

// --- View Helper / List Loading Functions ---

func (m model) createList() list.Model {
	l := list.New(nil, list.NewDefaultDelegate(), m.width, m.height-4)
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)
	l.KeyMap.Quit.SetKeys() // Disable built-in quit keys
	return l
}

// loadMainMenu prepares the list for the main menu view.
func (m model) loadMainMenu() model {
	items := []list.Item{
		item{title: "➕ New Note", desc: "Create a new note"},
		item{title: "📝 View Notes", desc: "Browse and manage your notes"},
		item{title: "🔍 Search Notes", desc: "Search through titles, content, and tags"},
		item{title: "📁 Manage Folders", desc: "Create and organize folders"},
		item{title: "🏷️  Manage Tags", desc: "Create and organize tags"},
		item{title: "📋 Templates", desc: "Use pre-built note templates"},
		item{title: "❌ Exit", desc: "Quit the application"},
	}
	m.list = m.createList()
	m.list.Title = "QuickNotes - Beautiful CLI Note Taking"
	m.list.SetItems(items)
	return m
}

// loadNoteList prepares the list for the note list view.
func (m model) loadNoteList() model {
	items := []list.Item{}
	for _, note := range m.data.Notes {
		desc := fmt.Sprintf("📁 %s | 🏷️ %s | %s", note.Folder, strings.Join(note.Tags, ", "), note.CreatedAt.Format("2006-01-02"))
		items = append(items, item{title: note.Title, desc: desc, id: note.ID})
	}
	m.list = m.createList()
	m.list.Title = "Your Notes"
	m.list.SetItems(items)
	return m
}

// loadFolderList prepares the list for the folder management view.
func (m model) loadFolderList() model {
	items := []list.Item{}
	for i, folder := range m.data.Folders {
		noteCount := countNotesInFolder(m.data.Notes, folder)
		desc := fmt.Sprintf("%d notes", noteCount)
		items = append(items, item{title: folder, desc: desc, id: i})
	}
	items = append(items, item{title: "➕ Add New Folder", desc: "Create a new folder", id: -1})
	m.list = m.createList()
	m.list.Title = "Folder Management"
	m.list.SetItems(items)
	return m
}

// loadTagList prepares the list for the tag management view.
func (m model) loadTagList() model {
	items := []list.Item{}
	for i, tag := range m.data.Tags {
		noteCount := countNotesWithTag(m.data.Notes, tag)
		desc := fmt.Sprintf("%d notes", noteCount)
		items = append(items, item{title: tag, desc: desc, id: i})
	}
	items = append(items, item{title: "➕ Add New Tag", desc: "Create a new tag", id: -1})
	m.list = m.createList()
	m.list.Title = "Tag Management"
	m.list.SetItems(items)
	return m
}

// loadTemplateList prepares the list for the templates view.
func (m model) loadTemplateList() model {
	items := []list.Item{}
	for i, template := range m.data.Templates {
		tags := strings.Join(template.Tags, ", ")
		desc := fmt.Sprintf("Tags: %s", tags)
		items = append(items, item{title: template.Name, desc: desc, id: i})
	}
	m.list = m.createList()
	m.list.Title = "Note Templates"
	m.list.SetItems(items)
	return m
}

// searchNotes filters notes based on a query.
func (m model) searchNotes(query string) []Note {
	var results []Note
	query = strings.ToLower(query)

	for _, note := range m.data.Notes {
		if strings.Contains(strings.ToLower(note.Title), query) ||
			strings.Contains(strings.ToLower(note.Content), query) ||
			containsTag(note.Tags, query) {
			results = append(results, note)
		}
	}
	return results
}