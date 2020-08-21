// Copyright 2020 Insolar Network Ltd.
// All rights reserved.
// This material is licensed under the Insolar License version 1.0,
// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.

package main

const (
	PublishCheckerMock = `
		// Copyright 2020 Insolar Network Ltd.
		// All rights reserved.
		// This material is licensed under the Insolar License version 1.0,
		// available at https://github.com/insolar/assured-ledger/blob/master/LICENSE.md.
		// Code generated by TypePublishGenerator. DO NOT EDIT.

		package checker

		import (
			"context"
			"time"

			"github.com/ThreeDotsLabs/watermill/message"
			"github.com/gojuno/minimock/v3"

			"github.com/insolar/assured-ledger/ledger-core/insolar/payload"
			"github.com/insolar/assured-ledger/ledger-core/vanilla/atomickit"
			"github.com/insolar/assured-ledger/ledger-core/vanilla/synckit"
		)

		// ============================================================================

	{{ range $pos, $msg := .Messages }}
		type {{ $msg }}Definition struct {
			touched       bool
			count         atomickit.Int
			countBefore   atomickit.Int
			expectedCount int
			handler       {{ $msg }}Handler
		}
		type {{ $msg }}Handler func(*payload.{{ $msg }} ) bool
		type Pub{{ $msg }}Mock struct{ parent *Typed }

		func (p Pub{{ $msg }}Mock) ExpectedCount(count int) Pub{{ $msg }}Mock {
			p.parent.Handlers.{{ $msg }}.touched = true
			p.parent.Handlers.{{ $msg }}.expectedCount = count
			return p
		}

		func (p Pub{{ $msg }}Mock) Set(handler {{ $msg }}Handler) Pub{{ $msg }}Mock {
			p.parent.Handlers.{{ $msg }}.touched = true
			p.parent.Handlers.{{ $msg }}.handler = handler
			return p
		}

		func (p Pub{{ $msg }}Mock) SetResend(resend bool) Pub{{ $msg }}Mock {
			p.parent.Handlers.{{ $msg }}.touched = true
			p.parent.Handlers.{{ $msg }}.handler = func(*payload.{{ $msg }}) bool { return resend }
			return p
		}

		func (p Pub{{ $msg }}Mock) Count() int {
			return p.parent.Handlers.{{ $msg }}.count.Load()
		}

		func (p Pub{{ $msg }}Mock) CountBefore() int {
			return p.parent.Handlers.{{ $msg }}.countBefore.Load()
		}

		func (p Pub{{ $msg }}Mock) Wait(ctx context.Context, count int) synckit.SignalChannel {
			return waitCounterIndefinitely(ctx, &p.parent.Handlers.{{ $msg }}.count, count)
		}

		// ============================================================================
	{{ end }}

		type TypedHandlers struct {
		{{ range $pos, $msg := .Messages -}}
			{{ $msg }} {{ $msg }}Definition
		{{ end }}

			BaseMessage struct {
				handler func(message *message.Message)
			}
		}

		type Typed struct {
			t             minimock.Tester
			timeout       time.Duration
			ctx           context.Context
			defaultResend bool
			resend        func(ctx context.Context, msg *message.Message)

			Handlers TypedHandlers

		{{ range $pos, $msg := .Messages -}}
			{{ $msg }} Pub{{ $msg }}Mock
		{{ end }}
		}

		func NewTyped(ctx context.Context, t minimock.Tester, sender Sender) *Typed {
			checker := &Typed{
				t:             t,
				ctx:           ctx,
				defaultResend: false,
				timeout:       10 * time.Second,
				resend:        sender.SendMessage,

				Handlers: TypedHandlers{
				{{ range $pos, $msg := .Messages -}}
					{{ $msg }}: {{ $msg }}Definition{expectedCount: -1},
				{{ end }}
				},
			}

			{{ range $pos, $msg := .Messages -}}
				checker.{{ $msg }} = Pub{{ $msg }}Mock{parent: checker}
			{{ end }}

			if controller, ok := t.(minimock.MockController); ok {
				controller.RegisterMocker(checker)
			}

			return checker
		}

		func (p *Typed) CheckMessages(topic string, messages ...*message.Message) error {
			for _, msg := range messages {
				p.checkMessage(p.ctx, msg)
			}

			return nil
		}

		func (p *Typed) checkMessage(ctx context.Context, msg *message.Message) {
			basePayload, err := payload.UnmarshalFromMeta(msg.Payload)
			if err != nil {
				return
			}

			var resend bool

			switch payload := basePayload.(type) {
		{{- range $pos, $msg := .Messages }}
			case *payload.{{ $msg }}:
				hdlStruct := &p.Handlers.{{ $msg }}

				resend = p.defaultResend

				oldCount := hdlStruct.countBefore.Add(1)

				if hdlStruct.handler != nil {
					done := make(synckit.ClosableSignalChannel)

					go func() {
						defer func() { _ = synckit.SafeClose(done) }()

						resend = hdlStruct.handler(payload)
					}()

					select {
					case <-done:
					case <-time.After(p.timeout):
						p.t.Error("timeout: failed to check message {{ $msg }} (position: %s)", oldCount)
					}
				} else if !p.defaultResend && !hdlStruct.touched {
					p.t.Fatalf("unexpected %T payload", payload)
					return
				}

				hdlStruct.count.Add(1)
		{{ end }}

			default:
				p.t.Fatalf("unexpected %T payload", basePayload)
				return
			}

			if resend {
				p.resend(ctx, msg)
			}
		}

		func (p *Typed) SetDefaultResend(flag bool) *Typed {
			p.defaultResend = flag
			return p
		}

		func (p *Typed) minimockDone() bool {
			ok := true

		{{ range $pos, $msg := .Messages -}}
			{
				fn := func () bool {
					hdl := &p.Handlers.{{ $msg }}

					switch {
					case hdl.expectedCount < 0:
						return true
					case p.defaultResend:
						return true
					case hdl.expectedCount == 0:
						return true
					}

					return hdl.count.Load() == hdl.expectedCount
				}

				ok = ok && fn()
			}
		{{ end }}

			return ok
		}

		// MinimockFinish checks that all mocked methods have been called the expected number of times
		func (p *Typed) MinimockFinish() {
			if !p.minimockDone() {
				p.t.Fatal("failed conditions on Typed")
			}
		}

		// MinimockWait waits for all mocked methods to be called the expected number of times
		func (p *Typed) MinimockWait(timeout time.Duration) {
			timeoutCh := time.After(timeout)
			for {
				if p.minimockDone() {
					return
				}
				select {
				case <-timeoutCh:
					p.MinimockFinish()
					return
				case <-time.After(10 * time.Millisecond):
				}
			}
		}
	`
)
