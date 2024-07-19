func InitLog(level string, local *LocalLog) {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		fmt.Printf("Init loglevel fail, use default info level. error: %v\n", err)
		lvl = logrus.InfoLevel
	}

	logrus.SetLevel(lvl)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&XlinkFormatter{})

	if !local.Enable {
		return
	}
  ...
}

// 自定义日志格式
func (x *XlinkFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	levelText := strings.ToUpper(entry.Level.String())
	if len(levelText) > LevelLen {
		levelText = levelText[0:LevelLen]
	}
	levelText = fmt.Sprintf("[%-5s]", levelText)

	entry.Message = strings.TrimSuffix(entry.Message, "\n")

	b.WriteString(levelText)
	b.WriteString("[")
	b.WriteString(entry.Time.Format("2006-01-02 15:04:05"))
	b.WriteString("]")
	b.WriteString("[")
	b.WriteString(path.Base(entry.Caller.File))
	b.WriteString(":")
	b.WriteString(strconv.Itoa(entry.Caller.Line))
	b.WriteString("]")
	b.WriteString(entry.Caller.Function)
	b.WriteString(" ")
	b.WriteString(entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}
