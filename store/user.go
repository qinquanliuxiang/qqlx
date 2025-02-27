package store

import (
	"context"
	"fmt"
	"os/user"
	"qqlx/base"
	"qqlx/base/apierrs"
	"qqlx/model"

	"gorm.io/gorm"
)

type UserStore struct {
	store *gorm.DB
}

func NewUserStore(store *gorm.DB) *UserStore {
	return &UserStore{
		store: store,
	}
}

func (u *UserStore) Create(ctx context.Context, user *model.User) (err error) {
	if user == nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to create user, user is nil"))
	}
	err = u.store.WithContext(ctx).Create(&user).Error
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to create user, %w", err))
	}
	return nil
}

func (u *UserStore) Delete(ctx context.Context, user *model.User) (err error) {
	err = u.store.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete user, %w", err))
	}
	return nil
}

func (u *UserStore) Save(ctx context.Context, user *model.User) (err error) {
	if user == nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to save user, user is nil"))
	}
	if err = u.store.WithContext(ctx).Save(&user).Error; err != nil {
		return apierrs.NewSaveError(fmt.Errorf("failed to save user, %w", err))
	}
	return nil
}

func (u *UserStore) GetUserByID(ctx context.Context, id uint, options ...base.UserQueryOption) (user *model.User, err error) {
	query := u.store.WithContext(ctx).Model(&user).Where("id = ?", id)
	for _, option := range options {
		option(query)
	}

	if err = query.Take(&user).Error; err != nil {
		return nil, apierrs.NewGetError(fmt.Errorf("failed to get user, %w", err))
	}
	return user, nil
}

func (u *UserStore) GetUserByName(ctx context.Context, name string, options ...base.UserQueryOption) (user *model.User, err error) {
	query := u.store.WithContext(ctx).Model(&user).Where("name = ?", name)
	for _, option := range options {
		option(query)
	}
	if err = query.Take(&user).Error; err != nil {
		return nil, apierrs.NewGetError(fmt.Errorf("failed to get user, %w", err))
	}
	return user, nil
}

func (u *UserStore) GetUserByEmail(ctx context.Context, email string, options ...base.UserQueryOption) (user *model.User, err error) {
	query := u.store.WithContext(ctx).Model(&user).Where("email = ?", email)
	for _, option := range options {
		option(query)
	}
	if err = query.Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserStore) List(ctx context.Context, page, pageSize int) (total int64, users []*model.User, err error) {
	// 计数查询
	query := u.store.WithContext(ctx).Model(&user.User{})
	if err = query.Count(&total).Error; err != nil {
		return 0, nil, apierrs.NewListError(fmt.Errorf("failed to count users, %w", err))

	}

	// 数据查询
	if err = query.
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Find(&users).Error; err != nil {
		return 0, nil, apierrs.NewListError(fmt.Errorf("failed to list users, %w", err))
	}
	return total, users, nil
}

type GetUserStore struct {
	*UserStore
}

func NewGetUserStore(userStore *UserStore) *GetUserStore {
	return &GetUserStore{
		UserStore: userStore,
	}
}
