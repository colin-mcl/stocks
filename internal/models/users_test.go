package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

type newUserRequest struct {
	username       string
	email          string
	hashedPassword []byte
	firstName      string
	lastName       string
}

type newUserCases struct {
	name   string
	input  newUserRequest
	output *User
	err    error
}

// 256 char string for testing
var longString string = "KpK$$=cQ7Y+_ttVFAbv2+z2_PRUeFu%(H?+y[GTY[GF(qQx007[;]XU&GMGp@q$eb!R7J7Rb#rb2Q9FV,uDVY3jTRwmqSvP23VJv9i!}fAj}B-,a{j,:*FuE?E(06Np,He=Jnq6[G*5mqGnmZWL/e9u3ehU(zha:+GdL%%.HjC$w.gke.$p)zhSK(P:_+LgSu{a=*}e%#N%_NQ$FnHcN%8*4N7UziLv-vA{n9{hRebiSHrg9MG#(ZWt_TrU(3_Z}"

// 72 char string for testing longest password
var longPassword string = "jfq!JuS_.w#Stu#?[m6Jj/]m,!M,?*V0@_c0]}YT{]pk}45zbu)S2}S[/hGv]vnDxmFTa@gQ"

func TestNewUserErr(t *testing.T) {

	// bad hashed password for testing
	badPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)

	for _, scenario := range []newUserCases{
		{
			name: "empty username",
			input: newUserRequest{
				username:       "",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "last",
			},
			err: ErrEmptyField,
		},
		{
			name: "empty email",
			input: newUserRequest{
				username:       "default",
				email:          "",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "last",
			},
			err: ErrEmptyField,
		},
		{
			name: "empty password",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: nil,
				firstName:      "first",
				lastName:       "last",
			},
			err: ErrEmptyField,
		},
		{
			name: "empty firstname",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "",
				lastName:       "last",
			},
			err: ErrEmptyField,
		},
		{
			name: "empty lastname",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "",
			},
			err: ErrEmptyField,
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			u, e := NewUser(
				scenario.input.username,
				scenario.input.email,
				scenario.input.hashedPassword,
				scenario.input.firstName,
				scenario.input.lastName)
			assert.Nil(t, u)
			assert.EqualError(t, e, scenario.err.Error())
		})
	}
}

// tests successful new user cases
func TestNewUserSuc(t *testing.T) {

	// bad hashed password for testing
	badPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), 12)
	longPassword, _ := bcrypt.GenerateFromPassword([]byte(longPassword), 12)

	for _, scenario := range []newUserCases{
		{
			name: "valid account",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "last",
			},
		},
		{
			name: "long email",
			input: newUserRequest{
				username:       "username",
				email:          longString,
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "last",
			},
		},
		{
			name: "long username",
			input: newUserRequest{
				username:       longString,
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       "last",
			},
		},
		{
			name: "long password",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: longPassword,
				firstName:      "first",
				lastName:       "last",
			},
		},
		{
			name: "long firstname",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      longString,
				lastName:       "last",
			},
		},
		{
			name: "long lastname",
			input: newUserRequest{
				username:       "username",
				email:          "default",
				hashedPassword: badPassword,
				firstName:      "first",
				lastName:       longString,
			},
		},
	} {
		t.Run(scenario.name, func(t *testing.T) {
			u, e := NewUser(
				scenario.input.username,
				scenario.input.email,
				scenario.input.hashedPassword,
				scenario.input.firstName,
				scenario.input.lastName)
			assert.Nil(t, e)
			assert.Equal(t, scenario.input.username, u.Username)
			assert.Equal(t, scenario.input.email, u.Email)
			assert.Equal(t, scenario.input.hashedPassword, u.HashedPassword)
			assert.Equal(t, scenario.input.firstName, u.FirstName)
			assert.Equal(t, scenario.input.lastName, u.LastName)
		})
	}
}

// func TestAuthenticateUser(t *testing.T) {
// 	// Tests a non existent user
// 	id, err := users.Authenticate("fakeEmail", "badPassword")

// 	require.Equal(t, id, -1)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidCredentials.Error())

// 	// Test an existing user with the wrong password
// 	id, err = users.Authenticate("colin.mcl@gmail.com", "blah")

// 	require.Equal(t, id, -1)
// 	require.Error(t, err)
// 	require.EqualError(t, err, ErrInvalidCredentials.Error())

// 	// Test an existing user with the correct password
// 	id, err = users.Authenticate("colin.mcl@gmail.com", "Password123!")

// 	require.Equal(t, id, 1)
// 	require.Nil(t, err)
// }

// func TestInsertUser(t *testing.T) {
// 	// Inserting user with email that already exists
// 	id, err := users.Insert("bad", "user", "badusername",
// 		"colin.mcl@gmail.com", "pass")

// 	require.Equal(t, id, -1)
// 	require.Error(t, err)

// 	// Check that non existant user doesn't exist
// 	exists, _ := users.Exists(0)
// 	require.False(t, exists)

// 	// Inserting real user and check it exists
// 	password := time.Now().String()
// 	username := "realuser#" + password
// 	email := username + "@email.com"
// 	id, err = users.Insert("real", "user", username, email, password)

// 	require.NotEqual(t, id, -1)
// 	require.Nil(t, err)

// 	exists, err = users.Exists(id)
// 	require.True(t, exists)
// 	require.Nil(t, err)
// }

// func TestExists(t *testing.T) {
// 	exists, err := users.Exists(0)
// 	require.Nil(t, err)
// 	require.False(t, exists)

// 	exists, err = users.Exists(1)
// 	require.Nil(t, err)
// 	require.True(t, exists)
// }
