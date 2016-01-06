package ntile_test

import (
	"testing"

	"github.com/suonlight/ntile"
)

func TestNtile(t *testing.T) {
	// empty element with 3 tiles
	var xs = []interface{}{}
	var ntiles = ntile.Ntile(xs, 3)

	if len(ntiles[0]) != 0 || len(ntiles[1]) != 0 || len(ntiles[2]) != 0 {
		t.Error()
	}

	// 1 element with 3 tiles
	xs = []interface{}{4}
	ntiles = ntile.Ntile(xs, 3)
	if len(ntiles[0]) != 1 || len(ntiles[1]) != 0 || len(ntiles[2]) != 0 {
		t.Error()
	}

	if ntiles[0][0] != 4 {
		t.Error()
	}

	// 2 elements with 3 tiles
	xs = []interface{}{4, 23}
	ntiles = ntile.Ntile(xs, 3)
	if len(ntiles[0]) != 1 || len(ntiles[1]) != 1 || len(ntiles[2]) != 0 {
		t.Error()
	}

	if ntiles[0][0] != 4 || ntiles[1][0] != 23 {
		t.Error()
	}

	// 3 elements with 3 tiles
	xs = []interface{}{4, 23, 54}
	ntiles = ntile.Ntile(xs, 3)

	if len(ntiles[0]) != 1 || len(ntiles[1]) != 1 || len(ntiles[2]) != 1 {
		t.Error()
	}

	if ntiles[0][0] != 4 || ntiles[1][0] != 23 || ntiles[2][0] != 54 {
		t.Error()
	}

	// 4 elements with 3 tiles
	xs = []interface{}{4, 23, 54, 72}
	ntiles = ntile.Ntile(xs, 3)

	if ntiles[0][0] != 4 || ntiles[0][1] != 23 {
		t.Error()
	}

	if ntiles[1][0] != 54 || ntiles[2][0] != 72 {
		t.Error()
	}

	// 5 elements with 3 tiles
	xs = []interface{}{4, 23, 54, 72, 82}
	ntiles = ntile.Ntile(xs, 3)

	if ntiles[0][0] != 4 || ntiles[0][1] != 23 {
		t.Error()
	}

	if ntiles[1][0] != 54 || ntiles[1][1] != 72 {
		t.Error()
	}

	if ntiles[2][0] != 82 {
		t.Error()
	}

	// 6 elements with 3 tiles
	xs = []interface{}{4, 23, 54, 72, 82, 90}
	ntiles = ntile.Ntile(xs, 3)

	if ntiles[0][0] != 4 || ntiles[0][1] != 23 {
		t.Error()
	}

	if ntiles[1][0] != 54 || ntiles[1][1] != 72 {
		t.Error()
	}

	if ntiles[2][0] != 82 || ntiles[2][1] != 90 {
		t.Error()
	}

	// 7 elements with 3 tiles
	xs = []interface{}{4, 23, 54, 72, 82, 90, 102}
	ntiles = ntile.Ntile(xs, 3)

	if ntiles[0][0] != 4 || ntiles[0][1] != 23 || ntiles[0][2] != 54 {
		t.Error()
	}

	if ntiles[1][0] != 72 || ntiles[1][1] != 82 {
		t.Error()
	}

	if ntiles[2][0] != 90 || ntiles[2][1] != 102 {
		t.Error()
	}
}

func BenchmarkNtileAt(b *testing.B) {
	var xs = make([]interface{}, b.N)
	for i := 0; i < b.N; i++ {
		xs[i] = i
	}
	ntile.NtileAt(xs, 5000, 5000)
}
