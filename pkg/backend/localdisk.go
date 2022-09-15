/*
 * Copyright (c) 2022. Nydus Developers. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package backend

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/errdefs"
	"github.com/opencontainers/go-digest"
	"github.com/pkg/errors"
)

type LocalDiskBackend struct {
	device_path string
}

func newLocalDiskBackend(rawConfig []byte) (*LocalDiskBackend, error) {
	var configMap map[string]string
	if err := json.Unmarshal(rawConfig, &configMap); err != nil {
		return nil, errors.Wrap(err, "parse LocalDisk storage backend configuration")
	}

	device_path, ok := configMap["device_path"]
	if !ok {
		return nil, fmt.Errorf("no `device_path` option is specified")
	}

	return &LocalDiskBackend{
		device_path: device_path,
	}, nil
}

//TODO
func (b *LocalDiskBackend) Push(ctx context.Context, ra content.ReaderAt, blobDigest digest.Digest) error {

	return nil
}

//TODO
func (b *LocalDiskBackend) Check(blobDigest digest.Digest) (string, error) {

	return "", errdefs.ErrNotFound
}

func (b *LocalDiskBackend) Type() string {
	return BackendTypeLocalDisk
}
