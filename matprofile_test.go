// +build matprofile

package gocv

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	if MatProfile.Count() != 0 {
		var b bytes.Buffer
		MatProfile.WriteTo(&b, 1)
		fmt.Printf("Not all Mat's in tests were closed: %v", b.String())
		os.Exit(1)
	}
	os.Exit(ret)
}

func TestMatProfile(t *testing.T) {
	if MatProfile.Count() != 0 {
		var b bytes.Buffer
		MatProfile.WriteTo(&b, 1)
		t.Errorf("Mat profile should start with 0 entries. A test failure here likely means that some other test is not closing all Mats. Here are the current profile entries:\n%v", b.String())
	}
	mat := NewMat()
	if MatProfile.Count() != 1 {
		t.Errorf("Mat profile should == 1 after NewMat but instead was %v", MatProfile.Count())
	}
	mat2 := NewMat()
	if MatProfile.Count() != 2 {
		t.Errorf("Mat profile should == 2 after NewMat but instead was %v", MatProfile.Count())
	}
	mat.Close()
	mat2.Close()
	if MatProfile.Count() != 0 {
		t.Errorf("Mat profile should == 0 after Close but instead was %v", MatProfile.Count())
	}
}
