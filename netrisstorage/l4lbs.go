/*
Copyright 2021. Netris, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package netrisstorage

import (
	"sync"

	"github.com/netrisai/netriswebapi/v1/types/l4lb"
)

// L4LBStorage .
type L4LBStorage struct {
	sync.Mutex
	L4LBs []*l4lb.LoadBalancer
}

// NewL4LBStorage .
func NewL4LBStorage() *L4LBStorage {
	return &L4LBStorage{}
}

// GetAll .
func (p *L4LBStorage) GetAll() []*l4lb.LoadBalancer {
	p.Lock()
	defer p.Unlock()
	return p.getAll()
}

func (p *L4LBStorage) getAll() []*l4lb.LoadBalancer {
	return p.L4LBs
}

func (p *L4LBStorage) storeAll(items []*l4lb.LoadBalancer) {
	p.L4LBs = items
}

// FindByName .
func (p *L4LBStorage) FindByName(name string) (*l4lb.LoadBalancer, bool) {
	p.Lock()
	defer p.Unlock()
	return p.findByName(name)
}

func (p *L4LBStorage) findByName(name string) (*l4lb.LoadBalancer, bool) {
	for _, item := range p.L4LBs {
		if item.Name == name {
			return item, true
		}
	}
	return nil, false
}

// FindByID .
func (p *L4LBStorage) FindByID(id int) (*l4lb.LoadBalancer, bool) {
	p.Lock()
	defer p.Unlock()
	item, ok := p.findByID(id)
	if !ok {
		_ = p.download()
		return p.findByID(id)
	}
	return item, ok
}

func (p *L4LBStorage) findByID(id int) (*l4lb.LoadBalancer, bool) {
	for _, item := range p.L4LBs {
		if item.ID == id {
			return item, true
		}
	}
	return nil, false
}

// Download .
func (p *L4LBStorage) download() error {
	items, err := Cred.L4LB().Get()
	if err != nil {
		return err
	}
	p.storeAll(items)
	return nil
}

// Download .
func (p *L4LBStorage) Download() error {
	p.Lock()
	defer p.Unlock()
	return p.download()
}
