package main

/*

5. Longest Palindromic Substring

https://leetcode.com/problems/longest-palindromic-substring/

#string #dynamic-programming

*/

import "fmt"

func expandAroundCenter(s string, length int, left int, right int) (int, int) {
	resultLeft := 0
	resultRight := 0

	for (left >= 0) && (right <= length-1) {
		if s[left] == s[right] {
			resultLeft = left
			resultRight = right
			left--
			right++
		} else {
			break
		}
	}

	return resultLeft, resultRight
}

func longestPalindrome(s string) string {
	length := len(s)

	if length <= 1 {
		return s
	}

	resultLeft := 0
	resultRight := 0

	for index := 0; index < length; index++ {
		left, right := expandAroundCenter(s, length, index, index)
		if (right - left) > (resultRight - resultLeft) {
			resultLeft, resultRight = left, right
		}
	}

	for index := 0; index < length-1; index++ {
		left, right := expandAroundCenter(s, length, index, index+1)
		if (right - left) > (resultRight - resultLeft) {
			resultLeft, resultRight = left, right
		}
	}

	return s[resultLeft : resultRight+1]
}

func main() {
	fmt.Println(longestPalindrome("babad")) // bab                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          // bab
	fmt.Println(longestPalindrome("cbbd"))  // bb                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       // bb
	fmt.Println(longestPalindrome("a"))     // a
	fmt.Println(longestPalindrome("ac"))    // a
	fmt.Println(longestPalindrome(
		"babaddtattarrattatddetartrateedredividerb",
	)) // ddtattarrattatdd
	fmt.Println(longestPalindrome(
		"uwqrvqslistiezghcxaocjbhtktayupazvowjrgexqobeymperyxtfkchujjkeef" +
			"mdngfabycqzlslocjqipkszmihaarekosdkwvsirzxpauzqgnftcuflzyqwftwde" +
			"izwjhloqwkhevfovqwyvwcrosexhflkcudycvuelvvqlbzxoajisqgwgzhioomuc" +
			"fmkmyaqufqggimzpvggdohgxheielsqucemxrkmmagozxhvxlwvtbbcegkvvdrgk" +
			"qszgajebbobxnossfrafglxvryhvyfcibfkgpbsorqprfujfgbmbctsenvbzcvyp" +
			"cjubsnjrjvyznbswqawodghmigdwgijfytxbgpxreyevuprpztmjejkaqyhppchu" +
			"uytkdsteroptkouuvmkvejfunmawyuezxvxlrjulzdikvhgxajohpzrshrnngesa" +
			"rimyopgqydcmsaciegqlpqnclpwcjqmhtmtwwtbkmtnntdllqbyyhfxsjyhugnjb" +
			"ebtxeljytoxvqvrxygmtogndrhlcmbmgiueliyfkkcuypvvzkomjrfhuhhnfbxeu" +
			"vssvvllgukdolffukzwqaimxkngnjnmsbvwkajyxqntsqjkjqvwxnlxwjfiaofej" +
			"tjcveqstqhdzgqistxwsgrovvwgorjaoosremqbzllgbgrwtqdggxnyvkivlcvnv",
	)) // qjkjq
	fmt.Println(longestPalindrome(
		"vnjwvalrbypfcbqnmopltjnoifmzwgvpzqzsdtvawndpjtpmpjbjionjifqtvvoc" +
			"peaftvhpdgjjfafunfndztdjkcxyihtsyppendfzzjeyxlbwpdygiqmdqcdbmgyj" +
			"igrmfkswcwryaydjilqqxvcnyvviesuncslvzikawwqykqwdfibggezufqihcjke" +
			"bapmgkvwixywgdextafxycnipjglsndkyjoqfyfljfkkvoieksmavdlmlhhnstes" +
			"ibffiopqvlyuidvrawndbzonwzbsjmpeqoglmdbinkovqpzfkxihzitdopnomseq" +
			"hmrrkcsvrzziphwpuhjngeotwcrebcmbtirkgeavojtmpakcewmexhxacngknokx" +
			"svtqobdgckutpexswgwqzbosjpxauyflnylfcxsucsehqvakbpvfmkelmkspsqxn" +
			"utwfwacpqqvovdqafeylobneojdsgqowcbxfsvuqusdbylcgcvgrofgvzubakjml" +
			"bffjhrafvnqttwuyhokzpmhlludpbowuxzrebxsdusalljfjgjkucwzpmndqncyk" +
			"vfnbrxcrcaxwisjpstejjqbpwegpxyrtyafxklgralnkwxkmjpuqfixzkonznmgu" +
			"yizlancpxdzcfkgiotyelegprbaytdhbutbuihkxnbtuqrtezaskfqsmrznfohhl" +
			"qp",
	)) // zqz
}
