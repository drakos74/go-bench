package code

type obj interface {
	GetI() int64
	GetF() float64
}

type micro struct {
	a  int64
	aa float64
}

func (s *micro) GetI() int64 {
	return s.a
}

func (s *micro) GetF() float64 {
	return s.aa
}

type tiny struct {
	a   int64
	aa  float64
	aaa string
}

func (s *tiny) GetI() int64 {
	return s.a
}

func (s *tiny) GetF() float64 {
	return s.aa
}

type mini struct {
	a   int64
	b   int64
	c   int64
	aa  float64
	bb  float64
	cc  float64
	aaa string
	bbb string
	ccc string
}

func (s *mini) GetI() int64 {
	return s.a
}

func (s *mini) GetF() float64 {
	return s.aa
}

type mega struct {
	a   int64
	b   int64
	c   int64
	d   int64
	e   int64
	f   int64
	g   int64
	h   int64
	i   int64
	j   int64
	k   int64
	aa  float64
	bb  float64
	cc  float64
	dd  float64
	ee  float64
	ff  float64
	gg  float64
	hh  float64
	ii  float64
	jj  float64
	kk  float64
	aaa string
	bbb string
	ccc string
	ddd string
	eee string
	fff string
	ggg string
	hhh string
	iii string
	jjj string
	kkk string
}

func (s *mega) GetI() int64 {
	return s.a
}

func (s *mega) GetF() float64 {
	return s.aa
}

type giga struct {
	a  int64
	b  int64
	c  int64
	d  int64
	e  int64
	f  int64
	g  int64
	h  int64
	i  int64
	j  int64
	k  int64
	l  int64
	m  int64
	n  int64
	o  int64
	p  int64
	q  int64
	r  int64
	s  int64
	t  int64
	u  int64
	v  int64
	x  int64
	y  int64
	w  int64
	z  int64
	aa float64
	bb float64
	cc float64
	dd float64
	ee float64
	ff float64
	gg float64
	hh float64
	ii float64
	jj float64
	kk float64
	ll float64
	mm float64
	nn float64
	oo float64
	pp float64
	qq float64
	rr float64
	ss float64
	tt float64
	uu float64
	vv float64
	xx float64
	yy float64
	ww float64
	zz float64
}

type tera struct {
	a   int64
	b   int64
	c   int64
	d   int64
	e   int64
	f   int64
	g   int64
	h   int64
	i   int64
	j   int64
	k   int64
	l   int64
	m   int64
	n   int64
	o   int64
	p   int64
	q   int64
	r   int64
	s   int64
	t   int64
	u   int64
	v   int64
	x   int64
	y   int64
	w   int64
	z   int64
	aa  float64
	bb  float64
	cc  float64
	dd  float64
	ee  float64
	ff  float64
	gg  float64
	hh  float64
	ii  float64
	jj  float64
	kk  float64
	ll  float64
	mm  float64
	nn  float64
	oo  float64
	pp  float64
	qq  float64
	rr  float64
	ss  float64
	tt  float64
	uu  float64
	vv  float64
	xx  float64
	yy  float64
	ww  float64
	zz  float64
	aaa string
	bbb string
	ccc string
	ddd string
	eee string
	fff string
	ggg string
	hhh string
	iii string
	jjj string
	kkk string
	lll string
	mmm string
	nnn string
	ooo string
	ppp string
	qqq string
	rrr string
	sss string
	ttt string
	uuu string
	vvv string
	xxx string
	yyy string
	www string
	zzz string
}

func (s *tera) GetI() int64 {
	return s.a
}

func (s *tera) GetF() float64 {
	return s.aa
}

