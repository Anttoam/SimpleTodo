package usecase

import (
	"testing"

	"github.com/Anttoam/golang-htmx-todos/domain"
	"github.com/Anttoam/golang-htmx-todos/pkg/utils"
	"github.com/stretchr/testify/require"
)

func RandomUser(t *testing.T) (user domain.User, password string) {
	password = "password"
	hashedPassword, err := utils.HashPassword(password)
	require.NoError(t, err)

	user = domain.User{
		ID:       1,
		Name:     "test",
		Email:    "test@example.com",
		Password: hashedPassword,
	}
	return
}
