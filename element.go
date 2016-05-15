package main

import (
	"github.com/nsf/termbox-go"
)

type attr struct {
	x, y   int
	fg, bg termbox.Attribute
	ch     rune
}

type elt struct {
	val int
	at  []attr
}

var (
	elts = map[int][]attr{
		2: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorRed, '2'},
			{5, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorRed, ' '}},

		4: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorGreen, '4'},
			{5, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorGreen, ' '}},

		8: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorYellow, '8'},
			{5, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorYellow, ' '}},

		16: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorBlue, '1'},
			{4, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorBlue, '6'},
			{5, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorBlue, ' '}},

		32: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorMagenta, '3'},
			{4, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorMagenta, '2'},
			{5, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '}},

		64: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorCyan, '6'},
			{4, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorCyan, '4'},
			{5, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{6, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorCyan, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorCyan, ' '}},

		128: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorRed, '1'},
			{4, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorRed, '2'},
			{5, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorRed, '8'},
			{6, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorRed, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorRed, ' '}},

		256: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorGreen, '2'},
			{4, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorGreen, '5'},
			{5, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorGreen, '6'},
			{6, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorGreen, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorGreen, ' '}},

		512: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorYellow, '5'},
			{4, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorYellow, '1'},
			{5, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorYellow, '2'},
			{6, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorYellow, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorYellow, ' '}},

		1024: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorBlue, '1'},
			{3, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorBlue, '0'},
			{4, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorBlue, '2'},
			{5, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorBlue, '4'},
			{6, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorBlue, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorBlue, ' '}},

		2048: []attr{
			{1, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{1, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{2, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 3, termbox.ColorWhite, termbox.ColorMagenta, '2'},
			{3, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{3, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 3, termbox.ColorWhite, termbox.ColorMagenta, '0'},
			{4, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{4, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 3, termbox.ColorWhite, termbox.ColorMagenta, '4'},
			{5, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{5, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 3, termbox.ColorWhite, termbox.ColorMagenta, '8'},
			{6, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{6, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{7, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{8, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 1, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 2, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 3, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 4, termbox.ColorWhite, termbox.ColorMagenta, ' '},
			{9, 5, termbox.ColorWhite, termbox.ColorMagenta, ' '}}}
)
