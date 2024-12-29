package utils

import "github.com/ronaldalds/res/internal/models"

func ContainsAll(listX, listY []models.Role) bool {
	// Criar um mapa para os itens de X
	itemMap := make(map[uint]bool)
	for _, item := range listX {
		itemMap[item.ID] = true
	}

	// Verificar se todos os itens de Y estão no mapa de X
	for _, item := range listY {
		if !itemMap[item.ID] {
			return false // Item de Y não está em X
		}
	}

	return true // Todos os itens de Y estão em X
}

func ExtrairPermissionUser(user models.User) []string {
	var permissions []string
	for _, role := range user.Roles {
		for _, permission := range role.Permissions {
			permissions = append(permissions, permission.Code)
		}
	}
	return permissions
}

func Contains(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
