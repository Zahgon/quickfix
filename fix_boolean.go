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

// FIXBoolean is a FIX Boolean value, implements FieldValue.
type FIXBoolean bool

// Bool converts the FIXBoolean value to bool.
func (f FIXBoolean) Bool() bool { _ = "STUB: not implemented"; return false }

func (f *FIXBoolean) Read(bytes []byte) error { _ = "STUB: not implemented"; return nil }

func (f FIXBoolean) Write() []byte { _ = "STUB: not implemented"; return nil }
