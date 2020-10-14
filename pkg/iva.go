package iva

import (
	"log"
	"multisourceTest/pkg/videosource"
	"sync"
)

type IVA struct {
	mu         sync.RWMutex
	vsRegistry map[string]*videosource.VS
}

func NewIVA() *IVA {
	return &IVA{vsRegistry: make(map[string]*videosource.VS)}
}

func (iva *IVA) AddSource(sourceURL string) {
	log.Printf("start source %s\n", sourceURL)
	iva.mu.Lock()
	defer iva.mu.Unlock()
	iva.vsRegistry[sourceURL] = videosource.NewVS(sourceURL)
	iva.vsRegistry[sourceURL].Process(sourceURL)
}

func (iva *IVA) RemoveSource(sourceURL string) {
	log.Printf("end source %s\n", sourceURL)
	iva.mu.Lock()
	defer iva.mu.Unlock()
	if _, ok := iva.vsRegistry[sourceURL]; ok {
		iva.vsRegistry[sourceURL].EndProcess()
		delete(iva.vsRegistry, sourceURL)
	}
}
