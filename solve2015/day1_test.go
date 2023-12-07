package solve2015

import (
	"fmt"
	"testing"
)

func TestDay1(t *testing.T) {
	fmt.Println("testing with: (())")
	if res := Day1("(())"); res != 0 {
		t.Errorf("(()) should give 0, go: %d", res)
	} else {
		fmt.Println("Passed")
	}

	fmt.Println("testing with: ()()")
	if res := Day1("()()"); res != 0 {
		t.Errorf("()() should give 0, go: %d", res)
	} else {
		fmt.Println("Passed")
	}

	fmt.Println("testing with: (((")
	if res := Day1("((("); res != 3 {
		t.Errorf("((( should give 3, go: %d", res)
	} else {
		fmt.Println("Passed")
	}

	fmt.Println("testing with: ))(((((")
	if res := Day1("))((((("); res != 3 {
		t.Errorf("))((((( should give 3, go: %d", res)
	} else {
		fmt.Println("Passed")
	}

	fmt.Println("testing with: (()(()(")
	if res := Day1("(()(()("); res != 3 {
		t.Errorf("(()(()( should give 3, go: %d", res)
	} else {
		fmt.Println("Passed")
	}

	fmt.Println("testing with: )())())")
	if res := Day1(")())())"); res != -3 {
		t.Errorf(")())()) should give 3, go: %d", res)
	} else {
		fmt.Println("Passed")
	}
}

func TestDay1Part2(t *testing.T) {
	fmt.Println("testing with: )")
	if res := Day1Part2(")"); res != 1 {
		t.Errorf(") should give 1, got: %d", res)
	} else {
		fmt.Println("passed")
	}
	fmt.Println("testing with: ()())")
	if res := Day1Part2("()())"); res != 5 {
		t.Errorf("()()) should give 5, got: %d", res)
	} else {
		fmt.Println("passed")
	}
}
