package sugar

import (
	"context"
	"testing"
)

type testResp struct {
	code int
}

type testRawCaller struct{}

func (c *testRawCaller) Do(code int, args ...interface{}) (*testResp, error) {
	return &testResp{code: code}, nil
}

type testCall struct {
	raw  *testRawCaller
	code int
}

func (c *testCall) Call(ctx context.Context) (*testResp, error) {
	return c.raw.Do(c.code)
}

func Test_Gather(t *testing.T) {
	futureCreator := NewFutureCreator[*testResp]()
	future := futureCreator.Create()
	future.Add(&testCall{raw: &testRawCaller{}, code: 200})
	future.Add(&testCall{raw: &testRawCaller{}, code: 403})
	future.Add(&testCall{raw: &testRawCaller{}, code: 401})
	responses, err := future.Gather(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}
	if responses[0].Response.code != 200 || responses[1].Response.code != 403 || responses[2].Response.code != 401 {
		t.Fatalf("unexpected responses: %v", responses)
	}
}
