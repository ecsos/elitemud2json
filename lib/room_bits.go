package lib

// Room bit definitions
const (
	ROOM_DARK    = 1 << iota
	ROOM_DEATH
	ROOM_NOMOB
	ROOM_INDOORS
	ROOM_LAWFUL
	ROOM_NEUTRAL
	ROOM_CHAOTIC
	ROOM_NOMAGIC
	ROOM_TUNNEL
	ROOM_PRIVATE
	ROOM_GODROOM
	ROOM_BFS_MARK
	ROOM_ZERO_MANA
	ROOM_DISPELL
	ROOM_SILENT
	ROOM_IN_AIR
	ROOM_OCS
	ROOM_PKOK
	ROOM_ARENA
	ROOM_REGEN
	ROOM_NO_TELEPORT
	ROOM_NO_SCRY
	ROOM_NO_FLEE
	ROOM_DAMAGE
	ROOM_NOTRACK
	ROOM_NOSWEEP
	ROOM_NOSCOUT
	ROOM_NOSLEEP
	ROOM_NOSUMMON
	ROOM_NOQUIT
	ROOM_NODROP
)

// RoomChars converts a letter-style bit to the corresponding bit name
var RoomChars = map[rune]string{
	'a': "DARK",
	'b': "DEATH",
	'c': "NOMOB",
	'd': "INDOORS",
	'e': "LAWFUL",
	'f': "NEUTRAL",
	'g': "CHAOTIC",
	'h': "NOMAGIC",
	'i': "TUNNEL",
	'j': "PRIVATE",
	'k': "GODROOM",
	'l': "BFS_MARK",
	'm': "ZERO_MANA",
	'n': "DISPELL",
	'o': "SILENT",
	'p': "IN_AIR",
	'q': "OCS",
	'r': "PKOK",
	's': "ARENA",
	't': "REGEN",
	'u': "NO_TELEPORT",
	'v': "NO_SCRY",
	'w': "NO_FLEE",
	'x': "DAMAGE",
	'y': "NOTRACK",
	'z': "NOSWEEP",
	'A': "NOSCOUT",
	'B': "NOSLEEP",
	'C': "NOSUMMON",
	'D': "NOQUIT",
	'E': "NODROP",
}

var RoomBits = map[int]string{
	ROOM_DARK:        "DARK",
	ROOM_DEATH:       "DEATH",
	ROOM_NOMOB:       "NOMOB",
	ROOM_INDOORS:     "INDOORS",
	ROOM_LAWFUL:      "LAWFUL",
	ROOM_NEUTRAL:     "NEUTRAL",
	ROOM_CHAOTIC:     "CHAOTIC",
	ROOM_NOMAGIC:     "NOMAGIC",
	ROOM_TUNNEL:      "TUNNEL",
	ROOM_PRIVATE:     "PRIVATE",
	ROOM_GODROOM:     "GODROOM",
	ROOM_BFS_MARK:    "BFS_MARK",
	ROOM_ZERO_MANA:   "ZERO_MANA",
	ROOM_DISPELL:     "DISPELL",
	ROOM_SILENT:      "SILENT",
	ROOM_IN_AIR:	  "IN_AIR",
	ROOM_OCS:         "OCS",
	ROOM_PKOK:        "PKOK",
	ROOM_ARENA:	      "ARENA",
	ROOM_REGEN:       "REGEN",
	ROOM_NO_TELEPORT: "NO_TELEPORT",
	ROOM_NO_SCRY:     "NO_SCRY",
	ROOM_NO_FLEE:     "NO_FLEE",
	ROOM_DAMAGE:      "DAMAGE",
	ROOM_NOTRACK: 	  "NOTRACK",
	ROOM_NOSWEEP:     "NOSWEEP",
	ROOM_NOSCOUT:     "NOSCOUT",
	ROOM_NOSLEEP:     "NOSLEEP",
	ROOM_NOSUMMON:    "NOSUMMON",
	ROOM_NOQUIT:      "NOQUIT",
	ROOM_NODROP:      "NODROP",
}

// BitVectorToNames converts a room's bitvector into a list of bit names
func BitVectorToNames(vector string) ([]string, error) {
	return BitsToNames(vector, ROOM_BFS_MARK, RoomBits, RoomChars)
}
