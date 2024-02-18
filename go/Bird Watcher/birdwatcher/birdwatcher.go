package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalBirds := 0

	for _, birds := range birdsPerDay {
		totalBirds += birds
	}

	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsOfThatWeek := []int{}
	birdsPerDayLen := len(birdsPerDay)
	numOfWeeks := birdsPerDayLen / 7
	weekInit := 0
	weekEnd := 7
	
	for i := 1; i <= numOfWeeks; i++ {	
		if i == week {
			birdsOfThatWeek = append(birdsOfThatWeek, birdsPerDay[weekInit:weekEnd]...)
			break
		}

		weekInit += 7
		weekEnd += 7
	}
	
	return TotalBirdCount(birdsOfThatWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	lenn := len(birdsPerDay)

	for i := 0; i < lenn; i += 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}

	return birdsPerDay
}
