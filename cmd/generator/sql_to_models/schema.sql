-- *_masters    マスターデータを管理
-- *_resources  資産
-- *_logs       履歴(historyは使用しない)

-- --------------------------------------------------------------------------------------------------------------------------
-- マスターデータ
-- --------------------------------------------------------------------------------------------------------------------------
-- 戦闘用キャラクターのマスターデータ
CREATE TABLE IF NOT EXISTS character_masters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `name` VARCHAR(255) NOT NULL DEFAULT "",
    `level` INT NOT NULL DEFAULT 0,
    `reality` INT NOT NULL DEFAULT 0,
    `experience_point` BIGINT NOT NULL DEFAULT 0,
    `attack` INT NOT NULL DEFAULT 0,
    `defense` INT NOT NULL DEFAULT 0,
    `speed` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- イベント開催用のマスターデータ
CREATE TABLE IF NOT EXISTS event_masters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `since` DATETIME,
    `until` DATETIME,
    `title` VARCHAR(255) NOT NULL DEFAULT "",
    `description` TEXT NOT NULL DEFAULT "",
    `event_type` INT NOT NULL DEFAULT 0,
    `parent_event_id` VARCHAR(255),
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ギルドに関するもの
CREATE TABLE IF NOT EXISTS guild_masters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `level` INT NOT NULL DEFAULT 0,
    `max_members` INT NOT NULL DEFAULT 0,
    -- `upgrade_cost` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- 施設のマスターデータ
CREATE TABLE IF NOT EXISTS facility_masters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `size_x` INT NOT NULL DEFAULT 1,
    `size_y` INT NOT NULL DEFAULT 1,
    `name` VARCHAR(255) NOT NULL DEFAULT "",
    `description` TEXT NOT NULL DEFAULT "",
    `facility_type` INT NOT NULL DEFAULT 0,
    `parent_facility_id` VARCHAR(255),
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ユーザーを含むレベルという概念のものを強化する時に必要なコストを定義
CREATE TABLE IF NOT EXISTS upgrade_cost_masters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `item_id` VARCHAR(255) NOT NULL DEFAULT "",
    `amount` INT NOT NULL DEFAULT 1,
    `name` VARCHAR(255) NOT NULL DEFAULT "",
    `level` INT NOT NULL DEFAULT 0,
    `description` TEXT NOT NULL DEFAULT "",
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- --------------------------------------------------------------------------------------------------------------------------
-- ユーザーに関するもの
-- --------------------------------------------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS users(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `name` VARCHAR(255) DEFAULT "",
    `level` BIGINT NOT NULL,
    `introduction` TEXT NOT NULL DEFAULT "",
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ユーザーのリソース状況について
-- 対象ユーザーとリソース種別によってunique
CREATE TABLE IF NOT EXISTS user_resources(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL,
    `balance` BIGINT NOT NULL,
    `resource_type` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- PRIMARY KEY (`id`)
    -- UNIQUE `user_resource_type` (`user_id`, `resource_type`)
);

-- ユーザーの戦闘用キャラ
CREATE TABLE IF NOT EXISTS user_characters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL,
    `character_master_id` VARCHAR(255) NOT NULL,
    `level` BIGINT NOT NULL DEFAULT 1,
    `experience_point` BIGINT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ユーザーの施設の所有状況について
CREATE TABLE IF NOT EXISTS user_facilities(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL,
    `facility_master_id` VARCHAR(255) NOT NULL,
    `level` BIGINT NOT NULL DEFAULT 1,
    `name` VARCHAR(255) NOT NULL,
    `is_setting` TINYINT(1) DEFAULT 0,
    `last_approve_at` DATETIME,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ユーザーの施設の配置状況
CREATE TABLE IF NOT EXISTS user_facility_mapping(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL,
    `mapping_index` VARCHAR(255) NOT NULL,
    `mapping` TEXT,
    `is_over` TINYINT(1) DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- PRIMARY KEY (`id`)
    -- index (`user_id`)
    -- UNIQUE user_mapping_index(`user_id`, `mapping_index`)
);

-- --------------------------------------------------------------------------------------------------------------------------
-- ギルド情報
-- --------------------------------------------------------------------------------------------------------------------------
-- 作成されたギルド
CREATE TABLE IF NOT EXISTS guilds(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `name` VARCHAR(255) DEFAULT "",
    `level` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- ギルドのメンバー一覧
CREATE TABLE IF NOT EXISTS guild_members(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `guild_id` BIGINT NOT NULL,
    `user_id` VARCHAR(255) NOT NULL UNIQUE,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    index (`guild_id`)
    PRIMARY KEY (`id`)
);

-- ギルドの資産に関するもの
CREATE TABLE IF NOT EXISTS guild_resources(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL UNIQUE,
    `balance` BIGINT NOT NULL,
    `resource_type` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- UNIQUE `resource_balance` (`resource_type`,`balance`),
    -- PRIMARY KEY (`id`)
);

-- ギルド収支ログ
CREATE TABLE IF NOT EXISTS guild_resource_transaction_logs(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `user_id` VARCHAR(255) NOT NULL UNIQUE,
    `balance` BIGINT NOT NULL,
    `resource_type` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- --------------------------------------------------------------------------------------------------------------------------
-- ギルドバトルイベント
-- --------------------------------------------------------------------------------------------------------------------------
-- ギルドバトルのマッチング情報
CREATE TABLE IF NOT EXISTS guild_battle_mattings(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `event_id` VARCHAR(255) NOT NULL DEFAULT "",
    `winner_guild_id` VARCHAR(255) NOT NULL DEFAULT "",
    `guild_id_a` VARCHAR(255) NOT NULL DEFAULT "",
    `guild_id_b` VARCHAR(255) NOT NULL DEFAULT "",
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,    
    PRIMARY KEY (`id`)
);

-- ギルドバトルに設定された戦闘用キャラクター達
CREATE TABLE IF NOT EXISTS guild_battle_matting_characters(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `matting_id` VARCHAR(255) NOT NULL DEFAULT "",
    `character_id` VARCHAR(255) NOT NULL DEFAULT "",
    `user_id` VARCHAR(255) NOT NULL DEFAULT "",
    `user_character_index` INT NOT NULL DEFAULT 0,
    `guild_id` VARCHAR(255) NOT NULL DEFAULT "",
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- unique matting_user_character_index (`matting_id`, `user_id`, `user_character_index`),
    -- PRIMARY KEY (`id`),
    -- INDEX (`matting_id`, `parent_user_id`)
);

-- ギルドバトルに設定されたリソース状況
CREATE TABLE IF NOT EXISTS guild_battle_matting_resources(
    `id` VARCHAR(255) NOT NULL UNIQUE,
    `matting_id` VARCHAR(255) NOT NULL DEFAULT "",
    `resource_type` VARCHAR(255) NOT NULL DEFAULT "",
    `guild_id` VARCHAR(255) NOT NULL DEFAULT "",
    `amount` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    -- PRIMARY KEY (`id`),
    -- INDEX (`matting_id`),
);
