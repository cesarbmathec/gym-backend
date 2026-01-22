package models

import (
	"time"

	"github.com/gorilla/websocket"
)

type Chat struct {
	ID          uint      `gorm:"primaryKey;column:id_chat" json:"id_chat"`
	ChatName    string    `gorm:"not null;size:100;column:chat_name" json:"chat_name"`
	Description string    `gorm:"type:text;column:description" json:"description"`
	IsGroupChat bool      `gorm:"default:true;column:is_group_chat" json:"is_group_chat"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

type ChatParticipant struct {
	ID          uint   `gorm:"primaryKey;column:id_participant" json:"id_participant"`
	ChatID      uint   `gorm:"column:chat_id" json:"chat_id"`
	UserID      uint   `gorm:"column:user_id" json:"user_id"`
	IsChatAdmin bool   `gorm:"default:false;column:is_chat_admin" json:"is_chat_admin"`
	UserType    string `gorm:"size:20;column:user_type" json:"user_type"`
}

type ChatMessage struct {
	ID       uint      `gorm:"primaryKey;column:id_message" json:"id_message"`
	ChatID   uint      `gorm:"column:chat_id" json:"chat_id"`
	UserID   uint      `gorm:"column:user_id" json:"user_id"`
	UserType string    `gorm:"size:20;column:user_type" json:"user_type"`
	Message  string    `gorm:"type:text;column:message" json:"message"`
	SentAt   time.Time `gorm:"column:sent_at" json:"sent_at"`
	IsRead   bool      `gorm:"default:false;column:is_read" json:"is_read"`
}

// Agregar al final:
type ChatSession struct {
	ClientID uint
	Username string
	Conn     *websocket.Conn
	RoomID   string
}
