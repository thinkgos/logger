package logger

import (
	"context"
	"runtime"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

const loggerPackage = "github.com/thinkgos/logger"

type CallerCore struct {
	level        AtomicLevel
	Skip         int
	SkipPackages []string
	Caller       func(depth int, skipPackages ...string) Field
}

func NewCallerCore() *CallerCore {
	return &CallerCore{
		level:        NewAtomicLevelAt(ErrorLevel),
		Skip:         0,
		SkipPackages: nil,
		Caller:       DefaultCallerFile,
	}
}

// AddSkip add the number of callers skipped by caller annotation.
func (c *CallerCore) AddSkip(callerSkip int) *CallerCore {
	c.Skip += callerSkip
	return c
}

// AddSkipPackage add the caller skip package.
func (c *CallerCore) AddSkipPackage(vs ...string) *CallerCore {
	c.SkipPackages = append(c.SkipPackages, vs...)
	return c
}

// SetLevel set the caller level.
func (c *CallerCore) SetLevel(lv Level) *CallerCore {
	c.level.SetLevel(lv)
	return c
}

// Level returns the minimum enabled log level.
func (c *CallerCore) Level() Level {
	return c.level.Level()
}

// Enabled returns true if the given level is at or above this level.
func (c *CallerCore) Enabled(lvl Level) bool {
	return c.level.Enabled(lvl)
}

// UseExternalLevel use external level, which controller by user.
func (c *CallerCore) UseExternalLevel(l AtomicLevel) *CallerCore {
	c.level = l
	return c
}

// UnderlyingLevel get underlying level.
func (c *CallerCore) UnderlyingLevel() AtomicLevel {
	return c.level
}

// DefaultCallerFile caller file.
func DefaultCallerFile(depth int, skipPackages ...string) Field {
	var file string
	var line int
	var ok bool

	for i := depth; i < depth+10; i++ {
		_, file, line, ok = runtime.Caller(i)
		if ok && !skipPackage(file, skipPackages...) {
			break
		}
	}
	return zap.String("file", file+":"+strconv.Itoa(line))
}

// DefaultCaller caller.
func DefaultCaller(depth int, skipPackages ...string) Field {
	var file string
	var line int
	var ok bool

	for i := depth; i < depth+10; i++ {
		_, file, line, ok = runtime.Caller(i)
		if ok && !skipPackage(file, skipPackages...) {
			break
		}
	}
	idx := strings.LastIndexByte(file, '/')
	return zap.String("caller", file[idx+1:]+":"+strconv.Itoa(line))
}

// File returns a Valuer that returns a pkg/file:line description of the caller.
func File(depth int, skipPackages ...string) HookFunc {
	return func(context.Context) Field {
		return DefaultCallerFile(depth, skipPackages...)
	}
}

// Caller returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int, skipPackages ...string) HookFunc {
	return func(context.Context) Field {
		return DefaultCaller(depth, skipPackages...)
	}
}

func skipPackage(file string, skipPackages ...string) bool {
	if strings.HasSuffix(file, "_test.go") {
		return false
	}
	if strings.Contains(file, loggerPackage) {
		return true
	}
	for _, p := range skipPackages {
		if strings.Contains(file, p) {
			return true
		}
	}
	return false
}
