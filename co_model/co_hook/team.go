package co_hook

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_enum"
	"github.com/jack353249002/exam-message-send-modules/co_model"
)

type InviteJoinTeamHookFunc func(ctx context.Context, state sys_enum.InviteType, invite *sys_model.InviteRes, teamInfo co_model.ITeamRes) (bool, error)
