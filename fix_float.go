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

package quickfix

// FIXFloat is a FIX Float Value, implements FieldValue.
type FIXFloat float64

// Float64 converts the FIXFloat value to float64.
func (f FIXFloat) Float64() float64 { _ = "STUB: not implemented"; return 0 }

func (f *FIXFloat) Read(bytes []byte) error { _ = "STUB: not implemented"; return nil }

// `strconv` allows values like "+100.00", which is not allowed for FIX float types.

func (f FIXFloat) Write() []byte { _ = "STUB: not implemented"; return nil }

func isDecimal(b byte) bool { _ = "STUB: not implemented"; return false }
