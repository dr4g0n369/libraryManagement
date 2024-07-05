package helper

import (
	"testing"

	"github.com/dr4g0n369/libraryManagement/pkg/types"
)

func TestCreateToken(t *testing.T) {
	user := types.Login{
		Id:       2,
		Username: "random",
		Role:     "user",
	}

	want := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Miwicm9sZSI6InVzZXIiLCJ1c2VybmFtZSI6InJhbmRvbSJ9.C6yR5XBxVpgNl_YQjJfPLOQD53VDVZvfiNaHST4cgcs"
	got, err := CreateToken(&user)
	if err != nil {
		t.Errorf(err.Error())
	}

	if got != want {
		t.Errorf("Incorrect token: %v\nExpected token: %v", got, want)
	}
}
