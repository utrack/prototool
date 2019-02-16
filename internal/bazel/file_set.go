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
	"fmt"
	"path/filepath"
	"sort"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/uber/prototool/internal/protoc"
	"github.com/uber/prototool/internal/settings"
)

// fileSet is a set of file for a single FileDescriptorSet.
//
// All files will be within the same directory.
type fileSet struct {
	// Config is the config.
	Config settings.Config `json:"-"`
	// AbsBaseDirPath is the absolute path of the directory that prefixes
	// all relative paths in the fileSet. This is the path of the configuration file.
	AbsBaseDirPath string
	// RelDirPath is the path relative to AbsBaseDirPath of the directory
	// that all files in the fileSet are within.
	RelDirPath string
	// RelFilePathsToGenerate are the paths of the files that are within
	// RelDirPath to generate for. All directories will be equal to RelDirPath.
	// This list will be sorted.
	RelFilePathsToGenerate []string
	// RelFilePathToFile are the paths of all files in the FileDescriptorSet.
	// This will include files not in this directory.
	// All paths in RelFilePathsToGenerate will be in this map.
	RelFilePathToFile map[string]*file
	// FilesToGenerate are the files within RelDirPath to generate for.
	// This is effectively duplicate information to the above two fields
	// however as we are building this library out, we keep them all.
	FilesToGenerate []*file
}

// file represents a single file.
type file struct {
	// FileSet is the containing fileSet.
	FileSet *fileSet `json:"-"`
	// RelPath is the path releative to AbsBaseDirPath.
	RelPath string
	// RelDepPaths are the paths of the dependencies relative to AbsBaseDirPath.
	// This list will be sorted.
	RelDepPaths []string
}

// newFileSet returns a new fileSet.
func newFileSet(fileDescriptorSet *protoc.FileDescriptorSet) (*fileSet, error) {
	fileSet := &fileSet{
		Config:            fileDescriptorSet.ProtoSet.Config,
		AbsBaseDirPath:    fileDescriptorSet.ProtoSet.Config.DirPath,
		RelFilePathToFile: make(map[string]*file),
	}
	// this should never return error as we validated this in ProtoSetProvider
	// this could however be "."
	relDirPath, err := filepath.Rel(fileDescriptorSet.ProtoSet.Config.DirPath, fileDescriptorSet.DirPath)
	if err != nil {
		return nil, err
	}
	fileSet.RelDirPath = relDirPath

	relFilePathToFileDescriptorProto, err := getCleanedFileDescriptorProtoMap(fileDescriptorSet)
	if err != nil {
		return nil, err
	}
	for relFilePath, fileDescriptorProto := range relFilePathToFileDescriptorProto {
		fileSet.RelFilePathToFile[relFilePath] = newFile(fileSet, relFilePath, fileDescriptorProto)
	}

	relFilePathToGenerateMap := make(map[string]struct{})
	// ProtoFiles are the files to generate for
	for _, protoFile := range fileDescriptorSet.ProtoFiles {
		// this should never return error as we validated this in ProtoSetProvider
		// this could however be "."
		relFilePath, err := filepath.Rel(fileDescriptorSet.ProtoSet.Config.DirPath, protoFile.Path)
		if err != nil {
			return nil, err
		}
		fileSet.RelFilePathsToGenerate = append(fileSet.RelFilePathsToGenerate, relFilePath)
		if _, ok := relFilePathToGenerateMap[relFilePath]; ok {
			return nil, fmt.Errorf("duplicate ProtoFile %s, this is a system error", relFilePath)
		}
		relFilePathToGenerateMap[relFilePath] = struct{}{}
		file, ok := fileSet.RelFilePathToFile[relFilePath]
		if !ok {
			return nil, fmt.Errorf("no FileDescriptorProto for file %s, this is a system error", relFilePath)
		}
		fileSet.FilesToGenerate = append(fileSet.FilesToGenerate, file)
	}
	sort.Strings(fileSet.RelFilePathsToGenerate)
	sort.Slice(fileSet.FilesToGenerate, func(i int, j int) bool {
		return fileSet.FilesToGenerate[i].RelPath < fileSet.FilesToGenerate[j].RelPath
	})

	return fileSet, nil
}

func newFile(fileSet *fileSet, relPath string, fileDescriptorProto *descriptor.FileDescriptorProto) *file {
	file := &file{
		FileSet: fileSet,
		RelPath: relPath,
	}
	file.RelDepPaths = append(file.RelDepPaths, fileDescriptorProto.Dependency...)
	sort.Strings(file.RelDepPaths)
	return file
}

func getCleanedFileDescriptorProtoMap(fileDescriptorSet *protoc.FileDescriptorSet) (map[string]*descriptor.FileDescriptorProto, error) {
	m := make(map[string]*descriptor.FileDescriptorProto, len(fileDescriptorSet.File))
	for _, fileDescriptorProto := range fileDescriptorSet.File {
		cleanFileDescriptorProto(fileDescriptorProto)
		if err := validateFileDescriptorProto(fileDescriptorProto); err != nil {
			return nil, err
		}
		name := fileDescriptorProto.GetName()
		if name == "" {
			return nil, fmt.Errorf("nil FileDescriptorProto name")
		}
		if _, ok := m[name]; ok {
			return nil, fmt.Errorf("duplicate FileDescriptorProto name: %s", name)
		}
		m[name] = fileDescriptorProto
	}
	return m, nil
}

func cleanFileDescriptorProto(fileDescriptorProto *descriptor.FileDescriptorProto) {
	if fileDescriptorProto.Name != nil {
		fileDescriptorProto.Name = proto.String(filepath.Clean(fileDescriptorProto.GetName()))
	}
	for i, dependency := range fileDescriptorProto.Dependency {
		fileDescriptorProto.Dependency[i] = filepath.Clean(dependency)
	}
}

func validateFileDescriptorProto(fileDescriptorProto *descriptor.FileDescriptorProto) error {
	if err := validateFilePath(fileDescriptorProto.GetName()); err != nil {
		return err
	}
	for _, dependency := range fileDescriptorProto.Dependency {
		if err := validateFilePath(dependency); err != nil {
			return err
		}
	}
	return nil
}

func validateFilePath(filePath string) error {
	if filePath == "" {
		return fmt.Errorf("empty filePath")
	}
	if filepath.IsAbs(filePath) {
		return fmt.Errorf("%s cannot be an absolute path", filePath)
	}
	if filepath.Ext(filePath) != ".proto" {
		return fmt.Errorf("%s is not a valid proto file", filePath)
	}
	return nil
}
