package main

import (
	`fmt`
	`math`
	`runtime`
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v;
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main()  {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(math.Pow(2,3))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today+0:
		fmt.Println("today")
	case today+1:
		fmt.Println("tomorrow")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}