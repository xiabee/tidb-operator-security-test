// Copyright 2019 PingCAP, Inc.
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

package tidbmonitor

import (
	"fmt"
	"time"

	perrors "github.com/pingcap/errors"
	"github.com/pingcap/tidb-operator/pkg/controller"
	"github.com/pingcap/tidb-operator/pkg/metrics"
	"github.com/pingcap/tidb-operator/pkg/monitor/monitor"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
)

// Controller syncs TidbMonitor
type Controller struct {
	deps    *controller.Dependencies
	control ControlInterface
	// tidbMonitor that need to be synced.
	queue workqueue.RateLimitingInterface
}

// NewController creates a tidbmonitor controller.
func NewController(deps *controller.Dependencies) *Controller {
	c := &Controller{
		deps:    deps,
		control: NewDefaultTidbMonitorControl(deps, monitor.NewMonitorManager(deps)),
		queue: workqueue.NewNamedRateLimitingQueue(
			controller.NewControllerRateLimiter(1*time.Second, 100*time.Second),
			"tidbmonitor",
		),
	}

	tidbMonitorInformer := deps.InformerFactory.Pingcap().V1alpha1().TidbMonitors()
	statefulsetInformer := deps.KubeInformerFactory.Apps().V1().StatefulSets()
	controller.WatchForObject(tidbMonitorInformer.Informer(), c.queue)
	controller.WatchForController(statefulsetInformer.Informer(), c.queue, func(ns, name string) (runtime.Object, error) {
		return c.deps.TiDBMonitorLister.TidbMonitors(ns).Get(name)
	}, nil)

	return c
}

// Name returns the name of the controller
func (c *Controller) Name() string {
	return "tidbmonitor"
}

func (c *Controller) Run(workers int, stopCh <-chan struct{}) {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	klog.Info("Starting tidbmonitor controller")
	defer klog.Info("Shutting down tidbmonitor controller")

	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}

	<-stopCh
}

func (c *Controller) worker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem dequeues items, processes them, and marks them done. It enforces that the syncHandler is never
// invoked concurrently with the same key.
func (c *Controller) processNextWorkItem() bool {
	metrics.ActiveWorkers.WithLabelValues(c.Name()).Add(1)
	defer metrics.ActiveWorkers.WithLabelValues(c.Name()).Add(-1)

	key, quit := c.queue.Get()
	if quit {
		return false
	}
	defer c.queue.Done(key)
	if err := c.sync(key.(string)); err != nil {
		if perrors.Find(err, controller.IsRequeueError) != nil {
			klog.Infof("TidbMonitor: %v, still need sync: %v, requeuing", key.(string), err)
		} else {
			utilruntime.HandleError(fmt.Errorf("TidbMonitor: %v, sync failed, err: %v", key.(string), err))
		}
		c.queue.AddRateLimited(key)
	} else {
		c.queue.Forget(key)
	}
	return true
}

func (c *Controller) sync(key string) (err error) {
	startTime := time.Now()
	defer func() {
		duration := time.Since(startTime)
		metrics.ReconcileTime.WithLabelValues(c.Name()).Observe(duration.Seconds())

		if err == nil {
			metrics.ReconcileTotal.WithLabelValues(c.Name(), metrics.LabelSuccess).Inc()
		} else if perrors.Find(err, controller.IsRequeueError) != nil {
			metrics.ReconcileTotal.WithLabelValues(c.Name(), metrics.LabelRequeue).Inc()
		} else {
			metrics.ReconcileTotal.WithLabelValues(c.Name(), metrics.LabelError).Inc()
			metrics.ReconcileErrors.WithLabelValues(c.Name()).Inc()
		}

		klog.V(4).Infof("Finished syncing TidbMonitor %q (%v)", key, duration)
	}()

	ns, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		return err
	}
	tm, err := c.deps.TiDBMonitorLister.TidbMonitors(ns).Get(name)
	if errors.IsNotFound(err) {
		klog.Infof("TidbMonitor has been deleted %v", key)
		return nil
	}
	if err != nil {
		return err
	}

	return c.control.ReconcileTidbMonitor(tm)
}
