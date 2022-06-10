package throttle

var throttles = map[string]int{}

// ShouldThrottle
//  for systems that don't need to be evaluated on
//  every update loop, we can slow them down with a throttle
func ShouldThrottle(key string, interval int) bool {
	if _, ok := throttles[key]; !ok {
		throttles[key] = 0
	}
	throttles[key]++
	return throttles[key]%interval != 0
}
