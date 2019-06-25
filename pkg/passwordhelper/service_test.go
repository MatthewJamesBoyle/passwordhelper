package passwordhelper_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"password-helper/pkg/passwordhelper"
)

func TestService_GetCharAt(t *testing.T) {
	invalidTestCases := []struct {
		testDefinition string
		input          struct {
			password string
			indexes  []int
		}
		expected struct {
			result *string
			error  error
		}
	}{
		{
			testDefinition: "Returns error given empty password",
			input: struct {
				password string
				indexes  []int
			}{password: "", indexes: []int{1}}, expected: struct {
			result *string
			error  error
		}{
			result: nil, error: passwordhelper.ErrEmptyPass},
		},
		{
			testDefinition: "Returns error given indexes of 0 or less (1 value)",
			input: struct {
				password string
				indexes  []int
			}{password: "some-valid-password", indexes: []int{0}}, expected: struct {
			result *string
			error  error
		}{
			result: nil, error: passwordhelper.ErrInvalidIndex},
		},
		{
			testDefinition: "Returns error given indexes of 0 or less (multiple values,error in middle)",
			input: struct {
				password string
				indexes  []int
			}{password: "some-valid-password", indexes: []int{3, 4, 5, -4, 6, 7, 8}}, expected: struct {
			result *string
			error  error
		}{
			result: nil, error: passwordhelper.ErrInvalidIndex},
		},
		{
			testDefinition: "Returns error given any provided indexes are > length of string",
			input: struct {
				password string
				indexes  []int
			}{password: "length7", indexes: []int{3, 4, 5, 2, 3, 7, 8}}, expected: struct {
			result *string
			error  error
		}{
			result: nil, error: passwordhelper.ErrInvalidIndex},
		},
	}

	s := passwordhelper.Service{}
	for _, test := range invalidTestCases {
		t.Run(test.testDefinition, func(t *testing.T) {
			res, err := s.CharsAt(test.input.password, test.input.indexes...)
			assert.Equal(t, test.expected.error, err)
			assert.Equal(t, test.expected.result, res)
		})
	}

	// Have to split up positive and negative test cases due to the pointer...
	// Is it worth it? Could I have just returned ""? Probably
	t.Run("Returns p,a,r,t,y and no error given partypooper password and correct indexes", func(t *testing.T) {
		s := passwordhelper.Service{}
		res, err := s.CharsAt("partypooper", 1, 2, 3, 4, 5)
		require.NoError(t, err)
		assert.Equal(t, "p,a,r,t,y", *res)
	})
	t.Run("returns correct result given 1 char pass and 1 index requested", func(t *testing.T) {
		s := passwordhelper.Service{}
		res, err := s.CharsAt("p", 1)
		require.NoError(t, err)
		assert.Equal(t, "p", *res)
	})
	t.Run("maintains case of requests chars", func(t *testing.T) {
		s := passwordhelper.Service{}
		res, err := s.CharsAt("IaMD1ff3RentCaz0s", 1,4,7,8)
		require.NoError(t, err)
		assert.Equal(t, "I,D,f,3", *res)
	})

}
