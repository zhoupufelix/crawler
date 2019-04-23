package fetcher

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func Fetch(url string)([]byte,error){
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}
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