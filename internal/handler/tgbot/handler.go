package handler

// import (
// 	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
// 	repository "github.com/kasaderos/notification/internal/repository/user"
// )

// // TelegramBotHandler handles Telegram bot interactions
// type TelegramBotHandler struct {
// 	bot                     *tgbotapi.BotAPI
// 	notificationRuleService *service.NotificationRuleService
// 	userRepo                repository.UserRepository
// }

// // NewTelegramBotHandler creates a new Telegram bot handler
// func NewTelegramBotHandler(
// 	bot *tgbotapi.BotAPI,
// 	notificationRuleService *service.NotificationRuleService,
// 	userRepo repository.UserRepository,
// ) *TelegramBotHandler {
// 	return &TelegramBotHandler{
// 		bot:                     bot,
// 		notificationRuleService: notificationRuleService,
// 		userRepo:                userRepo,
// 	}
// }

// // HandleUpdate processes incoming Telegram updates
// func (h *TelegramBotHandler) HandleUpdate(ctx context.Context, update tgbotapi.Update) {
// 	if update.Message == nil {
// 		return
// 	}

// 	message := update.Message
// 	userID := message.From.ID

// 	// Ensure user exists in database
// 	user, err := h.userRepo.GetByTelegramID(ctx, int64(userID))
// 	if err != nil {
// 		// Create new user if not found
// 		user = &model.User{
// 			ID:         uuid.New().String(),
// 			Name:       fmt.Sprintf("%s %s", message.From.FirstName, message.From.LastName),
// 			TelegramID: int64(userID),
// 		}
// 		if err := h.userRepo.Create(ctx, user); err != nil {
// 			log.Printf("Failed to create user: %v", err)
// 			return
// 		}
// 	}

// 	// Handle commands
// 	if message.IsCommand() {
// 		h.handleCommand(ctx, message, user)
// 		return
// 	}

// 	// Handle regular messages (for rule creation/updates)
// 	h.handleMessage(ctx, message, user)
// }

// // handleCommand processes bot commands
// func (h *TelegramBotHandler) handleCommand(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	switch message.Command() {
// 	case "start":
// 		h.sendWelcomeMessage(message.Chat.ID)
// 	case "create_rule":
// 		h.handleCreateRule(ctx, message, user)
// 	case "update_rule":
// 		h.handleUpdateRule(ctx, message, user)
// 	case "delete_rule":
// 		h.handleDeleteRule(ctx, message, user)
// 	case "list_rules":
// 		h.handleListRules(ctx, message, user)
// 	default:
// 		h.sendMessage(message.Chat.ID, "Unknown command. Use /start to see available commands.")
// 	}
// }

// // handleCreateRule handles rule creation
// func (h *TelegramBotHandler) handleCreateRule(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	// Parse rule parameters from message text
// 	// Format: /create_rule domains:domain1,domain2 keywords:keyword1,keyword2 prompt:some prompt
// 	text := strings.TrimPrefix(message.Text, "/create_rule ")

// 	rule := &model.NotificationRule{
// 		Rule: model.Rule{
// 			Keywords: []string{},
// 			Sources: struct {
// 				Domains []string `json:"domains" db:"domains"`
// 			}{Domains: []string{}},
// 			Prompt: "",
// 			Schedule: struct {
// 				Time string `json:"time" db:"time"`
// 				ASAP bool   `json:"asap" db:"asap"`
// 			}{Time: "10:00", ASAP: false},
// 		},
// 	}

// 	// Parse parameters (simplified parsing)
// 	parts := strings.Split(text, " ")
// 	for _, part := range parts {
// 		if strings.HasPrefix(part, "domains:") {
// 			domains := strings.TrimPrefix(part, "domains:")
// 			rule.Rule.Sources.Domains = strings.Split(domains, ",")
// 		} else if strings.HasPrefix(part, "keywords:") {
// 			keywords := strings.TrimPrefix(part, "keywords:")
// 			rule.Rule.Keywords = strings.Split(keywords, ",")
// 		} else if strings.HasPrefix(part, "prompt:") {
// 			rule.Rule.Prompt = strings.TrimPrefix(part, "prompt:")
// 		}
// 	}

// 	// Create the rule
// 	if err := h.notificationRuleService.CreateRule(ctx, user.ID.String(), rule); err != nil {
// 		h.sendMessage(message.Chat.ID, fmt.Sprintf("Failed to create rule: %v", err))
// 		return
// 	}

// 	h.sendMessage(message.Chat.ID, "Rule created successfully!")
// }

// // handleUpdateRule handles rule updates
// func (h *TelegramBotHandler) handleUpdateRule(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	// TODO: Implement rule update logic
// 	h.sendMessage(message.Chat.ID, "Rule update functionality coming soon!")
// }

// // handleDeleteRule handles rule deletion
// func (h *TelegramBotHandler) handleDeleteRule(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	// TODO: Implement rule deletion logic
// 	h.sendMessage(message.Chat.ID, "Rule deletion functionality coming soon!")
// }

// // handleListRules handles listing user's rules
// func (h *TelegramBotHandler) handleListRules(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	// TODO: Implement rule listing logic
// 	h.sendMessage(message.Chat.ID, "Rule listing functionality coming soon!")
// }

// // handleMessage processes regular messages
// func (h *TelegramBotHandler) handleMessage(ctx context.Context, message *tgbotapi.Message, user *model.User) {
// 	// TODO: Implement message processing for rule creation/updates
// 	h.sendMessage(message.Chat.ID, "Please use commands to interact with the bot. Use /start to see available commands.")
// }

// // sendWelcomeMessage sends welcome message with available commands
// func (h *TelegramBotHandler) sendWelcomeMessage(chatID int64) {
// 	welcomeText := `Welcome to Events Notification Bot! 📰

// Available commands:
// /create_rule - Create a new notification rule
// /update_rule - Update an existing notification rule
// /delete_rule - Delete a notification rule
// /list_rules - List your notification rules

// Example usage:
// /create_rule domains:tengrievents.kz keywords:Iran,Kazakhstan prompt:Events about Iran and Kazakhstan`

// 	h.sendMessage(chatID, welcomeText)
// }

// // sendMessage sends a message to the specified chat
// func (h *TelegramBotHandler) sendMessage(chatID int64, text string) {
// 	msg := tgbotapi.NewMessage(chatID, text)
// 	if _, err := h.bot.Send(msg); err != nil {
// 		log.Printf("Failed to send message: %v", err)
// 	}
// }
