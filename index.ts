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

export interface ModuleMemory {
	Write(offset:number, v:Uint8Array): void
	Read(offset:number, byteCount:number): Uint8Array
}

export type Resizer = (name:string, size:number) => number

export type InstallableFunc = (mem: ModuleMemory, resize: Resizer, params: number[]) => void

// Extension is an interface that must be implemented by all Scale Extensions
// for use with the Host. Guest implementations do not use this interface.
export interface Extension {
    Init(): Map<string, InstallableFunc>
    Reset(): void
}
