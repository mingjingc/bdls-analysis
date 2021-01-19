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

import "errors"

var (
	// Config Related
	ErrConfigEpoch              = errors.New("Config.Epoch is nil")
	ErrConfigStateNil           = errors.New("Config.CurrentState is nil")
	ErrConfigStateCompare       = errors.New("Config.StateCompare function has not set")
	ErrConfigStateValidate      = errors.New("Config.StateValidate function has not set")
	ErrConfigPrivateKey         = errors.New("Config.PrivateKey has not set")
	ErrConfigParticipants       = errors.New("Config.Participants must contain at least 4 participants")
	ErrConfigPubKeyToCoordinate = errors.New("Config.must contain at least 4 participants")

	// common errors related to every message
	ErrMessageVersion            = errors.New("the message has different version")
	ErrMessageValidator          = errors.New("the message has been rejected by external validator")
	ErrMessageIsEmpty            = errors.New("the message being verified is empty")
	ErrMessageUnknownMessageType = errors.New("unrecognized message type")
	ErrMessageSignature          = errors.New("cannot verify the signature of this message")
	ErrMessageUnknownParticipant = errors.New("the message is from unknown partcipants")

	// <roundchange> related
	ErrRoundChangeHeightMismatch  = errors.New("the <roundchange> message has another height than expected")
	ErrRoundChangeRoundLower      = errors.New("the <roundchange> message has lower round than expected")
	ErrRoundChangeStateValidation = errors.New("the state data validation failed <roundchange> message")

	// <lock> related
	ErrLockEmptyState              = errors.New("the state is empty in <lock> message")
	ErrLockStateValidation         = errors.New("the state data validation failed <lock> message")
	ErrLockHeightMismatch          = errors.New("the <lock> message has another height than expected")
	ErrLockRoundLower              = errors.New("the <lock> message has lower round than expected")
	ErrLockNotSignedByLeader       = errors.New("the <lock> message is not signed by leader")
	ErrLockProofUnknownParticipant = errors.New("the proofs in <lock> message has unknown participant")
	ErrLockProofTypeMismatch       = errors.New("the proofs in <lock> message is not <roundchange>")
	ErrLockProofHeightMismatch     = errors.New("the proofs in <lock> message has mismatched height")
	ErrLockProofRoundMismatch      = errors.New("the proofs in <lock> message has mismatched round")
	ErrLockProofStateValidation    = errors.New("the proofs in <lock> message has invalid state data")
	ErrLockProofInsufficient       = errors.New("the <lock> message has insufficient <roundchange> proofs to the proposed state")

	// <select> related
	ErrSelectStateValidation         = errors.New("the state data validation failed <select> message")
	ErrSelectHeightMismatch          = errors.New("the <select> message has another height than expected")
	ErrSelectRoundLower              = errors.New("the <select> message has lower round than expected")
	ErrSelectNotSignedByLeader       = errors.New("the <select> message is not signed by leader")
	ErrSelectStateMismatch           = errors.New("the <select> message has nil state but proof contains non-nil state")
	ErrSelectProofUnknownParticipant = errors.New("the proofs in <select> message has unknown participant")
	ErrSelectProofTypeMismatch       = errors.New("the proofs in <select> message is not <roundchange>")
	ErrSelectProofHeightMismatch     = errors.New("the proofs in <select> message has mismatched height")
	ErrSelectProofRoundMismatch      = errors.New("the proofs in <select> message has mismatched round")
	ErrSelectProofStateValidation    = errors.New("the proofs in <select> message has invalid state data")
	ErrSelectProofNotTheMaximal      = errors.New("the proposed state is not the maximal one in the <select> message")
	ErrSelectProofInsufficient       = errors.New("the <select> message has insufficient overall proofs")
	ErrSelectProofExceeded           = errors.New("the <select> message overall state proposals exceeded maximal")

	// <decide> Related
	ErrDecideHeightLower             = errors.New("the <decide> message has lower height than expected")
	ErrDecideEmptyState              = errors.New("the state is empty in <decide> message")
	ErrDecideStateValidation         = errors.New("the state data validation failed <decide> message")
	ErrDecideNotSignedByLeader       = errors.New("the <decide> message is not signed by leader")
	ErrDecideProofUnknownParticipant = errors.New("the proofs in <decide> message has unknown participant")
	ErrDecideProofTypeMismatch       = errors.New("the proofs in <decide> message is not <commit>")
	ErrDecideProofHeightMismatch     = errors.New("the proofs in <decide> message has mismatched height")
	ErrDecideProofRoundMismatch      = errors.New("the proofs in <decide> message has mismatched round")
	ErrDecideProofStateValidation    = errors.New("the proofs in <decide> message has invalid state data")
	ErrDecideProofInsufficient       = errors.New("the <decide> message has insufficient <commit> proofs to the proposed state")

	// <lock-release> related
	ErrLockReleaseStatus = errors.New("received <lock-release> message in non LOCK-RELEASE state")

	// <commit> related
	ErrCommitEmptyState      = errors.New("the state is empty in <commit> message")
	ErrCommitStateMismatch   = errors.New("the state in <commit> message does not match what leader has locked")
	ErrCommitStateValidation = errors.New("the state data validation failed <commit> message")
	ErrCommitStatus          = errors.New("received <commit> message in non COMMIT state")
	ErrCommitHeightMismatch  = errors.New("the <commit> messge has another height than expected")
	ErrCommitRoundMismatch   = errors.New("the <commit> message is from another round")

	// <decide> verification
	ErrMismatchedTargetState = errors.New("the state in <decide> message does not match the provided target state")
)
