package main

import "fmt"

// 1. THE DATA MODEL
// A struct is like a template for a record. Every 'Contact' will have these 4 fields.
type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

// 2. GLOBAL STATE
// contactList: A slice to keep the contacts in the order they were added.
var contactList []Contact

// contactIndexByName: A map used as an "index". 
// It maps a Name (string) to its position (int) in the contactList slice.
var contactIndexByName map[string]int

// nextID: A simple counter to give every contact a unique ID number.
var nextID = 1

// 3. THE INIT FUNCTION
// This special function runs automatically before main(). 
// It's the perfect place to initialize our map and slice.
func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

// 4. ADDING DATA
func addContact(name, email, phone string) {
	// Check if name already exists in our map index.
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	// Create a new instance of the Contact struct.
	newContact := Contact{
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}
	nextID++

	// Step A: Add to the slice.
	contactList = append(contactList, newContact)
	
	// Step B: Update the map index. 
	// The index is 'len(contactList) - 1' because append just added it to the end.
	contactIndexByName[name] = len(contactList) - 1
	
	fmt.Printf("Contact added: %v\n", name)
}

// 5. SEARCHING (Using the Index)
// This returns a *Contact (a pointer). 
// If found, we point to the data in the slice. If not, we return nil.
func findContact(name string) *Contact {
	index, exists := contactIndexByName[name]
	if exists {
		return &contactList[index]
	}
	return nil
}

// 6. DISPLAYING DATA
func ListContacts() {
	fmt.Println("--- Listing Contacts ---")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return
	}
	// 'range' gives us both the index (i) and a copy of the contact.
	for i, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}
	fmt.Println("")
}

func main() {
	addContact("Alice Wonderland", "alice@example.com", "111-2222")
	addContact("Bob The Builder", "bob@example.com", "333-4444")
	addContact("Charlie Brown", "charlie@example.com", "555-6666")
	
	// This will trigger the "already exists" logic.
	addContact("Alice Wonderland", "alice.new@example.com", "777-8888") 

	ListContacts()

	// Testing the find function.
	// Note: It looks for "Bob", but the contact is "Bob The Builder", so it will be nil.
	bob := findContact("Bob")
	if bob == nil {
		fmt.Println("No Bob contact found.")
	} else {
		fmt.Println("Bob contact found.", bob.Name)
	}
}
