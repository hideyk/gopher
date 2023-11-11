package iteration

import "fmt"

func doWork() {
	// Infinite
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// Three component
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration", i+1)
	}

	// While equivalent
	n := 0
	for n <= 5 {
		fmt.Println("Iteration", n)
		n++
	}

	// Do while equivalent
	n = 1
	for ok := true; ok; ok = n != 5 {
		fmt.Println("Iteration", n)
		n++
	}

	// For each
	s1 := []int{1, 2, 3}
	mp := map[string]string{"key1": "value1", "key2": "value2"}
	for index, value := range s1 {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}

	for key, val := range mp {
		fmt.Printf("Key %s, value: %s\n", key, val)
	}
}
