/* /////////////////////////////////////////////////////////////////////////
 * File:        version.go
 *
 * Purpose:     Version file for Diagnosticism.Go
 *
 * Created:     5th March 2019
 * Updated:     22nd July 2020
 *
 * Home:        https://github.com/synesissoftware/Diagnosticism.Go
 *
 * Copyright (c) 2019-2020, Matthew Wilson
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 * - Redistributions of source code must retain the above copyright notice,
 *   this list of conditions and the following disclaimer.
 * - Redistributions in binary form must reproduce the above copyright
 *   notice, this list of conditions and the following disclaimer in the
 *   documentation and/or other materials provided with the distribution.
 * - Neither the names of Matthew Wilson and Synesis Information Systems nor
 *   the names of any contributors may be used to endorse or promote
 *   products derived from this software without specific prior written
 *   permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS
 * IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO,
 * THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR
 * PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
 * CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
 * EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
 * PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 * LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
 * NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 * ////////////////////////////////////////////////////////////////////// */


package diagnosticism

import (

	"fmt"
)

const (

	VersionMajor int16		=	0
	VersionMinor int16		=	6
	VersionPatch int16		=	0
	VersionBuild int16		=	1
	Version int64			=	(int64(VersionMajor) << 48) + (int64(VersionMinor) << 32) + (int64(VersionPatch) << 16) + (int64(VersionBuild) << 0)

	VersionRevision int16	=	VersionPatch

)

var VersionString = fmt.Sprintf("%d.%d.%d.%d", VersionMajor, VersionMinor, VersionPatch, VersionBuild)

/* ///////////////////////////// end of file //////////////////////////// */


