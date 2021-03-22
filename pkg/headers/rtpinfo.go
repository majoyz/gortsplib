package headers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/majoyz/gortsplib/pkg/base"
)

// RTPInfoEntry is an entry of an RTP-Info header.
type RTPInfoEntry struct {
	URL            *base.URL
	SequenceNumber uint16
	Timestamp      uint32
}

// RTPInfo is a RTP-Info header.
type RTPInfo []*RTPInfoEntry

// Read decodes a RTP-Info header.
func (h *RTPInfo) Read(v base.HeaderValue) error {
	if len(v) == 0 {
		return fmt.Errorf("value not provided")
	}

	if len(v) > 1 {
		return fmt.Errorf("value provided multiple times (%v)", v)
	}

	for _, tmp := range strings.Split(v[0], ",") {
		e := &RTPInfoEntry{}

		for _, kv := range strings.Split(tmp, ";") {
			tmp := strings.SplitN(kv, "=", 2)
			if len(tmp) != 2 {
				return fmt.Errorf("unable to parse key-value (%v)", kv)
			}

			k, v := tmp[0], tmp[1]
			switch k {
			case "url":
				vu, err := base.ParseURL(v)
				if err != nil {
					return err
				}
				e.URL = vu

			case "seq":
				vi, err := strconv.ParseUint(v, 10, 16)
				if err != nil {
					return err
				}
				e.SequenceNumber = uint16(vi)

			case "rtptime":
				vi, err := strconv.ParseUint(v, 10, 32)
				if err != nil {
					return err
				}
				e.Timestamp = uint32(vi)

			default:
				return fmt.Errorf("invalid key: %v", k)
			}
		}

		*h = append(*h, e)
	}

	return nil
}

// Clone clones a RTPInfo.
func (h RTPInfo) Clone() *RTPInfo {
	nh := &RTPInfo{}
	for _, e := range h {
		*nh = append(*nh, &RTPInfoEntry{
			URL:            e.URL,
			SequenceNumber: e.SequenceNumber,
			Timestamp:      e.Timestamp,
		})
	}
	return nh
}

// Write encodes a RTP-Info header.
func (h RTPInfo) Write() base.HeaderValue {
	rets := make([]string, len(h))

	for i, e := range h {
		rets[i] = "url=" + e.URL.String() +
			";seq=" + strconv.FormatUint(uint64(e.SequenceNumber), 10) +
			";rtptime=" + strconv.FormatUint(uint64(e.Timestamp), 10)
	}

	return base.HeaderValue{strings.Join(rets, ",")}
}
