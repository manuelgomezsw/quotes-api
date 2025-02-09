package cache

import (
	"log"
	"math/rand"
	"reflect"
	"time"
)

var rng *rand.Rand

func init() {
	rand.NewSource(time.Now().UnixNano())
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// GetRandomItem obtiene un ID aleatorio de una colección, utilizando la caché.
func GetRandomItem(cachePrefix string, getItemByID func(int64) (interface{}, error), loadMinMaxFromDB func() (int64, int64, error)) (interface{}, error) {
	// Obtener valores mínimo y máximo desde la caché
	minID, maxID, err := getMinMaxIDFromCache(cachePrefix, loadMinMaxFromDB)
	if err != nil {
		return nil, err
	}

	rangeVal := maxID - minID + 1

	// Intentar encontrar un elemento válido
	for {
		randomID := rng.Int63n(rangeVal) + minID

		item, err := getItemByID(randomID)
		if err != nil {
			return nil, err
		}

		// Verificar si el item retornado es "vacío".
		zeroValue := reflect.Zero(reflect.TypeOf(item)).Interface()
		if reflect.DeepEqual(item, zeroValue) {
			continue
		}

		return item, nil
	}
}

// getMinMaxIDFromCache obtiene los valores mínimo y máximo desde la caché.
func getMinMaxIDFromCache(cachePrefix string, loadMinMaxFromDB func() (int64, int64, error)) (int64, int64, error) {
	var minID, maxID int64

	// Intentar obtener valores desde la caché
	if err := Get(cachePrefix+"_minID", &minID); err == nil {
		if err := Get(cachePrefix+"_maxID", &maxID); err == nil {
			return minID, maxID, nil
		}
	}

	// Si los valores no están en caché, cargarlos desde la BD
	minID, maxID, err := loadMinMaxFromDB()
	if err != nil {
		return 0, 0, err
	}

	// Guardar los valores en caché
	if err := Set(cachePrefix+"_minID", minID); err != nil {
		log.Printf("No se pudo almacenar %s_minID en caché: %v", cachePrefix, err)
	}
	if err := Set(cachePrefix+"_maxID", maxID); err != nil {
		log.Printf("No se pudo almacenar %s_maxID en caché: %v", cachePrefix, err)
	}

	return minID, maxID, nil
}
