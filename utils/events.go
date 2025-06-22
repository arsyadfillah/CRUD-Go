package utils

import (
	"fmt"
	"math/rand"
)

type details struct {
	EventId, Price                          int
	Name, Date, Location, Description, Time string
}

const maxx = 20

type Events [maxx]details

func AddEvents(isEvents *Events, n *int) {

	var continueAdding string
	for continueAdding != "no" {
		if *n > maxx {
			*n = maxx

		}

		fmt.Print("Enter Event Name:")
		fmt.Scan(&isEvents[*n].Name)

		fmt.Print("Enter Event Date:")
		fmt.Scan(&isEvents[*n].Date)

		fmt.Print("Enter Event Time:")
		fmt.Scan(&isEvents[*n].Time)

		fmt.Print("Enter Event Location:")
		fmt.Scan(&isEvents[*n].Location)

		fmt.Print("Enter Event Description:")
		fmt.Scan(&isEvents[*n].Description)

		fmt.Print("Enter Price:")
		fmt.Scan(&isEvents[*n].Price)

		isEvents[*n].EventId = rand.Intn(190)

		*n++
		fmt.Println("Do you want to continue adding events? (yes/no)")
		fmt.Scan(&continueAdding)

	}
}

func ShowEvents(e Events, n int) bool {
	var continueAdding string
	for i := 0; i < n; i++ {
		fmt.Println("Id:", e[i].EventId)
		fmt.Println("Event Name:", e[i].Name)
		fmt.Println("Event Date:", e[i].Date)
		fmt.Println("Event Time:", e[i].Time)
		fmt.Println("Event Location:", e[i].Location)
		fmt.Println("Event Description:", e[i].Description)
		fmt.Println("Event Price:", e[i].Price)
		fmt.Println()

	}
	fmt.Println("Do you want to back? (yes/no)")
	fmt.Scan(&continueAdding)

	return continueAdding == "yes"
}

func UpdateEvents(ev *Events, n int) {
	var option, options string
	var num, value int
	var found bool

	fmt.Println("Which attributes would you like to change?")
	fmt.Scan(&option)

	if option != "Price" {
		fmt.Println("input the desired value")
		fmt.Scan(&options)
	}

	fmt.Println("input the id of the event")
	fmt.Scan(&num)

	for i := 0; i < n; i++ {
		if ev[i].EventId == num {
			switch option {
			case "Name":
				ev[i].Name = options
			case "Date":
				ev[i].Date = options
			case "Time":
				ev[i].Time = options
			case "Location":
				ev[i].Location = options
			case "Description":
				ev[i].Description = options
			case "Price":
				fmt.Println("Input the new price:")
				fmt.Scan(&value)
				ev[i].Price = value
			default:
				fmt.Println("Invalid attribute")
				return
			}
			found = true
		}
	}
	if found {
		fmt.Print("Event has been updated successfully\n")
	} else {
		fmt.Print("The id you inputted is not in the array\n")
	}
}

func DeleteEvent(ev *Events, n *int) {
	var num int
	var notFound bool

	notFound = true
	fmt.Println("Input the id of the event:")
	fmt.Scan(&num)

	for i := 0; i < *n; i++ {
		if ev[i].EventId == num {
			for j := i; j < *n-1; j++ {
				ev[j] = ev[j+1]
			}
			*n--
			i--
			notFound = false
		}
	}
	if notFound {
		fmt.Println("The id you're looking for is not in the array")
	} else {
		fmt.Println("Data has been deleted")
	}
}

func SearchEvent(ev Events, n int) bool {
	var num, left, right, mid, idx int
	var found bool
	var continuee string
	EventIDSortingAsc(&ev, n)
	fmt.Println("input the id of the event")
	fmt.Scan(&num)
	left = 0
	right = n - 1
	idx = -1
	for left <= right && idx == -1 {
		mid = (left + right) / 2
		if ev[mid].EventId == num {
			fmt.Println("Id:", ev[mid].EventId)
			fmt.Println("Event Name:", ev[mid].Name)
			fmt.Println("Event Date:", ev[mid].Date)
			fmt.Println("Event Time:", ev[mid].Time)
			fmt.Println("Event Location:", ev[mid].Location)
			fmt.Println("Event Description:", ev[mid].Description)
			fmt.Println("Event Description:", ev[mid].Price)
			idx = mid
			found = true
		} else if ev[mid].EventId < num {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	if !found {
		fmt.Println("The id you're looking for is not in the array")
	}
	fmt.Println("Do you want to back? (yes/no)")
	fmt.Scan(&continuee)
	return true
}

func EventSortingAsc(ev *Events, n int) {
	//var max int
	//max = 1
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ev[j].Price < ev[idx].Price {
				ev[j], ev[idx] = ev[idx], ev[j]
			}
		}
	}

}
func EventIDSortingAsc(ev *Events, n int) {
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if ev[j].EventId < ev[idx].EventId {
				ev[j], ev[idx] = ev[idx], ev[j]
			}
		}
	}

}
func EventSortingDesc(ev *Events, n int) {
	pass := n - 1
	i := 1
	for i <= pass {
		tempp := ev[i]
		j := i - 1
		for j >= 0 && ev[j].Price < tempp.Price {
			ev[j+1] = ev[j]
			j--
		}
		ev[j+1] = tempp
		i++
	}
}
