package alias

type Alias interface {
	SetAlias(key string, value string) Alias
	UnsetAlias(key string) Alias
}

type alias struct {
	table map[string]string
}

func NewAlias() Alias {
	return &alias{table: map[string]string{}}
}

func (a *alias) SetAlias(key string, value string) Alias {
	a.table[key] = value
	return a
}

func (a *alias) UnsetAlias(key string) Alias {
	delete(a.table, key)
	return a
}

