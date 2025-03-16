package lib

// SectorType is a conversion between CircleMUD's sector type number and a human-readable string.
var SectorType = map[string]string{
	"0": "INSIDE",       // Indoors (small number of move points needed).
	"1": "CITY",         // The streets of a city.
	"2": "FIELD",        // An open field.
	"3": "FOREST",       // A dense forest.
	"4": "HILLS",        // Low foothills.
	"5": "MOUNTAIN",     // Steep mountain regions.
	"6": "WATER_SWIM",   // Water (swimmable).
	"7": "WATER_NOSWIM", // Unswimmable water - boat required for passage.
	"8": "UNDERWATER",   // Underwater.
	"9": "FLYING",       // Wheee!
	"10": "VOID",
	"11": "DESERT",
	"12": "FROZEN_WASTE",
	"13": "FROZEN_MOUNTAINS",
	"14": "FROZEN_UNDERWATER",
	"15": "FROZEN_OCEAN",
}

// ExitDir is the conversion between CircleMUD's direction number and a human-readable string.
var ExitDir = map[string]string{
	"0": "north",
	"1": "east",
	"2": "south",
	"3": "west",
	"4": "up",
	"5": "down",
}

var RoomProcs = map[string]string{
	"trans": "trans",
	"ttrans": "ttrans",
	"echo": "echo",
	"push": "push",
	"pushall": "pushall",
}

// DoorFlags is the conversion between CircleMUD's door flags and a human-readable string.
var DoorFlags = map[string]string{
	"0": "NONE",
	"1": "NORMAL",
	"2": "CLOSED",
	"3" : "HIDDENALTAR",
	"4": "LOCKED",
	"5": "GATE",
	"7": "SMALLWOODENDOOR",
	"8": "UNUSED1",
	"9": "DEATH1",
	"11": "WOODEN",
	"16": "UNUSED2",
	"31": "WALL",
	"32": "PICKPROOF",
	"33": "3BRIDGELAVA",
	"35": "WILLOWDOOR",
	"39": "SECRETDOOR",
	"64": "TRAP1",
	"128": "UNUSED3",
	"256": "NOBASH",
	"257": "ELEGANTDOOR",
	"259": "METALPLATEMARK",
	"263": "SECRETFIREPLACE",
	"289": "PORTRAITDOOR",
	"291": "ALTARGREYMASS",
	"293": "FRENCHBALCONY",
	"295":"BENEATHICEFLOOR",
	"299": "KURRELDOOR",
	"315": "PENTAGRAM",
	"318": "PERSONALDOOR",
	"512": "UNUSED4",
	"803": "SMALLHOVERDOOR",
	"807": "HOBBITHOME",
	"1024": "UNUSED5",
	"1830": "MOONGATE",
	"2048": "TRAP2",
	"2112": "OPENSQUARE",
	"2113": "DIRTDOOR",
	"2375": "WINGEDDOOR",

}
