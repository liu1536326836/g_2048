package main

import (
	"github.com/nsf/termbox-go"
)

func DrawSurface() {
	// 绘制边框
	termbox.SetCell(20, 0, 0x250C, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(60, 0, 0x2510, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(20, 24, 0x2514, termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(60, 24, 0x2518, termbox.ColorYellow, termbox.ColorDefault)

	for i := 21; i < 60; i++ {
		termbox.SetCell(i, 0, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 6, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 12, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 18, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(i, 24, 0x2500, termbox.ColorYellow, termbox.ColorDefault)
	}

	for i := 1; i < 24; i++ {
		termbox.SetCell(20, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(30, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(40, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(50, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
		termbox.SetCell(60, i, 0x2502, termbox.ColorYellow, termbox.ColorDefault)
	}
}

func DrawElement(a *array) {
	for i, tn := range a {
		for j, v := range tn {
			if v != 0 {
				printElement(a, elts[v], i, j)
			}
		}
	}
}

func printElement(a *array, e []attr, row, col int) {
	for _, k := range e {
		termbox.SetCell((col+2)*10+k.x, row*6+k.y, k.ch, k.fg, k.bg)
	}

}

// 刷新
func Flush() {
	termbox.Flush()
}

// 清屏
func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func Draw(a *array) {
	Clear()
	DrawSurface()
	DrawElement(a)
	Flush()
}

func PrintGameover() {
	Clear()

	termbox.SetCell(36, 11, 'G', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 11, 'A', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 11, 'M', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 11, 'E', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 11, 'O', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 11, 'V', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 11, 'E', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(43, 11, 'R', termbox.ColorYellow, termbox.ColorDefault)

	termbox.SetCell(32, 12, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(33, 12, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(34, 12, 'r', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(35, 12, 'l', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(36, 12, '+', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 12, 'R', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 12, ':', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 12, ' ', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 12, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 12, 'o', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 12, 'n', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(43, 12, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(44, 12, 'i', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(45, 12, 'n', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(46, 12, 'u', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(47, 12, 'e', termbox.ColorYellow, termbox.ColorDefault)

	termbox.SetCell(34, 13, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(35, 13, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(36, 13, 'r', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 13, 'l', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 13, '+', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 13, 'Q', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 13, ':', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 13, ' ', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 13, 'E', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(43, 13, 'x', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(44, 13, 'i', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(45, 13, 't', termbox.ColorYellow, termbox.ColorDefault)

	Flush()
}

func PrintWin() {
	Clear()

	termbox.SetCell(36, 11, 'W', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 11, 'I', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 11, 'N', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 11, 'N', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 11, 'I', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 11, 'N', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 11, 'G', termbox.ColorYellow, termbox.ColorDefault)

	termbox.SetCell(32, 12, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(33, 12, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(34, 12, 'r', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(35, 12, 'l', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(36, 12, '+', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 12, 'R', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 12, ':', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 12, ' ', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 12, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 12, 'o', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 12, 'n', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(43, 12, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(44, 12, 'i', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(45, 12, 'n', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(46, 12, 'u', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(47, 12, 'e', termbox.ColorYellow, termbox.ColorDefault)

	termbox.SetCell(34, 13, 'C', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(35, 13, 't', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(36, 13, 'r', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(37, 13, 'l', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(38, 13, '+', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(39, 13, 'Q', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(40, 13, ':', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(41, 13, ' ', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(42, 13, 'E', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(43, 13, 'x', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(44, 13, 'i', termbox.ColorYellow, termbox.ColorDefault)
	termbox.SetCell(45, 13, 't', termbox.ColorYellow, termbox.ColorDefault)

	Flush()
}
