package repository

import (
	"fmt"
	"os"
	"quotes-api/internal/util/mysql"
	"strings"
)

const (
	basePathSqlQueries = "sql/tags"

	fileSqlCreateQuoteTags  = "CreateQuoteTags.sql"
	fileSqlCreateReviewTags = "CreateReviewTags.sql"
	fileSqlDeleteQuoteTags  = "DeleteQuoteTags.sql"
	fileSqlDeleteReviewTags = "DeleteReviewTags.sql"
)

// CreateTags crea los tags en base al ID y la cadena de tags.
func CreateTags(quoteID int64, reviewID int, tags []string) error {
	fileSQL, err := getFileSQLCreateTags(quoteID, reviewID)
	if err != nil {
		return err
	}

	// Leer la consulta base desde el archivo.
	query, err := getSQLQuery(fileSQL)
	if err != nil {
		return err
	}

	var values []string
	var args []interface{}
	id := getUniqueID(quoteID, reviewID)

	// Separar y limpiar los tags.
	for _, tag := range tags {
		tag = strings.TrimSpace(tag)
		if tag == "" {
			continue
		}
		values = append(values, "(?, ?)")
		args = append(args, id, tag)
	}

	// Si no hay tags válidos, salimos sin error.
	if len(values) == 0 {
		return nil
	}

	// Construir la query final.
	bulkInsert := query + strings.Join(values, ",")
	_, err = mysql.ClientDB.Exec(bulkInsert, args...)
	return err
}

// DeleteTags elimina los tags asociados al ID dado.
func DeleteTags(quoteID int64, reviewID int) error {
	fileSQL, err := getFileSQLDeleteTags(quoteID, reviewID)
	if err != nil {
		return err
	}

	id := getUniqueID(quoteID, reviewID)
	query, err := getSQLQuery(fileSQL)
	if err != nil {
		return err
	}

	_, err = mysql.ClientDB.Exec(query, id)
	return err
}

func getUniqueID(quoteID int64, reviewID int) interface{} {
	var id interface{}
	if quoteID > 0 {
		id = quoteID
	} else if reviewID > 0 {
		id = reviewID
	} else {
		return fmt.Errorf("no se proporcionó un quoteID ni reviewID válido")
	}

	return id
}

// getSQLQuery lee el archivo SQL y devuelve su contenido.
func getSQLQuery(fileName string) (string, error) {
	path := fmt.Sprintf("%s/%s", basePathSqlQueries, fileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error leyendo el archivo SQL %s: %w", path, err)
	}
	return string(data), nil
}

// getFileSQLCreateTags determina qué archivo usar para crear tags.
func getFileSQLCreateTags(quoteID int64, reviewID int) (string, error) {
	if quoteID > 0 {
		return fileSqlCreateQuoteTags, nil
	} else if reviewID > 0 {
		return fileSqlCreateReviewTags, nil
	}
	return "", fmt.Errorf("no se proporcionó un quoteID o reviewID válido")
}

// getFileSQLDeleteTags determina qué archivo usar para borrar tags.
func getFileSQLDeleteTags(quoteID int64, reviewID int) (string, error) {
	if quoteID > 0 {
		return fileSqlDeleteQuoteTags, nil
	} else if reviewID > 0 {
		return fileSqlDeleteReviewTags, nil
	}
	return "", fmt.Errorf("no se proporcionó un quoteID o reviewID válido")
}
