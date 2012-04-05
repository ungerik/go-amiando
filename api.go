package amiando

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

///////////////////////////////////////////////////////////////////////////////
// Api

func NewApi(key string) *Api {
	return &Api{Key: key}
}

type Api struct {
	Key  string
	http http.Client
}

func (self *Api) httpGet(url string) (body []byte, err error) {
	response, err := self.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (self *Api) Call(resourceFormat string, resourceArg interface{}, result ErrorReporter) (err error) {
	result.Reset()

	sep := "?"
	if strings.Contains(resourceFormat, "?") {
		sep = "&"
	}
	resourceFormat = "http://www.amiando.com/api/" + resourceFormat + sep + "apikey=%s&version=1&format=json"
	url := fmt.Sprintf(resourceFormat, resourceArg, self.Key)

	j, err := self.httpGet(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, result)
	if err != nil {
		return err
	}

	return result.Err()
}

func (self *Api) Call_debug(resourceFormat string, resourceArg interface{}, result ErrorReporter) (err error) {
	result.Reset()

	sep := "?"
	if strings.Contains(resourceFormat, "?") {
		sep = "&"
	}
	resourceFormat = "http://www.amiando.com/api/" + resourceFormat + sep + "apikey=%s&version=1&format=json"
	url := fmt.Sprintf(resourceFormat, resourceArg, self.Key)

	fmt.Println("URL: ", url)
	j, err := self.httpGet(url)
	if err != nil {
		return err
	}
	fmt.Println("Result:\n", PrettifyJSON(j))

	// Catch nasty problems in json.Unmarshal
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
			fmt.Println(err)
			debug.PrintStack()
			os.Exit(-1)
		}
	}()

	err = json.Unmarshal(j, result)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", result)

	return result.Err()
}

func (self *Api) Payment(id ID, out interface{}) (err error) {
	type Result struct {
		ResultBase
		Payment interface{} `json:"payment"`
	}
	result := Result{Payment: out}
	return self.Call("payment/%v", id, &result)
}

func (self *Api) TicketIDsOfPayment(paymentID ID) (ids []ID, err error) {
	type Result struct {
		ResultBase
		Tickets []ID `json:"tickets"`
	}
	var result Result
	err = self.Call("payment/%v/tickets", paymentID, &result)
	if err != nil {
		return nil, err
	}
	return result.Tickets, nil
}

func (self *Api) Ticket(id ID, out interface{}) (err error) {
	type Result struct {
		ResultBase
		Ticket interface{} `json:"ticket"`
	}
	result := Result{Ticket: out}
	return self.Call("ticket/%v", id, &result)
}

func (self *Api) User(id ID, out interface{}) (err error) {
	type Result struct {
		ResultBase
		User interface{} `json:"user"`
	}
	result := Result{User: out}
	return self.Call("user/%v", id, &result)
}
