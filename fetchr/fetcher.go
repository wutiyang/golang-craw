package fetchr

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Microsecond)

func Fetch(url string) ([]byte, error)  {
	<- rateLimiter

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
		//panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code ", resp.StatusCode)

		return nil, fmt.Errorf("wrong status code :%d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

//func determineEncode(r io.Reader) encoding.Encoding {
//	bytes, err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//		return ""
//	}
//
//	e,_,_ := charset.DetermineEncoding(bytes, "")
//	return e
//}