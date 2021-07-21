package handlers

import (
	"github.com/go-openapi/strfmt"
	"mnimidamonbackend/adapter/restapi/modelapi"
	"mnimidamonbackend/domain/model"
)

func MapToUser(um *model.User) *modelapi.User {
	if um == nil {
		return nil
	}

	return &modelapi.User{
		UserID:   int64(um.ID),
		Username: um.Username,
	}
}

func MapToUsers(ums []*model.User) []*modelapi.User {
	if ums == nil {
		return nil
	}

	var us []*modelapi.User

	for _, um := range ums {
		u := MapToUser(um)
		us = append(us, u)
	}

	return us
}

func MapToGroup(gm *model.Group) *modelapi.Group {
	if gm == nil {
		return nil
	}

	return &modelapi.Group{
		GroupID: int64(gm.ID),
		Name:    gm.Name,
	}
}

func MapToGroups(gms []*model.Group) []*modelapi.Group {
	if gms == nil {
		return nil
	}

	var gs []*modelapi.Group

	for _, gm := range gms {
		g := MapToGroup(gm)
		gs = append(gs, g)
	}

	return gs
}

func MapToComputer(cm *model.Computer) *modelapi.Computer {
	if cm == nil {
		return nil
	}

	return &modelapi.Computer{
		ComputerID: int64(cm.ID),
		Name:       cm.Name,
		OwnerID:    int64(cm.OwnerID),
	}
}

func MapToInvite(im *model.Invite) *modelapi.Invite {
	if im == nil {
		return nil
	}

	return &modelapi.Invite{
		Date:  strfmt.Date(im.CreatedAt),
		Group: MapToGroup(im.Group),
		User:  MapToUser(im.User),
	}
}

func MapToInvites(ims []*model.Invite) []*modelapi.Invite {
	if ims == nil {
		return nil
	}

	var is []*modelapi.Invite

	for _, im := range ims {
		i := MapToInvite(im)
		is = append(is, i)
	}

	return is
}