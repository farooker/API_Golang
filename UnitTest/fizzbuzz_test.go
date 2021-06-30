package fizzbuzz

import (
	"testing"
)

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := Fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}

func TestInput2ShouldBeDisplay2(t *testing.T) {
	v := Fizzbuzz(2)
	if "2" != v {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}

func TestInput3ShouldBeDisplayFizz(t *testing.T) {
	v := Fizzbuzz(3)
	if "Fizz" != v {
		t.Error("fizzbuzz of 3 should be 'Fizz' but have", v)
	}
}

func TestInput4ShouldBeDisplay4(t *testing.T) {
	v := Fizzbuzz(4)
	if "4" != v {
		t.Error("fizzbuzz of 4 should be '4' but have", v)
	}
}

func TestInput5ShouldBeDisplayBuzz(t *testing.T) {
	v := Fizzbuzz(5)
	if "Buzz" != v {
		t.Error("fizzbuzz of 5 should be 'Buzz' but have", v)
	}
}

func TestInput15ShouldBeDisplayFizzBuzz(t *testing.T) {
	v := Fizzbuzz(15)
	if "FizzBuzz" != v {
		t.Error("fizzbuzz of 15 should be 'FizzBuzz' but have", v)
	}
}

func TestIConnectMongo(t *testing.T) {
	var  url = "mongodb://localhost:27017/"
	v := Connected(url)
	if true!= v {
		t.Error("Connect Mogno  should be 'TRUE'", v)
	}
}

