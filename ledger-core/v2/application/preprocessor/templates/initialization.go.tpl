//
// Copyright 2019 Insolar Technologies GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by insgocc. DO NOT EDIT.
// source template in logicrunner/preprocessor/templates

package {{ .Package }}

import (
{{ range $name, $path := .CustomImports }}
    {{ $name }} {{ $path }}
{{- end }}

{{ range $contract := .Contracts }}
    {{ $contract.ImportName }} "{{ $contract.ImportPath }}"
{{- end }}
)

func InitializeContractMethods() map[string]XXX_insolar.ContractWrapper {
    return map[string]XXX_insolar.ContractWrapper{
{{- range $contract := .Contracts }}
        "{{ $contract.Name }}": {{ $contract.ImportName }}.Initialize(),
{{- end }}
    }
}

func shouldLoadRef(strRef string) XXX_reference.Global {
    ref, err := XXX_reference.GlobalFromString(strRef)
    if err != nil {
        panic(errors.Wrap(err, "Unexpected error, bailing out"))
    }
    return ref
}

func InitializeCodeRefs() map[XXX_reference.Global]string {
    rv := make(map[XXX_reference.Global]string, {{ len .Contracts }})

    {{ range $contract := .Contracts -}}
    rv[shouldLoadRef("{{ $contract.CodeReference }}")] = "{{ $contract.Name }}"
    {{ end }}

    return rv
}

func InitializePrototypeRefs() map[XXX_reference.Global]string {
    rv := make(map[XXX_reference.Global]string, {{ len .Contracts }})

    {{ range $contract := .Contracts -}}
    rv[shouldLoadRef("{{ $contract.PrototypeReference }}")] = "{{ $contract.Name }}"
    {{ end }}

    return rv
}

func InitializeCodeDescriptors() []XXX_descriptor.Code {
    rv := make([]XXX_descriptor.Code, 0, {{ len .Contracts }})

    {{ range $contract := .Contracts -}}
    // {{ $contract.Name }}
    rv = append(rv, XXX_descriptor.NewCode(
        /* code:        */ nil,
        /* machineType: */ XXX_machine.Builtin,
        /* ref:         */ shouldLoadRef("{{ $contract.CodeReference }}"),
    ))
    {{ end }}
    return rv
}

func InitializePrototypeDescriptors() []XXX_descriptor.Prototype {
    rv := make([]XXX_descriptor.Prototype, 0, {{ len .Contracts }})

    {{ range $contract := .Contracts }}
    { // {{ $contract.Name }}
        pRef := shouldLoadRef("{{ $contract.PrototypeReference }}")
        cRef := shouldLoadRef("{{ $contract.CodeReference }}")
        rv = append(rv, XXX_descriptor.NewPrototype(
            /* head:         */ pRef,
            /* state:        */ pRef.GetLocal(),
            /* code:         */ cRef,
        ))
    }
    {{ end }}
    return rv
}
