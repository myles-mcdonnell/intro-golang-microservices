package logrusx

import (
	"fmt"
	"github.com/myles-mcdonnell/jsonx"
	log "github.com/sirupsen/logrus"
)

//JSONFormatter is a custom JSON formatter that can serialise any struct
type JSONFormatter struct {
	TimestampFormat string
	Indent          bool
}

// Format renders a single log entry
func (f *JSONFormatter) Format(entry *log.Entry) ([]byte, error) {
	data := make(log.Fields, len(entry.Data))
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	var (
		serialized []byte
		err        error
	)

	if f.Indent {
		serialized, err = jsonx.MarshalIndentWithOptions(data, "", "  ", jsonx.MarshalOptions{SkipUnserializableFields: true})
	} else {
		serialized, err = jsonx.MarshalWithOptions(data, jsonx.MarshalOptions{SkipUnserializableFields: true})
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}

	return append(serialized, '\n'), nil
}
