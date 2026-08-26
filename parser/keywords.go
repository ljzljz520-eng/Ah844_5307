package parser

var SQLKeywords = []string{"CREATE", "TABLE", "ALTER", "DROP", "SELECT", "FROM", "WHERE", "INSERT", "INTO", "VALUES", "UPDATE", "SET", "DELETE", "JOIN", "LEFT", "RIGHT", "INNER", "OUTER", "ON", "AS", "AND", "OR", "NOT", "NULL", "PRIMARY", "KEY", "FOREIGN", "REFERENCES", "UNIQUE", "CHECK", "DEFAULT", "INDEX", "VIEW", "TRIGGER", "PROCEDURE", "FUNCTION", "BEGIN", "END", "COMMIT", "ROLLBACK", "GRANT", "REVOKE", "GROUP", "BY", "ORDER", "HAVING", "LIMIT", "OFFSET", "UNION", "ALL", "DISTINCT", "ASC", "DESC", "CASE", "WHEN", "THEN", "ELSE", "CAST", "CONVERT", "WITH", "RECURSIVE", "RETURNING", "CONSTRAINT", "SCHEMA", "DATABASE", "IF", "EXISTS", "TEMPORARY", "TEMP", "CASCADE", "RESTRICT", "MATERIALIZED", "WINDOW", "OVER", "PARTITION", "SEQUENCE", "SERIAL", "IDENTITY", "AUTOINCREMENT", "ENGINE", "CHARSET", "COLLATE", "COMMENT", "COLUMN", "RENAME", "MODIFY", "ADD", "DROP", "ANALYZE", "VACUUM", "EXPLAIN", "DESCRIBE", "SHOW", "USE", "DECLARE", "WHILE", "LOOP", "EXCEPTION", "RAISE", "RETURN", "OPEN", "CLOSE", "FETCH", "EXECUTE", "PREPARE", "DEALLOCATE", "LOCK", "SHARE", "NOWAIT", "WAIT", "READ", "WRITE", "ISOLATION", "SERIALIZABLE", "REPEATABLE", "COMMITTED", "UNCOMMITTED", "BOOLEAN", "INTEGER", "DECIMAL", "NUMERIC", "VARCHAR", "TEXT", "DATE", "TIME", "TIMESTAMP", "JSON", "XML", "BINARY", "UUID", "ARRAY", "ENUM", "GEOMETRY", "GEOGRAPHY", "MONEY", "ROWID", "CURRENT_DATE", "CURRENT_TIMESTAMP", "TRUE", "FALSE"}

func IsKeyword(s string) bool {
	for _, k := range SQLKeywords {
		if s == k {
			return true
		}
	}
	return false
}
func KeywordCount() int { return len(SQLKeywords) }
func KeywordAt(i int) string {
	if i < 0 || i >= len(SQLKeywords) {
		return ""
	}
	return SQLKeywords[i]
}
