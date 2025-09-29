package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Store notes with: Title, Content, Tags, CreatedAt, UpdatedAt
type Note struct {
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

// LoadNotes reads notes from JSON file
func LoadNotes() (*Notes, error) {
	file, err := os.Open(filename)
	if err != nil {
		// If file doesn't exist, return empty list
		if os.IsNotExist(err) {
			return &Notes{Notes: []Note{}}, nil
		}
		return nil, err
	}
	defer file.Close()

	var notes Notes
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notes); err != nil {
		return nil, err
	}

	return &notes, nil
}

// SaveNotes writes notes to JSON file
func (n *Notes) SaveNotes() error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty print JSON
	return encoder.Encode(n)
}

var notesStore *Notes 

// Create a new note
func newNote(title, content, tags string) (*Notes, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}
	if content == "" {
		return nil, errors.New("content cannot be empty")
	}
	if tags == "" {
		return nil, errors.New("tags cannot be empty")
	}

	note := Note{
		Title:     title,
		Content:   content,
		Tags:      tags,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	notesStore.Notes = append(notesStore.Notes, note)
	return notesStore, nil
}
// view  note by index 

func viewNote(id string) (*Notes, error){
	// fmt.Printf("%s %T\n",id,id)
	if id == ""{
		return nil, errors.New("id cannot be empty")
	}

	parsedId,err := strconv.Atoi(id)
	if err != nil{
		return nil,err
	}

	if parsedId > len(notesStore.Notes)-1{
		return nil, errors.New("invalid number")
	}
	for i,note := range notesStore.Notes{
		if parsedId == i {
		fmt.Printf("%d. %s\n   %s\n   Tags: %s\n   Created: %s\n\n",
				i+1, note.Title, note.Content, note.Tags, note.CreatedAt.Format(time.RFC1123))
		}
	}
	return nil,nil

}


func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	var err error
	notesStore, err = LoadNotes() 
	if err != nil {
		fmt.Println("Error loading notes:", err)
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 5 {
			fmt.Println("Usage: note add <title> <content> <tags>")
			return
		}
		title := os.Args[2]
		content := os.Args[3]
		tags := os.Args[4]

		if _, err := newNote(title, content, tags); err != nil {
			fmt.Println("Error creating note:", err)
			return
		}

		if err := notesStore.SaveNotes(); err != nil {
			fmt.Println("Error saving:", err)
		} else {
			fmt.Println("✨ Note added!")
		}

	case "list":
		if len(notesStore.Notes) == 0 {
			fmt.Println("No notes found.")
			return
		}
		for i, note := range notesStore.Notes {
			fmt.Printf("%d. %s\n   %s\n   Tags: %s\n   Created: %s\n\n",
				i+1, note.Title, note.Content, note.Tags, note.CreatedAt.Format(time.RFC1123))
		}
	case "view":
		title := os.Args[2]
		fmt.Printf("Found: note %s found.\n\n",title)

		if _,err :=viewNote(title); err != nil{
			fmt.Println("error:", err)
		}

	default:
		printUsage()
	}
}


// printUsage displays usage instructions for the CLI application.
func printUsage() {
	fmt.Println("Usage: note-app <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  add <title> <content> <tags>   Add a new note")
	fmt.Println("  list                           List all notes")
	fmt.Println("  view <id>                      View a note")
	fmt.Println("  edit <id>                      Edit a note")
	fmt.Println("  search <tag>                   Search for notes")
	fmt.Println("  help                           Show this help message")
}