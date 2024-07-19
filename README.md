# fy


```go
func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		//ForceColors:      true, // 强制输出带色彩的日志
		FullTimestamp:    true, // 显示完整的时间戳
		QuoteEmptyFields: true,
	})
}
```

https://www.bookstack.cn/read/GoExpertProgramming/chapter08-README.md
https://www.topgoer.com/%E6%95%B0%E6%8D%AE%E5%BA%93%E6%93%8D%E4%BD%9C/gorm/gorm%E6%9F%A5%E8%AF%A2.html

https://github.com/graham924/share-code-operator-study/tree/main

```go
func waitSignal(f func()) {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	switch <-sigs {
	case syscall.SIGINT, syscall.SIGTERM:
		log.Println("Shutdown quickly, bye...")
	case syscall.SIGQUIT:
		log.Println("Shutdown gracefully, bye...")
		// do graceful shutdown
	}
	f()
}

func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}


```
