package types

type Id int

type Update struct {
	Id                int      `json:"update_id"`
	Message           *Message `json:"message"`
	EditedMessage     *Message `json:"edited_message"`
	ChannelPost       *Message `json:"channel_post"`
	EditedChannelPost *Message `json:"edited_channel_post"`
}

type UpdateResult struct {
	Ok      bool
	Updates []*Update `json:"result"`
}

type User struct {
	Id                      Id     `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	Username                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineRequests  bool   `json:"supports_inline_queries"`
}

type Chat struct {
	Id                                 Id               `json:"id"`
	Type                               string           `json:"type"`
	Title                              string           `json:"title"`
	FirstName                          string           `json:"first_name"`
	LastName                           string           `json:"last_name"`
	Username                           string           `json:"username"`
	IsForum                            bool             `json:"is_forum"`
	Photo                              *ChatPhoto       `json:"photo"`
	ActiveUsernames                    []string         `json:"active_usernames"`
	EmojiStatusCustomEmojiId           string           `json:"emoji_status_custom_emoji_id"`
	Bio                                string           `json:"bio"`
	HasPrivateForwards                 bool             `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool             `json:"has_restricted_voice_and_video_messages"`
	JoinToSendMessages                 bool             `json:"join_to_send_messages"`
	JoinByRequest                      bool             `json:"join_by_request"`
	Description                        string           `json:"description"`
	InviteLink                         string           `json:"invite_link"`
	PinnedMessage                      *Message         `json:"pinned_message"`
	Permissions                        *ChatPermissions `json:"permissions"`
	SlowModeDelay                      int              `json:"slow_mode_delay"`
	MessageAutoDeleteTime              int              `json:"message_auto_delete_time"`
	HasProtectedContent                bool             `json:"has_protected_content"`
	StickerSetName                     string           `json:"sticker_set_name"`
	CanSetStickerSet                   bool             `json:"can_set_sticker_set"`
	LinkedChatId                       Id               `json:"linked_chat_id"`
	Location                           *ChatLocation    `json:"location"`
}

type Message struct {
	Id                            Id                             `json:"message_id"`
	ThreadId                      int                            `json:"message_thread_id"`
	From                          *User                          `json:"from"`
	SenderChat                    *Chat                          `json:"sender_chat"`
	Date                          int                            `json:"date"`
	Chat                          *Chat                          `json:"chat"`
	ForwardFrom                   *User                          `json:"forward_from"`
	ForwardFromChat               *Chat                          `json:"forward_from_chat"`
	ForwardFromMessageId          Id                             `json:"forward_from_message_id"`
	ForwardSignature              string                         `json:"forward_signature"`
	ForwardSenderName             string                         `json:"forward_sender_name"`
	ForwardDate                   int                            `json:"forward_date"`
	IsTopicMessage                bool                           `json:"is_topic_message"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward"`
	ReplyToMessage                *Message                       `json:"reply_to_message"`
	ViaBot                        *User                          `json:"via_bot"`
	EditDate                      int                            `json:"edit_date"`
	HasProtectedContent           bool                           `json:"has_protected_content"`
	MediaGroupId                  string                         `json:"media_group_id"`
	AuthorSignature               string                         `json:"author_signature"`
	Text                          string                         `json:"text"`
	Entities                      []*MessageEntity               `json:"entities"`
	Animation                     *Animation                     `json:"animation"`
	Audio                         *Audio                         `json:"audio"`
	Document                      *Document                      `json:"document"`
	Photo                         []*PhotoSize                   `json:"photo"`
	Sticker                       *Sticker                       `json:"sticker"`
	Video                         *Video                         `json:"video"`
	VideoNote                     *VideoNote                     `json:"video_note"`
	Voice                         *Voice                         `json:"voice"`
	Caption                       string                         `json:"caption"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities"`
	Contact                       *Contact                       `json:"contact"`
	Dice                          *Dice                          `json:"dice"`
	Game                          *Game                          `json:"game"`
	Poll                          *Poll                          `json:"poll"`
	Venue                         *Venue                         `json:"venue"`
	Location                      *Location                      `json:"location"`
	NewChatMembers                []*User                        `json:"new_chat_members"`
	LeftChatMember                *User                          `json:"left_chat_member"`
	NewChatTitle                  string                         `json:"new_chat_title"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo"`
	GroupChatCreated              bool                           `json:"group_chat_created"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created"`
	ChannelChatCreated            bool                           `json:"channel_chat_created"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatId               Id                             `json:"migrate_to_chat_id"`
	MigrateGromChatId             Id                             `json:"migrate_from_chat_id"`
	PinnedMessage                 *Message                       `json:"pinned_message"`
	Invoice                       *Invoice                       `json:"invoice"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment"`
	ConnectedWebsite              string                         `json:"connected_website"`
	PassportData                  *PassportData                  `json:"passport_data"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created"`
	ForumTopicClosed              bool                           `json:"forum_topic_closed"`
	ForumTopicReopened            bool                           `json:"forum_topic_reopened"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              bool                           `json:"video_chat_started"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited"`
	WebAppData                    *WebAppData                    `json:"web_app_data"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup"`
}

