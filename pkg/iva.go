package iva

import (
	"multisourceTest/pkg/videosource"
	"sync"
)

type IVA struct {
	mu         sync.RWMutex
	vsRegistry map[string]*videosource.VS
}

func (iva *IVA) AddSource() {

}

func (iva *IVA) RemoveSource() {

}
