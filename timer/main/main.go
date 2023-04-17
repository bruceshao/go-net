/*
 * Copyright (C) THL A29 Limited, a Tencent company. All rights reserved.
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("timer start")
	t := time.After(time.Second)
	select {
	case <-t:
		fmt.Println("it is time")
	}
	time.Sleep(10 * time.Minute)
}
