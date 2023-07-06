package apidocs

const ApiTypeScript = `import http from '@/utils/request'
{{range $key, $value := .Function}}
/** {{ $value.Description }} */
export function {{ $key }}({{if $value.Body }}{{$value.Body}}?: object{{ end }}): Promise<IApiResponseData<any>> {
	return http.request<IApiResponseData<any>>({
		url: '{{ $value.Url }}',
		method: '{{ $value.Method }}',
		{{if $value.Body }}data: {{$value.Body}},{{ end }}
	})
}
{{ end }}
`
