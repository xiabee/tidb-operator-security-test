// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package tidbcluster

import (
	"fmt"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/pingcap/tidb-operator/pkg/apis/pingcap/v1alpha1"
	"github.com/pingcap/tidb-operator/pkg/controller"
	apps "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/cache"
)

func TestTidbClusterControllerEnqueueTidbCluster(t *testing.T) {
	g := NewGomegaWithT(t)
	tc := newTidbCluster()
	tcc := NewController(controller.NewFakeDependencies())
	tcc.control = NewFakeTidbClusterControlInterface()
	tcc.enqueueTidbCluster(tc)
	g.Expect(tcc.queue.Len()).To(Equal(1))
}

func TestTidbClusterControllerEnqueueTidbClusterFailed(t *testing.T) {
	g := NewGomegaWithT(t)
	tcc := NewController(controller.NewFakeDependencies())
	tcc.control = NewFakeTidbClusterControlInterface()
	tcc.enqueueTidbCluster(struct{}{})
	g.Expect(tcc.queue.Len()).To(Equal(0))
}

func TestTidbClusterControllerAddStatefulSet(t *testing.T) {
	g := NewGomegaWithT(t)
	type testcase struct {
		name                    string
		modifySet               func(*v1alpha1.TidbCluster) *apps.StatefulSet
		addTidbClusterToIndexer bool
		expectedLen             int
	}

	testFn := func(test *testcase, t *testing.T) {
		t.Log("test: ", test.name)

		tc := newTidbCluster()
		set := test.modifySet(tc)

		fakeDeps := controller.NewFakeDependencies()
		tcc := NewController(fakeDeps)
		tcc.control = NewFakeTidbClusterControlInterface()
		tcIndexer := fakeDeps.InformerFactory.Pingcap().V1alpha1().TidbClusters().Informer().GetIndexer()
		if test.addTidbClusterToIndexer {
			err := tcIndexer.Add(tc)
			g.Expect(err).NotTo(HaveOccurred())
		}
		tcc.addStatefulSet(set)
		g.Expect(tcc.queue.Len()).To(Equal(test.expectedLen))
	}

	tests := []testcase{
		{
			name: "normal",
			modifySet: func(tc *v1alpha1.TidbCluster) *apps.StatefulSet {
				return newStatefulSet(tc)
			},
			addTidbClusterToIndexer: true,
			expectedLen:             1,
		},
		{
			name: "have deletionTimestamp",
			modifySet: func(tc *v1alpha1.TidbCluster) *apps.StatefulSet {
				set := newStatefulSet(tc)
				set.DeletionTimestamp = &metav1.Time{Time: time.Now().Add(30 * time.Second)}
				return set
			},
			addTidbClusterToIndexer: true,
			expectedLen:             1,
		},
		{
			name: "without controllerRef",
			modifySet: func(tc *v1alpha1.TidbCluster) *apps.StatefulSet {
				set := newStatefulSet(tc)
				set.OwnerReferences = nil
				return set
			},
			addTidbClusterToIndexer: true,
			expectedLen:             0,
		},
		{
			name: "without tidbcluster",
			modifySet: func(tc *v1alpha1.TidbCluster) *apps.StatefulSet {
				return newStatefulSet(tc)
			},
			addTidbClusterToIndexer: false,
			expectedLen:             0,
		},
	}

	for i := range tests {
		testFn(&tests[i], t)
	}
}

func TestTidbClusterControllerUpdateStatefulSet(t *testing.T) {
	g := NewGomegaWithT(t)
	type testcase struct {
		name                    string
		updateSet               func(*apps.StatefulSet) *apps.StatefulSet
		addTidbClusterToIndexer bool
		expectedLen             int
	}

	testFn := func(test *testcase, t *testing.T) {
		t.Log("test: ", test.name)

		tc := newTidbCluster()
		set1 := newStatefulSet(tc)
		set2 := test.updateSet(set1)

		fakeDeps := controller.NewFakeDependencies()
		tcc := NewController(fakeDeps)
		tcc.control = NewFakeTidbClusterControlInterface()
		tcIndexer := fakeDeps.InformerFactory.Pingcap().V1alpha1().TidbClusters().Informer().GetIndexer()
		if test.addTidbClusterToIndexer {
			err := tcIndexer.Add(tc)
			g.Expect(err).NotTo(HaveOccurred())
		}
		tcc.updateStatefulSet(set1, set2)
		g.Expect(tcc.queue.Len()).To(Equal(test.expectedLen))
	}

	tests := []testcase{
		{
			name: "normal",
			updateSet: func(set1 *apps.StatefulSet) *apps.StatefulSet {
				set2 := *set1
				set2.ResourceVersion = "1000"
				return &set2
			},
			addTidbClusterToIndexer: true,
			expectedLen:             1,
		},
		{
			name: "same resouceVersion",
			updateSet: func(set1 *apps.StatefulSet) *apps.StatefulSet {
				set2 := *set1
				return &set2
			},
			addTidbClusterToIndexer: true,
			expectedLen:             0,
		},
		{
			name: "without controllerRef",
			updateSet: func(set1 *apps.StatefulSet) *apps.StatefulSet {
				set2 := *set1
				set2.ResourceVersion = "1000"
				set2.OwnerReferences = nil
				return &set2
			},
			addTidbClusterToIndexer: true,
			expectedLen:             0,
		},
		{
			name: "without tidbcluster",
			updateSet: func(set1 *apps.StatefulSet) *apps.StatefulSet {
				set2 := *set1
				set2.ResourceVersion = "1000"
				return &set2
			},
			addTidbClusterToIndexer: false,
			expectedLen:             0,
		},
	}

	for i := range tests {
		testFn(&tests[i], t)
	}
}

