package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

// 回文ではない文字列を生成します。
func randomNonPalindrome(rng *rand.Rand) string {
	// 一度回文を生成した後、ランダムな 1 rune を変更します。
	n := rng.Intn(23) + 2 // random length: 2 -- 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		var r rune
		for !unicode.IsLetter(r) {
			// The conversion r -> upper -> lower doesn't preserve
			// the value of r in some cases, e.g., µ Μ, ſ S, ı I
			r = unicode.ToLower(rune(rng.Intn(0x1000))) // random rune up to '\u0999'
		}
		runes[i] = r
		runes[n-1-i] = r
	}
	palindrome := string(runes)

	for palindrome == string(runes) {
		var r rune
		for !unicode.IsLetter(r) {
			// The conversion r -> upper -> lower doesn't preserve
			// the value of r in some cases, e.g., µ Μ, ſ S, ı I
			r = unicode.ToLower(rune(rng.Intn(0x1000))) // random rune up to '\u0999'
		}
		pos := rng.Intn(n / 2)
		runes[rng.Intn(2)*((n-1)-2*pos)+pos] = r
	}

	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindromes(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) = true", []rune(p))
		}
	}
}
