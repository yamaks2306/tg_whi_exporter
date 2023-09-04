package pg

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/yamaks2306/tg_whi_exporter/config"
	"github.com/yamaks2306/tg_whi_exporter/docker"
	"github.com/yamaks2306/tg_whi_exporter/util"
)

func GetDbUsersCount(config_pg_docker config.ConfigPgDocker, config_pg config.ConfigPg) int {
	pg_ip := docker.GetPgContainerIP(config_pg_docker.PgContainerName, config_pg_docker.PgContainerNetwork)
	port, err := strconv.ParseInt(config_pg.DbPort, 0, 64)
	util.CheckError(err)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", pg_ip, port, config_pg.DbUser, config_pg.DbPassword, config_pg.DbName)
	db, err := sql.Open("postgres", psqlconn)
	util.CheckError(err)
	defer db.Close()

	err = db.Ping()
	util.CheckError(err)

	rows, err := db.Query(`SELECT count(*) FROM users`)
	util.CheckError(err)
	defer rows.Close()

	var count int

	for rows.Next() {
		err = rows.Scan(&count)
		util.CheckError(err)
	}

	return count
}
