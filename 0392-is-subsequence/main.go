package main

func isSubsequence(s string, t string) bool {
	lengthS, lengthT := len(s), len(t)
	iS, iT := 0, 0
	for iS < lengthS && iT < lengthT {
		if s[iS] == t[iT] {
			iS++
		}
		iT++
	}
	return iS == lengthS
}
