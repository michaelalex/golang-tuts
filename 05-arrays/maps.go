package arrays

func MakeMap() map[string]int {
	newMap := make(map[string]int)
	newMap["John"] = 56
	newMap["Anna"] = 55
	return newMap
}

func GetItemFromMap(key string, thisMap map[string]int) (string, int) {
	return key, thisMap[key]
}

func GetMapLength(thisMap map[string]int) int {
	return len(thisMap)
}

func RemoveKeyFromMap(key string, thisMap map[string]int) {
	delete(thisMap, key)
}
