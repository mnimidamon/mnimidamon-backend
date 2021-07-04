package modelsql

import (
	"mnimidamonbackend/domain/model"
	"time"
)

// Represents a Group which has GroupMembers and a list of active Invites.
type Group struct {
	Entity

	Name string `gorm:"unique; index; size:15"`

	GroupMembers []User   `gorm:"many2many:group_members;"`
	Invites      []Invite `gorm:"many2many:group_invites;"`
}

func NewGroupFromBusinessModel(gm *model.Group) *Group {
	if gm == nil {
		return nil
	}

	g := &Group{
		Entity: Entity{
			ID:        gm.ID,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		Name: gm.Name,
		GroupMembers: nil,
		Invites: nil,
	}

	if gm.Invites != nil {
		var gInvites []Invite
		for _, mi := range gm.Invites {
			i := NewInviteFromBusinessModel(&mi)
			gInvites = append(gInvites, *i)
		}
		g.Invites = gInvites
	}

	if gm.GroupMembers != nil {
		var gMembers []User
		for _, mgm := range gm.GroupMembers {
			gm := NewUserFromBusinessModel(&mgm)
			gMembers = append(gMembers, *gm)
		}
		g.GroupMembers = gMembers
	}

	return g
}

func (g *Group) NewBusinessModel() *model.Group {
	if g == nil {
		return nil
	}

	gm := new(model.Group)
	g.CopyToBusinessModel(gm)
	return gm
}

func (g *Group) CopyToBusinessModel(gm *model.Group) {
	if g == nil {
		gm = nil
		return
	}

	gm.ID = g.ID
	gm.Name = g.Name

	if g.GroupMembers != nil {
		var gmMembers []model.User
		for _, m := range g.GroupMembers {
			mm := m.NewBusinessModel()
			gmMembers = append(gmMembers, *mm)
		}
		gm.GroupMembers = gmMembers
	}

	if g.Invites != nil {
		var gmInvites []model.Invite
		for _, i := range g.Invites {
			im := i.NewBusinessModel()
			gmInvites = append(gmInvites, *im)
		}
		gm.Invites = gmInvites
	}
}
