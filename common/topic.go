package common

import "web/pubsub"

const (
	TopicIncreaseLikeCountIdea    pubsub.Topic = "TopicIncreaseLikeCountIdea"
	TopicIncreaseDisLikeCountIdea pubsub.Topic = "TopicIncreaseDisLikeCountIdea"
	TopicIncreaseCommentCountIdea pubsub.Topic = "TopicIncreaseCommentCountIdea"
	TopicDecreaseCommentCountIdea pubsub.Topic = "TopicDecreaseCommentCountIdea"
	TopicIncreaseViewCountIdea    pubsub.Topic = "TopicIncreaseViewCountIdea"
	TopicDecreaseLikeCountIdea    pubsub.Topic = "TopicDecreaseLikeCountIdea"
	TopicDecreaseDisLikeCountIdea pubsub.Topic = "TopicDecreaseDisLikeCountIdea"

	TopicDecreaseReplyCountComment pubsub.Topic = "TopicDecreaseReplyCountComment"
	TopicIncreaseReplyCountComment pubsub.Topic = "TopicIncreaseReplyCountComment"
)
