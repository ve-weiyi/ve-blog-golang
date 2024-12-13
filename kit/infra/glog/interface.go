package glog

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})

	Infof(fmt string, args ...interface{})
	Errorf(fmt string, args ...interface{})
}
