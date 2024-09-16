package helpers_test

import (
	"os"
	"testing"
	"time"

	"github.com/kenmobility/git-api-service/pkg/helpers"
	"github.com/stretchr/testify/assert"
)

// Test RandomInt function
func TestRandomInt(t *testing.T) {
	min := int64(10)
	max := int64(100)

	for i := 0; i < 10; i++ {
		randInt := helpers.RandomInt(min, max)
		assert.GreaterOrEqual(t, randInt, min)
		assert.LessOrEqual(t, randInt, max)
	}
}

// Test IsLocal function
func TestIsLocal(t *testing.T) {
	// Test when APP_ENV is not set
	os.Setenv("APP_ENV", "")
	assert.True(t, helpers.IsLocal())

	// Test when APP_ENV is set to "local"
	os.Setenv("APP_ENV", "local")
	assert.True(t, helpers.IsLocal())

	// Test when APP_ENV is set to "production"
	os.Setenv("APP_ENV", "production")
	assert.False(t, helpers.IsLocal())
}

// Test Getenv function
func TestGetenv(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "value")
	assert.Equal(t, "value", helpers.Getenv("TEST_ENV_VAR"))

	os.Unsetenv("TEST_ENV_VAR")
	assert.Equal(t, "default_value", helpers.Getenv("TEST_ENV_VAR", "default_value"))
	assert.Equal(t, "", helpers.Getenv("TEST_ENV_VAR"))
}

// Test IsRepositoryNameValid function
func TestIsRepositoryNameValid(t *testing.T) {
	assert.True(t, helpers.IsRepositoryNameValid("owner/repo"))
	assert.False(t, helpers.IsRepositoryNameValid("invalid_repo_name"))
}

// Test ValidateInput function
type SampleInput struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func TestValidateInput(t *testing.T) {
	input := SampleInput{
		Name:  "",
		Email: "invalid-email",
	}

	errors := helpers.ValidateInput(input)
	assert.Contains(t, errors, "Name field is required")
}

// Test RandomString function
func TestRandomString(t *testing.T) {
	randomStr := helpers.RandomString(6)
	assert.Len(t, randomStr, 6)

	randomStr = helpers.RandomString(10)
	assert.Len(t, randomStr, 10)
}

// Test RandomRepositoryName function
func TestRandomRepositoryName(t *testing.T) {
	repoName := helpers.RandomRepositoryName()
	assert.Contains(t, repoName, "/")
	assert.Len(t, repoName, 13) // 6 characters for owner, 6 for repo, 1 for slash
}

// Test RandomRepositoryUrl function
func TestRandomRepositoryUrl(t *testing.T) {
	repoUrl := helpers.RandomRepositoryUrl()
	assert.Contains(t, repoUrl, "https://github.com/")
}

// Test RandomFetchStartDate function
func TestRandomFetchStartDate(t *testing.T) {
	startDate := helpers.RandomFetchStartDate()
	expectedStartDate := time.Now().AddDate(0, -8, 0)
	assert.WithinDuration(t, expectedStartDate, startDate, time.Hour*24)
}

// Test RandomFetchEndDate function
func TestRandomFetchEndDate(t *testing.T) {
	endDate := helpers.RandomFetchEndDate()
	assert.WithinDuration(t, time.Now(), endDate, time.Second*1)
}

// Test RandomWords function
func TestRandomWords(t *testing.T) {
	randomWords := helpers.RandomWords(3)
	words := len(randomWords) - len(randomWords)/5 + 1
	assert.GreaterOrEqual(t, words, 3)
}
