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

type Executor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

type Queries[T Executor] struct {
	db T
}

func New[T Executor](db T) *Queries[T] {
	return &Queries[T]{db: db}
}

func (q *Queries[T]) WithTx(tx *sql.Tx) *Queries[*sql.Tx] {
	return &Queries[*sql.Tx]{db: tx}
}

type HashTable[K comparable, V any] struct {
	table    map[K]V
	capacity int
}

func NewHashTable[K comparable, V any](capacity int) *HashTable[K, V] {
	return &HashTable[K, V]{
		table:    make(map[K]V, capacity),
		capacity: capacity,
	}
}

func (h *HashTable[K, V]) Set(key K, value V) {
	h.table[key] = value
}

func (h *HashTable[K, V]) Get(key K) (V, bool) {
	val, ok := h.table[key]
	return val, ok
}

func (h *HashTable[K, V]) Delete(key K) {
	delete(h.table, key)
}

func (q *Queries[T]) Hash() string {
	// Implementation depends on what makes each query unique in your context
	return uuid.New().String()
}
