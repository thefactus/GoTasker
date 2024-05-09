package database

import (
	"todolist/schemas"
	"todolist/types"
)

func (s *service) GetLists() ([]schemas.List, error) {
	var lists []schemas.List
	if err := s.db.Find(&lists).Error; err != nil {
		return nil, err
	}
	return lists, nil
}

func (s *service) CreateList(payload types.CreateListPayload) (*schemas.List, error) {
	list := schemas.List{
		Title: payload.Title,
	}

	// Create the list in the database
	if err := s.db.Create(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (s *service) UpdateList(listID string, payload types.UpdateListPayload) (*schemas.List, error) {
	var list schemas.List
	if err := s.db.First(&list, listID).Error; err != nil {
		return nil, err
	}

	list.Title = payload.Title

	if err := s.db.Save(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}

func (s *service) DeleteList(listID string) error {
	var list schemas.List
	if err := s.db.First(&list, listID).Error; err != nil {
		return err
	}

	if err := s.db.Delete(&list).Error; err != nil {
		return err
	}

	return nil
}
