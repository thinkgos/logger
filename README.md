# logger

zap logger with lumberjack

[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/thinkgos/logger?tab=doc)
[![codecov](https://codecov.io/gh/thinkgos/logger/branch/main/graph/badge.svg)](https://codecov.io/gh/thinkgos/logger)
[![Tests](https://github.com/thinkgos/logger/actions/workflows/ci.yml/badge.svg)](https://github.com/thinkgos/logger/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/thinkgos/logger)](https://goreportcard.com/report/github.com/thinkgos/logger)
[![License](https://img.shields.io/github/license/thinkgos/logger)](https://raw.githubusercontent.com/thinkgos/logger/main/LICENSE)
[![Tag](https://img.shields.io/github/v/tag/thinkgos/logger)](https://github.com/thinkgos/logger/tags)

## Features

## Usage

[example](./examples/main.go)

### Installation

Use go get.

```bash
    go get github.com/thinkgos/logger
```

Then import the package into your own code.

```bash
    import "github.com/thinkgos/logger"
```

### 示例

[embedmd]:# (examples/main.go go)
```go
package main

import (
	"context"

	"github.com/thinkgos/logger"
)

type ctxKey struct{}

func ExampleHook(ctx context.Context) logger.Field {
	s, ok := ctx.Value(ctxKey{}).(string)
	if !ok {
		return logger.Skip()
	}
	return logger.String("ctx_key1", s)
}
func TmpHook(ctx context.Context) logger.Field {
	return logger.String("tmp_key1", "tmp_val1")
}

func main() {
	l := logger.NewLogger(
		logger.WithLevel(logger.DebugLevel.String()),
		logger.WithFormat(logger.FormatJson),
	).
		SetCallerLevel(logger.WarnLevel)
	logger.ReplaceGlobals(l)
	logger.SetDefaultValuer(ExampleHook)

	ctx := context.WithValue(context.Background(), ctxKey{}, "ctx_val1")

	logger.WithValuer(TmpHook).DebugContext(ctx, "Debug1")
	logger.WithNewValuer(TmpHook).DebugContext(ctx, "Debug2")

	logger.SetLevel(logger.WarnLevel)
	logger.DebugContext(ctx, "Debug3")

	err := logger.SetLevelWithText(logger.DebugLevel.String())
	_ = err
	logger.DebugContext(ctx, "Debug4")

	// 先改成warn等级方便测试
	logger.SetLevel(logger.WarnLevel)
	if logger.Enabled(logger.InfoLevel) {
		logger.InfoContext(ctx, "Info1")
	}
	if logger.V(logger.InfoLevel) {
		logger.InfoContext(ctx, "Info2")
	}
	logger.SetLevel(logger.DebugLevel)

	logger.Named("叫个名字").DebugContext(ctx, " Debug5")

	logger.With(
		logger.String("name", "jack"),
		logger.Int("age", 18),
	).InfoContext(ctx, " Debug6")

	// `log.Print`风格的日志
	logger.DebugContext(ctx, "Debug")
	logger.InfoContext(ctx, "Info")
	logger.WarnContext(ctx, "Warn")
	logger.ErrorContext(ctx, "Error")
	logger.DPanicContext(ctx, "DPanic")

	// `log.Printf`风格的日志
	logger.DebugfContext(ctx, "Debugf: %s", "debug")
	logger.InfofContext(ctx, "Infof: %s", "info")
	logger.WarnfContext(ctx, "Warnf: %s", "warn")
	logger.ErrorfContext(ctx, "Errorf: %s", "error")
	logger.DPanicfContext(ctx, "DPanicf: %s", "dPanic")

	// 松散键值对风格的日志
	logger.DebugwContext(ctx, "Debugw", "k1", "v1", logger.String("k2", "v2"))
	logger.InfowContext(ctx, "Infow", "k1", "v1", logger.String("k2", "v2"))
	logger.WarnwContext(ctx, "Warnw", "k1", "v1", logger.String("k2", "v2"))
	logger.InfowContext(ctx, "Infow", "k1", "v1", logger.String("k2", "v2"))
	logger.ErrorwContext(ctx, "Errorw", "k1", "v1", logger.String("k2", "v2"))
	logger.DPanicwContext(ctx, "DPanicw", "k1", "v1", logger.String("k2", "v2"))

	// 纯结构型的日志
	logger.DebugxContext(ctx, "Debugx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.InfoxContext(ctx, "Infox", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.WarnxContext(ctx, "Warnx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.InfoxContext(ctx, "Infox", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.ErrorxContext(ctx, "Errorx", logger.String("k1", "v1"), logger.String("k2", "v2"))
	logger.DPanicxContext(ctx, "DPanicx", logger.String("k1", "v1"), logger.String("k2", "v2"))
}
```

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
