
Squiring code structure
``` lua
web
├── api/v1 -- Main API
|   ├── sys_initdb.go -- ico
|   └── sys_user.go --  
├── config -- Configuration file setting operation structure
|   ├── auto_code.go -- ico captcha.go
|   ├── ... -- ico captcha.go
|   └── zap.go -- core
├── core -- Main structure code
|   ├── server_other.go -- ico captcha.go
|   ├── ... -- ico captcha.go
|   └── zap.go -- 
├── docs -- Document system
|   ├── docs.go -- ico captcha.go
|   ├── swagger.json -- json
|   └── swagger.yaml -- yaml  
├── global -- global
├── initialize -- initialize 
├── middleware -- Intermediate key
├── model -- global
│   ├── request  -- All request MODEL structures
|   |   ├── common.go 
|   |   ├── ...
|   |   └── sys_user.go -- yaml  
|   ├── response  -- Return data
|   |   ├── common.go 
|   |   ├── ...
|   |   └── sys_user.go -- yaml  
├── packfile -- File writing
├── resource -- resource
├── router -- routing
├── service -- Service layer
├── source -- File directory operation 
├── utils
├── config.yaml  -- 
├── Dockerfile  -- Docker configuration
├── go.mod    -- mod Configure
├── go.sum -- sum
├── latest_log  -- vue-cli Configure
└── main.go  -- package.json
```