type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionId        string     `json:"shipping_option_id"`
	OrderInfo               *OrderInfo `json:"order_info"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
}

type OrderInfo struct {
	Name            string           `json:"name"`
	PhoneNumber     string           `json:"phone_number"`
	Email           string           `json:"email"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials        `json:"credentials"`
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data"`
	PhoneNumber string          `json:"phone_number"`
	Email       string          `json:"email"`
	Files       []*PassportFile `json:"files"`
	FrontSide   []*PassportFile `json:"front_side"`
	ReverseSide []*PassportFile `json:"reverse_side"`
	Selfie      []*PassportFile `json:"selfie"`
	Translation []*PassportFile `json:"translation"`
	Hash        string          `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type PassportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

type Sticker struct {
	FileId           string        `json:"file_id"`
	FileUniqueId     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumb            *PhotoSize    `json:"thumb"`
	Emoji            string        `json:"emoji"`
	SetName          string        `json:"set_name"`
	PremiumAnimation *File         `json:"premium_animation"`
	MaskPosition     *MaskPosition `json:"mask_position"`
	CustomEmojiId    string        `json:"custom_emoji_id"`
	FileSize         int           `json:"file_size"`
}

type StickerSet struct {
	Name        string     `json:"name"`
	Title       string     `json:"title"`
	StickerType string     `json:"sticker_type"`
	IsAnimated  bool       `json:"is_animated"`
	IsVideo     bool       `json:"is_video"`
	Thumb       *PhotoSize `json:"thumb"`
	Stickers    []*Sticker `json:"stickers"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

type MessageEntity struct {
	Type          string `json:"type"`
	Offset        int    `json:"offset"`
	Length        int    `json:"length"`
	Url           string `json:"url"`
	User          *User  `json:"user"`
	Language      string `json:"language"`
	CustomEmojiId string `json:"custom_emoji_id"`
}

type PhotoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

type Animation struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type Audio struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer"`
	Title        string     `json:"title"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

type Document struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
	Thumb        *PhotoSize `json:"thumb"`
}

type Video struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileName     string     `json:"file_name"`
	MimeType     string     `json:"mime_type"`
	FileSize     int        `json:"file_size"`
}

type VideoNote struct {
	FileId       string     `json:"file_id"`
	FileUniqueId string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb"`
	FileSize     int        `json:"file_size"`
}

type Voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int    `json:"file_size"`
}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      Id     `json:"user_id"`
	VCard       string `json:"vcard"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voter_count"`
}

type PollAnswer struct {
	PollId       string `json:"poll_id"`
	User         *User  `json:"user"`
	OptionIdList []int  `json:"option_ids"`
}

type Poll struct {
	Id                    string           `json:"id"`
	Question              string           `json:"question"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionId       int              `json:"correct_option_id"`
	Explanation           string           `json:"explanation"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities"`
	OpenPeriod            int              `json:"open_period"`
	CloseDate             int              `json:"close_date"`
}

type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int     `json:"live_period"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximity_alert_radius"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Adress          string    `json:"address"`
	FoursquareId    string    `json:"foursquare_id"`
	FoursquareType  string    `json:"foursquare_type"`
	GooglePlaceId   string    `json:"google_place_id"`
	GooglePlaceType string    `json:"google_place_type"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}

type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"`
}

type UserProfilePhotos struct {
	TotalCount int          `json:"total_count"`
	Photos     []*PhotoSize `json:"photos"`
}

type File struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FilePath     string `json:"file_path"`
}

type WebAppInfo struct {
	Url string `json:"url"`
}

type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool                `json:"resize_keyboard"`
	OneTimeKeyboard       bool                `json:"one_time_keyboard"`
	InputFieldPlaceholder string              `json:"input_field_placeholder"`
	Selective             bool                `json:"selective"`
}

type KeyboardButton struct {
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact"`
	RequestLocation bool                    `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`
	WebApp          *WebAppInfo             `json:"web_app"`
}

