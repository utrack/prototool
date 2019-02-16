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
	"path/filepath"

	filepkg "github.com/uber/prototool/internal/file"
	"github.com/uber/prototool/internal/settings"
	"github.com/uber/prototool/internal/strs"
)

type bazelFileSet struct {
	fileSet *fileSet
}

func newBazelFileSet(fileSet *fileSet) *bazelFileSet {
	return &bazelFileSet{
		fileSet: fileSet,
	}
}

func (b *bazelFileSet) AbsBaseDirPath() string {
	return b.fileSet.AbsBaseDirPath
}

func (b *bazelFileSet) RelDirPath() string {
	return b.fileSet.AbsBaseDirPath
}

func (b *bazelFileSet) Files() []*bazelFile {
	files := make([]*bazelFile, len(b.fileSet.FilesToGenerate))
	for i, file := range b.fileSet.FilesToGenerate {
		files[i] = newBazelFile(b, file)
	}
	return files
}

func (b *bazelFileSet) ShouldGenerateRule(ruleName string) bool {
	_, ok := b.getBazelGenRule(ruleName)
	return ok
}

// *** proto_library ***

// ProtoLibraryName returns the proto_library name.
func (b *bazelFileSet) ProtoLibraryName() string {
	return "protos"
}

// ProtoLibrarySrcs returns the proto_library srcs.
func (b *bazelFileSet) ProtoLibrarySrcs() []string {
	return []string{}
}

// ProtoLibraryDeps returns the proto_library deps.
func (b *bazelFileSet) ProtoLibraryDeps() []string {
	files := b.Files()
	deps := make([]string, 0, len(files))
	for _, file := range files {
		deps = append(deps, ":"+file.ProtoLibraryName())
	}
	return deps
}

// ProtoLibraryVisibility returns the proto_library visibility.
func (b *bazelFileSet) ProtoLibraryVisibility() []string {
	return getVisibility(b.fileSet.Config.Bazel.Gen.ProtoLibraryVisibilityAlias)
}

// *** cc_proto_library ***

// CCProtoLibraryName returns the cc_proto_library name.
func (b *bazelFileSet) CCProtoLibraryName() string {
	return "cc_protos"
}

// CCProtoLibraryDeps returns the cc_proto_library deps.
func (b *bazelFileSet) CCProtoLibraryDeps() []string {
	return []string{
		":" + b.ProtoLibraryName(),
	}
}

// CCProtoLibraryVisibility returns the cc_proto_library visibility.
func (b *bazelFileSet) CCProtoLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("cc_proto_library")
}

// *** cc_library ***

// CCLibraryName returns the cc_library name.
func (b *bazelFileSet) CCLibraryName() string {
	return "lib_cc_protos"
}

// CCLibraryDeps returns the cc_library deps.
func (b *bazelFileSet) CCLibraryDeps() []string {
	return []string{
		":" + b.CCProtoLibraryName(),
	}
}

// CCLibraryVisibility returns the cc_library visibility.
func (b *bazelFileSet) CCLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("cc_proto_library")
}

// *** java_proto_library ***

// JavaProtoLibraryName returns the java_proto_library name.
func (b *bazelFileSet) JavaProtoLibraryName() string {
	return "java_protos"
}

// JavaProtoLibraryDeps returns the java_proto_library deps.
func (b *bazelFileSet) JavaProtoLibraryDeps() []string {
	return []string{
		":" + b.ProtoLibraryName(),
	}
}

// JavaProtoLibraryVisibility returns the java_library visibility.
func (b *bazelFileSet) JavaProtoLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("java_proto_library")
}

// *** java java_library ***

// JavaJavaLibraryName returns the java_library name.
func (b *bazelFileSet) JavaJavaLibraryName() string {
	return "lib_java_protos"
}

// JavaJavaLibraryDeps returns the java_library deps.
func (b *bazelFileSet) JavaJavaLibraryDeps() []string {
	return []string{
		":" + b.JavaProtoLibraryName(),
	}
}

// JavaJavaLibraryVisibility returns the java_proto_library visibility.
func (b *bazelFileSet) JavaJavaLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("java_proto_library")
}

// *** java_lite_proto_library ***

// JavaLiteProtoLibraryName returns the java_lite_proto_library name.
func (b *bazelFileSet) JavaLiteProtoLibraryName() string {
	return "java_lite_protos"
}

// JavaLiteProtoLibraryDeps returns the java_lite_proto_library deps.
func (b *bazelFileSet) JavaLiteProtoLibraryDeps() []string {
	return []string{
		":" + b.ProtoLibraryName(),
	}
}

