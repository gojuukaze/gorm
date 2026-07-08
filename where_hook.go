package gorm

const UpdateWhereHookKey = ".gorm.:update_where_hook"

// SqlHookType "{table_name}_{action}" : ["where sql", values]
// ex:
//
//	{
//		"user_update": [ "id=? and name=?", 1, "bob" ],
//	}
type SqlHookType map[string][]any
