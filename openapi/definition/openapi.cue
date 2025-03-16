package definition

#openapi: {
	openapi: "3.0.3"
	info:    #info
	servers?: [...#server]
	paths:       #paths
	components?: #components
	security?: [...#security_requirement]
	tags?: [...#tag]
	externalDocs?: #external_docs
}

#info: {
	title:           string
	description?:    string
	termsOfService?: string
	contact?:        #contact
	license?:        #license
	version:         string
}

#contact: {
	name?:  string
	url?:   string
	email?: string
}

#license: {
	name: string
	url?: string
}

#server: {
	url:          string
	description?: string
	variables?: [string]: #server_variable
}

#server_variable: {
	enum?: [...string]
	default:      string
	description?: string
}

#components: {
	schemas?: [string]:         #schema | #reference
	responses?: [string]:       #response | #reference
	parameters?: [string]:      #parameter | #reference
	examples?: [string]:        #example | #reference
	requestBodies?: [string]:   #request_body | #reference
	headers?: [string]:         #header | #reference
	securitySchemes?: [string]: #security_scheme | #reference
	links?: [string]:           #link | #reference
	callbacks?: [string]:       #callback | #reference
}

#paths: [string]: #path

#path: {
	$ref?:        string
	summary?:     string
	description?: string
	get?:         #operation
	put?:         #operation
	post?:        #operation
	delete?:      #operation
	options?:     #operation
	head?:        #operation
	patch?:       #operation
	trace?:       #operation
	servers?: [...#server]
}

#operation: {
	tags?: [string, ...]
	summary?:      string
	description?:  string
	externalDocs?: #external_docs
	operationId?:  string
	parameters?: [...#parameter | #reference]
	requestBody?: #request_body | #reference
	responses: [string]:  #response
	callbacks?: [string]: #callback | #reference
	deprecated?: bool
	security?: [...#security_requirement]
	servers?: [...#server]
}

#external_docs: {
	description?: string
	url:          string
}

#parameter: {
	name:             string
	in:               "query" | "header" | "path" | "cookie"
	description?:     string
	required?:        bool
	deprecated?:      bool
	allowEmptyValue?: bool
	style?:           string
	explode?:         bool
	allowReserved?:   bool
	schema?:          #schema | #reference
	example?:         _
	examples?: [string]: #example | #reference
}

#request_body: {
	description?: string
	content: [string]: #media_type
	required?: bool
}

#media_type: {
	schema?:  #schema | #reference
	example?: _
	examples?: [string]: #example | #reference
	encoding?: [string]: #encoding
}

#encoding: {
	contentType?: string
	headers?: [string]: #header | #reference
	style?:         string
	explode?:       bool
	allowReserved?: bool
}

#response: {
	description: string
	headers?: [string]: #header | #reference
	content?: [string]: #media_type
	links?: [string]:   #link | #reference
}

#callback: {
	[string]: #path
}

#example: {
	summary?:       string
	description?:   string
	value?:         _
	externalValue?: string
}

#link: {
	operationRef?: string
	operationId?:  string
	parameters: [string]: _
	requestBody: _
	description: string
	server:      #server
}

#header: {
	description?:     string
	required?:        bool
	deprecated?:      bool
	allowEmptyValue?: bool
	style?:           string
	explode?:         bool
	allowReserved?:   bool
	schema?:          #schema | #reference
	example?:         _
	examples?: [string]: #example | #reference
}

#tag: {
	name:          string
	description?:  string
	externalDocs?: #external_docs
}

#reference: $ref: string

#oneOf: { 
	oneOf: [...#reference]
}

// TODO: schema is not complete (e.g. multipleOf, maximum, and etc)
#schema: #schema2 | #oneOf

#schema2: {
	type:         "null" | "string" | "number" | "integer" | "boolean" | "array" | "object"
	title?:       string
	description?: string
	items?:       #schema | #reference
	properties?: [string]: #schema | #reference
	format?:   "int32" | "int64" | "float" | "double" | "byte" | "binary" | "date" | "date-time" | "password" | "email" | "uri"
	example?:  _
	examples?:  _
	nullable?: bool
	required?: [...string]
	readOnly?:     bool
	writeOnly?:    bool
	externalDocs?: #external_docs
	deprecated?:   bool
	default?:      _
	enum?: [...string]
}

#security_scheme: #security_scheme_api_key | #security_scheme_http | #security_scheme_oauth2 | #security_scheme_openidconnect

#security_scheme_api_key: {
	type:         "apiKey"
	description?: string
	name:         string
	in:           "query" | "header" | "cookie"
}

#security_scheme_http: {
	type:          "http"
	description?:  string
	scheme:        string
	bearerFormat?: string
}

#security_scheme_oauth2: {
	type:         "oauth2"
	description?: string
	flows:        #oauth_flows
}

#security_scheme_openidconnect: {
	type:             "openIdConnect"
	description?:     string
	openIdConnectUrl: string
}

#oauth_flows: {
	implicit?:          #oauth_flow_implict
	password?:          #oauth_flow_password
	clientCredentials?: #oauth_flow_client_credentials
	authorizationCode?: #oauth_flow_authorization_code
}

#oauth_flow_implict: {
	authorizationUrl: string
	refreshUrl?:      string
	scopes: [string]: string
}

#oauth_flow_password: {
	tokenUrl:    string
	refreshUrl?: string
	scopes: [string]: string
}

#oauth_flow_client_credentials: {
	tokenUrl:    string
	refreshUrl?: string
	scopes: [string]: string
}

#oauth_flow_authorization_code: {
	authorizationUrl: string
	tokenUrl:         string
	refreshUrl?:      string
	scopes: [string]: string
}

#security_requirement: [string]: [...string]
