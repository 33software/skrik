package repository

import (
	"skrik/internal/entities"

	"gorm.io/gorm"
)
type ChatRepositoryInterface interface {
    CreateRoom(room *entities.Room) (*entities.Room, error)
    GetRoomByID(roomID uint) (*entities.Room, error)
    SaveMessage(message *entities.Message) error
    GetMessagesByRoomID(roomID uint) ([]entities.Message, error)
}



type ChatRepository struct{
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository{
	return &ChatRepository{db: db}
}

func (cr *ChatRepository) CreateRoom (room *entities.Room) (*entities.Room, error){
	if err := cr.db.Create(room).Error; err != nil {
		return nil, err
	}
	return room, nil

}
func (cr *ChatRepository) GetRoomByID (roomID uint) (*entities.Room, error) {
	var room entities.Room
	err := cr.db.First(&room, roomID).Error 
	if err != nil {
		return nil, err
	}
	return &room, nil
}
func (cr *ChatRepository) SaveMessage(message *entities.Message) error{
	return cr.db.Create(message).Error
}
func (cr *ChatRepository) GetMessagesByRoomID(roomID uint) ([]entities.Message, error) {
	var messages []entities.Message
	err := cr.db.Where("RoomID=?", roomID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}