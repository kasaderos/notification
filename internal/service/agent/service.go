package agent

type Service struct {
	repo                Repository
	subscriptionService SubscriptionService
	notificationService NotificationService
	userEventQueue      UserEventQueue
}

func NewService(
	repo Repository,
	subscriptionService SubscriptionService,
	notificationService NotificationService,
	userEventQueue UserEventQueue,
) *Service {
	return &Service{
		repo:                repo,
		subscriptionService: subscriptionService,
		notificationService: notificationService,
		userEventQueue:      userEventQueue,
	}
}
