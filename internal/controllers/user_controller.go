package controllers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ronaldalds/res/internal/handlers"
	"github.com/ronaldalds/res/internal/i18n"
	"github.com/ronaldalds/res/internal/schemas"
	"github.com/ronaldalds/res/internal/utils"
	"github.com/ronaldalds/res/internal/validators"
)

func (con *Controller) CreateUserHandler(c *fiber.Ctx) error {
	validate := validators.NewValidator()
	var body schemas.CreateUser
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(i18n.ERR_INVALID_BODY, err.Error()))
	}
	if validationErrors := validate.ValidateStruct(&body); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	if err := validate.ValidatePassword(body.Password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(i18n.ERR_CRYPTING_PASSWORD_FAILED, err.Error()))
	}
	body.Password = hashedPassword

	creator, err := utils.GetJwtHeaderPayload(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := con.Service.CreateUser(creator.Claims.Sub, body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}
	res := &schemas.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Active:    user.Active,
		RoleNames: roleNames,
		Phone1:    user.Phone1,
		Phone2:    user.Phone2,
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (con *Controller) RegisterPermissiontHandler(c *fiber.Ctx) error {
	errors := handlers.NewError()
	validate := validators.NewValidator()
	var body schemas.CreatePermissionRequest
	if err := c.BodyParser(&body); err != nil {
		errors.AddDetailErr("bodyParser", fmt.Sprintf(i18n.ERR_INVALID_BODY, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if validationErrors := validate.ValidateStruct(&body); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	permission, err := con.Service.CreatePermission(body)
	if err != nil {
		errors.AddDetailErr("Register", err.Error())
		return c.Status(fiber.StatusCreated).JSON(errors)
	}
	res := &schemas.CreatePermissionResponse{
		ID:   permission.ID,
		Code: permission.Code,
		Name: permission.Name,
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (con *Controller) RegisterRoleHandler(c *fiber.Ctx) error {
	errors := handlers.NewError()
	validate := validators.NewValidator()
	var body schemas.CreateRoleRequest
	if err := c.BodyParser(&body); err != nil {
		errors.AddDetailErr("bodyParser", fmt.Sprintf(i18n.ERR_INVALID_BODY, err.Error()))
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	if validationErrors := validate.ValidateStruct(&body); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	role, err := con.Service.CreateRole(body)
	if err != nil {
		errors.AddDetailErr("Register", err.Error())
		return c.Status(fiber.StatusCreated).JSON(errors)
	}
	// Extrair os codes das permissões
	var permissionCodes []string
	for _, permission := range role.Permissions {
		permissionCodes = append(permissionCodes, permission.Code)
	}

	// Preparar a resposta
	res := &schemas.CreateRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Permissions: permissionCodes, // Adicionar apenas os codes
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func (con *Controller) UpdateUserHandler(c *fiber.Ctx) error {
	validate := validators.NewValidator()
	idParam := c.Params("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, i18n.ERR_INVALID_ID_PARAMS)
	}
	var body schemas.UpdateUser
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(i18n.ERR_INVALID_BODY, err.Error()))
	}
	if validationErrors := validate.ValidateStruct(&body); validationErrors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(validationErrors)
	}
	editor, err := utils.GetJwtHeaderPayload(c)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	user, err := con.Service.UpdateUser(editor.Claims.Sub, uint(id), body)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var roleNames []string
	for _, role := range user.Roles {
		roleNames = append(roleNames, role.Name)
	}
	res := &schemas.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		Email:     user.Email,
		Active:    user.Active,
		RoleNames: roleNames,
		Phone1:    user.Phone1,
		Phone2:    user.Phone2,
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
