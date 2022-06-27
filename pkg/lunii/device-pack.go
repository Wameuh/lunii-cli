package lunii

import (
	"github.com/google/uuid"
)

// respresentation of a lunii device pack
type DevicePack struct {
	StageNodes []PackStageNode
	ListNodes  []ListNode
	Metadata   Metadata
	Uuid       uuid.UUID
	Ref        string
}

type Metadata struct {
	Title       string
	Description string
}

type PackStageNode struct {
	Image           []byte
	Audio           []byte
	OkTransition    *PackTransition
	HomeTransition  *PackTransition
	ControlSettings *PackControlSettings
}

type PackListNode struct {
	StartsAt           int
	StageNodeIndexList []int
}

type PackControlSettings struct {
	Wheel    bool
	Ok       bool
	Home     bool
	Pause    bool
	Autoplay bool
}

type PackTransition struct {
	ActionNode  string `json:"actionNode"`
	OptionIndex int    `json:"optionIndex"`
}