type gigantic interface {
	get_a() int64
	get_b() int64
	get_c() int64
	get_d() int64
	get_e() int64
	get_f() int64
	get_g() int64
	get_h() int64
	get_i() int64
	get_j() int64
	get_k() int64
	get_l() int64
	get_m() int64
	get_n() int64
	get_o() int64
	get_p() int64
	get_q() int64
	get_r() int64
	get_s() int64
	get_t() int64
	get_u() int64
	get_v() int64
	get_x() int64
	get_y() int64
	get_w() int64
	get_z() int64
	get_aa() float64
	get_bb() float64
	get_cc() float64
	get_dd() float64
	get_ee() float64
	get_ff() float64
	get_gg() float64
	get_hh() float64
	get_ii() float64
	get_jj() float64
	get_kk() float64
	get_ll() float64
	get_mm() float64
	get_nn() float64
	get_oo() float64
	get_pp() float64
	get_qq() float64
	get_rr() float64
	get_ss() float64
	get_tt() float64
	get_uu() float64
	get_vv() float64
	get_xx() float64
	get_yy() float64
	get_ww() float64
	get_zz() float64
}

func (o *giga) get_a() int64 {
	return o.a
}

func (o *giga) get_b() int64 {
	return o.a
}

func (o *giga) get_c() int64 {
	return o.a
}

func (o *giga) get_d() int64 {
	return o.a
}

func (o *giga) get_e() int64 {
	return o.a
}

func (o *giga) get_f() int64 {
	return o.a
}

func (o *giga) get_g() int64 {
	return o.a
}

func (o *giga) get_h() int64 {
	return o.a
}

func (o *giga) get_i() int64 {
	return o.a
}

func (o *giga) get_j() int64 {
	return o.a
}

func (o *giga) get_k() int64 {
	return o.a
}

func (o *giga) get_l() int64 {
	return o.a
}

func (o *giga) get_m() int64 {
	return o.a
}

func (o *giga) get_n() int64 {
	return o.a
}

func (o *giga) get_o() int64 {
	return o.a
}

func (o *giga) get_p() int64 {
	return o.a
}

func (o *giga) get_q() int64 {
	return o.a
}

func (o *giga) get_r() int64 {
	return o.a
}

func (o *giga) get_s() int64 {
	return o.a
}

func (o *giga) get_t() int64 {
	return o.a
}

func (o *giga) get_u() int64 {
	return o.a
}

func (o *giga) get_v() int64 {
	return o.a
}

func (o *giga) get_x() int64 {
	return o.a
}

func (o *giga) get_y() int64 {
	return o.a
}

func (o *giga) get_w() int64 {
	return o.a
}

func (o *giga) get_z() int64 {
	return o.a
}

func (o *giga) get_aa() float64 {
	return o.aa
}

func (o *giga) get_bb() float64 {
	return o.bb
}

func (o *giga) get_cc() float64 {
	return o.cc
}

func (o *giga) get_dd() float64 {
	return o.dd
}

func (o *giga) get_ee() float64 {
	return o.ee
}

func (o *giga) get_ff() float64 {
	return o.ff
}

func (o *giga) get_gg() float64 {
	return o.gg
}

func (o *giga) get_hh() float64 {
	return o.hh
}

func (o *giga) get_ii() float64 {
	return o.ii
}

func (o *giga) get_jj() float64 {
	return o.jj
}

func (o *giga) get_kk() float64 {
	return o.kk
}

func (o *giga) get_ll() float64 {
	return o.ll
}

func (o *giga) get_mm() float64 {
	return o.nn
}

func (o *giga) get_nn() float64 {
	return o.nn
}

func (o *giga) get_oo() float64 {
	return o.oo
}

func (o *giga) get_pp() float64 {
	return o.pp
}

func (o *giga) get_qq() float64 {
	return o.gg
}

func (o *giga) get_rr() float64 {
	return o.rr
}

func (o *giga) get_ss() float64 {
	return o.ss
}

func (o *giga) get_tt() float64 {
	return o.tt
}

func (o *giga) get_uu() float64 {
	return o.uu
}

func (o *giga) get_vv() float64 {
	return o.vv
}

func (o *giga) get_xx() float64 {
	return o.xx
}

func (o *giga) get_yy() float64 {
	return o.yy
}

func (o *giga) get_ww() float64 {
	return o.ww
}

func (o *giga) get_zz() float64 {
	return o.zz
}
