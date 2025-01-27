package usecases

import (
	"skrik/internal/entities"
	"skrik/internal/repository"
	"strconv"
	"gorm.io/gorm"
)

type ChatUsecaseInterface interface {
    CreateRoom(name string) (*entities.Room, error)
    GetRoomByID(roomID uint) (*entities.Room, error)
    SaveMessage(roomID string, userID uint, content string) error
    GetMessagesByRoomID(roomID uint) ([]entities.Message, error)
	CreateRoomIfNotExists(roomID string) (*entities.Room, error)
}
type ChatUsecase struct {
	repo repository.ChatRepositoryInterface
}
func NewChatUsecase(cr repository.ChatRepositoryInterface) *ChatUsecase {
	return &ChatUsecase{repo: cr}
}

func (cu *ChatUsecase) CreateRoom (name string) (*entities.Room, error) {
	room := &entities.Room{Name: name}
	_, err := cu.repo.CreateRoom(room)
	if err != nil {
		return nil, err
	}
	return room, nil
}
func (cu *ChatUsecase) GetRoomByID(roomID uint) (*entities.Room, error){
	
	return cu.repo.GetRoomByID(roomID)
}

func (cu *ChatUsecase) SaveMessage(roomID string, userID uint, content string) error{
	roomid, err := strconv.Atoi(roomID)
	if err != nil {
		return err
	}
	message := &entities.Message{Content: content, RoomID: uint(roomid), UserID: userID}
	return cu.repo.SaveMessage(message)
}
func (cu *ChatUsecase) GetMessagesByRoomID(roomID uint) ([]entities.Message, error){
	return cu.repo.GetMessagesByRoomID(roomID)
}

func (cu *ChatUsecase) CreateRoomIfNotExists(roomID string) (*entities.Room, error) {
	roomid, err := strconv.Atoi(roomID)
	if err != nil {
		return nil, err
	}
	room, _ := cu.repo.GetRoomByID(uint(roomid))
	if room != nil {
		return room, nil
	}
	newRoom := &entities.Room{Model: gorm.Model{ID: uint(roomid)}}
	return cu.repo.CreateRoom(newRoom)
	
}
