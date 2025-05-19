package world

import (
	"os"
	"slices"

	"github.com/ItsMalma/craftgame/phys"
)

type World struct {
	width, height, depth int

	blocks      []byte
	lightDepths []int

	listeners []Listener
}

func New(width, height, depth int) (*World, error) {
	w := new(World)
	w.width = width
	w.height = height
	w.depth = depth
	w.blocks = make([]byte, width*height*depth)
	w.lightDepths = make([]int, width*height)
	w.listeners = []Listener{}

	for x := range width {
		for y := range depth {
			for z := range height {
				block := 0
				if y <= depth*2/3 {
					block = 1
				}

				w.blocks[(y*w.height+z)*w.width+x] = byte(block)
			}
		}
	}

	w.CalcLightDepths(0, 0, width, height)

	if err := w.Load(); err != nil {
		return nil, err
	}

	return w, nil
}

func (w *World) Width() int {
	return w.width
}

func (w *World) Height() int {
	return w.height
}

func (w *World) Depth() int {
	return w.depth
}

func (w *World) Load() error {
	worldFile, err := os.Open("world.dat")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer worldFile.Close()

	if _, err = worldFile.Read(w.blocks); err != nil {
		return err
	}

	w.CalcLightDepths(0, 0, w.width, w.height)

	for _, listener := range w.listeners {
		listener.AllChanged()
	}

	return nil
}

func (w *World) Save() error {
	worldFile, err := os.OpenFile("world.dat", os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer worldFile.Close()

	if _, err = worldFile.Write(w.blocks); err != nil {
		return err
	}

	return nil

}

func (w *World) CalcLightDepths(x0, y0, x1, y1 int) {
	for x := x0; x < x0+x1; x++ {
		for z := y0; z < y0+y1; z++ {
			oldDepth := w.lightDepths[x+z*w.width]

			y := w.depth - 1
			for y > 0 && !w.IsLightBlocker(x, y, z) {
				y--
			}

			w.lightDepths[x+z*w.width] = y
			if oldDepth != y {
				yl0 := y
				if oldDepth < y {
					yl0 = oldDepth
				}
				yl1 := y
				if oldDepth > y {
					yl1 = oldDepth
				}

				for _, listener := range w.listeners {
					listener.LightColumnChanged(x, z, yl0, yl1)
				}
			}
		}
	}
}

func (w *World) AddListener(l Listener) {
	w.listeners = append(w.listeners, l)
}

func (w *World) RemoveListener(l Listener) {
	w.listeners = slices.DeleteFunc(w.listeners, func(e Listener) bool {
		return e == l
	})
}

func (w *World) IsTile(x, y, z int) bool {
	if x >= 0 && y >= 0 && z >= 0 && x < w.width && y < w.depth && z < w.height {
		return w.blocks[(y*w.height+z)*w.width+x] == 1
	} else {
		return false
	}
}

func (w *World) IsSolidTile(x, y, z int) bool {
	return w.IsTile(x, y, z)
}

func (w *World) IsLightBlocker(x, y, z int) bool {
	return w.IsSolidTile(x, y, z)
}

func (w *World) GetCubes(aabb phys.AABB) []phys.AABB {
	aabbs := make([]phys.AABB, 0)
	x0 := int(aabb.X0)
	x1 := int(aabb.X1 + 1)
	y0 := int(aabb.Y0)
	y1 := int(aabb.Y1 + 1)
	z0 := int(aabb.Z0)
	z1 := int(aabb.Z1 + 1)

	if x0 < 0 {
		x0 = 0
	}
	if y0 < 0 {
		y0 = 0
	}
	if z0 < 0 {
		z0 = 0
	}

	if x1 > w.width {
		x1 = w.width
	}
	if y1 > w.depth {
		y1 = w.depth
	}
	if z1 > w.height {
		z1 = w.height
	}

	for x := x0; x < x1; x++ {
		for y := y0; y < y1; y++ {
			for z := z0; z < z1; z++ {
				if w.IsSolidTile(x, y, z) {
					aabbs = append(aabbs, phys.AABB{
						X0: float32(x),
						Y0: float32(y),
						Z0: float32(z),
						X1: float32(x + 1),
						Y1: float32(y + 1),
						Z1: float32(z + 1),
					})
				}
			}
		}
	}

	return aabbs
}

func (w *World) GetBrightness(x, y, z int) float32 {
	const dark = 0.8
	const light = 1.0

	if x >= 0 && y >= 0 && z >= 0 && x < w.width && y < w.depth && z < w.height {
		if y < w.lightDepths[x+z*w.width] {
			return dark
		} else {
			return light
		}
	} else {
		return light
	}
}

func (w *World) SetTile(x, y, z int, type_ byte) {
	if x >= 0 && y >= 0 && z >= 0 && x < w.width && y < w.depth && z < w.height {
		w.blocks[(y*w.height+z)*w.width+x] = type_
		w.CalcLightDepths(x, z, 1, 1)

		for _, listener := range w.listeners {
			listener.TileChanged(x, y, z)
		}
	}
}
