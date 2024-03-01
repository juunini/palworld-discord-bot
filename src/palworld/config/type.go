package palworld_config

type Config struct {
	Difficulty                          string
	DayTimeSpeedRate                    float64
	NightTimeSpeedRate                  float64
	ExpRate                             float64
	PalCaptureRate                      float64
	PalSpawnNumRate                     float64
	PalDamageRateAttack                 float64
	PalDamageRateDefense                float64
	PlayerDamageRateAttack              float64
	PlayerDamageRateDefense             float64
	PlayerStomachDecreaceRate           float64
	PlayerStaminaDecreaceRate           float64
	PlayerAutoHPRegeneRate              float64
	PlayerAutoHpRegeneRateInSleep       float64
	PalStomachDecreaceRate              float64
	PalStaminaDecreaceRate              float64
	PalAutoHPRegeneRate                 float64
	PalAutoHpRegeneRateInSleep          float64
	BuildObjectDamageRate               float64
	BuildObjectDeteriorationDamageRate  float64
	CollectionDropRate                  float64
	CollectionObjectHpRate              float64
	CollectionObjectRespawnSpeedRate    float64
	EnemyDropItemRate                   float64
	DeathPenalty                        string
	EnablePlayerToPlayerDamage          bool
	EnableFriendlyFire                  bool
	EnableInvaderEnemy                  bool
	ActiveUNKO                          bool
	EnableAimAssistPad                  bool
	EnableAimAssistKeyboard             bool
	DropItemMaxNum                      int
	DropItemMaxNum_UNKO                 int
	BaseCampMaxNum                      int
	BaseCampWorkerMaxNum                int
	DropItemAliveMaxHours               float64
	AutoResetGuildNoOnlinePlayers       bool
	AutoResetGuildTimeNoOnlinePlayers   float64
	GuildPlayerMaxNum                   int
	PalEggDefaultHatchingTime           float64
	WorkSpeedRate                       float64
	IsMultiplay                         bool
	IsPvP                               bool
	CanPickupOtherGuildDeathPenaltyDrop bool
	EnableNonLoginPenalty               bool
	EnableFastTravel                    bool
	IsStartLocationSelectByMap          bool
	ExistPlayerAfterLogout              bool
	EnableDefenseOtherGuildPlayer       bool
	CoopPlayerMaxNum                    int
	ServerPlayerMaxNum                  int
	ServerName                          string
	ServerDescription                   string
	AdminPassword                       string
	ServerPassword                      string
	PublicPort                          int
	PublicIP                            string
	RCONEnabled                         bool
	RCONPort                            int
	Region                              string
	UseAuth                             bool
	BanListURL                          string
}

const (
	DIFFICULTY_CUSTOM = "None"
	DIFFICULTY_CASUAL = "Casual"
	DIFFICULTY_NORMAL = "Normal"
	DIFFICULTY_HARD   = "Hard"

	DEATH_PENALTY_DROP_ALL                = "All"
	DEATH_PENALTY_DROP_ITEM_AND_EQUIPMENT = "ItemAndEquipment"
	DEATH_PENALTY_DROP_ITEM               = "Item"
	DEATH_PENALTY_NONE                    = "None"
)
