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
//	"testing"
//
//	"github.com/yubo/client-go/kubernetes/scheme"
//	v1 "github.com/yubo/client-go/tools/clientcmd/api/v1"
//	"github.com/yubo/golib/runtime"
//	runtimejson "github.com/yubo/golib/runtime/serializer/json"
//)
//
//// getEncoder mimics how github.com/yubo/client-go/rest.createSerializers creates a encoder
//func getEncoder() runtime.Encoder {
//	jsonSerializer := runtimejson.NewSerializer(runtimejson.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, false)
//	directCodecFactory := scheme.Codecs.WithoutConversion()
//	return directCodecFactory.EncoderForVersion(jsonSerializer, v1.SchemeGroupVersion)
//}
//
//func TestEncodeDecodeRoundTrip(t *testing.T) {
//}
