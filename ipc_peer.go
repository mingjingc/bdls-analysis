// BSD 3-Clause License
//
// Copyright (c) 2020, Sperax
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package bdls

import (
	"crypto/ecdsa"
	fmt "fmt"
	math "math"
	rand "math/rand"
	"net"
	"sync"
	"time"
	"unsafe"

	"github.com/Sperax/bdls/timer"
)

// fake address for IPCPeer
type fakeAddress string

func (fakeAddress) Network() string  { return "ipc" }
func (f fakeAddress) String() string { return string(f) }

// IPCPeer represents an in-process peer for testing, which sends messages
// directly via function call, message delivery latency can be customizable
// to emulate variety of network latency. Delay is randomized with standard
// normal distribution based on given parameters.
type IPCPeer struct {
	c *Consensus
	sync.Mutex
	latency      time.Duration
	die          chan struct{}
	dieOnce      sync.Once
	msgCount     int64
	bytesCount   int64
	minLatency   time.Duration
	maxLatency   time.Duration
	totalLatency time.Duration
}

// NewIPCPeer creates IPC based peer with latency, latency is distributed with
// standard normal distribution.
func NewIPCPeer(c *Consensus, latency time.Duration) *IPCPeer {
	p := new(IPCPeer)
	p.c = c
	p.latency = latency
	p.die = make(chan struct{})
	p.minLatency = math.MaxInt64
	return p
}

// GetPublicKey returns peer's public key as identity
func (p *IPCPeer) GetPublicKey() *ecdsa.PublicKey { return &p.c.privateKey.PublicKey }

// RemoteAddr implements Peer.RemoteAddr, the address is p's memory address
func (p *IPCPeer) RemoteAddr() net.Addr { return fakeAddress(fmt.Sprint(unsafe.Pointer(p))) }

// GetMessageCount returns messages count this peer received
func (p *IPCPeer) GetMessageCount() int64 {
	p.Lock()
	defer p.Unlock()
	return p.msgCount
}

// GetBytesCount returns messages bytes count this peer received
func (p *IPCPeer) GetBytesCount() int64 {
	p.Lock()
	defer p.Unlock()
	return p.bytesCount
}

// Propose a state, awaiting to be finalized at next height.
func (p *IPCPeer) Propose(s State) {
	p.Lock()
	defer p.Unlock()
	p.c.Propose(s)
}

// GetLatestState returns latest state
func (p *IPCPeer) GetLatestState() (height uint64, round uint64, data State) {
	p.Lock()
	defer p.Unlock()
	return p.c.CurrentState()
}

// GetLatencies returns actual generated latency
func (p *IPCPeer) GetLatencies() (min time.Duration, max time.Duration, total time.Duration) {
	p.Lock()
	defer p.Unlock()
	return p.minLatency, p.maxLatency, p.totalLatency
}

// Send implements Peer.Send
func (p *IPCPeer) Send(msg []byte) error {
	delay := p.delay()
	txDelay := func() {
		p.Lock()
		defer p.Unlock()

		if p.minLatency > delay {
			p.minLatency = delay
		}

		if p.maxLatency < delay {
			p.maxLatency = delay
		}
		p.totalLatency += delay
		p.msgCount++
		p.bytesCount += int64(len(msg))

		err := p.c.ReceiveMessage(msg, time.Now())
		if err != nil {
			//		log.Println(err)
		}
	}

	timer.SystemTimedSched.Put(txDelay, time.Now().Add(delay))
	return nil
}

// delay is randomized with standard normal distribution
func (p *IPCPeer) delay() time.Duration {
	return time.Duration(0.1*rand.NormFloat64()*float64(p.latency)) + p.latency
}

// Update will call itself perodically
func (p *IPCPeer) Update() {
	p.Lock()
	defer p.Unlock()

	select {
	case <-p.die:
	default:
		// call consensus update
		_ = p.c.Update(time.Now())
		timer.SystemTimedSched.Put(p.Update, time.Now().Add(20*time.Millisecond))
	}
}

// Close this peer
func (p *IPCPeer) Close() {
	p.dieOnce.Do(func() {
		close(p.die)
	})
}
