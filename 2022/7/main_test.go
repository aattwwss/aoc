package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

var dirA, dirB, dirC, dirD, dirE Dir

//	func TestInit(t *testing.T) {
//		dirA = Dir{Prev: nil, Files: []int{1, 1, 2, 3, 11}}
//		dirB = Dir{Prev: &dirA, Dirs: []*Dir{&dirA}}
//		dirC = Dir{Files: []int{1, 2, 3}}
//		dirD = Dir{Files: []int{4, 5, 6}}
//		dirE = Dir{Files: []int{7, 8, 9}}
//	}
func TestGetDirSize(t *testing.T) {
	dirA := Dir{Prev: nil, Files: []int{1, 1, 2, 3}}
	got := dirA.GetDirSize()
	want := 7
	if got != want {
		t.Errorf("files only: got %v, wanted %v", got, want)
	}
	dirB := Dir{Dirs: []*Dir{&dirA}}
	got = dirB.GetDirSize()
	want = 7
	if got != want {
		t.Errorf("1 dir only: got %v, wanted %v", got, want)
	}

	dirC := Dir{Files: []int{1, 2, 3}}
	dirD := Dir{Files: []int{4, 5, 6}}
	dirB.Dirs = append(dirB.Dirs, &dirC, &dirD)
	got = dirB.GetDirSize()
	want = 28
	if got != want {
		t.Errorf("3 dir only: got %v, wanted %v", got, want)
	}
	dirE := Dir{Files: []int{1}}
	dirD.Dirs = append(dirD.Dirs, &dirE)
	got = dirB.GetDirSize()
	want = 29
	if got != want {
		t.Errorf("4 dir only: got %v, wanted %v", got, want)
	}
}

func TestGetDirSizeMap(t *testing.T) {
	dirA := Dir{Name: "a", Files: []int{1, 1, 1}}
	dirB := Dir{Name: "b", Files: []int{2, 2, 2}}
	dirC := Dir{Name: "c", Files: []int{3, 3, 3}}
	dirD := Dir{Name: "d", Files: []int{4, 4, 4}}
	dirA.Dirs = []*Dir{&dirB}
	dirB.Dirs = []*Dir{&dirC, &dirD}
	m := dirA.GetDirSizeMap()
	mar, _ := json.Marshal(m)
	fmt.Println(string(mar))
	got := len(m)
	want := 4
	if got != want {
		t.Errorf("map length error: got %v, wanted %v", got, want)
	}
}
