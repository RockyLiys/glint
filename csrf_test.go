package main

import (
	"fmt"
	"glint/config"
	"glint/pkg/pocs/csrf"
	"glint/plugin"
	"sync"
	"testing"
	"time"
)

func Test_CSRF(t *testing.T) {

	var PluginWg sync.WaitGroup
	data, _ := config.ReadResultConf("result.json")
	myfunc := []plugin.PluginCallback{}
	myfunc = append(myfunc, csrf.Csrfeval)
	pluginInternal := plugin.Plugin{
		PluginName:   "Csrf",
		PluginId:     plugin.Csrf,
		MaxPoolCount: 20,
		Callbacks:    myfunc,
		Timeout:      200 * time.Second,
	}
	pluginInternal.Init()
	PluginWg.Add(1)
	Progress := 0.
	args := plugin.PluginOption{
		PluginWg: &PluginWg,
		Progress: &Progress,
		IsSocket: false,
		Data:     data,
		TaskId:   999,
		// Sendstatus: &PliuginsMsg,
	}
	go func() {
		pluginInternal.Run(args)
	}()
	PluginWg.Wait()
	fmt.Println("exit...")
}
