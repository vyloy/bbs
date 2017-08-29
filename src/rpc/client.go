package rpc

import (
	"context"
	"github.com/skycoin/bbs/src/misc/boo"
	"github.com/skycoin/bbs/src/store/object"
	"net/rpc"
)

// Call represents a function that outputs method and input object of call.
type Call func() (method string, in interface{})

// Send sends the "call" to specified addresses.
func Send(ctx context.Context, addresses interface{}, req Call) (goal uint64, e error) {
	for _, address := range addresses.([]string) {

		var client *rpc.Client
		client, e = rpc.Dial("tcp", address)
		if e != nil {
			continue
		}

		methodName, in := req()
		call := client.Go(methodName, in, &goal, nil)

		select {
		case <-call.Done:
			e = call.Error

		case <-ctx.Done():
			e = ctx.Err()
		}

		return
	}

	return 0, boo.New(boo.NotFound,
		"successful submission address not found")
}

// NewThread is a call to create a new thread.
func NewThread(thread *object.Thread) Call {
	return func() (string, interface{}) {
		return "Gateway.NewThread", thread
	}
}

// NewPost is a call to create a new post.
func NewPost(post *object.Post) Call {
	return func() (string, interface{}) {
		return "Gateway.NewPost", post
	}
}

// NewVote is a call to create a new vote.
func NewVote(vote *object.Vote) Call {
	return func() (string, interface{}) {
		return "Gateway.NewVote", vote
	}
}