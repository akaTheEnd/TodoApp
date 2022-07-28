package grpc

import (
	"context"
	"todoApp/model"
)

func (s *Server) CreateItem(_ context.Context, req *CreateItemRequest) (*CreateItemResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &CreateItemResponse{
			Id:           -1,
			ErrorMessage: err.Error(),
		}, err
	}
	listId, err := s.services.TodoItem.Create(userId, int(req.ListId), model.TodoItem{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return &CreateItemResponse{
			Id:           -1,
			ErrorMessage: err.Error(),
		}, err
	}

	return &CreateItemResponse{
		Id:           int32(listId),
		ErrorMessage: "",
	}, err
}

func (s *Server) GetAllItems(_ context.Context, req *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &GetAllItemsResponse{
			Items:        []*Item{},
			ErrorMessage: err.Error(),
		}, err
	}
	items, err := s.services.TodoItem.GetAll(userId, int(req.ListId))
	if err != nil {
		return &GetAllItemsResponse{
			Items:        []*Item{},
			ErrorMessage: err.Error(),
		}, err
	}
	var responseItems []*Item
	for _, item := range items {
		responseItem := Item{
			Id:          int32(item.Id),
			Title:       item.Title,
			Description: item.Description,
			Done:        item.Done,
		}
		responseItems = append(responseItems, &responseItem)
	}

	return &GetAllItemsResponse{
		Items:        responseItems,
		ErrorMessage: "",
	}, err
}

func (s *Server) GetItemById(_ context.Context, req *GetItemByIdRequest) (*GetItemByIdResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &GetItemByIdResponse{
			Item:         &Item{},
			ErrorMessage: err.Error(),
		}, err
	}
	item, err := s.services.TodoItem.GetById(userId, int(req.ItemId))
	if err != nil {
		return &GetItemByIdResponse{
			Item:         &Item{},
			ErrorMessage: err.Error(),
		}, err
	}
	responseItem := &Item{
		Id:          int32(item.Id),
		Title:       item.Title,
		Description: item.Description,
		Done:        item.Done,
	}

	return &GetItemByIdResponse{
		Item:         responseItem,
		ErrorMessage: "",
	}, err
}

func (s *Server) UpdateItem(_ context.Context, req *UpdateItemRequest) (*UpdateItemResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &UpdateItemResponse{
			ErrorMessage: err.Error(),
		}, err
	}
	err = s.services.TodoItem.Update(userId, int(req.ItemId), model.UpdateItemInput{
		Title:       &req.Title,
		Description: &req.Description,
		Done:        &req.Done,
	})
	if err != nil {
		return &UpdateItemResponse{
			ErrorMessage: err.Error(),
		}, err
	}

	return &UpdateItemResponse{
		ErrorMessage: "",
	}, err
}

func (s *Server) DeleteItem(_ context.Context, req *DeleteItemRequest) (*DeleteItemResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &DeleteItemResponse{
			ErrorMessage: err.Error(),
		}, err
	}
	err = s.services.TodoItem.Delete(userId, int(req.ItemId))
	if err != nil {
		return &DeleteItemResponse{
			ErrorMessage: err.Error(),
		}, err
	}

	return &DeleteItemResponse{
		ErrorMessage: "",
	}, err
}
