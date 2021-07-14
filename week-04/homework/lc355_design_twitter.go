package homework

type Twitter struct {
	FollowersMap map[int]map[int]bool
	UsersTweetsMap map[int]*Tweet
}

type Tweet struct {
	TweetId int
	Timestamp int
	Next *Tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter {
		FollowersMap: make(map[int]map[int]bool, 0),
		UsersTweetsMap: make(map[int]*Tweet, 0),
	}
}

var GlobalTimestamp int
/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	GlobalTimestamp++
	next := this.UsersTweetsMap[userId]
	this.UsersTweetsMap[userId] = &Tweet{
		TweetId: tweetId,
		Timestamp: GlobalTimestamp,
		Next: next,
	}
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	userIds := []int{userId}
	for id, _ := range this.FollowersMap[userId] {
		userIds = append(userIds, id)
	}

	var tweetLists []*Tweet
	for _, id := range userIds {
		// 深拷贝
		head := this.UsersTweetsMap[id]
		if head == nil {
			continue
		}
		tweet := &Tweet {}
		newHead := tweet
		for head != nil {
			newHead.TweetId = head.TweetId
			newHead.Timestamp = head.Timestamp
			head = head.Next
			if head != nil {
				newHead.Next = &Tweet{}
			}
			newHead = newHead.Next
		}
		tweetLists = append(tweetLists, tweet)
	}


	head := mergeKLists(tweetLists)

	var tweetIds []int
	count := 10
	for count > 0 && head != nil {
		tweetIds = append(tweetIds, head.TweetId)
		head = head.Next
		count--
	}

	return tweetIds
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	_, ok := this.FollowersMap[followerId]
	if !ok {
		this.FollowersMap[followerId] = make(map[int]bool, 0)
	}
	this.FollowersMap[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	delete(this.FollowersMap[followerId], followeeId)
}

// lc23
func mergeKLists(tweetLists []*Tweet) *Tweet {
	if len(tweetLists) == 0 {
		return nil
	}

	return mergeTwoLists(tweetLists[0], mergeKLists(tweetLists[1:]))
}

// lc121
func mergeTwoLists(l1 *Tweet, l2 *Tweet) *Tweet {
	var head *Tweet

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Timestamp < l2.Timestamp {
		head = l2
		l2Next := l2.Next
		l2.Next = mergeTwoLists(l1, l2Next)
	} else {
		head = l1
		l1Next := l1.Next
		l1.Next = mergeTwoLists(l2, l1Next)
	}

	return head
}