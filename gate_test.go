package gate_test

import (
	"gate"
	"testing"
)

func TestANDOutputIsTrueWhenInputAEqualsTrueAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := true

	inputA := true
	inputB := true

	got := gate.AND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestANDOutputIsFalseWhenInputAEqualsFalseAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := false

	inputA := false
	inputB := false

	got := gate.AND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestANDOutputIsFalseWhenInputAEqualsFalseAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := false
	inputB := true

	got := gate.AND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestANDOutputIsFalseWhenInputAEqualsTrueAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true
	inputB := false

	got := gate.AND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestOROutputIsFalseWhenInputAAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := true

	inputA := true
	inputB := true

	got := gate.OR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestOROutputIsTrueWhenInputAAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := false

	inputA := false
	inputB := false

	got := gate.OR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestOROutputIsTrueWhenInputAEqualsFalseAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false
	inputB := true

	got := gate.OR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestOROutputIsTrueWhenInputAEqualsTrueAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := true
	inputB := false

	got := gate.OR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOTOutputIsFalseWhenInputAEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true

	got := gate.NOT(inputA)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOTOutputIsTrueWhenInputAEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false

	got := gate.NOT(inputA)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNANDOutputIsFalseWhenInputAEqualsTrueAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true
	inputB := true

	got := gate.NAND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNANDOutputIsTrueWhenInputAEqualsTrueAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := true
	inputB := false

	got := gate.NAND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNANDOutputIsTrueWhenInputAEqualsFalseAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false
	inputB := true

	got := gate.NAND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNANDOutputIsTrueWhenInputAEqualsFalseAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false
	inputB := false

	got := gate.NAND(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOROutputIsTrueWhenInputAEqualsFalseAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false
	inputB := false

	got := gate.NOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOROutputIsFalseWhenInputAEqualsTrueAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true
	inputB := true

	got := gate.NOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOROutputIsFalseWhenInputAEqualsTrueAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true
	inputB := false

	got := gate.NOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestNOROutputIsFalseWhenInputAEqualsFalseAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := false
	inputB := true

	got := gate.NOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestXOROutputIsFalseWhenInputAEqualsFalseAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := false

	inputA := false
	inputB := false

	got := gate.XOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestXOROutputIsFalseWhenInputAEqualsTrueAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := false

	inputA := true
	inputB := true

	got := gate.XOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestXOROutputIsTrueWhenInputAEqualsTrueAndInputBEqualsFalse(t *testing.T) {

	t.Parallel()

	want := true

	inputA := true
	inputB := false

	got := gate.XOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}

func TestXOROutputIsTrueWhenInputAEqualsFalseAndInputBEqualsTrue(t *testing.T) {

	t.Parallel()

	want := true

	inputA := false
	inputB := true

	got := gate.XOR(inputA, inputB)

	if want != got {
		t.Fatalf("want: %v, got: %v", want, got)
	}

}
