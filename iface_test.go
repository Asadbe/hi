package main

import "testing"

func TestRecArea(t *testing.T) {

	rec := rectangle{width: 6, height: 2}
	if geometry.area(rec) != 12 {
		t.Errorf("the function is incorrect")
	}

}
func TestCirArea(t *testing.T) {

	cir := circle{radius: 10}
	if geometry.area(cir) != 314.15927{
		t.Errorf("the function is incorrect")
	}

}

func TestRecPerimetr(t *testing.T) {

	rec := rectangle{width: 6, height: 2}
	if geometry.perimetr(rec) != 16 {
		t.Errorf("the function is incorrect")
	}

}

func TestCirPerimetr(t *testing.T) {

	cir := circle{radius: 10}
	if geometry.perimetr(cir) != 62.831856{
		t.Errorf("the function is incorrect")
	}

}