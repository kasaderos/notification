package workerpool

import (
	"context"
	"fmt"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type WorkerPool struct {
	workers int

	jobOrder int32
	jobs     []Job
}

type Job func(context.Context) error

func New(workers int, jobs []Job) *WorkerPool {
	return &WorkerPool{
		workers:  workers,
		jobs:     jobs,
		jobOrder: -1,
	}
}

func (wp *WorkerPool) Run(ctx context.Context) error {
	g, gctx := errgroup.WithContext(ctx)

	for i := 0; i < wp.workers; i++ {
		g.Go(func() error {
			job, ok := wp.nextJob()
			if !ok {
				return nil
			}

			return job(gctx)
		})
	}

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("wait workers: %w", err)
	}

	return nil
}

func (wp *WorkerPool) nextJob() (Job, bool) {
	idx := atomic.AddInt32(&wp.jobOrder, 1)
	if idx >= int32(len(wp.jobs)) {
		return nil, false
	}

	return wp.jobs[idx], true
}
