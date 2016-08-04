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
    "log"
    "strconv"
    "os"
)


func TestAPINetVersion(t *testing.T) {

    testUrl := os.Getenv("TEST_RPC_URL");

    client, clientErr := CreateClient(testUrl);
    assert.Nil(t, clientErr);

    version, err := client.Net.Version();

    assert.NotNil(t, version);
    assert.Equal(t, "100", version);
    assert.Nil(t, err);

}

func TestAPINetIsListening(t *testing.T) {

    testUrl := os.Getenv("TEST_RPC_URL");

    client, clientErr := CreateClient(testUrl);
    assert.Nil(t, clientErr);

    isListening, err := client.Net.IsListening();
    log.Println("Is lIstening:  " + strconv.FormatBool(isListening));

    assert.NotNil(t, isListening);
    assert.True(t, isListening);
    assert.Nil(t, err);


}

func TestAPINetPeerCount(t *testing.T) {

    testUrl := os.Getenv("TEST_RPC_URL");

    client, clientErr := CreateClient(testUrl);
    assert.Nil(t, clientErr);

    peerCount, _ := client.Net.PeerCount();
    log.Printf("Peer Count:  %v", peerCount);

    assert.True(t, peerCount >= 0);

}

