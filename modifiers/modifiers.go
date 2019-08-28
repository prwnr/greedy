package modifiers

//Monster modifiers
const (
	//MonsterBaseHealth is used in health calculation
	MonsterBaseHealth = 30
	//MonsterBaseAttack is used in attack calculation
	MonsterBaseAttack = 4
	//MonsterBaseExperience is used in exp calculation
	MonsterBaseExperience = 15
)

//CalculateMonsterHealth based on base health and current level
func CalculateMonsterHealth(l int) int {
	return MonsterBaseHealth*l + (int(float64(l-1) * 5))
}

//CalculateMonsterAttack multiplied by level
func CalculateMonsterAttack(l int) int {
	return MonsterBaseAttack * l
}

//CalculateMonsterExperience multiplied by level
func CalculateMonsterExperience(l int) int {
	return MonsterBaseExperience * l
}

//Hero modifiers
const (
	//HeroBaseExperience modifier used to calculate next level experience
	HeroBaseExperience = 80
	//HeroMaxLevel modifier sets maximum level available
	HeroMaxLevel = 10
	//HeroBaseHealth modifier added to each level health calculation
	HeroBaseHealth = 150
	//HeroBaseMana modifier added to each level mana calculation
	HeroBaseMana = 60
	//HeroBaseAttack modifier added to each level attack calculation
	HeroBaseAttack = 15
	//HeroRegenTimeout in seconds, defines how often health and mana regenerates
	HeroRegenTimeout = 8
	//HeroHealthRegen modifier defines health amount regenerated every HeroRegenTimeout seconds
	HeroHealthRegen = 4
	//HeroManaRegen modifier defines mana amount regenerated every HeroRegenTimeout seconds
	HeroManaRegen = 2
)

//CalculateHeroAttack by base attack + multiplied level
func CalculateHeroAttack(l int) int {
	return HeroBaseAttack + int(float64(l)*2.5)
}

//CalculateHeroHealth by base health + multiplied level
func CalculateHeroHealth(l int) int {
	return HeroBaseHealth + l*25
}

//CalculateHeroMana by base mana + multiplied level
func CalculateHeroMana(l int) int {
	return HeroBaseMana + l*12
}

//CalculateHeroLevelExperience by base experience with formula
func CalculateHeroLevelExperience(l int) int {
	i := (l - 1) * HeroBaseExperience
	return i + int(float64(i)*0.5)*(l-1)
}

//Location modifiers
const (
	LocationBaseTimeFrame  = 160
	LocationMonsterTarget  = "@"
	LocationBaseKillsCount = 2
)
