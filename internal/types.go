package internal

type Damage int

const (
	ShortRangeSensor Damage = iota
	ComputerDisplay
	LongRangeSensor
	Phasers
	WarpEngine
	PhotonTorpedoTubes
	Shield
)

type ship struct {
	Energy int
	X      int
	Y      int
}

// Klingon represents the state of a Klingon ship
type Klingon struct {
	ship
}

// Enterprise represents the state of the player's ship
type Enterprise struct {
	ship
	QuadrantX int
	QuadrantY int
	IsDocked  bool
	Torpedoes int
	System    [8]Damage
}

// Quadrant contains the game assets for the region
type Quadrant struct {
	Base     bool
	Klingons int
	Stars    int
}

// World represents the current state of the game
type World struct {
	NumKlingons     int
	NumBases        int
	NumDays         int
	NumCasualties   int
	NumDeadKlingons int
	Quadrants       [9][9]Quadrant
	Sectors         [9][9]int
}
