package tmpl

const Api = `import http from "@/utils/request"

interface {{.StructName}} {
	{{- range .Fields}}
    {{.FieldValueName}}: any
	{{- end}}
}

/** 增 */
export function create{{.StructName}}Api(data?: object): Promise<IApiResponseData<{{.StructName}}>> {
  return http.request<IApiResponseData<{{.StructName}}>>({
    url: "/api/v1/{{.JsonName}}",
    method: "post",
    data,
  })
}

/** 改 */
export function update{{.StructName}}Api(data?: object): Promise<IApiResponseData<{{.StructName}}>> {
  return http.request<IApiResponseData<{{.StructName}}>>({
    url: "/api/v1/{{.JsonName}}",
    method: "put",
    data,
  })
}

/** 删 */
export function delete{{.StructName}}Api(id: number): Promise<IApiResponseData<{{.StructName}}>> {
  return http.request<IApiResponseData<{{.StructName}}>>({
    url: ` + "`/api/v1/{{.JsonName}}/${id}`" + `,
    method: "delete",
  })
}

/** 查 */
export function find{{.StructName}}Api(id: number): Promise<IApiResponseData<{{.StructName}}>> {
  return http.request<IApiResponseData<{{.StructName}}>>({
    url: ` + "`/api/v1/{{.JsonName}}/${id}`" + `,
    method: "get",
  })
}

/** 删除 批量操作 */
export function delete{{.StructName}}ByIdsApi(ids: number[]): Promise<IApiResponseData<{{.StructName}}>> {
  return http.request<IApiResponseData<{{.StructName}}>>({
    url: "/api/v1/{{.JsonName}}/batch_delete",
    method: "delete",
    data: ids,
  })
}

/** 查询 分页列表 */
export function find{{.StructName}}ListApi(page?: Page): Promise<IApiResponseData<PageResult<{{.StructName}}>>> {
  return http.request<IApiResponseData<PageResult<{{.StructName}}>>>({
    url: "/api/v1/{{.JsonName}}/list",
    method: "post",
    data: page,
  })
}
`
