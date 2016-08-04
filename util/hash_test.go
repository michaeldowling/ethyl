/*
 * (C) Copyright 2016 Michael Dowling (http://github.com/michaeldowling) and others.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package util

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestEthereumHash(t *testing.T) {

    result, err := EthereumHash([]byte("baz(uint32,bool)"));
    assert.Nil(t, err);
    assert.NotNil(t, result);

    assert.Equal(t, "0xcdcd77c0992ec5bbfc459984220f8c45084cc24d9b6efed1fae540db8de801d2", result);

}
