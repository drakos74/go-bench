package code

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkMicro_Calc(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMicro(o))
	}
}

func BenchmarkMicroP_Calc(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMicroP(&o))
	}
}

func BenchmarkMicroIfc_Calc(b *testing.B) {

	o := createMicro()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcIfc(&o))
	}
}

func BenchmarkTiny_Calc(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcTiny(o))
	}
}

func BenchmarkTinyP_Calc(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcTinyP(&o))
	}
}

func BenchmarkTinyIfc_Calc(b *testing.B) {

	o := createTiny()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcIfc(&o))
	}
}

func BenchmarkMini_Calc(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMini(o))
	}
}

func BenchmarkMiniP_Calc(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMiniP(&o))
	}
}

func BenchmarkMiniIfc_Calc(b *testing.B) {

	o := createMini()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcIfc(&o))
	}
}

func BenchmarkMega_Calc(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMega(o))
	}
}

func BenchmarkMegaP_Calc(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcMegaP(&o))
	}
}

func BenchmarkMegaIfc_Calc(b *testing.B) {

	o := createMega()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcIfc(&o))
	}
}

func BenchmarkTera_Calc(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcTera(o))
	}
}

func BenchmarkTeraP_Calc(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcTeraP(&o))
	}
}

func BenchmarkTeraIfc_Calc(b *testing.B) {

	o := createTera()

	for i := 0; i < b.N; i++ {
		avoidCompilerOptimisation(calcIfc(&o))
	}
}

// avoid compiler optimisations

var F = 0.0

func avoidCompilerOptimisation(f float64, err error) {
	if err != nil {
		log.Fatal("Error")
	}
	F += f
}

// object creation

func createMicro() micro {
	return micro{
		a:  rand.Int63(),
		aa: rand.Float64(),
	}
}

func createTiny() tiny {
	return tiny{
		a:   rand.Int63(),
		aa:  rand.Float64(),
		aaa: strconv.Itoa(rand.Int()),
	}
}

func createMini() mini {
	return mini{
		a:   rand.Int63(),
		b:   rand.Int63(),
		c:   rand.Int63(),
		aa:  rand.Float64(),
		bb:  rand.Float64(),
		cc:  rand.Float64(),
		aaa: strconv.Itoa(rand.Int()),
		bbb: strconv.Itoa(rand.Int()),
		ccc: strconv.Itoa(rand.Int()),
	}
}

func createMega() mega {
	return mega{
		a:   rand.Int63(),
		b:   rand.Int63(),
		c:   rand.Int63(),
		d:   rand.Int63(),
		e:   rand.Int63(),
		f:   rand.Int63(),
		g:   rand.Int63(),
		h:   rand.Int63(),
		i:   rand.Int63(),
		j:   rand.Int63(),
		k:   rand.Int63(),
		aa:  rand.Float64(),
		bb:  rand.Float64(),
		cc:  rand.Float64(),
		dd:  rand.Float64(),
		ee:  rand.Float64(),
		ff:  rand.Float64(),
		gg:  rand.Float64(),
		hh:  rand.Float64(),
		ii:  rand.Float64(),
		jj:  rand.Float64(),
		kk:  rand.Float64(),
		aaa: strconv.Itoa(rand.Int()),
		bbb: strconv.Itoa(rand.Int()),
		ccc: strconv.Itoa(rand.Int()),
		ddd: strconv.Itoa(rand.Int()),
		eee: strconv.Itoa(rand.Int()),
		fff: strconv.Itoa(rand.Int()),
		ggg: strconv.Itoa(rand.Int()),
		hhh: strconv.Itoa(rand.Int()),
		iii: strconv.Itoa(rand.Int()),
		jjj: strconv.Itoa(rand.Int()),
		kkk: strconv.Itoa(rand.Int()),
	}
}

