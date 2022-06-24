/*
Copyright 2022 epiccakeking

This file is part of Ring Match.

Ring Match is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

Ring Match is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with Ring Match. If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed res
var res embed.FS
var scale float64 = 1

type Sprite struct{ *ebiten.Image }

// Draw the sprite to the screen/image. Rotation r is in radians.
func (s Sprite) Draw(screen *ebiten.Image, x, y, r float64) {
	var o ebiten.DrawImageOptions
	sX, sY := s.Size()
	o.GeoM.Translate(-float64(sX)/2, -float64(sY)/2)
	o.GeoM.Rotate(r)
	o.GeoM.Translate(x, y)
	o.GeoM.Scale(scale, scale)
	screen.DrawImage(s.Image, &o)
}

// Load a sprite from res
func ResSprite(path string) Sprite {
	f, err := res.Open(path)
	if err != nil {
		panic(err)
	}
	i, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return Sprite{ebiten.NewImageFromImage(i)}
}
