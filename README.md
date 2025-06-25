# 🌈 QuickNotes

A beautiful, fast, and intuitive CLI note-taking application built with Go and Bubble Tea. QuickNotes provides a colorful terminal-based interface for creating, organizing, and managing your notes with ease.

## ✨ Features

- **📝 Rich Note Management**: Create, edit, view, and delete notes with full text editing capabilities
- **📁 Folder Organization**: Organize notes into custom folders for better structure
- **🏷️ Tag System**: Add and manage tags for easy categorization and filtering
- **🔍 Powerful Search**: Search through note titles, content, and tags instantly
- **📋 Templates**: Use pre-built templates for common note types (meetings, journals, etc.)
- **🎨 Beautiful UI**: Colorful, modern terminal interface with intuitive navigation
- **💾 Auto-Save**: Automatic data persistence with JSON storage
- **⚡ Fast Performance**: Lightweight and responsive CLI experience

## 🚀 Installation

### Prerequisites

- Go 1.19 or higher
- Git

### Build from Source

1. Clone the repository:
```bash
git clone https://github.com/2004-nikhil/quicknotes.git
cd quicknotes
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o quicknotes cmd/quicknotes/main.go
```

4. Run QuickNotes:
```bash
./quicknotes
```

### Install Globally 

To install QuickNotes globally on your system:

```bash
go install github.com/2004-nikhil/quicknotes/cmd/quicknotes@latest
```

Then you can run it from anywhere:
```bash
quicknotes
```

## 📖 Usage

### Main Menu Navigation

When you launch QuickNotes, you'll see the main menu with these options:

- **➕ New Note**: Create a new note from scratch
- **📝 View Notes**: Browse and manage existing notes
- **🔍 Search Notes**: Search through your notes
- **📁 Manage Folders**: Create and organize folders
- **🏷️ Manage Tags**: Create and organize tags
- **📋 Templates**: Use pre-built note templates
- **❌ Exit**: Quit the application

### Keyboard Shortcuts

#### Global Controls
- `↑/↓` or `j/k`: Navigate lists
- `Enter`: Select/confirm
- `Esc`: Go back/cancel

#### Note List View
- `Enter`: Edit selected note
- `d`: Delete selected note
- `q`: Return to main menu

#### Note Editor
- `Ctrl+S`: Save note
- `Esc`: Cancel editing (without saving)

#### Search
- Type your query and press `Enter` to search
- Search works across note titles, content, and tags

#### Folder/Tag Management
- `Enter`: Select item or create new folder/tag
- `d`: Delete selected folder/tag (except default folders)
- `q`: Return to main menu

## 📁 Data Storage

QuickNotes stores your data in a JSON file located at:
- **Linux/macOS**: `~/.quicknotes/data.json`
- **Windows**: `%USERPROFILE%\.quicknotes\data.json`

The data file contains:
- All your notes with metadata
- Custom folders and tags
- Application settings
- Note templates

## 📋 Default Templates

QuickNotes comes with several built-in templates:

1. **Meeting Notes** - For capturing meeting discussions and action items
2. **Daily Journal** - For personal journaling and daily reflections
3. **Project Planning** - For project planning and task organization
4. **Quick Idea** - For capturing spontaneous ideas and thoughts


## 🔧 Development

### Project Structure

```
quicknotes/
├── cmd/quicknotes/          # Application entry point
│   └── main.go
├── internal/tui/            # Terminal UI package
│   ├── app.go              # Main application runner
│   ├── data.go             # Data structures and persistence
│   ├── model.go            # Bubble Tea model
│   ├── updates.go          # Update logic and event handling
│   ├── views.go            # UI rendering and view logic
│   ├── styles.go           # Color schemes and styling
│   └── utils.go            # Utility functions
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies
└── README.md              # This file
```

### Dependencies

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling and layout

### Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🐛 Troubleshooting

### Common Issues

**Application won't start**
- Ensure Go 1.19+ is installed
- Check if all dependencies are installed with `go mod tidy`

**Data not persisting**
- Check if the `~/.quicknotes` directory has write permissions
- Ensure sufficient disk space is available

**Display issues**
- Try resizing your terminal window
- Ensure your terminal supports colors and Unicode characters

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Charm](https://charm.sh/) for the amazing Bubble Tea framework
- The Go community for excellent tooling and libraries
- All contributors who help improve QuickNotes

---

**Happy note-taking! 📝✨**