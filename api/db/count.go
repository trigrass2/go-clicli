package db

import (
	"github.com/cliclitv/go-clicli/api/def"
	"database/sql"
)

func GetCommentCount(pid int) (*def.Count, error) {
	stmtCount, err := dbConn.Prepare("SELECT COUNT(*) FROM comments WHERE pid = ?")
	if err != nil {
		return nil, err
	}

	var pv, cv int
	err = stmtCount.QueryRow(pid).Scan(&cv)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	res := &def.Count{Pid: pid, Pv: pv, Cv: cv}

	defer stmtCount.Close()

	return res, nil
}
