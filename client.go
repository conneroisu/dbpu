package dbpu

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
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
//
// Base URL is the base URL for API requests.
// Region URL is the base URL for region requests.
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

// Executor is an interface for executing SQL queries.
type Executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Queries is a set of queries.
type Queries[T Executor] struct {
	db T
}

// New returns a new Queries.
func New[T Executor](db T) *Queries[T] {
	return &Queries[T]{db: db}
}

// WithTx returns a new Queries with the transaction.
func (q *Queries[T]) WithTx(tx *sql.Tx) *Queries[*sql.Tx] {
	return &Queries[*sql.Tx]{db: tx}
}

// HashTable is a hash table for storing databases.
// The key is the unique hash of the database in the table.
// The value is the resulting database at the hash.
type HashTable[K comparable, V any] struct {
	table    map[K]V
	capacity int
}

// NewHashTable returns a new hash table.
// The capacity is the initial capacity of the hash table.
// However, the capacity will grow as needed.
func NewHashTable[K comparable, V any](capacity int) *HashTable[K, V] {
	return &HashTable[K, V]{
		table:    make(map[K]V, capacity),
		capacity: capacity,
	}
}

// Set sets the value for the key.
// If the key is already in the hash table, the value is updated.
func (h *HashTable[K, V]) Set(key K, value V) {
	h.table[key] = value
}

// Get returns the value for the key.
// If the key is not in the hash table, the second return value is false.
func (h *HashTable[K, V]) Get(key K) (V, bool) {
	val, ok := h.table[key]
	return val, ok
}

// Delete removes the key from the hash table.
// If the key is not in the hash table, this is a no-op.
func (h *HashTable[K, V]) Delete(key K) {
	delete(h.table, key)
}

// Hash returns a unique hash for the query.
// This is used to identify the query in the table.
func (q *Queries[T]) Hash() string {
	return uuid.New().String()
}

// GetAll returns all the databases in the hash table.
// The returned slice is a copy of the databases in the hash table.
func (h *HashTable[K, V]) GetAll() []V {
	var dbs []V
	for _, db := range h.table {
		dbs = append(dbs, db)
	}
	return dbs
}
