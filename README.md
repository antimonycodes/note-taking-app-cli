# 📝 Notes CLI

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub Issues](https://img.shields.io/github/issues/yourusername/notes-cli)](https://github.com/yourusername/notes-cli/issues)
[![GitHub Stars](https://img.shields.io/github/stars/yourusername/notes-cli)](https://github.com/yourusername/notes-cli/stargazers)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

A lightweight, fast command-line note-taking application written in Go. Manage your notes efficiently with tags, timestamps, and persistent JSON storage.

## ✨ Features

- **Create & Edit Notes** - Quickly create new notes or edit existing ones
- **Tag System** - Organize notes with multiple tags for easy categorization
- **Search & Filter** - Find notes by tags or content
- **Persistent Storage** - All notes saved locally in JSON format
- **Timestamps** - Automatic tracking of creation and modification times
- **Export** - Export individual notes to Markdown files
- **Fast & Lightweight** - Built with Go for optimal performance

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/notes-cli.git
cd notes-cli

# Build the application
go build -o notes

# Optional: Install globally
go install
```

### Basic Usage

```bash
# Create a new note
./notes new "My First Note"

# List all notes
./notes list

# View a specific note
./notes view 1

# Edit a note
./notes edit 1

# Add tags to a note
./notes tag 1 golang productivity

# Search notes by tag
./notes find golang

# Export note to markdown
./notes export 1 my-note.md

# Delete a note
./notes delete 1
```

## 📖 Commands

| Command              | Description             | Example                     |
| -------------------- | ----------------------- | --------------------------- |
| `new <title>`        | Create a new note       | `notes new "Meeting Notes"` |
| `list`               | Display all notes       | `notes list`                |
| `view <id>`          | View a specific note    | `notes view 1`              |
| `edit <id>`          | Edit note content       | `notes edit 1`              |
| `tag <id> <tags...>` | Add tags to a note      | `notes tag 1 work urgent`   |
| `find <tag>`         | Find notes by tag       | `notes find work`           |
| `export <id> <file>` | Export note to .md file | `notes export 1 note.md`    |
| `delete <id>`        | Delete a note           | `notes delete 1`            |

## 📂 Project Structure

```
notes-cli/
├── main.go           # Main application entry point
├── notes.json        # Storage file (auto-generated)
├── go.mod            # Go module file
├── README.md         # This file
└── LICENSE           # License file
```

## 🗂️ Data Structure

Notes are stored in `notes.json` with the following structure:

```json
{
  "notes": [
    {
      "id": 1,
      "title": "My Note",
      "content": "Note content here",
      "tags": ["golang", "cli"],
      "created_at": "2025-09-29T10:30:00Z",
      "updated_at": "2025-09-29T10:30:00Z"
    }
  ]
}
```

## 🛠️ Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/antimonycodes/-Note-Taking-App-CLI
cd notes-cli

# Run tests (when available)
go test ./...

# Build binary
go build -o notes

# Run without building
go run main.go list
```

### Code Structure

The application is organized into the following components:

- **Note struct** - Represents individual notes with metadata
- **NoteList struct** - Manages the collection of notes
- **Storage functions** - Handle JSON encoding/decoding
- **Command handlers** - Process CLI commands
- **Helper functions** - Utility functions for common operations

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -m 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Open a Pull Request

### Development Guidelines

- Write clear, descriptive commit messages
- Add comments for complex logic
- Follow Go best practices and conventions
- Test your changes thoroughly
- Update documentation as needed

## 🐛 Known Issues

- None currently. Please report issues on GitHub!

## 📝 Future Enhancements

- [ ] Full-text search across note content
- [ ] Note encryption for sensitive information
- [ ] Cloud sync support (Dropbox, Google Drive)
- [ ] Rich text formatting in terminal
- [ ] Note templates
- [ ] Bulk operations (delete, tag multiple notes)
- [ ] Archive/unarchive functionality
- [ ] Sort notes by date, title, or tags
- [ ] Import notes from other formats
- [ ] Note linking and backlinks

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👤 Author

**Antimony**

- GitHub: [@antimonycodes](https://github.com/antimonycodes)
- X: [@tobilobasb](https://x.com/tobilobasb)

## 🙏 Acknowledgments

- Built as part of learning Go and CLI application development
- Thanks to the Go community for excellent documentation and tools

## 📚 Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

---

**Happy Note Taking! 📝✨**

If you find this project helpful, please consider giving it a ⭐️ on GitHub!
