package monoton_test

import (
	"testing"

	"github.com/mustafaturan/monoton/v2"
	"github.com/mustafaturan/monoton/v2/sequencer"
)

func BenchmarkNext(b *testing.B) {
	b.ReportAllocs()

	m, _ := monoton.New(sequencer.NewMillisecond(), 0, 0)
	for n := 0; n < b.N; n++ {
		_ = m.Next()
	}
}

func BenchmarkNextBytes(b *testing.B) {
	b.ReportAllocs()

	m, _ := monoton.New(sequencer.NewMillisecond(), 0, 0)
	for n := 0; n < b.N; n++ {
		_ = m.NextBytes()
	}
}