type KeyboardButtonPollType struct {
	Type string `json:"type"`
}

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string      `json:"text"`
	Url                          string      `json:"url"`
	CallbackData                 string      `json:"callback_data"`
	WebApp                       *WebAppInfo `json:"web_app"`
	LoginUrl                     *LoginUrl   `json:"login_url"`
	SwitchInlineQuery            string      `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string      `json:"switch_inline_query_current_chat"`
	CallbackGame                 bool        `json:"callback_game"`
	Pay                          bool        `json:"pay"`
}

type LoginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

type CallbackQuery struct {
	Id              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`
	InlineMessageId string   `json:"inline_message_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder"`
	Selective             bool   `json:"selective"`
}

type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BigFileUniqueId   string `json:"big_file_unique_id"`
}

type ChatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 *User  `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name"`
	ExpireDate              int    `json:"expire_date"`
	MemberLimit             int    `json:"member_limit"`
	PendingJsonRequestCount int    `json:"pending_join_request_count"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages"`
	CanEditMessages     bool `json:"can_edit_messages"`
	CanPinMessages      bool `json:"can_pin_messages"`
	CanManageTopics     bool `json:"can_manage_topics"`
}

type ChatMember struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

type ChatMemberOwner struct {
	Status      string `json:"status"`
	User        *User  `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}

type ChatMemberAdministrator struct {
	Status              string `json:"status"`
	User                *User  `json:"user"`
	CanBeEdited         bool   `json:"can_be_edited"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages"`
	CanEditMessages     bool   `json:"can_edit_messages"`
	CanPinMessages      bool   `json:"can_pin_messages"`
	CanManageTopics     bool   `json:"can_manage_topics"`
	CustomTitle         string `json:"custom_title"`
}

type ChatMemberMember struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

type ChatMemberRestricted struct {
	Status                string `json:"status"`
	User                  *User  `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendMediaMessages  bool   `json:"can_send_media_messages"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate             int    `json:"until_date"`
}

type ChatMemberLeft struct {
	Status string `json:"status"`
	User   *User  `json:"user"`
}

type ChatMemberBanned struct {
	Status    string `json:"status"`
	User      *User  `json:"user"`
	UntilDate int    `json:"until_date"`
}

type ChatMemberUpdated struct {
	Chat          *Chat           `json:"chat"`
	From          *User           `json:"from"`
	Date          int             `json:"date"`
	OldChatMember *ChatMember     `json:"old_chat_member"`
	NewChatMember *ChatMember     `json:"new_chat_member"`
	InviteLink    *ChatInviteLink `json:"invite_link"`
}

type ChatJoinRequest struct {
	Chat       *Chat           `json:"chat"`
	From       *User           `json:"from"`
	Date       int             `json:"date"`
	Bio        string          `json:"bio"`
	InviteLink *ChatInviteLink `json:"invite_link"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendMediaMessages  bool `json:"can_send_media_messages"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

type ForumTopic struct {
	MessageThreadId   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope struct {
	Type   string `json:"type"`
	ChatId string `json:"chat_id"`
	UserId int    `json:"user_id"`
}

type MenuButtonCommands struct {
	Type string `json:"type"`
}

type MenuButtonWebApp struct {
	Type   string      `json:"type"`
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}

type MenuButtonDefault struct {
	Type string `json:"type"`
}

type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"`
	ReplyAfter      int `json:"retry_after"`
}

type InputMediaPhoto struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
}

type InputMediaVideo struct {
	Type              string           `json:"type"`
	Media             string           `json:"media"`
	Thumb             string           `json:"thumb"`
	Caption           string           `json:"caption"`
	ParseMode         string           `json:"parse_mode"`
	CaptionEntities   []*MessageEntity `json:"caption_entities"`
	Width             int              `json:"width"`
	Height            int              `json:"height"`
	Duration          int              `json:"duration"`
	SupportsStreaming bool             `json:"supports_streaming"`
}

type InputMediaAnimation struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Width           int              `json:"width"`
	Height          int              `json:"height"`
	Duration        int              `json:"duration"`
}

type InputMediaAudio struct {
	Type            string           `json:"type"`
	Media           string           `json:"media"`
	Thumb           string           `json:"thumb"`
	Caption         string           `json:"caption"`
	ParseMode       string           `json:"parse_mode"`
	CaptionEntities []*MessageEntity `json:"caption_entities"`
	Duration        int              `json:"duration"`
	Performer       string           `json:"performer"`
	Title           string           `json:"title"`
}

type InputMediaDocument struct {
	Type                        string           `json:"type"`
	Media                       string           `json:"media"`
	Thumb                       string           `json:"thumb"`
	Caption                     string           `json:"caption"`
	ParseMode                   string           `json:"parse_mode"`
	CaptionEntities             []*MessageEntity `json:"caption_entities"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"`
}

type InputFile []byte

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities"`
	Animation    *Animation       `json:"animation"`
}
