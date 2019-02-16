// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bazel

import (
	"github.com/uber/prototool/internal/protoc"
	"go.uber.org/zap"
)

// Output contains the resulting files to write.
type Output struct {
	// Absolute.
	// Cleaned.
	FilePath string
	// Will contain no extra whitespace.
	Data []byte
}

// Generator generates bazel targets.
type Generator interface {
	// Generate generates the output files.
	//
	// Outputs will be sorted by FilePath.
	Generate(protoc.FileDescriptorSets) ([]*Output, error)
}

// GeneratorOption is an option for a new Generator.
type GeneratorOption func(*generator)

// GeneratorWithLogger returns a GeneratorOption that uses the given logger.
//
// The default is to use zap.NewNop().
func GeneratorWithLogger(logger *zap.Logger) GeneratorOption {
	return func(generator *generator) {
		generator.logger = logger
	}
}

// NewGenerator returns a new Generator.
func NewGenerator(options ...GeneratorOption) Generator {
	return newGenerator(options...)
}
