package dto

type Log struct {
	className  string
	methodName string
	paramName  string
	value      any
}

func newLog(className string, methodName string, paramName string, value any) *Log {
	return &Log{
		className:  className,
		methodName: methodName,
		paramName:  paramName,
		value:      value,
	}
}

func OfLog(className string, methodName string, paramName string, value any) *Log {
	return newLog(className, methodName, paramName, value)
}

func (l *Log) Value() any {
	return l.value
}

func (l *Log) SetValue(value any) {
	l.value = value
}

func (l *Log) ParamName() string {
	return l.paramName
}

func (l *Log) SetParamName(paramName string) {
	l.paramName = paramName
}

func (l *Log) MethodName() string {
	return l.methodName
}

func (l *Log) SetMethodName(methodName string) {
	l.methodName = methodName
}

func (l *Log) ClassName() string {
	return l.className
}

func (l *Log) SetClassName(className string) {
	l.className = className
}
