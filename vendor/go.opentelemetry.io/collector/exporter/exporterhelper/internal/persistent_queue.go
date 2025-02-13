// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "go.opentelemetry.io/collector/exporter/exporterhelper/internal"

import (
	"context"
	"fmt"
	"sync"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension/experimental/storage"
)

// Monkey patching for unit test
var (
	stopStorage = func(queue *persistentQueue) {
		queue.storage.stop()
	}
)

// persistentQueue holds the queue backed by file storage
type persistentQueue struct {
	stopWG   sync.WaitGroup
	stopOnce sync.Once
	stopChan chan struct{}
	storage  *persistentContiguousStorage
}

// buildPersistentStorageName returns a name that is constructed out of queue name and signal type. This is done
// to avoid conflicts between different signals, which require unique persistent storage name
func buildPersistentStorageName(name string, signal component.DataType) string {
	return fmt.Sprintf("%s-%s", name, signal)
}

type PersistentQueueSettings struct {
	Name        string
	Signal      component.DataType
	Capacity    uint64
	Logger      *zap.Logger
	Client      storage.Client
	Unmarshaler RequestUnmarshaler
	Marshaler   RequestMarshaler
}

// NewPersistentQueue creates a new queue backed by file storage; name and signal must be a unique combination that identifies the queue storage
func NewPersistentQueue(ctx context.Context, params PersistentQueueSettings) ProducerConsumerQueue {
	return &persistentQueue{
		stopChan: make(chan struct{}),
		storage:  newPersistentContiguousStorage(ctx, buildPersistentStorageName(params.Name, params.Signal), params),
	}
}

// StartConsumers starts the given number of consumers which will be consuming items
func (pq *persistentQueue) StartConsumers(numWorkers int, callback func(item Request)) {
	for i := 0; i < numWorkers; i++ {
		pq.stopWG.Add(1)
		go func() {
			defer pq.stopWG.Done()
			for {
				select {
				case req := <-pq.storage.get():
					callback(req)
				case <-pq.stopChan:
					return
				}
			}
		}()
	}
}

// Produce adds an item to the queue and returns true if it was accepted
func (pq *persistentQueue) Produce(item Request) bool {
	err := pq.storage.put(item)
	return err == nil
}

// Stop stops accepting items, shuts down the queue and closes the persistent queue
func (pq *persistentQueue) Stop() {
	pq.stopOnce.Do(func() {
		// stop the consumers before the storage or the successful processing result will fail to write to persistent storage
		close(pq.stopChan)
		pq.stopWG.Wait()
		stopStorage(pq)
	})
}

// Size returns the current depth of the queue, excluding the item already in the storage channel (if any)
func (pq *persistentQueue) Size() int {
	return int(pq.storage.size())
}
