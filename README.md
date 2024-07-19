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
