package main

import (
	"testing"
)

func TestQueryInBytes(t *testing.T) {	
	query := QueryInBytes("XRP-USD", false)

	expected := []byte{123, 34, 97,
		99 ,116, 105 ,111 ,110 ,34 ,58 ,34 ,115 ,117 ,
		98, 115, 99, 114, 105, 98, 101, 34, 44, 34, 112, 
		97 ,114 ,97 ,109 ,115 ,34 ,58 ,34 ,88 ,76 ,50 ,46,
		88, 82 ,80 ,45 ,85, 83, 68, 34, 125}


	if string(query)!= string(expected) {
		t.Errorf("Bytes returned did not match those expected")
	}
}