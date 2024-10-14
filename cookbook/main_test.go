package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/k1LoW/runn"
)

func TestSample(t *testing.T) {

	ctx := context.Background()
	ts := httptest.NewServer(NewRouter())
	t.Cleanup(func() {
		ts.Close()
	})

	opts := []runn.Option{
		runn.T(t),
		runn.Runner("req", ts.URL),
	}
	o, err := runn.Load("*.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.RunN(ctx); err != nil {
		t.Fatal(err)
	}
}
