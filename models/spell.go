package models

type SpellType uint16

const (
	SPELL_TYPE_PROJECTILE SpellType = 1
	SPELL_TYPE_AOE
	SPELL_TYPE_DEFENSIVE
)

type Spell struct {
	ID        int // Unique ID for the spell
	Name      string
	Damage    int
	Cooldown  int // Time in seconds
	ManaCost  int
	SpellType SpellType // "Projectile", "AoE", "Defensive", etc.
	CastTime  float32   // 0 for instant cast, > 0 for charged spells
	CastRange int
}
