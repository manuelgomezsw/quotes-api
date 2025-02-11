package repository

import (
	"fmt"
	"os"
	"quotes-api/internal/util/mysql"
)

const (
	basePathSqlQueries = "sql/misc"

	fileSqlGetAuthors = "GetAuthors.sql"
	fileSqlGetWorks   = "GetWorks.sql"
)

func GetAuthors() ([]string, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetAuthors))
	if err != nil {
		return nil, err
	}

	resultTopics, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return nil, err
	}

	var topics []string
	for resultTopics.Next() {
		var topic string

		err = resultTopics.Scan(&topic)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}

func GetWorks() ([]string, error) {
	query, err := os.ReadFile(fmt.Sprintf("%s/%s", basePathSqlQueries, fileSqlGetWorks))
	if err != nil {
		return nil, err
	}

	resultTopics, err := mysql.ClientDB.Query(string(query))
	if err != nil {
		return nil, err
	}

	var topics []string
	for resultTopics.Next() {
		var topic string

		err = resultTopics.Scan(&topic)
		if err != nil {
			return nil, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}
