package vault

import (
	"context"
	"testing"
)

func TestHasherService(t *testing.T) {
	srv := NewService()
	ctx := context.Background()
	h, err := srv.Hash(ctx, "password")
	if err != nil {
		t.Errorf("Hash: %s", err)
	}

	ok, err := srv.Validate(ctx, "password", h)
	if err != nil {
		t.Errorf("Validate: %s", err)
	}

	if !ok {
		t.Error("expected true from service")
	}
	ok, err = srv.Validate(ctx, "wrong_password", h)
	if err != nil {
		t.Errorf("Validate: %s", err)
	}
	if ok {
		t.Error("expected false from service")
	}
}
