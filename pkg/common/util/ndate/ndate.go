package ndate


import (
	
	"fmt"
	
	"strconv"
	"strings"
	"time"
)


const layout = "20060102"




   

func NextDate(now time.Time, date string, repeat string) (string, error){
	

	nowParse := now.Format(layout)
	
	dateParse, err := time.Parse(layout, date)
	
	if err != nil {
        fmt.Println("Error parsing date:", err)
        return "", err
    }
	//years
	if strings.HasPrefix(repeat, "y"){
		
		for {
			repeatDate := dateParse.AddDate(1, 0, 0).Format(layout)
			if repeatDate > nowParse {
				
				return repeatDate, nil
			}
			dateParse = dateParse.AddDate(1, 0, 0)
		}
	}
	//days
	if strings.HasPrefix(repeat, "d"){
		if len(strings.Split(repeat, " "))!= 2{
			return "", fmt.Errorf("invalid repeat format: %s", repeat)
		}
		days := strings.Split(repeat, " ")[1]
		daysInt, err := strconv.Atoi(days)
		if err != nil {
			return "", fmt.Errorf("invalid repeat format: %s", repeat)
		}
		
		if daysInt < 0 || daysInt >400 {
			return "", fmt.Errorf("invalid repeat format: %s", repeat)
		}
		
		for {
			repeatDate := dateParse.AddDate(0, 0, daysInt).Format(layout)
			if repeatDate > nowParse {
				return repeatDate, nil
			}
			dateParse = dateParse.AddDate(0, 0, daysInt)
			
		}
	}
	

	return "", fmt.Errorf("invalid repeat format: %s", repeat)
}



