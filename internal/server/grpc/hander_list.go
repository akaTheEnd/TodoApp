package grpc

import (
	"context"
	"todoApp/model"
)

func (s *Server) CreateList(_ context.Context, req *CreateListRequest) (*CreateListResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &CreateListResponse{
			Id:           -1,
			ErrorMessage: err.Error(),
		}, err
	}
	listId, err := s.services.TodoList.Create(userId, model.TodoList{
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return &CreateListResponse{
			Id:           -1,
			ErrorMessage: err.Error(),
		}, err
	}

	return &CreateListResponse{
		Id:           int32(listId),
		ErrorMessage: "",
	}, err
}

func (s *Server) GetAllLists(_ context.Context, req *GetAllListsRequest) (*GetAllListsResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &GetAllListsResponse{
			Lists:        []*List{},
			ErrorMessage: err.Error(),
		}, err
	}
	lists, err := s.services.TodoList.GetAll(userId)
	if err != nil {
		return &GetAllListsResponse{
			Lists:        []*List{},
			ErrorMessage: err.Error(),
		}, err
	}
	var responseLists []*List
	for _, list := range lists {
		responseList := List{
			Id:          int32(list.Id),
			Title:       list.Title,
			Description: list.Description,
		}
		responseLists = append(responseLists, &responseList)
	}

	return &GetAllListsResponse{
		Lists:        responseLists,
		ErrorMessage: "",
	}, err
}

func (s *Server) GetListById(_ context.Context, req *GetListByIdRequest) (*GetListByIdResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &GetListByIdResponse{
			List:         &List{},
			ErrorMessage: err.Error(),
		}, err
	}
	list, err := s.services.TodoList.GetById(userId, int(req.ListId))
	if err != nil {
		return &GetListByIdResponse{
			List:         &List{},
			ErrorMessage: err.Error(),
		}, err
	}
	responseList := &List{
		Id:          int32(list.Id),
		Title:       list.Title,
		Description: list.Description,
	}

	return &GetListByIdResponse{
		List:         responseList,
		ErrorMessage: "",
	}, err
}

func (s *Server) UpdateList(_ context.Context, req *UpdateListRequest) (*UpdateListResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &UpdateListResponse{
			ErrorMessage: err.Error(),
		}, err
	}
	err = s.services.TodoList.Update(userId, int(req.ListId), model.UpdateListInput{
		Title:       &req.Title,
		Description: &req.Description,
	})
	if err != nil {
		return &UpdateListResponse{
			ErrorMessage: err.Error(),
		}, err
	}

	return &UpdateListResponse{
		ErrorMessage: "",
	}, err
}

func (s *Server) DeleteList(_ context.Context, req *DeleteListRequest) (*DeleteListResponse, error) {
	userId, err := s.services.Authorization.ParseToken(req.Token)
	if err != nil {
		return &DeleteListResponse{
			ErrorMessage: err.Error(),
		}, err
	}
	err = s.services.TodoList.Delete(userId, int(req.ListId))
	if err != nil {
		return &DeleteListResponse{
			ErrorMessage: err.Error(),
		}, err
	}

	return &DeleteListResponse{
		ErrorMessage: "",
	}, err
}
