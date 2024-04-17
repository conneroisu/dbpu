package dbpu

import (
	"fmt"
	"net/http"
)

// Client is a client for the dbpu API.
type Client struct {
	*http.Client
	BaseURL   string // Base URL for API requests.
	RegionURL string // Base URL for region requests.
	DbFolder  string // Folder for db files.
	OrgToken  string // Token for organization.
	OrgName   string // Name of organization.
	GroupName string // Name of group.
	ApiToken  string // Token for API.
}

// NewClient returns a new client.
func NewClient() *Client {
	return &Client{
		Client:    http.DefaultClient,
		BaseURL:   "https://api.turso.tech/v1",
		RegionURL: "https://region.turso.io",
	}
}

// SetOrgToken sets the organization token of the dbpu client.
func (c *Client) SetOrgToken(token string) { c.OrgToken = token }

// SetOrgName sets the name of the organization to use in the dbpu client.
func (c *Client) SetOrgName(name string) { c.OrgName = name }

// SetGroupName sets the name of the group to use in the dbpu client.
func (c *Client) SetGroupName(name string) { c.GroupName = name }

// SetApiToken sets the API token of the dbpu client.
func (c *Client) SetApiToken(token string) { c.ApiToken = token }

type Pool struct {
	databases map[string]*Database
}

// A simple key-value pair for the hash table
type Entry struct {
	key   string
	value interface{} // Can store any data type
}

// HashTable struct
type HashTable struct {
	table    []Entry // Underlying storage -  a slice of entries
	capacity int
}

// Hash function (a simple one for this example)
func hash(key string, capacity int) int {
	sum := 0
	for _, char := range key {
		sum += int(char)
	}
	return sum % capacity
}

// Create a new hash table
func NewHashTable(capacity int) *HashTable {
	return &HashTable{
		table:    make([]Entry, capacity),
		capacity: capacity,
	}
}

// Add an item to the hash table
func (h *HashTable) Add(key string, value interface{}) {
	index := hash(key, h.capacity)
	h.table[index] = Entry{key, value}
}

// Get the value associated with a key
func (h *HashTable) Get(key string) (interface{}, bool) {
	index := hash(key, h.capacity)
	entry := h.table[index]

	if entry.key == key {
		return entry.value, true
	}

	return nil, false
}

func main() {
	// Create a hash table with initial capacity of 5
	table := NewHashTable(5)

	table.Add("hello", "world")
	table.Add("name", "Alice")

	value, found := table.Get("hello")
	if found {
		fmt.Println("Value of 'hello':", value)
	}
}
