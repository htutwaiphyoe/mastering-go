package recursion

import "fmt"

func TowerOfHanoi(n int, source, auxiliary, target string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", source, target)
		return
	}

	TowerOfHanoi(n-1, source, target, auxiliary)

	fmt.Printf("Move disk %d from %s to %s\n", n, source, target)

	TowerOfHanoi(n-1, auxiliary, source, target)
}
