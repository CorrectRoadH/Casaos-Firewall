package model

import (
	"time"
)

type Obj interface {
	GetSize() int64
	GetName() string
	ModTime() time.Time
	IsDir() bool

	// The internal information of the driver.
	// If you want to use it, please understand what it means
	GetID() string
	GetPath() string
}
