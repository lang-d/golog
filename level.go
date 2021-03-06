package golog

type level struct {
	name  string
	level uint8
}

func (this *level) string() string {

	return this.name
}

func (this *level) isEnable(lel *level) bool {
	if this.level <= lel.level {
		return true
	}
	return false
}

func (this *level) isEnableDetail(lel *level) bool {
	if this.level >= lel.level {
		return true
	}
	return false
}

func (this *level) isEnableError(lel *level) bool {
	if this.level >= lel.level {
		return true
	}
	return false
}

var (
	DEBUG     = &level{name: "debug", level: 0}
	INFO      = &level{name: "info", level: 1}
	WARN      = &level{name: "warn", level: 2}
	ERROR     = &level{name: "error", level: 3}
	EXCEPTION = &level{name: "exception", level: 4}
	FATAL     = &level{name: "fatal", level: 5}
)
