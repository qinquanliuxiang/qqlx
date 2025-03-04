package store

import (
	"context"
	"fmt"
	"qqlx/base"
	"qqlx/base/apierrs"
	"qqlx/model"

	"gorm.io/gorm"
)

type GeneralRoleStore struct {
	store *gorm.DB
}

func NewRoleStore(store *gorm.DB) *GeneralRoleStore {
	return &GeneralRoleStore{
		store: store,
	}
}

func (r *GeneralRoleStore) Create(ctx context.Context, role *model.Role) (err error) {
	if role == nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to create role, role is nil"))
	}
	err = r.store.WithContext(ctx).Create(&role).Error
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to create role, %w", err))
	}
	return nil
}

func (r *GeneralRoleStore) Save(ctx context.Context, role *model.Role) (err error) {
	if role == nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to save role, role is nil"))
	}
	err = r.store.WithContext(ctx).Save(&role).Error
	if err != nil {
		return apierrs.NewSaveError(fmt.Errorf("failed to save role, %w", err))
	}
	return nil
}

func (r *GeneralRoleStore) Delete(ctx context.Context, role *model.Role) (err error) {
	err = r.store.WithContext(ctx).Delete(&role).Error
	if err != nil {
		return apierrs.NewDeleteError(fmt.Errorf("failed to delete user, %w", err))
	}
	return nil
}

func (r *GeneralRoleStore) List(ctx context.Context, page int, pageSize int) (total int64, roles []*model.Role, err error) {
	query := r.store.WithContext(ctx).Model(&model.Role{})
	// 计数查询
	err = query.Count(&total).Error
	if err != nil {
		return 0, nil, apierrs.NewListError(fmt.Errorf("failed to count roles, %w", err))
	}

	// 数据查询
	err = query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Preload("Policys").
		Find(&roles).Error
	if err != nil {
		return 0, nil, apierrs.NewListError(fmt.Errorf("failed to list roles, %w", err))
	}

	return total, roles, nil
}

func (r *GeneralRoleStore) GetRoleByID(ctx context.Context, id int, options ...base.RoleQueryOption) (role *model.Role, err error) {
	query := r.store.WithContext(ctx).Model(&role).Where("id = ?", id)
	// 添加查询选项
	for _, option := range options {
		option(query)
	}
	if err = query.Take(&role).Error; err != nil {
		return nil, apierrs.NewGetError(fmt.Errorf("failed to get role, %w", err))

	}
	return role, nil
}

func (r *GeneralRoleStore) GetRoleByName(ctx context.Context, name string, options ...base.RoleQueryOption) (role *model.Role, err error) {
	query := r.store.WithContext(ctx).Model(&role).Where("name = ?", name)
	// 添加查询选项
	for _, option := range options {
		option(query)
	}

	if err = query.Take(&role).Error; err != nil {
		return nil, apierrs.NewGetError(fmt.Errorf("failed to get role, %w", err))
	}
	return role, nil
}

type RoleAssociationStore struct {
	store *gorm.DB
}

func NewRoleAssociationStore(store *gorm.DB) *RoleAssociationStore {
	return &RoleAssociationStore{
		store: store,
	}
}

func (r *RoleAssociationStore) AppendPolicy(ctx context.Context, role *model.Role, appendPolicy []*model.Policy) (err error) {
	err = r.store.WithContext(ctx).Model(&role).Association("Policys").Append(&appendPolicy)
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to append policy, %w", err))
	}
	return nil
}

func (r *RoleAssociationStore) ReplacePolicy(ctx context.Context, role *model.Role, policy []*model.Policy) (err error) {
	err = r.store.WithContext(ctx).Model(&role).Association("Policys").Replace(&policy)
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to replace policy, %w", err))
	}
	return nil
}

func (r *RoleAssociationStore) DeletePolicy(ctx context.Context, role *model.Role, policy []*model.Policy) (err error) {
	err = r.store.WithContext(ctx).Model(&role).Association("Policys").Delete(&policy)
	if err != nil {
		return apierrs.NewCreateError(fmt.Errorf("failed to delete policy, %w", err))
	}
	return nil
}
