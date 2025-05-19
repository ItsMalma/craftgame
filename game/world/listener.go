package world

type Listener interface {
	TileChanged(var1, var2, var3 int)
	LightColumnChanged(var1, var2, var3, var4 int)
	AllChanged()
}
