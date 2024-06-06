package wal

import "os"

type Options struct {
	//File Directory Path
	DirPath string
	// SegmentSize specifies the maximum size of each segment file in bytes.
	SegmentSize int64
	// When Flush Disk Logic
	// true is waits for disk flush, safe and slow
	// false is only waits for buffer cache, non-safe and fast
	DiskFlushSync bool
	// Depending on the settings, the amount of data written at one time is determined. If set too high, there is a risk of collision.
	BytesPerSync uint32
	// Split Seg File Extension
	DiskFileExtension string
	// add BlockCache
	BlockCache uint32
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

var DefaultOptions = Options{
	DirPath:           os.TempDir(),
	DiskFlushSync:     false,
	BytesPerSync:      0,
	SegmentSize:       GB,
	DiskFileExtension: ".SDF",
	BlockCache:        32 * 10 * KB,
}

//.SDF => Segment DiskFile
