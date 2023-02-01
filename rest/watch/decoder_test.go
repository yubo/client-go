/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package versioned_test

//import (
//	"io"
//	"testing"
//	"time"
//
//	"github.com/yubo/client-go/kubernetes/scheme"
//	restclientwatch "github.com/yubo/client-go/rest/watch"
//	"github.com/yubo/golib/runtime"
//	runtimejson "github.com/yubo/golib/runtime/serializer/json"
//	"github.com/yubo/golib/runtime/serializer/streaming"
//	"github.com/yubo/golib/util/wait"
//)
//
//// getDecoder mimics how github.com/yubo/client-go/rest.createSerializers creates a decoder
//func getDecoder() runtime.Decoder {
//	jsonSerializer := runtimejson.NewSerializer(runtimejson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, false)
//	directCodecFactory := scheme.Codecs.WithoutConversion()
//	return directCodecFactory.DecoderToVersion(jsonSerializer)
//}
//
//func TestDecoder_SourceClose(t *testing.T) {
//	out, in := io.Pipe()
//	decoder := restclientwatch.NewDecoder(streaming.NewDecoder(out, getDecoder()), getDecoder())
//
//	done := make(chan struct{})
//
//	go func() {
//		_, _, err := decoder.Decode()
//		if err == nil {
//			t.Errorf("Unexpected nil error")
//		}
//		close(done)
//	}()
//
//	in.Close()
//
//	select {
//	case <-done:
//		break
//	case <-time.After(wait.ForeverTestTimeout):
//		t.Error("Timeout")
//	}
//}
