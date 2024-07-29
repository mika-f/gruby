package mruby_test

import (
	"testing"

	mruby "github.com/zhulik/gruby"
)

func TestEnableDisableGC(t *testing.T) {
	t.Parallel()

	mrb := mruby.NewMrb()
	defer mrb.Close()

	mrb.FullGC()
	mrb.DisableGC()

	_, err := mrb.LoadString("b = []; a = []; a = []")
	if err != nil {
		t.Fatal(err)
	}

	orig := mrb.LiveObjectCount()
	mrb.FullGC()

	if orig != mrb.LiveObjectCount() {
		t.Fatalf("Object count was not what was expected after full GC: %d %d", orig, mrb.LiveObjectCount())
	}

	mrb.EnableGC()
	mrb.FullGC()

	if orig-1 != mrb.LiveObjectCount() {
		t.Fatalf("Object count was not what was expected after full GC: %d %d", orig-2, mrb.LiveObjectCount())
	}
}

func TestIsDead(t *testing.T) {
	t.Parallel()

	mrb := mruby.NewMrb()

	val, err := mrb.LoadString("$a = []")
	if err != nil {
		t.Fatal(err)
	}

	if val.IsDead() {
		t.Fatal("Value is already dead and should not be")
	}

	mrb.Close()

	if !val.IsDead() {
		t.Fatal("Value should be dead and is not")
	}
}
