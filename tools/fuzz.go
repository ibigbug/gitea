// Copyright 2020 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// +build gofuzz

package fuzz

import (
	"code.gitea.io/gitea/modules/markup"
	"code.gitea.io/gitea/modules/markup/markdown"
)

// Contains fuzzing functions executed by
// fuzzing engine https://github.com/dvyukov/go-fuzz
//
// The function must return 1 if the fuzzer should increase priority of the given input during subsequent fuzzing
// (for example, the input is lexically correct and was parsed successfully).
// -1 if the input must not be added to corpus even if gives new coverage and 0 otherwise.

func FuzzMarkdownRenderRaw(data []byte) int {
	_ = markdown.RenderRaw(data, "", false)
	return 1
}

func FuzzMarkupPostProcess(data []byte) int {
	var localMetas = map[string]string{
		"user": "go-gitea",
		"repo": "gitea",
	}
	_, err := markup.PostProcess(data, "https://example.com", localMetas, false)
	if err != nil {
		return 0
	}
	return 1
}