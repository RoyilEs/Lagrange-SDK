package utils

import "math/rand"

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func IsInGroupS(str int64, groupUids []int64) bool {
	for _, s := range groupUids {
		if s == str {
			return true
		}
	}
	return false
}

func IsInListToS(str any, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func IsAdmins(uid int64, adminUids []int64) bool {
	for _, s := range adminUids {
		if s == uid {
			return true
		}
	}
	return false
}
