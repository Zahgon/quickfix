// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package file

import (
	"os"

	"github.com/quickfixgo/quickfix"
)

func sessionIDFilenamePrefix(s quickfix.SessionID) string { _ = "STUB: not implemented"; return "" }

// openOrCreateFile opens a file for reading and writing, creating it if necessary.
func openOrCreateFile(fname string, perm os.FileMode) (f *os.File, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
