package code

import (
	"testing"
)

const iterations = 100000000

func BenchmarkMicro_Iter(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMicro(o))
		}
	}

}

func BenchmarkMicroP_Iter(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMicroP(&o))
		}
	}

}

func BenchmarkMicroIfc_Iter(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcIfc(&o))
		}
	}

}

func BenchmarkTiny_Iter(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcTiny(o))
		}
	}

}

func BenchmarkTinyP_Iter(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcTinyP(&o))
		}
	}

}

func BenchmarkTinyIfc_Iter(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcIfc(&o))
		}
	}

}

func BenchmarkMini_Iter(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMini(o))
		}
	}

}

func BenchmarkMiniP_Iter(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMiniP(&o))
		}
	}

}

func BenchmarkMiniIfc_Iter(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcIfc(&o))
		}
	}

}

func BenchmarkMega_Iter(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMega(o))
		}
	}

}

func BenchmarkMegaP_Iter(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcMegaP(&o))
		}
	}

}

func BenchmarkMegaIfc_Iter(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcIfc(&o))
		}
	}

}

func BenchmarkTera_Iter(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcTera(o))
		}
	}

}

func BenchmarkTeraP_Iter(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcTeraP(&o))
		}
	}

}

func BenchmarkTeraIfc_Iter(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		for a := 0; a < iterations; a++ {
			avoidCompilerOptimisation(calcIfc(&o))
		}
	}

}