func TestTidbClusterControllerSync(t *testing.T) {
	g := NewGomegaWithT(t)
	type testcase struct {
		name                     string
		addTcToIndexer           bool
		errWhenUpdateTidbCluster bool
		errExpectFn              func(*GomegaWithT, error)
	}

	testFn := func(test *testcase, t *testing.T) {
		t.Log(test.name)

		tc := newTidbCluster()
		fakeDeps := controller.NewFakeDependencies()
		tcc := NewController(fakeDeps)
		tcIndexer := fakeDeps.InformerFactory.Pingcap().V1alpha1().TidbClusters().Informer().GetIndexer()
		tcControl := NewFakeTidbClusterControlInterface()
		tcc.control = tcControl
		if test.addTcToIndexer {
			err := tcIndexer.Add(tc)
			g.Expect(err).NotTo(HaveOccurred())
		}
		key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(tc)
		g.Expect(err).NotTo(HaveOccurred())

		if test.errWhenUpdateTidbCluster {
			tcControl.SetUpdateTCError(fmt.Errorf("update tidb cluster failed"))
		}

		err = tcc.sync(key)

		if test.errExpectFn != nil {
			test.errExpectFn(g, err)
		}
	}

	tests := []testcase{
		{
			name:                     "normal",
			addTcToIndexer:           true,
			errWhenUpdateTidbCluster: false,
			errExpectFn: func(g *GomegaWithT, err error) {
				g.Expect(err).NotTo(HaveOccurred())
			},
		},
		{
			name:                     "can't found tidb cluster",
			addTcToIndexer:           false,
			errWhenUpdateTidbCluster: false,
			errExpectFn: func(g *GomegaWithT, err error) {
				g.Expect(err).NotTo(HaveOccurred())
			},
		},
		{
			name:                     "update tidb cluster failed",
			addTcToIndexer:           true,
			errWhenUpdateTidbCluster: true,
			errExpectFn: func(g *GomegaWithT, err error) {
				g.Expect(err).To(HaveOccurred())
				g.Expect(strings.Contains(err.Error(), "update tidb cluster failed")).To(Equal(true))
			},
		},
	}

	for i := range tests {
		testFn(&tests[i], t)
	}

}

func newTidbCluster() *v1alpha1.TidbCluster {
	return &v1alpha1.TidbCluster{
		TypeMeta: metav1.TypeMeta{
			Kind:       "TidbCluster",
			APIVersion: "pingcap.com/v1alpha1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-pd",
			Namespace: corev1.NamespaceDefault,
			UID:       types.UID("test"),
		},
		Spec: v1alpha1.TidbClusterSpec{
			PD: &v1alpha1.PDSpec{
				ComponentSpec: v1alpha1.ComponentSpec{
					Image: "pd-test-image",
				},
				ResourceRequirements: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse("10G"),
					},
				},
			},
			TiKV: &v1alpha1.TiKVSpec{
				Replicas: 1,
				ComponentSpec: v1alpha1.ComponentSpec{
					Image: "tikv-test-image",
				},
				ResourceRequirements: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse("10G"),
					},
				},
			},
			TiDB: &v1alpha1.TiDBSpec{
				ComponentSpec: v1alpha1.ComponentSpec{
					Image: "tidb-test-image",
				},
			},
			TiFlash: &v1alpha1.TiFlashSpec{
				Replicas: 1,
				ComponentSpec: v1alpha1.ComponentSpec{
					Image: "tiflash-test-image",
				},
				ResourceRequirements: corev1.ResourceRequirements{
					Requests: corev1.ResourceList{
						corev1.ResourceStorage: resource.MustParse("10G"),
					},
				},
			},
		},
	}
}

func newStatefulSet(tc *v1alpha1.TidbCluster) *apps.StatefulSet {
	return &apps.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "StatefulSet",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-statefulset",
			Namespace: corev1.NamespaceDefault,
			UID:       types.UID("test"),
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(tc, controller.ControllerKind),
			},
			ResourceVersion: "1",
		},
		Spec: apps.StatefulSetSpec{
			Replicas: &tc.Spec.PD.Replicas,
		},
	}
}
