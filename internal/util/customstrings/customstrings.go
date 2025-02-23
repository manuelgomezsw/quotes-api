package customstrings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// CustomString encapsula un string y permite aplicar modificaciones encadenadas.
type StringBuilder struct {
	s string
}

// NewStringBuilder crea una nueva instancia de StringBuilder con el string inicial.
func NewStringBuilder(initial string) *StringBuilder {
	return &StringBuilder{s: initial}
}

// Build devuelve el string resultante.
func (sb *StringBuilder) Build() string {
	return sb.s
}

// TrimSpace elimina espacios en blanco al inicio y final.
func (sb *StringBuilder) TrimSpace() *StringBuilder {
	sb.s = strings.TrimSpace(sb.s)
	return sb
}

// RemoveEndPeriod elimina el punto final si existe.
func (sb *StringBuilder) RemoveEndPeriod() *StringBuilder {
	if sb.s == "" {
		return sb
	}
	if strings.HasSuffix(sb.s, ".") {
		sb.s = sb.s[:len(sb.s)-1]
	}
	return sb
}

// RemoveSpecialCharacters elimina comillas, corchetes y convierte a minúsculas.
func (sb *StringBuilder) RemoveSpecialCharacters() *StringBuilder {
	if sb.s == "" {
		return sb
	}

	// Elimina dobles comillas.
	sb.s = strings.ReplaceAll(sb.s, "\"", "")
	// Elimina dobles comillas cursivas.
	sb.s = strings.ReplaceAll(sb.s, "“", "")
	sb.s = strings.ReplaceAll(sb.s, "”", "")
	// Elimina comillas simples.
	sb.s = strings.ReplaceAll(sb.s, "'", "")
	// Elimina corchetes.
	sb.s = strings.ReplaceAll(sb.s, "[", "")
	sb.s = strings.ReplaceAll(sb.s, "]", "")

	// Convertir a minúsculas y limpiar espacios.
	sb.s = strings.ToLower(sb.s)
	sb.s = strings.TrimSpace(sb.s)

	return sb
}

// CapitalizeFirst convierte la primera letra a mayúscula.
func (sb *StringBuilder) CapitalizeFirst() *StringBuilder {
	if sb.s == "" {
		return sb
	}
	r, size := utf8.DecodeRuneInString(sb.s)
	sb.s = string(unicode.ToUpper(r)) + sb.s[size:]
	return sb
}

// TruncateString devuelve los primeros n caracteres de s concatenando ... al final.
func TruncateString(s string, n int) string {
	// Si el número total de runas es menor o igual que n, se retorna s completo.
	if utf8.RuneCountInString(s) <= n {
		return s
	}

	// Convertir el string a slice de runas.
	runes := []rune(s)
	return string(runes[:n-3]) + "..."
}
