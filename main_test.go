package main

import (
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		t.Fatal("Could not read standard.txt:", err)
	}
	lines := strings.Split(string(data), "\n")

	tests := []struct{ input, want string }{
		{
			input: "",
			want:  "",
		},
		{
			input: "\\n",
			want:  "\n",
		},
		{
			input: "A",
			want: "           \n" +
				"    /\\     \n" +
				"   /  \\    \n" +
				"  / /\\ \\   \n" +
				" / ____ \\  \n" +
				"/_/    \\_\\ \n" +
				"           \n" +
				"           \n",
		},
		{
			input: "Hello",
			want: " _    _          _   _          \n" +
				"| |  | |        | | | |         \n" +
				"| |__| |   ___  | | | |   ___   \n" +
				"|  __  |  / _ \\ | | | |  / _ \\  \n" +
				"| |  | | |  __/ | | | | | (_) | \n" +
				"|_|  |_|  \\___| |_| |_|  \\___/  \n" +
				"                                \n" +
				"                                \n",
		},
		{
			input: "Hello\\nThere",
			want: " _    _          _   _          \n" +
				"| |  | |        | | | |         \n" +
				"| |__| |   ___  | | | |   ___   \n" +
				"|  __  |  / _ \\ | | | |  / _ \\  \n" +
				"| |  | | |  __/ | | | | | (_) | \n" +
				"|_|  |_|  \\___| |_| |_|  \\___/  \n" +
				"                                \n" +
				"                                \n" +
				" _______   _                           \n" +
				"|__   __| | |                          \n" +
				"   | |    | |__     ___   _ __    ___  \n" +
				"   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n" +
				"   | |    | | | | |  __/ | |    |  __/ \n" +
				"   |_|    |_| |_|  \\___| |_|     \\___| \n" +
				"                                       \n" +
				"                                       \n",
		},
	}

	for _, tt := range tests {
		got := asciiArt(tt.input, lines)
		if got != tt.want {
			t.Errorf("input: %q\ngot:  %q\nwant: %q", tt.input, got, tt.want)
		}
	}
}
