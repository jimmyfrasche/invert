//Package invert exists to invert indices as returned by the regexp Index
//family of functions.
//It also contains inverting helpers for FindAll(String)?(Index)?.
package invert

import "regexp"

//Indicies takes a slice of len-2 slices describing non-overlapping subsequences
//of a sequence and the length of that sequence and returns the inverse
//of those subsequences.
func Indicies(idc [][]int, slen int) (out [][]int) {
	ln := len(idc)
	if ln == 0 {
		return [][]int{{0, slen}}
	}

	lopen, ropen := idc[0][0] == 0, idc[ln-1][1] == slen
	if ln == 1 && lopen && ropen {
		return nil
	}

	new := make([][]int, ln+1)
	w := 0
	if !lopen {
		new[0] = []int{0, idc[0][0]}
		w = 1
	}

	for i := 0; i < ln-1; i++ {
		new[w] = []int{idc[i][1], idc[i+1][0]}
		w++
	}
	if !ropen {
		new[w] = []int{idc[ln-1][1], slen}
		w++
	}
	return new[:w]
}

//FindAllIndex returns the inverse of Regexp.FindAllIndex.
func FindAllIndex(r *regexp.Regexp, b []byte, n int) [][]int {
	is := r.FindAllIndex(b, n)
	return Indicies(is, len(b))
}

//FindAllStringIndex returns the inverse of Regexp.FindAllStringIndex.
func FindAllStringIndex(r *regexp.Regexp, s string, n int) [][]int {
	is := r.FindAllStringIndex(s, n)
	return Indicies(is, len(s))
}

//FindAll returns the inverse of Regexp.FindAll.
func FindAll(r *regexp.Regexp, b []byte, n int) (out [][]byte) {
	is := FindAllIndex(r, b, n)
	out = make([][]byte, 0, len(is))
	for _, i := range is {
		out = append(out, b[i[0]:i[1]])
	}
	return
}

//FindAllString returns the inverse of Regexp.FindAllString.
func FindAllString(r *regexp.Regexp, s string, n int) (out []string) {
	is := FindAllStringIndex(r, s, n)
	out = make([]string, 0, len(is))
	for _, i := range is {
		out = append(out, s[i[0]:i[1]])
	}
	return
}
