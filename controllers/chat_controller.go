package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/cesarbmathec/gym-backend/auth"
	"github.com/cesarbmathec/gym-backend/config"
	"github.com/cesarbmathec/gym-backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // CORS OK
	}
	chatRooms = make(map[string]map[uint]*websocket.Conn) // roomID → clientID → conn
	roomsMu   sync.RWMutex
)

func ChatWebSocket(c *gin.Context) {
	token := c.Query("token")
	roomID := c.Param("room")

	// 1. VALIDAR JWT
	claims, err := auth.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// 2. VERIFICAR CLIENTE EXISTE
	var client models.Client
	if err := config.DB.First(&client, claims.UserID).Error; err != nil || !client.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "Client not active"})
		return
	}

	// 3. UPGRADE a WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close()

	// 4. REGISTRAR conexión
	roomsMu.Lock()
	if chatRooms[roomID] == nil {
		chatRooms[roomID] = make(map[uint]*websocket.Conn)
	}
	chatRooms[roomID][claims.UserID] = ws
	roomsMu.Unlock()

	log.Printf("✅ Client %s (%d) joined room %s", claims.Username, claims.UserID, roomID)

	// 5. AUTH CONFIRM MESSAGE
	authMsg := gin.H{
		"type":     "auth_success",
		"user_id":  claims.UserID,
		"username": claims.Username,
		"room":     roomID,
		"online":   len(chatRooms[roomID]),
	}
	ws.WriteJSON(authMsg)

	// 6. ESCUCHAR MENSAJES
	for {
		var msg struct {
			Type    string `json:"type"`
			Content string `json:"content"`
		}
		err := ws.ReadJSON(&msg)
		if err != nil {
			break // Client desconectado
		}

		if msg.Type == "message" {
			broadcastMessage(roomID, claims.UserID, claims.Username, msg.Content)
		}
	}

	// 7. LIMPIAR conexión
	roomsMu.Lock()
	delete(chatRooms[roomID], claims.UserID)
	if len(chatRooms[roomID]) == 0 {
		delete(chatRooms, roomID)
	}
	roomsMu.Unlock()
	log.Printf("❌ Client %d left room %s", claims.UserID, roomID)
}

func broadcastMessage(roomID string, senderID uint, senderName, content string) {
	roomsMu.RLock()
	defer roomsMu.RUnlock()

	message := gin.H{
		"type":      "message",
		"sender_id": senderID,
		"sender":    senderName,
		"content":   content,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	for clientID, conn := range chatRooms[roomID] {
		if err := conn.WriteJSON(message); err != nil {
			log.Printf("Error sending to client %d: %v", clientID, err)
			continue
		}
	}
}
