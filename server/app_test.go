package main

import "testing"

func TestHello(t *testing.T){
	t.Run("Returns 'Hello, World' when input is 'World'", func(t *testing.T){
		actualResult := Hello("World")
		expectedResult := "Hello, World"

		// assert block
		if actualResult != expectedResult{
			t.Errorf("Actual: %s \nExpected: %s", actualResult, expectedResult)
		}
	})
}