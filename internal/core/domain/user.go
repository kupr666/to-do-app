package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/kupr666/to-do-app/internal/core/errors"
)

type User struct {
	ID      int
	Version int

	FullName    string
	PhoneNumber *string
}

func NewUser(
	id int,
	version int,
	fullName string,
	phoneNumber *string,
) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(
	fullName string,
	phoneNumber *string,
) User {
	return NewUser(
		uninitializedID,
		uninitializedVersion,
		fullName,
		phoneNumber,
	)
}

func (u *User) Validate() error {

	// FullName check
	fullNameLen := len([]rune(u.FullName))
	if fullNameLen < 3 && fullNameLen > 100 {
		return fmt.Errorf("invalid `fullName` len: %d: %w",
			fullNameLen,
			core_errors.ErrInvalidArgument)
	}

	// PhoneNumber check
	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf(
				"invalid `PhoneNumber` len: %d: %w",
				phoneNumberLen,
				core_errors.ErrInvalidArgument,
			)
		}

		re := regexp.MustCompile(`^\+[0-9]+$`)

		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf(
				"invalid `PhoneNumber` format: %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}
	return nil
}

type UserPatch struct {
	FullName    Nullable[string]
	PhoneNumber Nullable[string]
}

func NewUserPatch(
	fullName Nullable[string],
	phoneNumber Nullable[string],
) UserPatch {
	return UserPatch{
		FullName:    fullName,
		PhoneNumber: phoneNumber,
	}
}

func (u *UserPatch) Validate() error {
	if u.FullName.Set && u.FullName.Value == nil {
		return fmt.Errorf(
			"`FullName` can't be patched to null: %w",
			core_errors.ErrInvalidArgument,
		)
	}
	return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.FullName.Set {
		tmp.FullName = *patch.FullName.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
