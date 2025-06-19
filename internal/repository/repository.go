package repository

import "gorm.io/gorm"

type Repository struct {
	CharacterMastersRepository             CharacterMastersRepository
	UpgradeCostMastersRepository           UpgradeCostMastersRepository
	EventMastersRepository                 EventMastersRepository
	GuildBattleMattingCharactersRepository GuildBattleMattingCharactersRepository
	GuildBattleMattingResourcesRepository  GuildBattleMattingResourcesRepository
	GuildBattleMattingsRepository          GuildBattleMattingsRepository
	GuildMastersRepository                 GuildMastersRepository
	GuildMembersRepository                 GuildMembersRepository
	GuildResourceTransactionLogsRepository GuildResourceTransactionLogsRepository
	GuildResourcesRepository               GuildResourcesRepository
	UserCharactersRepository               UserCharactersRepository
	UserFacilitiesRepository               UserFacilitiesRepository
	UserFacilityMappingRepository          UserFacilityMappingRepository
	UserResourcesRepository                UserResourcesRepository
	UsersRepository                        UsersRepository
	FacilityMastersRepository              FacilityMastersRepository
	GuildsRepository                       GuildsRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		CharacterMastersRepository:             NewCharacterMastersRepository(db),
		UpgradeCostMastersRepository:           NewUpgradeCostMastersRepository(db),
		EventMastersRepository:                 NewEventMastersRepository(db),
		GuildBattleMattingCharactersRepository: NewGuildBattleMattingCharactersRepository(db),
		GuildBattleMattingResourcesRepository:  NewGuildBattleMattingResourcesRepository(db),
		GuildBattleMattingsRepository:          NewGuildBattleMattingsRepository(db),
		GuildMastersRepository:                 NewGuildMastersRepository(db),
		GuildMembersRepository:                 NewGuildMembersRepository(db),
		GuildResourceTransactionLogsRepository: NewGuildResourceTransactionLogsRepository(db),
		GuildResourcesRepository:               NewGuildResourcesRepository(db),
		UserCharactersRepository:               NewUserCharactersRepository(db),
		UserFacilitiesRepository:               NewUserFacilitiesRepository(db),
		UserFacilityMappingRepository:          NewUserFacilityMappingRepository(db),
		UserResourcesRepository:                NewUserResourcesRepository(db),
		UsersRepository:                        NewUsersRepository(db),
		FacilityMastersRepository:              NewFacilityMastersRepository(db),
		GuildsRepository:                       NewGuildsRepository(db),
	}
}
