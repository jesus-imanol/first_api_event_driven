package entities

type Match struct {
	Id int32 `json:"id"`
	SenderUser int32 `json:"sender_user"`
	ReceiverUser int32 `json:"receiver_user"`
	Status  string `json:"status"`
}
func NewMatch(senderUser int32,receiverUser int32) *Match {
	return &Match{
		SenderUser: senderUser,
		ReceiverUser: receiverUser,
		Status: "pending",
	}
}