package updates

type User struct {
	Id            int32  `json:"id"`
	Is_bot        bool   `json:"is_bool"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Username      string `json:"username"`
	Language_code string `json:"language_code"`
}
type Chat struct {
	Ld       int32  `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
}
type Message struct {
	Msg_id      int32  `json:"message_id"`
	From        User   `json:"from"`
	Sender_chat Chat   `json:"sender_chat"`
	Date        int32  `json:"date"`
	Chat        Chat   `json:"chat"`
	Text        string `json:"text"`
	Description string `json:"description"`
	Invite_link string `json:"invite_link"`
}
type InlineQuery struct {
	Id        int32  `json:"id"`
	From      User   `json:"from"`
	Query     string `json:"query"`
	Offset    string `json:"offset"`
	Chat_type string `json:"chat_type"`
}
type chatJoinRequest struct {
	Chat Chat   `json:"chat"`
	From User   `json:"from"`
	Date int32  `json:"date"`
	Bio  string `json:"bio"`
}
type ChatMember struct {
	Status       string `json:"status"`
	User         User   `json:"user"`
	Is_anonymous bool   `json:"is_anonymous"`
	Custom_title string `json:"custom_title"`
}
type ChatMemberUpdated struct {
	Chat            Chat       `json:"chat"`
	From            User       `json:"user"`
	Date            string     `json:"date"`
	New_chat_member ChatMember `json:"new_chat_member"`
}
type Result struct {
	Message           Message           `json:"message"`
	Iinline_query     InlineQuery       `json:"inline_query"`
	Chat_member       ChatMemberUpdated `json:"chat_member"`
	Chat_join_request chatJoinRequest   `json:"chat_join_request"`
}
type Update struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

type GetUpdates struct {
	Offset          int32    `json:"offset"`
	Limit           int32    `json:"limit"`
	Timeout         int32    `json:"timeout"`
	Allowed_updates []string `json:"allowed_updates"`
}
