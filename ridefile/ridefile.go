// Copyright 2013 Andrew O'Neill. All rights reserved.  Use of this file is
// governed by a BSD-style license that can be found in the LICENSE file.

// Package ridefile implements a simple ride file interface for reading
// exported ride data.
package ridefile

type RideFile interface {
	Read() []byte, error
}
