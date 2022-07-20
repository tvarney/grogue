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

func DefaultDefinitions() ([]Definition, []Definition) {
	return []Definition{
			{Name: "empty"},
			{Name: "{.material}"},
			{Name: "{.material}"},
			{Name: "rough {.material} wall"},
			{Name: "smooth {.material} wall"},
		}, []Definition{
			{Name: "empty"},
			{Name: "{.material}"},
			{Name: "{.material}"},
			{Name: "rough {.material} floor"},
			{Name: "smooth {.material} floor"},
		}
}
