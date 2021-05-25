package main

import "fmt"

/**
https://leetcode.com/problems/time-based-key-value-store/

Create a time-based key-value store class TimeMap, that supports two operations.

1. set(string key, string value, int timestamp)
Stores the key and value, along with the given timestamp.

2. get(string key, int timestamp)
Returns a value such that set(key, value, timestamp_prev) was called previously, with timestamp_prev <= timestamp.
If there are multiple such values, it returns the one with the largest timestamp_prev.
If there are no values, it returns the empty string ("").


Example 1:
Input: inputs = ["TimeMap","set","get","get","set","get","get"], inputs = [[],["foo","bar",1],["foo",1],["foo",3],["foo","bar2",4],["foo",4],["foo",5]]
Output: [null,null,"bar","bar",null,"bar2","bar2"]
Explanation:
TimeMap kv;
kv.set("foo", "bar", 1); // store the key "foo" and value "bar" along with timestamp = 1
kv.get("foo", 1);  // output "bar"
kv.get("foo", 3); // output "bar" since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 ie "bar"
kv.set("foo", "bar2", 4);
kv.get("foo", 4); // output "bar2"
kv.get("foo", 5); //output "bar2"

Example 2:
Input: inputs = ["TimeMap","set","set","get","get","get","get","get"], inputs = [[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
Output: [null,null,null,"","high","high","low","low"]

Note:
All key/value strings are lowercase.
All key/value strings have length in the range [1, 100]
The timestamps for all TimeMap.set operations are strictly increasing.
1 <= timestamp <= 10^7
TimeMap.set and TimeMap.get functions will be called a total of 120000 times (combined) per test case.

*/

func main() {
	fmt.Println("initializing test data...")

	timeMap := NewTimeMap()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)
	fmt.Println("Answer is:", timeMap.Get("love", 5)) //""
	fmt.Println("Answer is:", timeMap.Get("love", 10)) //high
	fmt.Println("Answer is:", timeMap.Get("love", 15)) //high
	fmt.Println("Answer is:", timeMap.Get("love", 20)) //low
	fmt.Println("Answer is:", timeMap.Get("love", 25)) //low
}

type TimeMap struct {
	keyMap map[string]*[]ValueEntry
}

type ValueEntry struct {
	value string
	timestamp int
}


/** Initialize your data structure here. */
func NewTimeMap() TimeMap {
	return TimeMap{keyMap: make(map[string]*[]ValueEntry)}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
	existingValues := this.keyMap[key]
	if existingValues == nil {
		existingValues = &[]ValueEntry{}
		this.keyMap[key] = existingValues
	}
	//Since the timestamps for all TimeMap.set operations are strictly increasing. this will be a sorted array
	*existingValues = append(*existingValues, ValueEntry{value, timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
	values := this.keyMap[key]
	if values == nil {
		return ""
	}
	//binary search in iterative manner
	start, end := 0, len(*values)-1
	result := ValueEntry{"", 0}
	for start <= end {
		mid := start + (end-start)/2
		if (*values)[mid].timestamp == timestamp {
			return (*values)[mid].value
		} else if (*values)[mid].timestamp > timestamp {
			end = mid - 1
		} else if (*values)[mid].timestamp < timestamp {
			if (*values)[mid].timestamp > result.timestamp { //maybe we don't need this check?
				result = (*values)[mid]
			}
			start = mid + 1
		}
	}
	return result.value
}
