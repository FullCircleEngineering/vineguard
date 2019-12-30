package ttndb

import (
	"encoding/json"
	"strconv"
	"time"
)

type LSB50Msgs struct {
	Msgs []LSB50Msg
}

type LSB50Msg struct {
	BatV     float64
	TempC    float64
	DeviceId string
	Time     time.Time
}

type LSB50MsgRaw struct {
	BatV     float64 `json:"BatV"`
	TempC    string  `json:"TempC"`
	DeviceId string  `json:"device_id"`
	Time     string  `json:"time"`
}

func (o *LSB50Msgs) UnmarshalJSON(j []byte) error {
	var (
		rawJSON  []LSB50MsgRaw
		outTempC float64
		outTime  time.Time
	)
	err := json.Unmarshal(j, &rawJSON)
	if err != nil {
		return err
	}
	o.Msgs = make([]LSB50Msg, len(rawJSON))
	for i, rawMsg := range rawJSON {
		if outTempC, err = strconv.ParseFloat(rawMsg.TempC, 64); err != nil {
			return err
		}
		if outTime, err = time.Parse(time.RFC3339Nano, rawMsg.Time); err != nil {
			return err
		}

		o.Msgs[i] = LSB50Msg{
			BatV:     rawMsg.BatV,
			TempC:    outTempC,
			DeviceId: rawMsg.DeviceId,
			Time:     outTime,
		}
	}
	return nil
}
