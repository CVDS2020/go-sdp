package sdp

import (
	"fmt"
	"gitee.com/sy_183/common/uns"
	"log"
	"strings"
	"testing"
	"time"
)

func TestSDPDecode(t *testing.T) {
	sdp := strings.ReplaceAll(`v=0
o=jdoe 2890844526 2890842807 IN IP4 10.47.16.5
s=SDP Seminar
i=A Seminar on the session description protocol
u=http://www.example.com/seminars/sdp.pdf
e=j.doe@example.com (Jane Doe)
p=12345
b=CT:154798
t=2873397496 2873404696
r=7d 1h 0 25h
k=clear:ab8c4df8b8f4as8v8iuy8re
a=recvonly
m=audio 49170 RTP/AVP 0
m=video 51372 RTP/AVP 99
b=AS:66781
k=prompt
a=rtpmap:99 h263-1998/90000
y=0000000000
f=v/0/0/0/0/0a/0/0/0
`, "\n", "\r\n")
	message, err := Decode(uns.StringToBytes(sdp))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(message)
}

func TestSDPEncode(t *testing.T) {
	msg := Message{
		Version: 0,
		Origin: Origin{
			Username:       "video_storage",
			SessionID:      int64(TimeToNTP(time.Now())),
			SessionVersion: int64(TimeToNTP(time.Now())),
			NetworkType:    "IN",
			AddressType:    "IP4",
			Address:        "192.168.1.129",
		},
		Name:   "Play",
		Timing: []Timing{{Start: time.Time{}, End: time.Time{}}},
		Attributes: Attributes{
			{Key: "recvonly"},
		},
		Medias: Medias{{
			Description: MediaDescription{
				Type:     "video",
				Port:     5004,
				Protocol: "RTP/AVP",
				Formats:  []string{"96", "98"},
			},
			Attributes: Attributes{
				{Key: "rtpmap", Value: "96 PS/90000"},
				{Key: "rtpmap", Value: "98 H264/90000"}},
		}},
	}
	session := msg.Append(nil)
	data := session.AppendTo(nil)
	fmt.Println(uns.BytesToString(data))
}
