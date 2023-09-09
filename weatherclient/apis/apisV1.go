package apis

import (
	"fmt"
	"os"
)

func GetRealTimeData() {
	fmt.Println("Get Real Time Data ", os.Getenv("API_KEY"))
}
