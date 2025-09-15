package main

import "encoding/binary"

// Node types
const (
	BNODE_NODE = 1 // internal nodes with pointers
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

// Header getters and setters
func (node BNode) btype() uint16 {
	return binary.LittleEndian.Uint16(node[0:2])
}

func (node BNode) nkeys() uint16 {
	return binary.LittleEndian.Uint16(node[2:4])
}

func (node BNode) setHeader(btype uint16, nkeys uint16) {
	binary.LittleEndian.PutUint16(node[0:2], btype)
	binary.LittleEndian.PutUint16(node[2:4], nkeys)
}

// Pointer getters and setters
func (node BNode) getPtr(idx uint16) uint64 {
	pos := 4 + 8*idx
	return binary.LittleEndian.Uint64(node[pos:])
}

func (node BNode) setPtr(idx uint16, val uint64) {
	pos := 4 + 8*idx
	binary.LittleEndian.PutUint64(node[pos:], val)
}

// Offset getters and setters
func (node BNode) getOffset(idx uint16) uint16 {
	if idx == 0 {
		return 0
	}
	pos := 4 + 8*node.nkeys() + 2*(idx-1)
	return binary.LittleEndian.Uint16(node[pos:])
}

func (node BNode) setOffset(idx uint16, offset uint16) {
	if idx == 0 {
		return
	}
	pos := 4 + 8*node.nkeys() + 2*(idx-1)
	binary.LittleEndian.PutUint16(node[pos:], offset)
}

// Key-value position and access
func (node BNode) kvPos(idx uint16) uint16 {
	kvStart := 4 + 8*node.nkeys() + 2*node.nkeys()
	return uint16(kvStart) + node.getOffset(idx)
}

func (node BNode) getKey(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos:])
	return node[pos+4:][:klen]
}

func (node BNode) getVal(idx uint16) []byte {
	pos := node.kvPos(idx)
	klen := binary.LittleEndian.Uint16(node[pos+0:])
	vlen := binary.LittleEndian.Uint16(node[pos+2:])
	return node[pos+4+klen:][:vlen]
}
