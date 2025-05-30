package admin

import "sync/atomic"

type metrics struct {
	FileserverHits atomic.Int32
}
