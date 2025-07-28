package main 

// Node types
const (
	BNODE_NODE = 1, // internal nodes with pointers
	BNODE_LEAF = 2 // leaf nodes with values
)

// Page constraints
const BTREE_PAGE_SIZE = 4096
const BTREE_MAX_KEY_SIZE = 1000
const BTREE_MAX_VAL_SIZE = 3000

// BNode represents a B+tree node as a byte slice
// Can be used for both leaf nodes (with values) and internal nodes (with child pointers)

// Binary layout:
// | type | nkeys |  pointers  |  offsets   | key-values | unused |
// |  2B  |   2B  | nkeys × 8B | nkeys × 2B |     ...    |        |
type BNode []byte

// Header section (4 bytes total)
// Bytes 0-1: node type (BNODE_NODE or BNODE_LEAF)
// Bytes 2-3: number of keys

// Pointers section starts at byte 4
// nkeys × 8 bytes: child page numbers (unused for leaf nodes)

// Offsets section starts at byte (4 + nkeys × 8)
// nkeys × 2 bytes: starting position of each key-value pair
// Note: first offset is always 0, so we store (nkeys-1) offsets

// Key-values section starts at byte (4 + nkeys × 8 + nkeys × 2)
// Variable length data:
// Each pair: | key_size(2B) | val_size(2B) | key_data | val_data |