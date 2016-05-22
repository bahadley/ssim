package generator

import (
	"encoding/json"
	"time"

	"github.com/bahadley/ssim/log"
)

type (
	SensorTuple struct {
		Sensor    string  `json:"sensor"`
		Type      string  `json:"type"`
		Timestamp int64   `json:"ts"`
		Data      float64 `json:"data"`
		Aggregate float64 `json:"-"` // Hidden field
	}
)

func Unmarshal(msg []byte, st *SensorTuple) error {
	err := json.Unmarshal(msg, &st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return err
}

func Marshal(st *SensorTuple) ([]byte, error) {
	st.Timestamp = time.Now().UnixNano()

	msg, err := json.Marshal(&st)
	if err != nil {
		log.Warning.Println(err.Error())
	}
	return msg, err
}
