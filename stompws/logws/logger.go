package logws

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}
