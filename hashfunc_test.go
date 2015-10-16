package hashfunc

import (
	"testing"
)

func Test_hashfunc(t *testing.T) {
	const testData = "abcdef"

	//SDBMHash
	if SDBMHash(testData) != 387127491 {
		t.Error("Error SDBMHash")
	}

	//RSHash
	if RSHash(testData) != 1359898585 {
		t.Error("Error RSHash")
	}

	//JSHash
	if JSHash(testData) != 389775750 {
		t.Error("Error JSHash")
	}

	//PJWHash
	if PJWHash(testData) != 108567222 {
		t.Error("Error PJWHash")
	}

	//ELFHash
	if ELFHash(testData) != 108567222 {
		t.Error("Error ELFHash")
	}

	//BKDRHash
	if BKDRHash(testData) != 314492959 {
		t.Error("Error BKDRHash")
	}

	//DJBHash
	if DJBHash(testData) != 1900596090 {
		t.Error("Error DJBHash")
	}

	//APHash
	if APHash(testData) != 813648430 {
		t.Error("Error APHash")
	}

}
