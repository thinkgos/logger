package logger

import (
	"sync"

	"go.uber.org/zap"
)

var fieldPool = &sync.Pool{
	New: func() any {
		return &fieldContainer{
			make([]zap.Field, 0, 32),
		}
	},
}

type fieldContainer struct {
	Fields []Field
}

func (c *fieldContainer) reset() *fieldContainer {
	c.Fields = c.Fields[:0]
	return c
}

// PoolGet selects an arbitrary item from the field Pool, removes it from the
// field Pool, and returns it to the caller.
// PoolGet may choose to ignore the field pool and treat it as empty.
// Callers should not assume any relation between values passed to PoolPut and
// the values returned by PoolGet.
//
// NOTE: This function should be call PoolPut to give back.
// NOTE: You should know `sync.Pool` work principle
// ```go
//
// fc := logger.PoolGet()
// defer logger.PoolPut(fc)
// fc.Fields = append(fc.Fields, logger.String("k1", "v1"))
// ... use fc.Fields
//
// ```
func PoolGet() *fieldContainer {
	c := fieldPool.Get().(*fieldContainer)
	return c.reset()
}

// PoolPut adds x to the pool.
// NOTE: See PoolGet.
func PoolPut(c *fieldContainer) {
	if c == nil {
		return
	}
	fieldPool.Put(c.reset())
}
