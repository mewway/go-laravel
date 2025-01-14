package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mewway/go-laravel/contracts/queue"
	"github.com/mewway/go-laravel/support/file"
	testingfile "github.com/mewway/go-laravel/testing/file"
)

type Test struct {
}

// Signature The name and signature of the job.
func (receiver *Test) Signature() string {
	return "test"
}

// Handle Execute the job.
func (receiver *Test) Handle(args ...any) error {
	return file.Create("test.txt", args[0].(string))
}

func TestDispatchSync(t *testing.T) {
	task := &Task{
		jobs: []queue.Jobs{
			{
				Job: &Test{},
				Args: []queue.Arg{
					{Type: "uint64", Value: "test"},
				},
			},
		},
	}

	err := task.DispatchSync()
	assert.Nil(t, err)
	assert.True(t, file.Exists("test.txt"))
	assert.True(t, testingfile.GetLineNum("test.txt") == 1)
	assert.Nil(t, file.Remove("test.txt"))
}
