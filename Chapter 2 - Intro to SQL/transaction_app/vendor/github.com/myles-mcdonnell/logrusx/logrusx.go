package logrusx

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

const (
	DEBUG Level = 1 << iota
	INFO
	ERROR

	Epoch     string = "Epoch"
	Label     string = "Label"
	Service   string = "Service"
	Hostname  string = "Hostname"
	IpAddress string = "IpAddress"

	DefaultStackTraceLogLevel = DEBUG | ERROR

	DefaultTimestampFormat = time.RFC3339
)

var (
	traceLevels *Level

	entryFactory LogEntryFactory

	hostname, ipAddress string
)

type (
	Field string

	Level byte

	Entry interface {
		Write(...interface{})
		WithField(string, interface{}) Entry
		WithContext(context.Context) Entry
		WithFields(log.Fields) Entry
		Level() Level
		Entry() *log.Entry
	}

	EntryFunc func() Entry

	EntryFromContextFunc func(context.Context) Entry

	entry struct {
		entry             *log.Entry
		level             Level
		instanceAugmenter EntryAugmenter
		contextHandler    ContextHandler
		timestampFormat   string
	}

	logEntryFactory struct {
		service           string
		instanceAugmenter EntryAugmenter
		contextHandler    func(context.Context, Entry)
		timestampFormat   string
	}

	LogEntryFactory interface {
		MakeEntry(level Level, label string) EntryFunc
	}

	EntryAugmenter func(Entry) Entry
	ContextHandler func(context.Context, Entry)
)

func Infof(label string, args ...interface{}) Entry {
	return entryFactory.MakeEntry(DEBUG, fmt.Sprintf(label, args...))()
}

func Info(label string) Entry {
	return entryFactory.MakeEntry(DEBUG, label)()
}

func Debugf(label string, args ...interface{}) Entry {
	return entryFactory.MakeEntry(DEBUG, fmt.Sprintf(label, args...))()
}

func Debug(label string) Entry {
	return entryFactory.MakeEntry(DEBUG, label)()
}

func Errorf(label string, args ...interface{}) Entry {
	return entryFactory.MakeEntry(ERROR, fmt.Sprintf(label, args...))()
}

func Error(label string) Entry {
	return entryFactory.MakeEntry(ERROR, label)()
}

func Init(level Level, out io.Writer, formatter log.Formatter, factory LogEntryFactory) {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(out)

	var err error
	hostname, err = os.Hostname()
	if err != nil {
		hostname = "failed to determine"
	}

	ipAddress = getIpAddress()

	switch level {
	case DEBUG:
		log.SetLevel(log.DebugLevel)
	case INFO:
		log.SetLevel(log.InfoLevel)
	case ERROR:
		log.SetLevel(log.ErrorLevel)
	}

	traceLevels = LevelPtr(DefaultStackTraceLogLevel)

	entryFactory = factory
}

func GetLogLevel() Level {
	switch log.GetLevel() {
	case log.DebugLevel:
		return DEBUG
	case log.InfoLevel:
		return INFO
	case log.ErrorLevel:
		return ERROR
	}

	return INFO
}

func StackTraceLevel() *Level {
	return traceLevels
}

func SetStackTraceLevels(stackTraceLevel Level) {
	traceLevels = LevelPtr(stackTraceLevel)
}

func ClearStackTraceLevels() {
	traceLevels = nil
}

func ParseLevel(level string) (*Level, error) {
	switch strings.ToUpper(strings.Trim(level, "")) {
	case "DEBUG":
		return LevelPtr(DEBUG), nil
	case "INFO":
		return LevelPtr(INFO), nil
	case "ERROR":
		return LevelPtr(ERROR), nil
	default:
		return nil, fmt.Errorf("unrecognised log level : %v", level)
	}
}

func LevelPtr(level Level) *Level {
	return &level
}

func (entry *entry) Level() Level {
	return entry.level
}

func (entry *entry) Entry() *log.Entry {
	return entry.entry
}

func (entry *entry) Write(args ...interface{}) {
	etry := addInstanceFields(entry, args)
	switch etry.Level() {
	case DEBUG:
		etry.Entry().Debug(args...)
	case ERROR:
		etry.Entry().Error(args...)
	case INFO:
		etry.Entry().Info(args...)
	}
}

func (entry *entry) WithField(key string, value interface{}) Entry {
	entry.entry = entry.entry.WithField(key, value)
	return entry
}

func (entry *entry) WithFields(fields log.Fields) Entry {
	entry.entry = entry.entry.WithFields(fields)
	return entry
}

func (entry *entry) WithContext(ctx context.Context) Entry {
	entry.contextHandler(ctx, entry)
	return entry
}

func addInstanceFields(entry *entry, args []interface{}) Entry {
	var etry Entry = entry
	if traceLevels != nil && ((*traceLevels & entry.level) != 0) {
		for index, arg := range args {
			if err, ok := arg.(error); ok {
				etry = etry.WithField(
					fmt.Sprintf("ErrorStack_%v", index),
					fmt.Sprintf("%+v", err))
			}
		}
	}

	return entry.instanceAugmenter(etry.WithField(Epoch, time.Now().Format(entry.timestampFormat)))
}

func newLogEntry(level Level, service string, label string, instanceAugmenter EntryAugmenter, contextHandler func(context.Context, Entry), timestampFormat string) Entry {
	return &entry{
		entry: log.NewEntry(log.StandardLogger()).
			WithField(Service, service).
			WithField(Label, label),
		level:             level,
		instanceAugmenter: instanceAugmenter,
		contextHandler:    contextHandler,
		timestampFormat:   timestampFormat,
	}
}

func (factory *logEntryFactory) MakeEntry(level Level, label string) EntryFunc {
	return func() Entry {
		return newLogEntry(level, factory.service, label, factory.instanceAugmenter, factory.contextHandler, factory.timestampFormat).
			WithField(Hostname, hostname).
			WithField(IpAddress, ipAddress)
	}
}

func getIpAddress() (address string) {
	address = ""
	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			address += addr.String()
		}
	}

	return
}

func Context(contextHandler ContextHandler) func(*logEntryFactory) {
	return func(factory *logEntryFactory) {
		factory.contextHandler = contextHandler
	}
}

func Augmenter(augmenter EntryAugmenter) func(*logEntryFactory) {
	return func(factory *logEntryFactory) {
		factory.instanceAugmenter = augmenter
	}
}

func NewLogEntryFactory(service string, options ...func(*logEntryFactory)) LogEntryFactory {
	factory := &logEntryFactory{service: service, timestampFormat: DefaultTimestampFormat}

	for _, opt := range options {
		opt(factory)
	}

	if factory.instanceAugmenter == nil {
		factory.instanceAugmenter = func(entry Entry) Entry { return entry }
	}

	if factory.contextHandler == nil {
		factory.contextHandler = func(context.Context, Entry) {}
	}

	return factory
}
