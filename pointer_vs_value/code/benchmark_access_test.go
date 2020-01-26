package code

import "testing"

func BenchmarkGigaObject_Calc(b *testing.B) {

	o := createGiga()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcGiga(o))
	}
}

func BenchmarkGigaPObject_Calc(b *testing.B) {

	o := createGiga()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcGigaP(&o))
	}
}

func BenchmarkGigaIfcObject_Calc(b *testing.B) {

	o := createGiga()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcGigaWithInterface(&o))
	}
}
