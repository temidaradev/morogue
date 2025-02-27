package states

import (
	"github.com/kettek/morogue/client/ifs"
	"github.com/kettek/morogue/net"
	"github.com/tinne26/etxt"
)

const (
	modeConnecting = "connecting"
	modeFailed     = "failed"
	modeLost       = "lost"
	modeSuccess    = "success"
)

// Connect is a state that manages and watches a connection to a
// morogue server. It should be noted that this doesn't sit in
// the state machine, but rather runs directly in the app struct.
// FIXME: At the moment, this doesn't handle reconnection and all
// that. What it should do is when a connection is lost, to pop
// the entire state machine, then attempt to connect and push
// the login state upon success.
type Connect struct {
	connection     net.Connection
	connectionChan chan error
	loopChan       chan error
	mode           string
	result         string
}

func (state *Connect) Begin(ctx ifs.RunContext) error {
	return state.Connect(ctx)
}

func (state *Connect) Connect(ctx ifs.RunContext) error {
	if state.mode == modeConnecting {
		return nil
	}
	state.mode = modeConnecting
	state.connectionChan = state.connection.Connect(ctx.Cfg.LastServer)
	return nil
}

func (state *Connect) Return(interface{}) error {
	return nil
}

func (state *Connect) Leave() error {
	return nil
}

func (state *Connect) End() (interface{}, error) {
	return nil, nil
}

func (state *Connect) Update(ctx ifs.RunContext) error {
	select {
	case err := <-state.connectionChan:
		if err != nil {
			state.mode = modeFailed
			state.result = err.Error()
			ctx.Sm.Push(NewServer(&state.mode, &state.connection, &state.connectionChan))
			return nil
		}
		state.mode = modeSuccess

		messageChan, loopChan := state.connection.Loop()
		state.loopChan = loopChan

		// Push the login state onto the stack.
		ctx.Sm.Push(NewLogin(state.connection, messageChan))
	case err := <-state.loopChan:
		if state.mode == modeSuccess {
			state.mode = modeLost
		} else {
			state.mode = modeFailed
		}
		state.result = err.Error()
		ctx.Sm.Push(NewServer(&state.mode, &state.connection, &state.connectionChan))
	default:
		//
	}

	return nil
}

func (state *Connect) Draw(ctx ifs.DrawContext) {
	// get screen center position and text content
	bounds := ctx.Screen.Bounds() // assumes origin (0, 0)
	x, y := 2, bounds.Dy()-4

	al := ctx.Txt.Renderer.GetAlign()
	ctx.Txt.Renderer.SetAlign(etxt.Baseline | etxt.Left)

	// draw the text
	ctx.Txt.Draw(ctx.Screen, state.mode, x, y)
	y += int(ctx.Txt.Utils().GetLineHeight())
	var last int
	for i := 0; i <= len(state.result); {
		if i >= len(state.result) {
			ctx.Txt.Draw(ctx.Screen, state.result[last:i], x, y)
			break
		} else if state.result[i] == '\n' {
			ctx.Txt.Draw(ctx.Screen, state.result[last:i], x, y)
			y += int(ctx.Txt.Utils().GetLineHeight())
		}
		i++
	}
	ctx.Txt.Renderer.SetAlign(al)
}
