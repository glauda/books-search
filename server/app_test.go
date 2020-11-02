

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

/*
func TestSaveRead(t *testing.T){
	t.Run("Returns 'this is a test'", func(t *testing.T){
		
		p1 := &Page{Title: "TestPage", Body: []byte("this is a test")}
		p1.save()
		p2, _ := loadPage("TestPage")
		
		actualResult := string(p2.Body)
		expectedResult := "this is a test"

		// assert block
		if actualResult != expectedResult{
			t.Errorf("Actual: %s \nExpected: %s", actualResult, expectedResult)
		}
	})
}
*/