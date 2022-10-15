package DiscordPerm

const (
	CREATE_INVITE              = 0x1
	KICK_MEMBERS               = 0x2  // Needs 2FA
	BAN_MEMBERS                = 0x4  // Needs 2FA
	ADMINISTRATOR              = 0x8  // Needs 2FA
	MANAGE_CHANNELS            = 0x10 // Needs 2FA
	MANAGE_SERVER              = 0x20 // Needs 2FA
	ADD_REACTIONS              = 0x40
	VIEW_AUDIT_LOG             = 0x80
	PRIORITY_SPEAKER           = 0x100
	VIDEO                      = 0x200
	VIEW_CHANNELS              = 0x400
	SEND_MESSAGES              = 0x800
	SEND_TTS_MESSAGES          = 0x1000
	MANAGE_MESSAGES            = 0x2000 // Needs 2FA
	EMBED_LINKS                = 0x4000
	ATTACH_FILES               = 0x8000
	READ_MESSAGE_HISTORY       = 0x10000
	MENTION_EVERYONE           = 0x20000
	USE_EXTERNAL_EMOJI         = 0x40000
	VIEW_SERVER_INSIGHTS       = 0x80000
	CONNECT                    = 0x100000
	MUTE_MEMBERS               = 0x400000
	DEAFEN_MEMBERS             = 0x800000
	SPEAK                      = 0x200000
	MOVE_MEMBERS               = 0x1000000
	USE_VOICE_ACTIVITY         = 0x2000000
	CHANGE_NICKNAME            = 0x4000000
	MANAGE_NICKNAMES           = 0x8000000
	MANAGE_ROLES               = 0x10000000 // Needs 2FA
	MANAGE_WEBHOOKS            = 0x20000000 // Needs 2FA
	MANAGE_EMOJIS_AND_STICKERS = 0x40000000 // Needs 2FA
	USE_APPLICATION_COMMANDS   = 0x80000000
	REQUEST_TO_SPEAK           = 0x100000000
	MANAGE_EVENTS              = 0x200000000 // Needs 2FA
	MANAGE_THREADS             = 0x400000000 // Needs 2FA
	CREATE_PUBLIC_THREADS      = 0x800000000
	CREATE_PRIVATE_THREADS     = 0x1000000000
	SEND_MESSAGES_IN_THREAD    = 0x4000000000
	USE_EXTERNAL_STICKER       = 0x2000000000
	START_ACTIVITIES           = 0x8000000000
)
