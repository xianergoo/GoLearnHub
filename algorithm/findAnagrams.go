package main

import "fmt"

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}

func findAnagrams(s, p string) (ans []int) {
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		fmt.Println(string(s[i]), string(ch), s[i], ch, s[i]-'a', ch-'a')
		sCount[s[i]-'a']++
		pCount[ch-'a']++
		fmt.Println(s[i]-'a', sCount[s[i]-'a'])
		fmt.Println(ch-'a', pCount[ch-'a'])
	}

	fmt.Println(sCount)
	fmt.Println(pCount)

	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i, ch := range s[:sLen-pLen] {
		sCount[ch-'a']--
		sCount[s[i+pLen]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
