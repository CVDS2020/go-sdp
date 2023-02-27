package sdp

import (
	"fmt"
	"gitee.com/sy_183/sdp/errors"
	"regexp"
	"strconv"
)

type RtpMap struct {
	Type   int
	Format string
	Rate   int
}

func NewRtpMap(typ int, format string, rate int) *RtpMap {
	return &RtpMap{
		Type:   typ,
		Format: format,
		Rate:   rate,
	}
}

func ParseRtpMap(s string) (*RtpMap, error) {
	rtpMap := new(RtpMap)
	matched := regexp.MustCompile(`^(\d+) +(\S*?)/(\d+)$`).FindStringSubmatch(s)
	if matched == nil {
		return nil, errors.New("invalid rtpmap format")
	}
	if parsed, err := strconv.ParseUint(matched[1], 10, 32); err != nil {
		return nil, err
	} else {
		rtpMap.Type = int(parsed)
	}
	rtpMap.Format = matched[2]
	if parsed, err := strconv.ParseUint(matched[3], 10, 32); err != nil {
		return nil, err
	} else {
		rtpMap.Rate = int(parsed)
	}
	return rtpMap, nil
}

func (m *RtpMap) String() string {
	return fmt.Sprintf("%d %s/%d", m.Type, m.Format, m.Rate)
}
