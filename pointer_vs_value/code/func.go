package code

import "errors"

// methods w/o pointer arguments

func calcMicro(o micro) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

func calcTiny(o tiny) (float64, error) { // pointer_vs_value/pointer_receiver/func.go:98:21: calcTiny o does not escape
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'") // pointer_vs_value/pointer_receiver/func.go:100:25: error(&errors.errorString literal) escapes to heap
	} // pointer_vs_value/pointer_receiver/func.go:100:25: &errors.errorString literal escapes to heap
	return float64(o.a) / o.aa, nil
}

func calcMini(o mini) (float64, error) { // pointer_vs_value/pointer_receiver/func.go:105:21: calcMini o does not escape
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'") // pointer_vs_value/pointer_receiver/func.go:107:25: error(&errors.errorString literal) escapes to heap
	} // pointer_vs_value/pointer_receiver/func.go:107:25: &errors.errorString literal escapes to heap
	return float64(o.a) / o.aa, nil
}

func calcMega(o mega) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

func calcTera(o tera) (float64, error) { // calcTera o does not escape
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'") // error(&errors.errorString literal) escapes to heap
	} // &errors.errorString literal escapes to heap
	return float64(o.a) / o.aa, nil
}

// methods with pointer arguments

func calcIfc(o obj) (float64, error) {
	if o.GetF() == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.GetI()) / o.GetF(), nil
}

func calcMicroP(o *micro) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

func calcTinyP(o *tiny) (float64, error) { // calcTinyP o does not escape
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'") // error(&errors.errorString literal) escapes to heap
	} // &errors.errorString literal escapes to heap
	return float64(o.a) / o.aa, nil
}

func calcMiniP(o *mini) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

func calcMegaP(o *mega) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

func calcTeraP(o *tera) (float64, error) {
	if o.aa == 0.0 {
		return 0.0, errors.New("cannot divide by '0.0'")
	}
	return float64(o.a) / o.aa, nil
}

// access ALL object properties
func calcGiga(o giga) (float64, error) { // calcTera o does not escape
	// we skip error checking here for simplicity
	return float64(
		o.a+o.b+o.c+o.d+o.e+o.f+o.g+o.h+o.i+
			o.j+o.k+o.l+o.m+o.n+o.o+o.p+o.q+
			o.r+o.s+o.t+o.u+o.v+o.y+o.x+o.w+o.z) /
		(o.aa + o.bb + o.cc + o.dd + o.ee + o.ff + o.gg + o.hh + o.ii +
			o.jj + o.kk + o.ll + o.mm + o.nn + o.oo + o.pp + o.qq +
			o.rr + o.ss + o.tt + o.uu + o.vv + o.yy + o.xx + o.ww + o.zz), nil
}

func calcGigaP(o *giga) (float64, error) {
	// we skip error checking here for simplicity
	return float64(
		o.a+o.b+o.c+o.d+o.e+o.f+o.g+o.h+o.i+
			o.j+o.k+o.l+o.m+o.n+o.o+o.p+o.q+
			o.r+o.s+o.t+o.u+o.v+o.y+o.x+o.w+o.z) /
		(o.aa + o.bb + o.cc + o.dd + o.ee + o.ff + o.gg + o.hh + o.ii +
			o.jj + o.kk + o.ll + o.mm + o.nn + o.oo + o.pp + o.qq +
			o.rr + o.ss + o.tt + o.uu + o.vv + o.yy + o.xx + o.ww + o.zz), nil
}

func calcGigaWithInterface(o gigantic) (float64, error) {
	// we skip error checking here for simplicity
	return float64(
		o.get_a()+o.get_b()+o.get_c()+o.get_d()+o.get_e()+o.get_f()+o.get_g()+o.get_h()+o.get_i()+
			o.get_j()+o.get_k()+o.get_l()+o.get_m()+o.get_n()+o.get_o()+o.get_p()+o.get_q()+
			o.get_r()+o.get_s()+o.get_t()+o.get_u()+o.get_v()+o.get_y()+o.get_x()+o.get_w()+o.get_z()) /
		(o.get_aa() + o.get_bb() + o.get_cc() + o.get_dd() + o.get_ee() + o.get_ff() + o.get_gg() + o.get_hh() + o.get_ii() +
			o.get_jj() + o.get_kk() + o.get_ll() + o.get_mm() + o.get_nn() + o.get_oo() + o.get_pp() + o.get_qq() +
			o.get_rr() + o.get_ss() + o.get_tt() + o.get_uu() + o.get_vv() + o.get_yy() + o.get_xx() + o.get_ww() + o.get_zz()), nil
}
