package val

import (
	"testing"
	"time"
)

// TestRetrieveIDFromURL ...
func TestRetrieveIDFromURL(t *testing.T) {
	goodInputs := [][2]string{
		{"/landmarks/5", "landmarks"},
		{"/landmarks/235/delete", "landmarks"},
		{"/landmarks/345333/", "landmarks"},
	}

	badInputs := [][2]string{
		{"/landmarks/", "landmarks"},
		{"/landmarks/ffff/delete", "landmarks"},
		{"/landmarks/12343/", "invitations"},
	}

	goodOutputs := []int64{
		5,
		235,
		345333,
	}

	for i, v := range goodInputs {
		output, _ := retrieveIDFromURL(v[0], v[1])
		if output == goodOutputs[i] {
			t.Logf("PASSED expected: %d, got: %d", goodOutputs[i], output)
		} else {
			t.Errorf("FAILED expected: %d, got: %d", goodOutputs[i], output)
		}
	}

	for _, v := range badInputs {
		_, err := retrieveIDFromURL(v[0], v[1])
		if err != nil {
			t.Logf("PASSED expected: error, got: error")
		} else {
			t.Errorf("FAILED expected: error, got: no error")
		}
	}
}

func TestCheckUserID(t *testing.T) {
	intInputs := []int64{123, 544444444444}
	strGoodInputs := []string{"123", "544444444444"}
	strBadInputs := []string{"fr3asd", ""}

	for i, v := range intInputs {
		err := checkUserID(strGoodInputs[i], v)
		if err != nil {
			t.Errorf("FAILED expected: no error, got: error")
		} else {
			t.Logf("PASSED expected: no error, got: no error")
		}
	}

	for i, v := range intInputs {
		err := checkUserID(strBadInputs[i], v)
		if err != nil {
			t.Logf("PASSED expected: error, got: error")
		} else {
			t.Errorf("FAILED expected: error, got: no error")
		}
	}
}

func TestCheckDateAfter(t *testing.T) {
	goodInputs := []time.Time{
		time.Now().AddDate(2, 0, 0),
		time.Now().AddDate(10, 3, 2),
	}

	badInputs := []time.Time{
		time.Now(),
		time.Now().AddDate(-1, 0, 0),
	}

	for _, v := range goodInputs {
		err := checkDateAfter(v, "")
		if err != nil {
			t.Errorf("FAILED expected: no error, got: error")
		} else {
			t.Logf("PASSED expected: no error, got: no error")
		}
	}

	for _, v := range badInputs {
		err := checkDateAfter(v, "")
		if err != nil {
			t.Logf("PASSED expected: error, got: error")
		} else {
			t.Errorf("FAILED expected: error, got: no error")
		}
	}
}
