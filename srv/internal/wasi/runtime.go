package wasi

import (
	"context"
	"os"

	"github.com/andrescosta/goico/pkg/io"
	"github.com/tetratelabs/wazero"
)

type WasmRuntime struct {
	cacheDir      *string
	cache         wazero.CompilationCache
	runtimeConfig wazero.RuntimeConfig
}

func NewWasmRuntime(ctx context.Context, tempDir string) (*WasmRuntime, error) {
	wruntime := &WasmRuntime{}

	// Creates a directory to store wazero cache
	if err := io.CreateDirIfNotExist(tempDir); err != nil {
		return nil, err
	}
	cacheDir, err := os.MkdirTemp(tempDir, "cache")
	if err != nil {
		return nil, err
	}
	wruntime.cacheDir = &cacheDir

	cache, err := wazero.NewCompilationCacheWithDir(cacheDir)
	if err != nil {
		wruntime.Close(ctx)
		return nil, err
	}
	wruntime.cache = cache

	runtimeConfig := wazero.NewRuntimeConfig().WithCompilationCache(cache)
	wruntime.runtimeConfig = runtimeConfig
	return wruntime, nil
}

func (r *WasmRuntime) Close(ctx context.Context) {
	if r.cacheDir != nil {
		os.RemoveAll(*r.cacheDir)
	}
	if r.cache != nil {
		r.cache.Close(ctx)
	}
}
