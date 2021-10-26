package util

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"

	"github.com/go-pay/wechat-sdk/pkg/xlog"
)

func TestHexToString(t *testing.T) {
	h := sha256.New()
	h.Write([]byte("hsSuSUsePBqSQw2rYMtf9Nvha603xX8f2BMQBcYRoJiMNwOqt/UEhrqekebG5ar0LFNAm5MD4Uz6zorRwiXJwbySJ/FEJHav4NsobBIU1PwdjbJWVQLFy7+YFkHB32OnQXWMh6ugW7Dyk2KS5BXp1f5lniKPp1KNLyNLlFlNZ2mgJCJmWvHj5AI7BLpWwoRvqRyZvVXo+9FsWqvBdxmAPA=="))
	toString := hex.EncodeToString(h.Sum(nil))
	xlog.Debugf("hex: %s", toString)
}
