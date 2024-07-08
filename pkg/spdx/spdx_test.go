/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package spdx_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/bom/pkg/spdx"
	"sigs.k8s.io/bom/pkg/spdx/spdxfakes"
)

var err = errors.New("synthetic error")

func TestPackageFromImageTarball(t *testing.T) {
	for _, tc := range []struct {
		prepare     func(*spdxfakes.FakeSpdxImplementation)
		shouldError bool
	}{
		{ // success
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.PackageFromImageTarballReturns(&spdx.Package{Entity: spdx.Entity{Name: "test"}}, nil)
			},
			shouldError: false,
		},
		{ // PackageFromImageTarball fails
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.PackageFromImageTarballReturns(nil, err)
			},
			shouldError: true,
		},
	} {
		sut := spdx.NewSPDX()
		sut.Options().AnalyzeLayers = false
		mock := &spdxfakes.FakeSpdxImplementation{}
		tc.prepare(mock)
		sut.SetImplementation(mock)
		// Run the test function
		pkg, err := sut.PackageFromImageTarball("mock.tar")
		if tc.shouldError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
			require.NotNil(t, pkg)
		}
	}
}

func TestExtractTarballTmp(t *testing.T) {
	for _, tc := range []struct {
		prepare     func(*spdxfakes.FakeSpdxImplementation)
		shouldError bool
	}{
		{ // success
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.ExtractTarballTmpReturns("/mock/path", nil)
			},
			shouldError: false,
		},
		{ // error
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.ExtractTarballTmpReturns("/mock/path", err)
			},
			shouldError: true,
		},
	} {
		sut := spdx.NewSPDX()
		mock := &spdxfakes.FakeSpdxImplementation{}
		tc.prepare(mock)
		sut.SetImplementation(mock)

		path, err := sut.ExtractTarballTmp("/mock/path")
		if tc.shouldError {
			require.Error(t, err)
		} else {
			require.NotEmpty(t, path)
			require.NoError(t, err)
		}
	}
}

func TestPullImagesToArchive(t *testing.T) {
	for _, tc := range []struct {
		prepare     func(*spdxfakes.FakeSpdxImplementation)
		shouldError bool
	}{
		{ // success
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.PullImagesToArchiveReturns(nil, nil)
			},
			shouldError: false,
		},
		{ // success
			prepare: func(mock *spdxfakes.FakeSpdxImplementation) {
				mock.PullImagesToArchiveReturns(nil, err)
			},
			shouldError: true,
		},
	} {
		sut := spdx.NewSPDX()
		sut.Options().AnalyzeLayers = false
		mock := &spdxfakes.FakeSpdxImplementation{}
		tc.prepare(mock)
		sut.SetImplementation(mock)

		_, err := sut.PullImagesToArchive("mock-image:latest", "/tmp")
		if tc.shouldError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
