package main

import "testing"

func TestSum(t *testing.T ) {
    result := sum(6, 5)

    if result != 10 {
        t.Errorf("Soma de 5 + 5 != %d", result)
    }

}