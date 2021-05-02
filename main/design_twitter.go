package main

import (
	"container/heap"
	. "dailychallenge/utils"
	"fmt"
	"time"
)

/**
https://leetcode.com/problems/design-twitter/

Design a simplified version of Twitter where users can post tweets, follow/unfollow another user,
and is able to see the 10 most recent tweets in the user's news feed.

Implement the Twitter class:

Twitter() Initializes your twitter object.
void postTweet(int userId, int tweetId) Composes a new tweet with ID tweetId by the user userId. Each call to this function will be made with a unique tweetId.
List<Integer> getNewsFeed(int userId) Retrieves the 10 most recent tweet IDs in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user themself. Tweets must be ordered from most recent to least recent.
void follow(int followerId, int followeeId) The user with ID followerId started following the user with ID followeeId.
void unfollow(int followerId, int followeeId) The user with ID followerId started unfollowing the user with ID followeeId.


Example 1:

Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]

Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // User 1 posts a new tweet (id = 5).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
twitter.follow(1, 2);    // User 1 follows user 2.
twitter.postTweet(2, 6); // User 2 posts a new tweet (id = 6).
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
twitter.unfollow(1, 2);  // User 1 unfollows user 2.
twitter.getNewsFeed(1);  // User 1's news feed should return a list with 1 tweet id -> [5], since user 1 is no longer following user 2.


Constraints:

1 <= userId, followerId, followeeId <= 500
0 <= tweetId <= 104
All the tweets have unique IDs.
At most 3 * 104 calls will be made to postTweet, getNewsFeed, follow, and unfollow.



 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);

 */
func main()  {
	fmt.Println("initializing test data...")
	obj := constructor()
	obj.PostTweet(1,5)
	fmt.Println("News feed for user 1 is ", obj.GetNewsFeed(1)) // [5]
	obj.Follow(1,2)
	obj.PostTweet(2,6)
	fmt.Println("News feed for user 1 is ", obj.GetNewsFeed(1)) //[6,5]
	obj.Unfollow(1,2)
	fmt.Println("News feed for user 1 is ", obj.GetNewsFeed(1)) // [5]
}


type Twitter struct {
	//storage for follower/followee relationship  key:userID, value: Set of users he/she follows
	FollowerRecords map[int]*Set
	//storage for tweets key:userId, value:a circular buffer of his/her tweets (latest 10)
	TweetRecords map[int]*CircularBuffer
}

type Tweet struct {
	tweetId int
	userId int
	timestamp int64
}

/** Initialize your data structure here. */
func constructor() Twitter {
	twitter := Twitter{}
	twitter.FollowerRecords = make(map[int]*Set)
	twitter.TweetRecords = make(map[int]*CircularBuffer)
	return twitter
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	tweets, existing := this.TweetRecords[userId]
	if !existing {
		tweets = NewCircularBuffer(10)
	}
	tweets.Add(Tweet{tweetId, userId, time.Now().UnixNano()})
	this.TweetRecords[userId] = tweets
}

/**
	Retrieve the 10 most recent tweet ids in the user's news feed.
	Each item in the news feed must be posted by users who the user followed or by the user herself.
	Tweets must be ordered from most recent to least recent.
*/
const LIMIT = 10

func (this *Twitter) GetNewsFeed(userId int) []int {
	minHeap := &TweetMinHeap{}
	heap.Init(minHeap)

	followeesSet := this.FollowerRecords[userId]
	all := []int{userId}
	if followeesSet != nil {
		all = append(all, followeesSet.GetAll()...)
	}

	for _, individualId := range all {
		tweets := this.TweetRecords[individualId]
		if tweets != nil {
			for _, tweet := range tweets.GetAll() {
				if minHeap.Len() < LIMIT {
					heap.Push(minHeap, tweet)
				} else if (*minHeap)[0].timestamp < tweet.timestamp {
					heap.Pop(minHeap)
					heap.Push(minHeap, tweet)
				} //otherwise ignore this tweet
			}
		}
	}

	result := make([]int, minHeap.Len())
	for i := len(result)-1; i >= 0; i-- {
		tweet := heap.Pop(minHeap).(Tweet)
		result[i] = tweet.tweetId
	}
	return result
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	followingIds, existing := this.FollowerRecords[followerId]
	if !existing {
		followingIds = NewSet()
		this.FollowerRecords[followerId] = followingIds
	}
	followingIds.Add(followeeId)
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	followingIds, existing := this.FollowerRecords[followerId]
	if !existing {
		fmt.Println("followee for ", followerId, " doesn't exist!")
	}
	followingIds.Remove(followeeId)
}

//--------------- circular buffer to store 10 most recent tweets for each user -------------
type CircularBuffer struct {
	Tweets []Tweet
	Limit int
	currentPos int
}

func NewCircularBuffer(Limit int) *CircularBuffer {
	return &CircularBuffer{[]Tweet{}, Limit,0}
}

func (c *CircularBuffer) Add(item Tweet) {
	if len(c.Tweets) < c.Limit {
		c.Tweets = append(c.Tweets, item)
	} else {
		c.Tweets[c.currentPos] = item
		c.currentPos = (c.currentPos + 1) % c.Limit
	}
}

func (c *CircularBuffer) GetAll() []Tweet {
	return c.Tweets //TODO: maybe we should return them in time order based on 'currentPos'
}


//--------------- min-heap to get top 10 most recent tweets ---------------
type TweetMinHeap []Tweet

func (h TweetMinHeap) Len() int           { return len(h) }
func (h TweetMinHeap) Less(i, j int) bool { return h[i].timestamp < h[j].timestamp }
func (h TweetMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


//Note that Push and Pop in this interface are for package "container/heap" implementation to call.
//To add and remove things from the heap, use heap.Push and heap.Pop.
//See TestHeap() below
func (h *TweetMinHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Tweet))
}

func (h *TweetMinHeap) Pop() interface{}{
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}