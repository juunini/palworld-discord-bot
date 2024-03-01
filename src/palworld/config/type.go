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

	DAY_TIME_SPEED_RATE_MIN                    = 0.1
	DAY_TIME_SPEED_RATE_MAX                    = 5
	NIGHT_TIME_SPEED_RATE_MIN                  = 0.1
	NIGHT_TIME_SPEED_RATE_MAX                  = 5
	EXP_RATE_MIN                               = 0.1
	EXP_RATE_MAX                               = 20
	PAL_CAPTURE_RATE_MIN                       = 0.5
	PAL_CAPTURE_RATE_MAX                       = 2
	PAL_SPAWN_NUM_RATE_MIN                     = 0.5
	PAL_SPAWN_NUM_RATE_MAX                     = 3
	PAL_DAMAGE_RATE_ATTACK_MIN                 = 0.1
	PAL_DAMAGE_RATE_ATTACK_MAX                 = 5
	PAL_DAMAGE_RATE_DEFENSE_MIN                = 0.1
	PAL_DAMAGE_RATE_DEFENSE_MAX                = 5
	PLAYER_DAMAGE_RATE_ATTACK_MIN              = 0.1
	PLAYER_DAMAGE_RATE_ATTACK_MAX              = 5
	PLAYER_DAMAGE_RATE_DEFENSE_MIN             = 0.1
	PLAYER_DAMAGE_RATE_DEFENSE_MAX             = 5
	PLAYER_STOMACH_DECREASE_RATE_MIN           = 0.1
	PLAYER_STOMACH_DECREASE_RATE_MAX           = 5
	PLAYER_STAMINA_DECREASE_RATE_MIN           = 0.1
	PLAYER_STAMINA_DECREASE_RATE_MAX           = 5
	PLAYER_AUTO_HP_REGENE_RATE_MIN             = 0.1
	PLAYER_AUTO_HP_REGENE_RATE_MAX             = 5
	PLAYER_AUTO_HP_REGENE_RATE_IN_SLEEP_MIN    = 0.1
	PLAYER_AUTO_HP_REGENE_RATE_IN_SLEEP_MAX    = 5
	PAL_STOMACH_DECREASE_RATE_MIN              = 0.1
	PAL_STOMACH_DECREASE_RATE_MAX              = 5
	PAL_STAMINA_DECREASE_RATE_MIN              = 0.1
	PAL_STAMINA_DECREASE_RATE_MAX              = 5
	PAL_AUTO_HP_REGENE_RATE_MIN                = 0.1
	PAL_AUTO_HP_REGENE_RATE_MAX                = 5
	PAL_AUTO_HP_REGENE_RATE_IN_SLEEP_MIN       = 0.1
	PAL_AUTO_HP_REGENE_RATE_IN_SLEEP_MAX       = 5
	BUILD_OBJECT_DAMAGE_RATE_MIN               = 0.1
	BUILD_OBJECT_DAMAGE_RATE_MAX               = 3
	BUILD_OBJECT_DETERIORATION_DAMAGE_RATE_MIN = 0
	BUILD_OBJECT_DETERIORATION_DAMAGE_RATE_MAX = 10
	COLLECTION_DROP_RATE_MIN                   = 0.5
	COLLECTION_DROP_RATE_MAX                   = 3
	COLLECTION_OBJECT_HP_RATE_MIN              = 0.5
	COLLECTION_OBJECT_HP_RATE_MAX              = 3
	COLLECTION_OBJECT_RESPAWN_SPEED_RATE_MIN   = 0.5
	COLLECTION_OBJECT_RESPAWN_SPEED_RATE_MAX   = 3
	ENEMY_DROP_ITEM_RATE_MIN                   = 0.5
	ENEMY_DROP_ITEM_RATE_MAX                   = 3
	DROP_ITEM_MAX_NUM_MIN                      = 0
	DROP_ITEM_MAX_NUM_MAX                      = 5000
	BASE_CAMP_WORKER_MAX_NUM_MIN               = 1
	BASE_CAMP_WORKER_MAX_NUM_MAX               = 20
	GUILD_PLAYER_MAX_NUM_MIN                   = 1
	GUILD_PLAYER_MAX_NUM_MAX                   = 100
	PAL_EGG_DEFAULT_HATCHING_TIME_MIN          = 0
	PAL_EGG_DEFAULT_HATCHING_TIME_MAX          = 240
)
