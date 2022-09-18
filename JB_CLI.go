package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*-------------------------Logic--------------------------------*/

type Notepad interface {
	create(*[]string)
	list()
	clear()
	update(*[]string)
	delete(*[]string)
	exit()
	help()

	set_capacity(uint)

	get_scanner() *bufio.Scanner

	get_stop_flag() bool
}

// Specify one of notepad types
type simple_notepad struct {

	// declaring notes storage as nil string slice
	notes_storage []string

	// notes storage capacity
	storage_cap uint

	// declaring the scanner with default ScanLines split function
	std_in_scanner *bufio.Scanner

	// loop stop flag
	stop_flag bool
}

// Methods
func (CLI *simple_notepad) create(stdin_words_slice *[]string) {

	// check if storage was cleaned
	if CLI.notes_storage == nil {
		CLI.notes_storage = make([]string, CLI.storage_cap)
	}

	// check if storage is full
	if uint(len(CLI.notes_storage)) == CLI.storage_cap {

		fmt.Println("[Error] Notepad is full")

	} else {

		// check if there are any other words besides 'create' command
		if len(*stdin_words_slice) > 1 {

			rest_line := strings.Join((*stdin_words_slice)[1:], " ")

			//check if an empty string was provided via stdin after command (space was clicked)
			if strings.TrimSpace(rest_line) != "" {

				CLI.notes_storage = append(CLI.notes_storage, rest_line)

				fmt.Println("[OK] The note was successfully created")

				return

			}
		}

		fmt.Println("[Error] Missing note argument")
	}

}
func (CLI *simple_notepad) list() {

	if len(CLI.notes_storage) != 0 {
		for idx, val := range CLI.notes_storage {
			fmt.Printf("[Info] %d: %s\n", idx+1, val)
		}
	} else {
		fmt.Println("[Info] Notepad is empty")
	}

}
func (CLI *simple_notepad) clear() {

	if CLI.notes_storage != nil {
		fmt.Println("[OK] All notes were successfully deleted")
		CLI.notes_storage = nil
		CLI.storage_cap = 0
	} else {
		fmt.Println("[Info] Notepad is already empty")
	}

}
func (CLI *simple_notepad) update(stdin_words_slice *[]string) {

	if len(*stdin_words_slice) > 1 {

		if len(*stdin_words_slice) >= 3 {

			position, err := strconv.Atoi((*stdin_words_slice)[1])

			if err != nil || position <= 0 {
				fmt.Printf("[Error] Invalid position: %s\n", (*stdin_words_slice)[1])
				return
			}

			switch {

			case uint(position) > CLI.storage_cap:

				fmt.Printf("[Error] Position %d is out of the boundary [1, %d]\n", position, CLI.storage_cap)

			case position > len(CLI.notes_storage):

				fmt.Println("[Error] There is nothing to update")

			default:

				rest_line := strings.Join((*stdin_words_slice)[2:], " ")

				if strings.TrimSpace(rest_line) != "" {

					CLI.notes_storage[position-1] = rest_line
					fmt.Printf("[OK] The note at position %d was successfully updated\n", position)
					return

				}
			}

		} else {
			fmt.Println("[Error] Missing note argument")
		}

	} else {
		fmt.Println("[Error] Missing position argument")
	}

}
func (CLI *simple_notepad) delete(stdin_words_slice *[]string) {

	if len(*stdin_words_slice) > 1 {

		position, err := strconv.Atoi((*stdin_words_slice)[1])

		if err != nil || position <= 0 {
			fmt.Printf("[Error] Invalid position: %s\n", (*stdin_words_slice)[1])
			return
		}

		switch {

		case uint(position) > CLI.storage_cap:

			fmt.Printf("[Error] Position %d is out of boundaries [1, %d]\n", position, CLI.storage_cap)

		case position > len(CLI.notes_storage):

			fmt.Println("[Error] There is nothing to delete")

		default:

			CLI.notes_storage = append(CLI.notes_storage[:position-1], CLI.notes_storage[position:]...)
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", position)

		}

	} else {
		fmt.Println("[Error] Missing position argument")
	}
}
func (CLI *simple_notepad) exit() {
	CLI.stop_flag = true
	CLI.notes_storage = nil
	fmt.Println("[Info] Bye!")
}
func (CLI *simple_notepad) help() {
	fmt.Printf("This is an `simple notepad`.\nFollowing commands are available:\n\t- create <note>: adds typed text as note to the notepad;\n\t- list: output all the existing notes;\n\t- clear: clean up notepad, erase all the notes;\n\t- update <position> <new note>: update note at pointed position;\n\t- delete <position>: erase note at pointed position;\n\t- exit: stop notepad\n")
}
func (CLI *simple_notepad) set_capacity(value uint) {
	CLI.notes_storage = make([]string, 0, value)
	CLI.storage_cap = value
}
func (CLI *simple_notepad) get_scanner() *bufio.Scanner {
	return CLI.std_in_scanner
}
func (CLI *simple_notepad) get_stop_flag() bool {
	return CLI.stop_flag
}

/*---------------------Operations--------------------------------*/
// Get user input parameter, validate and fill storage_cap variable
func fill_capacity(notepad Notepad) bool {

	fmt.Print("Enter the maximum number of notes: > ")

	var storage_cap uint = 0
	fmt.Scanf("%d", &storage_cap)

	if storage_cap <= 0 {

		fmt.Printf("[Error] Invalid capacity parameter type: must input natural number\n")
		return false

	} else if storage_cap > 100 {

		fmt.Printf("[Error] Notepad buffer oveflow: entered capacity -> %d; available max capacity -> 100\n", storage_cap)
		return false

	}

	notepad.set_capacity(storage_cap)
	return true

}

// Get data from standart input; return input words slice and command as part of this slice
func get_input(notepad Notepad) []string {

	fmt.Print("Enter a command and data: > ")

	var input_words_slice []string = nil
	scanner := notepad.get_scanner()

	// check if input is complete
	if scanner.Scan() {

		line := scanner.Text()

		input_words_slice = strings.Split(line, " ")

	}

	return input_words_slice
}

// Switch command type and validate
func switch_command(input_slice *[]string, notepad Notepad) {
	switch (*input_slice)[0] {

	case "create":

		notepad.create(input_slice)

	case "list":

		notepad.list()

	case "clear":

		notepad.clear()

	case "update":

		notepad.update(input_slice)

	case "delete":

		notepad.delete(input_slice)

	case "help":

		notepad.help()

	case "exit":

		notepad.exit()

	default:

		fmt.Println("[Error] Unknown command")

	}
}

/* Main function */
func main() {

	// create an instance of simple_notepad
	var notepad Notepad = &simple_notepad{nil, 0, bufio.NewScanner(os.Stdin), false}

	// ask for correct capacity input while it's not
	cap_corr_flag := false
	for !cap_corr_flag {
		cap_corr_flag = fill_capacity(notepad)
	}

	// main loop
	for !notepad.get_stop_flag() {

		input_slice := get_input(notepad)

		switch_command(&input_slice, notepad)

	}

}
