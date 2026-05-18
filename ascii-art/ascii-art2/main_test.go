package main

// import (
// 	"testing"
// 	"os"
// )
// func TestLoadBanner_AllBanners(t *testing.T) {
// 	banners := []string{
// 		"banner/standard.txt",
// 		"banner/shadow.txt",
// 		"banner/thinkertoy.txt",
// 	}
// 	for _, b := range banners {
// 		lines, err := LoadBanner(b)
// 		if err != nil {
// 			t.Errorf("%s: expected no error, got %v", b, err)
// 		}
// 		if len(lines) != 855 {
// 			t.Errorf("%s: expected 855 lines, got %d", b, len(lines))
// 		}
// 	}
// }

// func TestGetLine(t *testing.T) {
// 	ban, err := LoadBanner("banner/standard.txt")
// 	if err != nil {
// 		t.Fatalf("failed to load banner: %v", err)
// 	}

// 	for c := rune(32); c <= 126; c++ {
// 		lines := getLine(ban, c)
// 		if lines == nil || len(lines) != 8 {
// 			t.Errorf("char %q: expected 8 lines, got %v", c, lines)
// 		}
// 	}

// 	for _, c := range []rune{0, '\n', 31, 127, 200} {
// 		if lines := getLine(ban, c); lines != nil {
// 			t.Errorf("char %q: expected nil, got %v", c, lines)
// 		}
// 	}

// 	for i, line := range getLine(ban, 'A') {
// 		if line != ban[298+i] {
// 			t.Errorf("A line %d: expected %q, got %q", i, ban[298+i], line)
// 		}
// 	}
// 	for _, c := range []rune{32, 'A', '~'} {
// 		if lines := getLine(ban[:5], c); lines != nil {
// 			t.Errorf("truncated ban: char %q should return nil, got %v", c, lines)
// 		}
// 	}
// }

// func TestPrintArt(t *testing.T) {
// 	ban, err := LoadBanner("banner/standard.txt")
// 	if err != nil {
// 		t.Fatalf("failed to load banner: %v", err)
// 	}
// printArt([]string{"123"}, ban)
// printArt([]string{"Hello\nThere"}, ban) // \n as literal split
// 	printArt([]string{""}, ban)

// 	printArt([]string{"A"}, ban)

// 	printArt([]string{"Hello"}, ban)

// 	printArt([]string{"1"}, ban)

// 	printArt([]string{"!@#"}, ban)

// 	printArt([]string{"", ""}, ban)

// 	printArt([]string{"Hello", "", "There"}, ban)

// 	all := make([]string, 1)
// 	for c := rune(32); c <= 126; c++ {
// 		all[0] = string(c)
// 		printArt(all, ban)
// 	}

// 	printArt([]string{string(rune(0))}, ban)
// 	printArt([]string{string(rune(127))}, ban)
// 	printArt([]string{string(rune(200))}, ban)
// }

// func TestValidateInput(t *testing.T) {
// 	got, err := validateInput([]string{"Hello\nThere"})
// 	if err != nil || len(got) != 1 {
// 		t.Errorf("newline in string: expected valid, got %v %v", got, err)
// 	}

// 	got, err = validateInput([]string{"a", "b", "c"})
// 	if err != nil || len(got) != 3 {
// 		t.Errorf("expected 3 results, got %d", len(got))
// 	}

// 	for _, input := range [][]string{
// 		{"hello"},
// 		{"Hello World"},
// 		{"1234567890"},
// 		{"!@#$%^&*()"},
// 		{"Hello\nWorld"},
// 		{""},
// 		{},
// 	} {
// 		if _, err := validateInput(input); err != nil {
// 			t.Errorf("input %v: expected no error, got %v", input, err)
// 		}
// 	}

// 	for _, input := range [][]string{
// 		{string(rune(0))},
// 		{string(rune(31))},
// 		{string(rune(127))},
// 		{string(rune(200))},
// 		{"hello\x01world"},
// 	} {
// 		if _, err := validateInput(input); err == nil {
// 			t.Errorf("input %v: expected error, got nil", input)
// 		}
// 	}

// 	got, err = validateInput([]string{"a", "b"})
// 	if err != nil || len(got) != 2 || got[0] != "a" || got[1] != "b" {
// 		t.Errorf("expected [a b], got %v, err %v", got, err)
// 	}
// }
// func TestMain_NoArgs(t *testing.T) {
// 	os.Args = []string{"cmd"}
// 	main()
// }

// func TestMain_EmptyString(t *testing.T) {
// 	os.Args = []string{"cmd", ""}
// 	main()
// }

// func TestMain_SingleNewline(t *testing.T) {
// 	os.Args = []string{"cmd", "\\n"}
// 	main()
// }

// func TestMain_HelloWorld(t *testing.T) {
// 	os.Args = []string{"cmd", "Hello"}
// 	main()
// }

// func TestMain_WithNewline(t *testing.T) {
// 	os.Args = []string{"cmd", "Hello\\nWorld"}
// 	main()
// }

// func TestMain_Numbers(t *testing.T) {
// 	os.Args = []string{"cmd", "123"}
// 	main()
// }

// func TestMain_SpecialChars(t *testing.T) {
// 	os.Args = []string{"cmd", "!@#"}
// 	main()
// }

// func TestMain_TooManyArgs(t *testing.T) {
// 	os.Args = []string{"cmd", "hello", "extra"}
// 	main()
// }