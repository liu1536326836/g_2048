package main

import (
	"testing"
)

func TestDownTurnToUp(t *testing.T) {
	res := 16
	var g g2048
	g.a = array{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	g.Rollback()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res, g.a[i][j])
			}
			res--
		}
	}
}

func TestTurnRight(t *testing.T) {
	res := array{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}
	var g g2048
	g.a = array{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	g.TurnRight()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestTurnLeft(t *testing.T) {
	res := array{{4, 8, 12, 16}, {3, 7, 11, 15}, {2, 6, 10, 14}, {1, 5, 9, 13}}
	var g g2048
	g.a = array{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	g.TurnLeft()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp1(t *testing.T) {
	res := array{{8, 2, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var g g2048
	g.a = array{{0, 0, 0, 0}, {0, 2, 0, 0}, {4, 0, 0, 0}, {4, 0, 0, 0}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp2(t *testing.T) {
	res := array{{4, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var g g2048
	g.a = array{{2, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {2, 0, 0, 0}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp3(t *testing.T) {
	res := array{{2, 0, 0, 0}, {4, 0, 0, 0}, {16, 0, 0, 0}, {4, 0, 0, 0}}
	var g g2048
	g.a = array{{2, 0, 0, 0}, {4, 0, 0, 0}, {16, 0, 0, 0}, {4, 0, 0, 0}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp4(t *testing.T) {
	res := array{{4, 0, 0, 0}, {4, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var g g2048
	g.a = array{{2, 0, 0, 0}, {2, 0, 0, 0}, {0, 0, 0, 0}, {4, 0, 0, 0}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp5(t *testing.T) {
	res := array{{4, 0, 0, 0}, {8, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var g g2048
	g.a = array{{2, 0, 0, 0}, {2, 0, 0, 0}, {4, 0, 0, 0}, {4, 0, 0, 0}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveUp6(t *testing.T) {
	res := array{{4, 4, 2, 8}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
	var g g2048
	g.a = array{{0, 4, 0, 4}, {2, 0, 0, 0}, {0, 0, 0, 0}, {2, 0, 2, 4}}
	g.MoveUp()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveDown1(t *testing.T) {
	res := array{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 4}, {8, 0, 8, 2}}
	var g g2048
	g.a = array{{4, 0, 0, 4}, {0, 0, 0, 0}, {4, 0, 0, 0}, {0, 0, 8, 2}}
	g.MoveDown()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveLeft1(t *testing.T) {
	res := array{{8, 0, 0, 0}, {2, 0, 0, 0}, {4, 0, 0, 0}, {4, 0, 0, 0}}
	var g g2048
	g.a = array{{4, 0, 0, 4}, {0, 0, 0, 2}, {4, 0, 0, 0}, {0, 2, 0, 2}}
	g.MoveLeft()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestMoveRight1(t *testing.T) {
	res := array{{0, 0, 0, 8}, {0, 0, 0, 2}, {0, 0, 0, 4}, {0, 0, 0, 4}}
	var g g2048
	g.a = array{{4, 0, 0, 4}, {0, 0, 0, 2}, {4, 0, 0, 0}, {0, 2, 0, 2}}
	g.MoveRight()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if g.a[i][j] != res[i][j] {
				t.Errorf("数组第%d行%d列的元素应该为%d, 实际为%d", i, j, res[i][j], g.a[i][j])
			}
		}
	}
}

func TestRand(t *testing.T) {
	var g g2048
	g.a = array{{4, 0, 0, 0}, {0, 0, 0, 0}, {4, 0, 0, 0}, {0, 0, 0, 0}}
	g.Rand()
	t.Log(g.a)
}
