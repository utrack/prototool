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

var (
	// visibilityPublic is the visibility list for public.
	visibilityPublic = []string{"//visibility:public"}
	// visibilitySubpackages is the visibility list for subpackages.
	visibilitySubpackages = []string{"//:__subpackages__"}

	// defaultWKTWorkspace is the workspace for the Well-Known Types
	//
	// This is pulled out in case we want to make this an option later.
	defaultWKTWorkspace = "@com_google_protobuf"

	// wktToProtoLibraryDeps is the map from Well-Known Type file to
	// the proto_library dependencies for it.
	wktToProtoLibraryDeps = map[string][]string{
		"google/protobuf/any.proto": []string{
			":any_proto",
		},
		"google/protobuf/api.proto": []string{
			":api_proto",
			":source_context_proto",
			":type_proto",
		},
		"google/protobuf/compiler/plugin.proto": []string{
			":compiler_plugin_proto",
			":descriptor_proto",
		},
		"google/protobuf/descriptor.proto": []string{
			":descriptor_proto",
		},
		"google/protobuf/duration.proto": []string{
			":duration_proto",
		},
		"google/protobuf/empty.proto": []string{
			":empty_proto",
		},
		"google/protobuf/field_mask.proto": []string{
			":field_mask_proto",
		},
		"google/protobuf/source_context.proto": []string{
			":source_context_proto",
		},
		"google/protobuf/struct.proto": []string{
			":struct_proto",
		},
		"google/protobuf/timestamp.proto": []string{
			":timestamp_proto",
		},
		"google/protobuf/type.proto": []string{
			":any_proto",
			":source_context_proto",
			":type_proto",
		},
		"google/protobuf/wrappers.proto": []string{
			":wrappers_proto",
		},
	}
)
