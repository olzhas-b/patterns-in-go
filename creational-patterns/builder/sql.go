package builder

type Isql interface {
	Table(string) *SqlBuilder
	Select(...string) *SqlBuilder
	Where(string, ...any) *SqlBuilder
	BuildSQL() string
}
