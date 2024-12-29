package database

import (
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"github.com/ronaldalds/res/internal/models"
	"github.com/ronaldalds/res/internal/settings"
	"github.com/ronaldalds/res/internal/utils"
	"gorm.io/gorm"
)

func (gs *GormStore) CreateAdmin() error {
	var user models.User
	err := gs.DB.Where("username = ?", settings.Env.SuperUsername).First(&user).Error
	if err == nil {
		return fmt.Errorf("admin already exists")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fmt.Errorf("failed to check admin existence: %s", err.Error())
	}
	hashPassword, err := utils.HashPassword(settings.Env.SuperPass)
	if err != nil {
		return fmt.Errorf("failed to create admin: %s", err.Error())
	}
	admin := &models.User{
		FirstName:   settings.Env.SuperName,
		LastName:    "Admin",
		Username:    settings.Env.SuperUsername,
		Email:       settings.Env.SuperEmail,
		Password:    hashPassword,
		Active:      true,
		IsSuperUser: true,
		Phone1:      fmt.Sprintf("+%v", settings.Env.SuperPhone),
	}
	if err := gs.DB.Create(&admin).Error; err != nil {
		return fmt.Errorf("failed to create user: %s", err.Error())
	}
	return fmt.Errorf("admin created successfully")
}

func findMissingIDsByPermissions(ids []uint, permissions []models.Permission) []uint {
	// Criar um mapa dos IDs encontrados
	foundIDs := make(map[uint]struct{})
	for _, p := range permissions {
		foundIDs[p.ID] = struct{}{}
	}

	// Identificar os IDs ausentes
	var missingIDs []uint
	for _, id := range ids {
		if _, exists := foundIDs[id]; !exists {
			missingIDs = append(missingIDs, id)
		}
	}

	return missingIDs
}

func findMissingIDsByRoles(ids []uint, roles []models.Role) []uint {
	// Criar um mapa dos IDs encontrados
	foundIDs := make(map[uint]struct{})
	for _, p := range roles {
		foundIDs[p.ID] = struct{}{}
	}

	// Identificar os IDs ausentes
	var missingIDs []uint
	for _, id := range ids {
		if _, exists := foundIDs[id]; !exists {
			missingIDs = append(missingIDs, id)
		}
	}

	return missingIDs
}

func (gs *GormStore) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := gs.DB.
		Preload("Roles.Permissions").
		Where("id = ?", id).
		First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found for id: %d", id)
		}
		return nil, fmt.Errorf("failed to query database: %w", result.Error)
	}
	return &user, nil
}

func (gs *GormStore) GetUserByUsernameOrEmail(text string) (*models.User, error) {
	var user models.User
	result := gs.DB.
		Preload("Roles.Permissions").
		Where("username = ? OR email = ?", text, text).
		First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found for email: %s", text)
		}
		return nil, fmt.Errorf("failed to query database: %w", result.Error)
	}
	return &user, nil
}

func (gs *GormStore) CheckIfUserExistsByUsernameOrEmail(email, username string) error {
	var user models.User
	if err := gs.DB.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		return fmt.Errorf("no record found for email or username")
	}
	return nil
}

func (gs *GormStore) GetPermissionByIds(ids []uint) ([]models.Permission, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no permission IDs provided")
	}

	var permissions []models.Permission
	// Buscar as permissões pelos IDs fornecidos
	if err := gs.DB.Where("id IN ?", ids).Find(&permissions).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch permissions: %s", err.Error())
	}

	// Verificar se todas as permissões foram encontradas
	if len(permissions) != len(ids) {
		missingIDs := findMissingIDsByPermissions(ids, permissions)
		return nil, fmt.Errorf("permissions not found for IDs: %v", missingIDs)
	}

	return permissions, nil
}

func (gs *GormStore) GetRoleByIds(ids []uint) ([]models.Role, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no role IDs provided")
	}

	var roles []models.Role
	// Buscar as permissões pelos IDs fornecidos
	if err := gs.DB.Where("id IN ?", ids).Find(&roles).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch roles: %s", err.Error())
	}

	// Verificar se todas as permissões foram encontradas
	if len(roles) != len(ids) {
		missingIDs := findMissingIDsByRoles(ids, roles)
		return nil, fmt.Errorf("roles not found for IDs: %v", missingIDs)
	}

	return roles, nil
}

func (gs *GormStore) CheckIfPermissionExistsByCodeOrName(code, name string) error {
	var permission models.Permission
	result := gs.DB.Where("code = ? OR name = ?", code, name).First(&permission)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("no record found for code or name")
		}
		return fmt.Errorf("failed to query database: %w", result.Error)
	}
	return nil
}

func (gs *GormStore) CheckIfRoleExistsByIds(id []uint) error {
	var rolesCount int64
	if err := gs.DB.Model(&models.Role{}).
		Where("id IN ?", id).
		Count(&rolesCount).Error; err != nil {
		return fmt.Errorf("failed to validate roles: %s", err.Error())
	}
	if rolesCount != int64(len(id)) {
		return fmt.Errorf("some roles are invalid or do not exist")
	}
	return nil
}

func (gs *GormStore) CheckIfRoleExistsByName(name string) error {
	var role models.Role
	result := gs.DB.Where("name = ?", name).First(&role)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("no record found for name")
		}
		return fmt.Errorf("failed to query database: %w", result.Error)
	}
	return nil
}
