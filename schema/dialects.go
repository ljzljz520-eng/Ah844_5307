package schema

var TypeAliases = map[string]string{
	"INT": "INTEGER", "INTEGER": "INTEGER", "BIGINT": "INTEGER", "SMALLINT": "INTEGER", "TINYINT": "INTEGER", "NUMBER": "INTEGER", "NUMERIC": "INTEGER", "DECIMAL": "INTEGER",
	"VARCHAR": "TEXT", "VARCHAR2": "TEXT", "CHAR": "TEXT", "NCHAR": "TEXT", "NVARCHAR": "TEXT", "NVARCHAR2": "TEXT", "TEXT": "TEXT", "CLOB": "TEXT",
	"DATE": "TEXT", "DATETIME": "TEXT", "TIMESTAMP": "TEXT", "TIME": "TEXT", "BOOLEAN": "INTEGER", "BOOL": "INTEGER", "BIT": "INTEGER",
}
var ReservedWords = []string{"select", "from", "where", "join", "group", "order", "having", "limit", "offset", "insert", "update", "delete", "create", "alter", "drop", "table", "index", "view", "user", "role", "grant", "revoke", "primary", "foreign", "key", "references", "constraint", "unique", "check", "default", "null", "not", "and", "or", "in", "exists", "as", "on", "into", "values", "set", "case", "when", "then", "else", "end", "union", "all", "distinct", "asc", "desc", "database", "schema", "trigger", "procedure", "function", "sequence", "cursor", "with", "recursive", "returning", "conflict", "replace", "engine", "collate", "comment", "column", "rename", "add", "modify", "partition", "cluster", "owner", "authorization", "grantable", "cascade", "restrict", "temporary", "temp", "materialized", "recursive", "over", "partitioned", "window", "rank", "row_number", "lag", "lead", "coalesce", "nullif", "cast", "convert", "extract", "dateadd", "datediff", "current_date", "current_timestamp", "true", "false", "begin", "commit", "rollback", "savepoint", "release", "analyze", "vacuum", "explain", "describe", "show", "use", "go", "declare", "if", "else", "while", "loop", "exception", "raise", "return", "open", "close", "fetch", "execute", "prepare", "deallocate", "lock", "share", "nowait", "wait", "isolation", "serializable", "repeatable", "committed", "uncommitted", "read", "write", "binary", "json", "xml", "array", "enum", "geometry", "geography", "money", "uuid", "identity", "serial", "autoincrement", "rowid", "oid", "xmin", "xmax", "xmin"}

func AliasType(t string) string {
	if v, ok := TypeAliases[t]; ok {
		return v
	}
	return NormalizeType(t)
}
func IsReserved(word string) bool {
	for _, v := range ReservedWords {
		if v == word {
			return true
		}
	}
	return false
}
func QuoteIfNeeded(word string) string {
	if IsReserved(word) {
		return "\"" + word + "\""
	}
	return word
}
