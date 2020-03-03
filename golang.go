//////////////////////////////////////////////////////////////////////////////
//
// (C) Copyright 2019-2020 by Unectio, Inc.
//
// The information contained herein is confidential, proprietary to Unectio,
// Inc.
//
//////////////////////////////////////////////////////////////////////////////

package api

import (
	"time"
)


type Request struct {
	RunRequest			`json:",inline"`

	/*
	 * If the request content-type is application/json, then
	 * the body will be auto-unmarshalled and provided into
	 * the Main() via this pointer (cast it to *Body)
	 */
	B	interface{}		`json:"-"`
}

type Response struct {
	/*
	 * By default the callet sets http.StatusOK (200) as the
	 * response status. If you want to change it, set the needed
	 * value here.
	 */
	Status	int
	/*
	 * Setting this will make the event to be called again on the
	 * function after the specified duration with the same req
	 * values.
	 */
	Defer	time.Duration
}
