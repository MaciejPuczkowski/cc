package sugar

import (
	"context"
	"time"
)

type Result[Response any] struct {
	RefIndex int
	Response Response
	Err      error
}
type Caller[Response any] interface {
	Call(ctx context.Context) (Response, error)
}
type Future[Response any] struct {
	calls   []Caller[Response]
	timeout time.Duration
}

func (f *Future[Response]) Add(calls ...Caller[Response]) {
	f.calls = append(f.calls, calls...)
}

func (f Future[Response]) Gather(ctx context.Context, calls ...Caller[Response]) ([]Result[Response], error) {
	if len(f.calls) > 0 {
		f.Add(calls...)
	}
	results := make(chan Result[Response])
	for i, call := range f.calls {
		go f.GatherOne(ctx, i, call, results)
	}
	return f.collect(ctx, results)
}

func (f *Future[Response]) GatherOne(ctx context.Context, i int, call Caller[Response], results chan Result[Response]) {
	resp, err := call.Call(ctx)
	results <- Result[Response]{RefIndex: i, Response: resp, Err: err}
}

func (f *Future[Response]) collect(ctx context.Context, chanRes chan Result[Response]) ([]Result[Response], error) {
	results := make([]Result[Response], len(f.calls))
	f.fillInitialResults(results)
	timer := time.NewTimer(f.timeout)
	for _, _ = range f.calls {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case result := <-chanRes:
			results[result.RefIndex] = result
		case <-timer.C:
			return results, ErrTimeout
		}
	}
	return results, nil
}

func (f *Future[Response]) fillInitialResults(results []Result[Response]) {
	for _, result := range results {
		result.Err = ErrNotCollected
	}
}

type FutureCreator[Response any] struct {
	timeout time.Duration
}

func NewFutureCreator[Response any]() *FutureCreator[Response] {
	return &FutureCreator[Response]{timeout: 10 * time.Second}
}

func (f *FutureCreator[Response]) SetTimeout(timeout time.Duration) {
	f.timeout = timeout
}

func (f *FutureCreator[Response]) Create() *Future[Response] {
	return &Future[Response]{
		calls:   make([]Caller[Response], 0),
		timeout: f.timeout,
	}
}
