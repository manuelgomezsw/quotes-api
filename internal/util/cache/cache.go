package cache

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/allegro/bigcache/v3"
	"strconv"
	"time"
)

var (
	instance *bigcache.BigCache
)

// Cache define la interfáz común para cachés compatible con BigCache.
type Cache interface {
	Set(key interface{}, value interface{}) error
	Get(key interface{}) (interface{}, error)
}

// BigCacheWrapper adapta BigCache a la interfaz Cache.
type BigCacheWrapper struct {
	cache *bigcache.BigCache
}

// NewBigCacheWrapper inicializa la caché global de BigCache
func NewBigCacheWrapper() error {
	if instance == nil {
		config := bigcache.Config{
			Shards:             128,              // Número de shards para concurrencia
			LifeWindow:         24 * time.Hour,   // Tiempo de vida de los elementos
			CleanWindow:        30 * time.Minute, // Intervalo de limpieza
			MaxEntriesInWindow: 1000 * 10 * 60,   // Estimación de entradas en el período
			MaxEntrySize:       500,              // Tamaño máximo por entrada
			Verbose:            false,            // Logs de BigCache (desactivado)
			HardMaxCacheSize:   0,                // Sin límite de memoria forzado
		}

		newCacheInstance, err := bigcache.New(context.Background(), config)
		if err != nil {
			return err
		}
		instance = newCacheInstance
	}

	return nil
}

// Set almacena un valor en caché con una clave genérica
func Set(key interface{}, value interface{}) error {
	if instance == nil {
		return errors.New("NewBigCacheWrapper: instance not initialized")
	}

	// Convertir la clave a string
	keyStr, err := toStringKey(key)
	if err != nil {
		return err
	}

	// Serializar el valor a JSON antes de guardarlo
	valueJSON, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return instance.Set(keyStr, valueJSON)
}

// Get recupera un valor desde la caché y lo deserializa al tipo correcto
func Get(key interface{}, target interface{}) error {
	if instance == nil {
		return errors.New("NewBigCacheWrapper: instance not initialized")
	}

	// Convertir la clave a string
	keyStr, err := toStringKey(key)
	if err != nil {
		return err
	}

	// Obtener los datos desde BigCache
	valueJSON, err := instance.Get(keyStr)
	if err != nil {
		return err
	}

	// Deserializar en el objeto de destino
	return json.Unmarshal(valueJSON, target)
}

// toStringKey convierte claves a string (compatibilidad con BigCache)
func toStringKey(key interface{}) (string, error) {
	switch v := key.(type) {
	case string:
		return v, nil
	case int:
		return strconv.Itoa(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	default:
		return "", errors.New("clave no soportada")
	}
}
