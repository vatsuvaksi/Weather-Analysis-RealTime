package apis

// import (
// 	"testing"

// 	"github.com/valyala/fasthttp"
// )

// // var mutex *sync.Mutex

// func TestGetRealTimeData(t *testing.T) {

// 	response, _ := GetRealTimeData("48.8567,2.3508")

// 	expected := fasthttp.StatusOK
// 	if response.StatusCode() != expected {
// 		t.Errorf("Expected %d, got %d", expected, response.StatusCode())
// 		t.Fail()
// 	} else {
// 		t.Log("json received for GetRealTimeData(48.8567,2.3508) --> ", response.Body())
// 	}
// }
