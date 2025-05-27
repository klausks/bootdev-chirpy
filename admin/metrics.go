package admin

import "sync/atomic"

type metrics struct {
	fileserverHits atomic.Int32
}