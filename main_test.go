package main

import "testing"

func TestIsPrime(t *testing.T) {
	var n int
	result, msg := isPrime(n)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", n)
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	n = 1
	result, msg = isPrime(n)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", n)
	}

	if msg != "1 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	n = -1
	result, msg = isPrime(n)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", n)
	}

	if msg != "Negative numbers are not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}

	n = 2
	result, msg = isPrime(n)
	if !result {
		t.Errorf("with %d as test parameter, got false, but expected true", n)
	}

	if msg != "2 is a prime number!" {
		t.Error("wrong message returned:", msg)
	}
}
