package main

/////////////////////////////////////

// CONSTANTS FOR NEEDLE FORMAT ON DISK

/////////////////////////////////////

// NEEDLE: MAGICNUMBER|UUID|DATA|DATA_SIZE|CHECKSUM|SIZE

// Size of Needle's magic number
const needleMagicSize = 2

// Size of Needle's UUID
const needleIDSize = 16

// Size of Needle
const needleSize = 4

// Size of Needle's checksum
const needleChecksum = 4

// The total size of the Needle
const needleFixedOverhead = 26

// The Needle magic number literal
const needleMagicVal uint16 = 0xCAFE

/////////////////////////////////////

// CONSTANTS FOR OBJECT INDEX ON DISK

/////////////////////////////////////

// IDX: OBJECT_ID|OFFSET|SIZE

// Object ID field
const idxObjectID = 16

// Offset field
const idxOffset = 8

// Size field
const idxSize = 4

// The total size of the Index entry
const idxSizeTotal = 28

/////////////////////////////////////

// CONSTANTS FOR FILENAME CONSTRAINTS ON DISK

/////////////////////////////////////

// Data file prefix
const dataFilePrefix = "volume_"

// Index file prefix
const indexFilePrefix = "volume_"

// Data file suffix
const dataFileExt = ".dat"

// Index file suffix
const indexFileExt = ".idx"

/////////////////////////////////////

func makeFileMame(port string) (string, string) {
}