// JavaLiteProtoLibraryVisibility returns the java_lite_proto_library visibility.
func (b *bazelFileSet) JavaLiteProtoLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("java_lite_proto_library")
}

// *** java_lite java_library ***

// JavaLiteJavaLibraryName returns the java_library name.
func (b *bazelFileSet) JavaLiteJavaLibraryName() string {
	return "lib_java_lite_protos"
}

// JavaLiteJavaLibraryDeps returns the java_library deps.
func (b *bazelFileSet) JavaLiteJavaLibraryDeps() []string {
	return []string{
		":" + b.JavaLiteProtoLibraryName(),
	}
}

// JavaLiteJavaLibraryVisibility returns the java_proto_library visibility.
func (b *bazelFileSet) JavaLiteJavaLibraryVisibility() []string {
	return b.getBazelGenRuleVisibility("java_lite_proto_library")
}

// *** end ***

func (b *bazelFileSet) getBazelGenRuleVisibility(ruleName string) []string {
	bazelGenRule, ok := b.getBazelGenRule(ruleName)
	if !ok {
		return visibilityPublic
	}
	return getVisibility(bazelGenRule.VisibilityAlias)
}

func (b *bazelFileSet) getBazelGenRule(ruleName string) (settings.BazelGenRule, bool) {
	for _, rule := range b.fileSet.Config.Bazel.Gen.Rules {
		if rule.Name == ruleName {
			return rule, true
		}
	}
	return settings.BazelGenRule{}, false
}

type bazelFile struct {
	bazelFileSet *bazelFileSet
	file         *file
}

func newBazelFile(bazelFileSet *bazelFileSet, file *file) *bazelFile {
	return &bazelFile{
		bazelFileSet: bazelFileSet,
		file:         file,
	}
}

// ProtoLibraryName returns the proto_library name.
func (f *bazelFile) ProtoLibraryName() string {
	return protoLibraryName(f.file.RelPath)
}

// ProtoLibrarySrcs returns the proto_libary srcs.
func (f *bazelFile) ProtoLibrarySrcs() []string {
	return []string{filepath.Base(f.file.RelPath)}
}

// ProtoLibraryDeps returns the proto_library deps.
func (f *bazelFile) ProtoLibraryDeps() []string {
	var deps []string
	for _, nonWKTDepFilePath := range getNonWKTDepFilePaths(f.file.RelDepPaths) {
		if filepath.Dir(nonWKTDepFilePath) == filepath.Dir(f.file.RelPath) {
			deps = append(deps, protoLibraryDepName(filepath.Base(nonWKTDepFilePath)))
		} else {
			deps = append(deps, protoLibraryDepName(nonWKTDepFilePath))
		}
	}
	deps = append(deps, getWKTProtoLibraryDeps(f.file.RelDepPaths)...)
	return strs.SortUniq(deps)
}

// ProtoLibraryVisibility returns the proto_library visibility.
func (f *bazelFile) ProtoLibraryVisibility() []string {
	return getVisibility(f.bazelFileSet.fileSet.Config.Bazel.Gen.ProtoLibraryFileVisibilityAlias)
}

func getVisibility(visibilityAlias string) []string {
	switch visibilityAlias {
	case "public":
		return visibilityPublic
	case "subpackages":
		return visibilitySubpackages
	default:
		return visibilityPublic
	}
}

func protoLibraryName(filePath string) string {
	return filepkg.BaseStrippedFileName(filePath) + "_proto"
}

func protoLibraryDepName(filePath string) string {
	depName := ":" + protoLibraryName(filePath)
	if filepath.Base(filePath) == filePath {
		return depName
	}
	return "//" + filepath.Dir(filePath) + depName
}

// getNonWKTDepFilePaths retuns the non-WKT file paths.
func getNonWKTDepFilePaths(depFilePaths []string) []string {
	s := make([]string, 0, len(depFilePaths))
	for _, depFilePath := range depFilePaths {
		if _, ok := wktToProtoLibraryDeps[depFilePath]; !ok {
			s = append(s, depFilePath)
		}
	}
	return strs.SortUniq(s)
}

// getWKTProtoLibraryDeps gets the WKT @com_google_protobuf deps.
func getWKTProtoLibraryDeps(depFilePaths []string) []string {
	wktBazelDepMap := make(map[string]struct{})
	for _, depFilePath := range depFilePaths {
		if wktDeps, ok := wktToProtoLibraryDeps[depFilePath]; ok {
			for _, wktDep := range wktDeps {
				wktBazelDepMap[defaultWKTWorkspace+"//"+wktDep] = struct{}{}
			}
		}
	}
	return strs.MapToSortedSlice(wktBazelDepMap)
}
