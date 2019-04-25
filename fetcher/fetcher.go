package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Fetch(url string)([]byte,error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url,nil)
	if err != nil {
		return nil,err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("Wrong status code:%d ",resp.StatusCode)
	}
	all,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return all,err
}