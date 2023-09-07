package pg

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/yamaks2306/tg_whi_exporter/config"
	"github.com/yamaks2306/tg_whi_exporter/docker"
)

func GetDbUsersCount(config_pg_docker config.ConfigPgDocker, config_pg config.ConfigPg) (int, error) {
	pg_ip, err := docker.GetPgContainerIP(config_pg_docker.PgContainerName, config_pg_docker.PgContainerNetwork)
	if err != nil {
		return 0, err
	}

	port, err := strconv.ParseInt(config_pg.DbPort, 0, 64)
	if err != nil {
		return 0, err
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg_ip, port, config_pg.DbUser, config_pg.DbPassword, config_pg.DbName)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		return 0, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return 0, err
	}

	rows, err := db.Query(`SELECT count(*) FROM users`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
