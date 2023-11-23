package database

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync"

	"github.com/acerquetti/ports/domain"
)

type memoryDB struct {
	Ports map[domain.ID]*domain.Port
	mutex sync.Mutex
}

func NewMemoryDB(portsReader io.Reader) (*memoryDB, error) {
	ports, err := readPorts(portsReader)
	if err != nil {
		return nil, err
	}

	return &memoryDB{
		Ports: ports,
	}, nil
}

func (m *memoryDB) Save(port domain.Port) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Ports[port.ID] = &port

	return nil
}

func readPorts(portsReader io.Reader) (ports map[domain.ID]*domain.Port, err error) {
	dec := json.NewDecoder(portsReader)
	ports = make(map[domain.ID]*domain.Port)

	// read open bracket
	if _, err = dec.Token(); err != nil {
		return
	}

	// while the object contains keys
	for dec.More() {
		var t json.Token

		// read port key
		t, err = dec.Token()
		if err != nil {
			return
		}

		key, ok := t.(string)
		if !ok {
			return nil, fmt.Errorf("couldn't load key from token %v", t)
		}

		// read port data
		if dec.More() {
			portRaw := domain.PortRaw{
				ID: key,
			}

			if err := dec.Decode(&portRaw); err != nil {
				return nil, err
			}

			port, err := domain.NewPortFromRaw(portRaw)
			// if data does not follow domain constraints, just log which one and keep processing
			if err != nil {
				log.Printf("couldn't create port with key %v: %v", key, err)
			} else {
				ports[port.ID] = port
			}
		}
	}

	// read closing bracket
	if _, err = dec.Token(); err != nil {
		return nil, err
	}

	return
}
