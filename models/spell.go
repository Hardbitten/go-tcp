package models

import (
	"main/enums"
	"main/utils"
)

type Spell struct {
	ID        int             // Unique ID for the spell
	SpellType enums.SpellType // "Projectile", "AoE", "Defensive", etc.
	Caster    *Player
	Name      string
	Damage    int
	Speed     int
	Slow      int

	ManaCost  int
	Cooldown  int     // Time in seconds
	CastTime  float32 // 0 for instant cast, > 0 for charged spells
	CastRange int

	TargetPosition utils.Vector3
}
