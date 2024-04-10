package dto

type Log struct {
	className  string
	methodName string
	paramName  string
	value      interface{}
}

func NewLog(className string, methodName string, paramName string, value interface{}) *Log {
	return &Log{
		className:  className,
		methodName: methodName,
		paramName:  paramName,
		value:      value,
	}
}

func (l *Log) Value() interface{} {
	return l.value
}

func (l *Log) SetValue(value interface{}) {
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
