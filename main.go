package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Note struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      string    `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Notes struct {
	Notes []Note `json:"notes"`
}

const filename = "notes.json"

// Load notes from file
func LoadNotes() (*Notes, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return &Notes{Notes: []Note{}}, nil
		}
		return nil, err
	}
	defer file.Close()

	var notes Notes
	if err := json.NewDecoder(file).Decode(&notes); err != nil {
		return nil, err
	}
	return &notes, nil
}

// Save notes to file
func (n *Notes) SaveNotes() error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	return enc.Encode(n)
}

// Add a new note
func (n *Notes) Add(title, content, tags string) (*Note, error) {
	if title == "" || content == "" || tags == "" {
		return nil, errors.New("title, content, and tags cannot be empty")
	}
	note := Note{
		ID:        len(n.Notes) + 1, 
		Title:     title,
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	n.Notes = append(n.Notes, note)
	return &note, nil
}

// List all notes
func (n *Notes) List() {
	if len(n.Notes) == 0 {
		fmt.Println("No notes found.")
		return
	}
	for _, note := range n.Notes {
		fmt.Printf("%d. %s\n   %s\n   Tags: %s\n   Created: %s\n\n",
			note.ID, note.Title, note.Content, note.Tags, note.CreatedAt.Format(time.RFC1123))
	}
}

// View a note by ID
func (n *Notes) View(id int) error {
	for _, note := range n.Notes {
		if note.ID == id {
			fmt.Printf("%d. %s\n   %s\n   Tags: %s\n   Created: %s\n\n",
				note.ID, note.Title, note.Content, note.Tags, note.CreatedAt.Format(time.RFC1123))
			return nil
		}
	}
	return errors.New("note not found")
}

func (n *Notes) Delete(id int) error {
	for i, note := range n.Notes {
		if note.ID == id {
			n.Notes = append(n.Notes[:i], n.Notes[i+1:]...)
			return nil
		}
	}
	return errors.New("note not found")
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	notes, err := LoadNotes()
	if err != nil {
		fmt.Println("Error loading notes:", err)
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: note add <title> <content> <tags>")
			return
		}
		title, content, tags := os.Args[2], os.Args[3], os.Args[4]
		if _, err := notes.Add(title, content, tags); err != nil {
			fmt.Println("Error:", err)
			return
		}
		notes.SaveNotes()
		fmt.Println("Note added!")

	case "list":
		notes.List()

	case "view":
		if len(os.Args) < 3 {
			fmt.Println("Usage: note view <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := notes.View(id); err != nil {
			fmt.Println("Error:", err)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: note delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := notes.Delete(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			notes.SaveNotes()
			fmt.Println("Note deleted!")
		}

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: note <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <title> <content> <tags>   Add a new note")
	fmt.Println("  list                           List all notes")
	fmt.Println("  view <id>                      View a note")
	fmt.Println("  delete <id>                    Delete a note")
}
