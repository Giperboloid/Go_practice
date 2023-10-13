package kata


// Kata: Two Oldest Ages

// The two oldest ages function/method needs to be completed. 
// It should take an array of numbers as its argument and return the two highest numbers within the array.
// The returned value should be an array in the format [second oldest age,  oldest age].
// The order of the numbers passed in could be any order. The array will always include at least 2 items. 
// If there are two or more oldest age, then return both of them in array format.

// Example:
// [1, 2, 10, 8] --> [8, 10]
// [1, 5, 87, 45, 8, 8] --> [45, 87]
// [1, 3, 10, 0]) --> [3, 10]

func TwoOldestAges(ages []int) [2]int {
  
  var prevMax, curMax int = ages[0], ages[1]

  if ages[0] > ages[1] { 
    curMax = ages[0] 
    prevMax = ages[1]
  } else { 
    prevMax = ages[0] 
	  curMax = ages[1]
  }
  
  for idx, value := range ages {
    if idx >= 2 {
      if value > curMax {
         prevMax = curMax
         curMax = value
      } else { if value > prevMax { prevMax = value } }
    }
}
  return [2] int {prevMax, curMax}
}
