package handle

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Handle struct {
	JingdongUrl     string
	JingdongReferer string
	pvid            string
}

var ha Handle

func init() {
	ha.JingdongUrl = "https://dd-search.jd.com/?terminal=pc&newjson=1&ver=2&zip=1&t=1649056874120&curr_url=search.jd.com/Search"
	ha.JingdongReferer = "https://search.jd.com/Search?enc=utf-8&pvid="
}

func PVid() (id string) {
	return Uuid(uuid.New().String())
}

func Uuid(uuid string) (s string) {
	split := strings.Split(uuid, "-")
	for _, v := range split {
		s = s + v
	}
	return
}

func HandlerUrl(keyWord string, pvid string) (url string) {
	rand.Seed(time.Now().Unix())
	ha.pvid = pvid
	return ha.JingdongUrl + "&pvid=" + pvid + "&key=" + keyWord + "&callback=jQuery" + strconv.Itoa(int(rand.Uint32()/100))
}

func HandlerReferer() (referer string) {
	fmt.Println(ha.JingdongReferer + ha.pvid)
	return ha.JingdongReferer + ha.pvid
}
