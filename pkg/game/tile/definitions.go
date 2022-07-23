package tile

const (
	BlockEmpty ID = iota
	BlockStone
	BlockSoil
	BlockRoughWall
	BlockSmoothWall
)

const (
	FloorEmpty ID = iota
	FloorStone
	FloorSoil
	FloorRough
	FloorSmooth
)

// DefaultDefinitions returns the default block and floor tile definitions.
func DefaultDefinitions() ([]Definition, []Definition) {
	blocks := []Definition{
		{ID: "block-empty", Name: "empty"},
		{ID: "block-stone", Name: "{{.Solid.Name}}"},
		{ID: "block-soil", Name: "{{.Solid.Name}}"},
		{ID: "block-wall-rough", Name: "rough {{.Solid.Adjective}} wall"},
		{ID: "block-wall-smooth", Name: "smooth {{.Solid.Adjective}} wall"},
	}
	floors := []Definition{
		{ID: "floor-empty", Name: "empty"},
		{ID: "floor-stone", Name: "{{.Solid.Name}}"},
		{ID: "floor-soil", Name: "{{.Solid.Name}}"},
		{ID: "floor-rough", Name: "rough {{.Solid.Adjective}} floor"},
		{ID: "floor-smooth", Name: "smooth {{.Solid.Adjective}} floor"},
	}
	return blocks, floors
}
