package gate_test

import (
	"gate"
	"testing"
)

func TestAND(t *testing.T) {

	t.Parallel()

	type testCase struct {
		a    bool
		b    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, b: true, want: true, desc: "a=true, b=true"},
		{a: false, b: false, want: false, desc: "a=0, b=0"},
		{a: false, b: true, want: false, desc: "a=0, b=true"},
		{a: true, b: false, want: false, desc: "a=true, b=0"},
	}

	for _, tc := range tcs {
		got := gate.AND(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("want: %v, got: %v", tc.want, got)
		}

	}
}

func TestOR(t *testing.T) {

	t.Parallel()

	type testCase struct {
		a    bool
		b    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, b: true, want: false, desc: "a=true, b=true"},
		{a: false, b: false, want: false, desc: "a=false, b=false"},
		{a: false, b: true, want: true, desc: "a=false, b=true"},
		{a: true, b: false, want: true, desc: "a=true, b=0"},
	}

	for _, tc := range tcs {
		got := gate.OR(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("test: %s want: %v, got: %v", tc.desc, tc.want, got)
		}

	}
}

func TestNOT(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, want: false, desc: "a=true"},
		{a: false, want: true, desc: "a=false"},
	}

	for _, tc := range tcs {
		got := gate.NOT(tc.a)

		if tc.want != got {
			t.Fatalf("test: %s want: %v, got: %v", tc.desc, tc.want, got)
		}

	}
}

func TestNAND(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    bool
		b    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, b: true, want: false, desc: "a=true, b=true"},
		{a: false, b: false, want: true, desc: "a=0, b=0"},
		{a: false, b: true, want: true, desc: "a=0, b=true"},
		{a: true, b: false, want: true, desc: "a=true, b=0"},
	}

	for _, tc := range tcs {
		got := gate.NAND(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("test: %s want: %v, got: %v", tc.desc, tc.want, got)
		}

	}
}

func TestNOR(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    bool
		b    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, b: true, want: true, desc: "a=true, b=true"},
		{a: false, b: false, want: true, desc: "a=0, b=0"},
		{a: false, b: true, want: false, desc: "a=0, b=true"},
		{a: true, b: false, want: false, desc: "a=true, b=0"},
	}

	for _, tc := range tcs {
		got := gate.NOR(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("test: %s want: %v, got: %v", tc.desc, tc.want, got)
		}

	}
}

func TestXOR(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    bool
		b    bool
		want bool
		desc string
	}

	tcs := []testCase{
		{a: true, b: true, want: false, desc: "a=true, b=true"},
		{a: false, b: false, want: false, desc: "a=0, b=0"},
		{a: false, b: true, want: true, desc: "a=0, b=true"},
		{a: true, b: false, want: true, desc: "a=true, b=0"},
	}

	for _, tc := range tcs {
		got := gate.XOR(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("test: %s want: %v, got: %v", tc.desc, tc.want, got)
		}

	}
}
