package contrib

import (
	"log"
	"testing"
)

func TestNewLBPHFaceMark(t *testing.T) {
	mark := NewLBPHFaceMark()
	log.Println(mark)
}

func TestLBPHFaceMark_LoadModel(t *testing.T) {
	mark := NewLBPHFaceMark()
	log.Println(mark)
	mark.LoadModel("/Users/wushaojie/Documents/project/golang/go-opencv/lbfmodel.yaml")
	log.Println("success")
}

func TestLBPHFaceMark_Fit(t *testing.T) {
}
