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

package ethyl

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)


func TestCreateClient(t *testing.T) {

    testUrl := os.Getenv("TEST_RPC_URL")

    client, err := CreateClient(testUrl);
    assert.Nil(t, err);
    assert.NotNil(t, client);
    assert.Equal(t, testUrl, client.Url);

}


