package model

import "time"

type UserSet struct {
	Name    string
	UserIDs []uint8
	Size    int
	Date    time.Time
}

type DAU struct {
	DAU []UserSet
}

func NewDAU(aDAU []UserSet) *DAU {
	return &DAU{
		DAU: aDAU,
	}
}

func (aDAUSimple *DAU) EachUserSet(aProc func(UserSet)) {
	for _, tUserSet := range aDAUSimple.DAU {
		aProc(tUserSet)
	}
}