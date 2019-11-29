package db

var versions = map[int]string{
	1: `
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

DROP TABLE IF EXISTS gray_goose_bar.users ;
CREATE TABLE IF NOT EXISTS gray_goose_bar.users (
  tg_user_id INT UNSIGNED NOT NULL,
  first_name VARCHAR(45) NOT NULL,
  last_name VARCHAR(45) NULL,
  username VARCHAR(45) NULL,
  api_token VARCHAR(45) NOT NULL,
  player_id VARCHAR(45) NOT NULL,
  balance_gold INT UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (tg_user_id))
ENGINE = InnoDB;

DROP TABLE IF EXISTS gray_goose_bar.cw_player_profile ;
CREATE TABLE IF NOT EXISTS gray_goose_bar.cw_player_profile (
  tg_user_id INT UNSIGNED NOT NULL,
  atk INT NULL,
  def INT NULL,
  max_hp INT NULL,
  exp INT NULL,
  gold INT NOT NULL DEFAULT 0,
  lvl INT NULL,
  pouches INT NOT NULL DEFAULT 0,
  user_name VARCHAR(45) NOT NULL,
  class VARCHAR(6) NULL,
  castle VARCHAR(6) NULL,
  PRIMARY KEY (tg_user_id),
  CONSTRAINT CW_PLAYER_STATS_TG_USER_ID
    FOREIGN KEY (tg_user_id)
    REFERENCES gray_goose_bar.users (tg_user_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

DROP TABLE IF EXISTS gray_goose_bar.user_purchased_items ;
CREATE TABLE IF NOT EXISTS gray_goose_bar.user_purchased_items (
  tg_user_id INT UNSIGNED NOT NULL,
  item VARCHAR(45) NULL,
  count INT NOT NULL DEFAULT 0,
  PRIMARY KEY (tg_user_id),
  CONSTRAINT USER_PURCHASED_ITEMS_TG_ID
    FOREIGN KEY (tg_user_id)
    REFERENCES gray_goose_bar.users (tg_user_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE)
ENGINE = InnoDB;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
`,
}
