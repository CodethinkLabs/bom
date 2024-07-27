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

package provenance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const testProvenanceFile1 = "testdata/k8s-1.23.0-alpha.4-provenance.json"

func TestReadStatment(t *testing.T) {
	s, err := LoadStatement(testProvenanceFile1)
	require.NoError(t, err)
	require.NotNil(t, s)

	require.NotNil(t, s.Predicate)
	require.Len(t, s.Subject, 461)
	require.Equal(t, "pkg:github/puerco/release@provenance", s.Predicate.Builder.ID)
	require.Equal(t, "94db9bed6b7c56420e722d1b15db4610c9cacd3f", s.Predicate.Materials[0].Digest["sha1"])
	require.Equal(t, "git+https://github.com/kubernetes/kubernetes", s.Predicate.Materials[0].URI)
}