func createTera() tera {
	return tera{
		a:   rand.Int63(),
		b:   rand.Int63(),
		c:   rand.Int63(),
		d:   rand.Int63(),
		e:   rand.Int63(),
		f:   rand.Int63(),
		g:   rand.Int63(),
		h:   rand.Int63(),
		i:   rand.Int63(),
		j:   rand.Int63(),
		k:   rand.Int63(),
		l:   rand.Int63(),
		m:   rand.Int63(),
		n:   rand.Int63(),
		o:   rand.Int63(),
		p:   rand.Int63(),
		q:   rand.Int63(),
		r:   rand.Int63(),
		s:   rand.Int63(),
		t:   rand.Int63(),
		u:   rand.Int63(),
		v:   rand.Int63(),
		x:   rand.Int63(),
		y:   rand.Int63(),
		w:   rand.Int63(),
		z:   rand.Int63(),
		aa:  rand.Float64(),
		bb:  rand.Float64(),
		cc:  rand.Float64(),
		dd:  rand.Float64(),
		ee:  rand.Float64(),
		ff:  rand.Float64(),
		gg:  rand.Float64(),
		hh:  rand.Float64(),
		ii:  rand.Float64(),
		jj:  rand.Float64(),
		kk:  rand.Float64(),
		ll:  rand.Float64(),
		mm:  rand.Float64(),
		nn:  rand.Float64(),
		oo:  rand.Float64(),
		pp:  rand.Float64(),
		qq:  rand.Float64(),
		rr:  rand.Float64(),
		ss:  rand.Float64(),
		tt:  rand.Float64(),
		uu:  rand.Float64(),
		vv:  rand.Float64(),
		xx:  rand.Float64(),
		yy:  rand.Float64(),
		ww:  rand.Float64(),
		zz:  rand.Float64(),
		aaa: strconv.Itoa(rand.Int()),
		bbb: strconv.Itoa(rand.Int()),
		ccc: strconv.Itoa(rand.Int()),
		ddd: strconv.Itoa(rand.Int()),
		eee: strconv.Itoa(rand.Int()),
		fff: strconv.Itoa(rand.Int()),
		ggg: strconv.Itoa(rand.Int()),
		hhh: strconv.Itoa(rand.Int()),
		iii: strconv.Itoa(rand.Int()),
		jjj: strconv.Itoa(rand.Int()),
		kkk: strconv.Itoa(rand.Int()),
		lll: strconv.Itoa(rand.Int()),
		mmm: strconv.Itoa(rand.Int()),
		nnn: strconv.Itoa(rand.Int()),
		ooo: strconv.Itoa(rand.Int()),
		ppp: strconv.Itoa(rand.Int()),
		qqq: strconv.Itoa(rand.Int()),
		rrr: strconv.Itoa(rand.Int()),
		sss: strconv.Itoa(rand.Int()),
		ttt: strconv.Itoa(rand.Int()),
		uuu: strconv.Itoa(rand.Int()),
		vvv: strconv.Itoa(rand.Int()),
		xxx: strconv.Itoa(rand.Int()),
		yyy: strconv.Itoa(rand.Int()),
		www: strconv.Itoa(rand.Int()),
		zzz: strconv.Itoa(rand.Int()),
	}
}

func createGiga() giga {
	return giga{
		a:  rand.Int63(),
		b:  rand.Int63(),
		c:  rand.Int63(),
		d:  rand.Int63(),
		e:  rand.Int63(),
		f:  rand.Int63(),
		g:  rand.Int63(),
		h:  rand.Int63(),
		i:  rand.Int63(),
		j:  rand.Int63(),
		k:  rand.Int63(),
		l:  rand.Int63(),
		m:  rand.Int63(),
		n:  rand.Int63(),
		o:  rand.Int63(),
		p:  rand.Int63(),
		q:  rand.Int63(),
		r:  rand.Int63(),
		s:  rand.Int63(),
		t:  rand.Int63(),
		u:  rand.Int63(),
		v:  rand.Int63(),
		x:  rand.Int63(),
		y:  rand.Int63(),
		w:  rand.Int63(),
		z:  rand.Int63(),
		aa: rand.Float64(),
		bb: rand.Float64(),
		cc: rand.Float64(),
		dd: rand.Float64(),
		ee: rand.Float64(),
		ff: rand.Float64(),
		gg: rand.Float64(),
		hh: rand.Float64(),
		ii: rand.Float64(),
		jj: rand.Float64(),
		kk: rand.Float64(),
		ll: rand.Float64(),
		mm: rand.Float64(),
		nn: rand.Float64(),
		oo: rand.Float64(),
		pp: rand.Float64(),
		qq: rand.Float64(),
		rr: rand.Float64(),
		ss: rand.Float64(),
		tt: rand.Float64(),
		uu: rand.Float64(),
		vv: rand.Float64(),
		xx: rand.Float64(),
		yy: rand.Float64(),
		ww: rand.Float64(),
		zz: rand.Float64(),
	}
}
