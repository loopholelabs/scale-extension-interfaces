/*
	Copyright 2023 Loophole Labs

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

package interfaces

type ModuleMemory interface {
	Write(offset uint32, v []byte) bool
	Read(offset uint32, byteCount uint32) ([]byte, bool)
}

type Resizer func(name string, size uint64) (uint64, error)

type InstallableFunc func(mem ModuleMemory, resize Resizer, params []uint64)

// Extension is an interface that must be implemented by all Scale Extensions
// for use with the Host. Guest implementations do not use this interface.
type Extension interface {
	Init() map[string]InstallableFunc // Called to init. Returns map of functions to insert into runtime.
	Reset()                           // This is called by the runtime before each run.
}
