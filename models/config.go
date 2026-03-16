package models

// GetConfig retrieves a config value by key
func GetConfig(key string) (string, error) {
	var value string
	err := DB.QueryRow(`SELECT value FROM config WHERE key = ?`, key).Scan(&value)
	if err != nil {
		return "", err
	}
	return value, nil
}

// GetAllConfig returns all config as a map
func GetAllConfig() (map[string]string, error) {
	rows, err := DB.Query(`SELECT key, value FROM config`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]string)
	for rows.Next() {
		var k, v string
		if err := rows.Scan(&k, &v); err != nil {
			return nil, err
		}
		result[k] = v
	}
	return result, rows.Err()
}

// SetConfig sets a config value
func SetConfig(key, value string) error {
	_, err := DB.Exec(
		`INSERT INTO config (key, value) VALUES (?, ?) ON CONFLICT(key) DO UPDATE SET value = excluded.value`,
		key, value,
	)
	return err
}
