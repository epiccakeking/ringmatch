/*
Copyright 2022 epiccakeking

This file is part of Ring Match.

Ring Match is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Ring Match is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Ring Match. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"crypto/rand"
	"log"
	"math"
	"math/big"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func main() {
	ebiten.SetFullscreen(true)
	ebiten.RunGame(NewGame())
}

func NewGame() *Game {
	g := new(Game)
	for i := range g.Ring {
		g.Ring[i] = Ring{}
		for j := range g.Ring[i] {
			g.Ring[i][j] = i
		}
	}
	for i := 0; i < 100; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(8))
		if err != nil {
			panic(err)
		}
		log.Println(n)
		g.Rotate(int(n.Int64()), true)
	}
	return g
}

type Game struct {
	Ring [5]Ring
}

func (g *Game) Draw(s *ebiten.Image) {
	for r := range g.Ring {
		for b := range g.Ring[r] {
			Ball[g.Ring[r][b]].Draw(s, 24+float64(r*24)+16*math.Cos(float64(b)*math.Pi/3), 24*math.Sqrt(3)*float64(1+r%2)-16*math.Sin(float64(b)*math.Pi/3), 0)
		}
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return int(152 * scale), int(72 * math.Sqrt(3) * scale)
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.Rotate(0, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.Rotate(1, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key3) {
		g.Rotate(2, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key4) {
		g.Rotate(3, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key5) {
		g.Rotate(4, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key6) {
		g.Rotate(5, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key7) {
		g.Rotate(6, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.Key8) {
		g.Rotate(7, true)
	}

	return nil
}

// Rotate the nth ring. Rings 5, 6, and 7 correspond to the gap between the "real" rings
func (g *Game) Rotate(n int, clockwise bool) {
	if clockwise {
		switch n {
		case 5:
			var (
				a = &g.Ring[0]
				b = &g.Ring[1]
				c = &g.Ring[2]
			)
			t := a[0]
			a[0] = a[5]
			a[5] = b[2]
			b[2] = b[1]
			b[1] = c[4]
			c[4] = c[3]
			c[3] = t
		case 6:
			var (
				a = &g.Ring[1]
				b = &g.Ring[2]
				c = &g.Ring[3]
			)
			t := a[0]
			a[0] = c[3]
			c[3] = c[2]
			c[2] = b[5]
			b[5] = b[4]
			b[4] = a[1]
			a[1] = t
		case 7:
			var (
				a = &g.Ring[2]
				b = &g.Ring[3]
				c = &g.Ring[4]
			)
			t := a[0]
			a[0] = a[5]
			a[5] = b[2]
			b[2] = b[1]
			b[1] = c[4]
			c[4] = c[3]
			c[3] = t
		default:
			a := &g.Ring[n]
			t := a[0]
			for i := 0; i < 5; i++ {
				a[i] = a[i+1]
			}
			a[5] = t
		}
	} else {
		panic("not implemented")
	}
}

type Ring [6]int

var Ball = [5]Sprite{
	ResSprite("res/duck.png"),
	ResSprite("res/cat.png"),
	ResSprite("res/ship.png"),
	ResSprite("res/rabbit.png"),
	ResSprite("res/apple.png"),
}
