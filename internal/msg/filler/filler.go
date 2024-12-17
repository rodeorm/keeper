package filler

import (
	"github.com/rodeorm/keeper/internal/core"
)

// NewFiller создает новый Filler
// Каждый Filler может наполнять очередь
func NewFiller(queue *core.Queue, storage core.MessageStorager, prd int) *Filler {
	return &Filler{
		queue:          queue,
		messageStorage: storage,
		period:         prd,
	}
}
