package foodlikebiz

import (
	"Food-delivery/common"
	foodlikemodel "Food-delivery/module/foodlike/model"
	"context"
)

type ListUserLikeFoodStore interface {
	GetUserLikeFood(context context.Context, conditions map[string]interface{}, filter *foodlikemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserLikeFoodBiz struct {
	store ListUserLikeFoodStore
}

func NewListUserLikeFood(store ListUserLikeFoodStore) *listUserLikeFoodBiz {
	return &listUserLikeFoodBiz{store: store}
}

func (biz *listUserLikeFoodBiz) ListUserLike(context context.Context, filter *foodlikemodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUserLikeFood(context, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(foodlikemodel.EntityName, err)
	}
	return users, nil
}
